//go:build compareopus

package opuscc

import (
	"encoding/binary"
	"reflect"
	"runtime"
	"unsafe"
)

// Serialized native fixtures contain scalar state and null pointer slots. They
// are transport bytes, not Go decoder allocations. Install them into an actual
// scanned layout before temporarily attaching any Go mode/codebook owners.
type compareDecoderImage struct {
	owner reflect.Value
	bytes []byte
	slots [][2]int
}

func compareDecoderPointerSlots(t reflect.Type, base int, slots *[][2]int) {
	switch t.Kind() {
	case reflect.Pointer, reflect.UnsafePointer:
		*slots = append(*slots, [2]int{base, base + int(t.Size())})
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			compareDecoderPointerSlots(f.Type, base+int(f.Offset), slots)
		}
	case reflect.Array:
		for i := 0; i < t.Len(); i++ {
			compareDecoderPointerSlots(t.Elem(), base+i*int(t.Elem().Size()), slots)
		}
	}
}

func compareDecoderImageStorage(layout reflect.Type, data []byte) compareDecoderImage {
	if len(data) < int(layout.Size()) {
		panic("short serialized decoder fixture")
	}
	if (len(data)-int(layout.Size()))%4 != 0 {
		panic("serialized decoder fixture scalar extent mismatch")
	}
	fields := []reflect.StructField{{Name: "Image", Type: layout}}
	if tail := len(data) - int(layout.Size()); tail != 0 {
		fields = append(fields, reflect.StructField{Name: "Tail", Type: reflect.ArrayOf(tail, reflect.TypeFor[byte]())})
	}
	t := reflect.StructOf(fields)
	// Go rounds a pointer-bearing struct's size to its alignment. Some C
	// variable tails end four bytes earlier. Never expose that Go padding in
	// the transport view or change a fixture's supplied length/capacity.
	if int(t.Size()) != (len(data)+t.Align()-1)&^(t.Align()-1) {
		panic("serialized decoder fixture alignment mismatch")
	}
	image := compareDecoderImage{owner: reflect.New(t).Elem()}
	compareDecoderPointerSlots(layout, 0, &image.slots)
	image.checkNullPointers(data)
	// This byte view serializes a real typed allocation; it never supplies a
	// decoder pointer or creates a GC map for caller-provided byte storage.
	image.bytes = unsafe.Slice((*byte)(image.owner.Addr().UnsafePointer()), len(data))
	copy(image.bytes, data)
	return image
}

func (image compareDecoderImage) checkNullPointers(data []byte) {
	for _, slot := range image.slots {
		for _, b := range data[slot[0]:slot[1]] {
			if b != 0 {
				panic("serialized decoder fixture contains a non-null pointer")
			}
		}
	}
}

func (image compareDecoderImage) save(data []byte) {
	image.checkNullPointers(image.bytes)
	copy(data, image.bytes)
	runtime.KeepAlive(image.owner)
}

func compareMSDecoderGeometry(data []byte) (channels, streams, coupled int32) {
	header := reflect.TypeFor[OpusT_OpusMSDecoder]()
	layout, _ := header.FieldByName("Flayout")
	t := reflect.TypeFor[OpusT_ChannelLayout]()
	read := func(name string) int32 {
		field, _ := t.FieldByName(name)
		offset := int(layout.Offset + field.Offset)
		return int32(binary.NativeEndian.Uint32(data[offset : offset+4]))
	}
	channels, streams, coupled = read("Fnb_channels"), read("Fnb_streams"), read("Fnb_coupled_streams")
	if channels < 1 || channels > 255 || streams < 1 || streams > 255 || coupled < 0 || coupled > streams || streams+coupled > 255 {
		panic("noncanonical serialized multistream geometry")
	}
	return
}

func compareMSDecoderImage(data []byte) (*OpusT_OpusMSDecoder, compareDecoderImage) {
	_, streams, coupled := compareMSDecoderGeometry(data)
	image := compareDecoderImageStorage(opusMSDecoderStorageType(streams, coupled), data)
	views := opusMSStorageViewsAt(image.owner.Field(0))
	compareCheckMSImage(views, coupled)
	return views.Decoder, image
}

func compareCheckMSImage(views opusMSStorageViews, coupled int32) {
	for i, codec := range views.Codecs {
		channels := int32(1)
		if int32(i) < coupled {
			channels = 2
		}
		layout := opusDecoderStorageType(channels)
		silk, _ := layout.FieldByName("Silk")
		celt, _ := layout.FieldByName("Celt")
		if codec.Decoder.Fchannels != channels || int64(codec.Decoder.Fsilk_dec_offset) != int64(silk.Offset) || int64(codec.Decoder.Fcelt_dec_offset) != int64(celt.Offset) {
			panic("noncanonical serialized multistream child")
		}
	}
}

func compareOpusDecoderImage(data []byte) (*OpusT_OpusDecoder, compareDecoderImage) {
	header := reflect.TypeFor[OpusT_OpusDecoder]()
	channel, _ := header.FieldByName("Fchannels")
	channels := int32(binary.NativeEndian.Uint32(data[int(channel.Offset) : int(channel.Offset)+4]))
	if channels != 1 && channels != 2 {
		panic("noncanonical serialized decoder channels")
	}
	layout := opusDecoderStorageType(channels)
	image := compareDecoderImageStorage(layout, data)
	state := image.owner.Field(0).FieldByName("Decoder").Addr().Interface().(*OpusT_OpusDecoder)
	silk, _ := layout.FieldByName("Silk")
	celt, _ := layout.FieldByName("Celt")
	if int64(state.Fsilk_dec_offset) != int64(silk.Offset) || int64(state.Fcelt_dec_offset) != int64(celt.Offset) {
		panic("noncanonical serialized decoder offsets")
	}
	return state, image
}

func compareCustomDecoderImage(data []byte) (*OpusT_OpusCustomDecoder, compareDecoderImage) {
	image := compareDecoderImageStorage(reflect.TypeFor[OpusT_OpusCustomDecoder](), data)
	return image.owner.Field(0).Addr().Interface().(*OpusT_OpusCustomDecoder), image
}
