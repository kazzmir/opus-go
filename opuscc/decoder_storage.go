package opuscc

import "reflect"

// These layouts describe actual Go allocations, including the runtime GC map.
// Caller-supplied ABI byte buffers are not retroactively made pointer-scanned.
func opusDecoderStorageType(channels int32) reflect.Type {
	header := reflect.TypeFor[OpusT_OpusDecoder]()
	silk := reflect.TypeFor[OpusT_silk_decoder]()
	celt := reflect.TypeFor[OpusT_OpusCustomDecoder]()
	silkOffset := uint64(opusAlignSize8(int32(header.Size())))
	celtOffset := silkOffset + uint64(opusAlignSize8(int32(silk.Size())))
	needed := uint64(Opus_opus_decoder_get_size(nil, channels))
	fields := []reflect.StructField{{Name: "Decoder", Type: header}}
	appendPad := func(name string, bytes uint64) {
		if bytes != 0 {
			fields = append(fields, reflect.StructField{Name: name, Type: reflect.ArrayOf(int(bytes), reflect.TypeFor[byte]())})
		}
	}
	appendPad("HeaderPadding", silkOffset-uint64(header.Size()))
	fields = append(fields, reflect.StructField{Name: "Silk", Type: silk})
	appendPad("SilkPadding", celtOffset-silkOffset-uint64(silk.Size()))
	fields = append(fields, reflect.StructField{Name: "Celt", Type: celt})
	tail := needed - celtOffset - uint64(celt.Size())
	fields = append(fields, reflect.StructField{Name: "Tail", Type: reflect.ArrayOf(int(tail/4), reflect.TypeFor[float32]())})
	layout := reflect.StructOf(fields)
	sf, _ := layout.FieldByName("Silk")
	cf, _ := layout.FieldByName("Celt")
	if uint64(sf.Offset) != silkOffset || uint64(cf.Offset) != celtOffset || tail%4 != 0 || uint64(layout.Size()) < needed {
		panic("opus decoder Go storage layout mismatch")
	}
	return layout
}

// Preserve the allocator boundary's 16-byte alignment without casting an
// aligned byte image into a pointer-bearing struct. If the first allocation is
// not aligned, a scanned array of canonical images covers every alignment residue.
func opusDecoderAlignedStorage(layout reflect.Type) reflect.Value {
	owner := reflect.New(layout).Elem()
	if owner.Addr().Pointer()&15 == 0 {
		return owner
	}
	unit := layout.Align()
	padding := (unit + 16 - int(layout.Size()%16)) % 16
	fields := []reflect.StructField{{Name: "Image", Type: layout}}
	if padding != 0 {
		fields = append(fields, reflect.StructField{Name: "AlignmentPadding", Type: reflect.ArrayOf(padding, reflect.TypeFor[byte]())})
	}
	element := reflect.StructOf(fields)
	images := reflect.New(reflect.ArrayOf(16/unit, element)).Elem()
	for i := 0; i < images.Len(); i++ {
		image := images.Index(i).Field(0)
		if image.Addr().Pointer()&15 == 0 {
			return image
		}
	}
	panic("opus decoder Go storage alignment mismatch")
}

func opusDecoderStorage(channels int32) *OpusT_OpusDecoder {
	owner := opusDecoderAlignedStorage(opusDecoderStorageType(channels))
	return owner.FieldByName("Decoder").Addr().Interface().(*OpusT_OpusDecoder)
}
