//go:build compareopus

package opuscc

import (
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

func compareCustomDecoderImage(data []byte) (*OpusT_OpusCustomDecoder, compareDecoderImage) {
	image := compareDecoderImageStorage(reflect.TypeFor[OpusT_OpusCustomDecoder](), data)
	return image.owner.Field(0).Addr().Interface().(*OpusT_OpusCustomDecoder), image
}
