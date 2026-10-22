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

func opusMSDecoderStorageType(streams, coupled int32) reflect.Type {
	header := reflect.TypeFor[OpusT_OpusMSDecoder]()
	fields := []reflect.StructField{{Name: "Decoder", Type: header}}
	padding := uint64(opusAlignSize8(int32(header.Size()))) - uint64(header.Size())
	if padding != 0 {
		fields = append(fields, reflect.StructField{Name: "HeaderPadding", Type: reflect.ArrayOf(int(padding), reflect.TypeFor[byte]())})
	}
	appendChildren := func(name string, count, channels int32) {
		if count == 0 {
			return
		}
		image := opusDecoderStorageType(channels)
		stride := uint64(opusAlignSize8(Opus_opus_decoder_get_size(nil, channels)))
		childFields := []reflect.StructField{{Name: "Image", Type: image}}
		if uint64(image.Size()) < stride {
			childFields = append(childFields, reflect.StructField{Name: "StridePadding", Type: reflect.ArrayOf(int(stride-uint64(image.Size())), reflect.TypeFor[byte]())})
		}
		child := reflect.StructOf(childFields)
		if uint64(child.Size()) != stride {
			panic("opus multistream Go child stride mismatch")
		}
		fields = append(fields, reflect.StructField{Name: name, Type: reflect.ArrayOf(int(count), child)})
	}
	appendChildren("Coupled", coupled, 2)
	appendChildren("Mono", streams-coupled, 1)
	layout := reflect.StructOf(fields)
	if uint64(layout.Size()) != uint64(Opus_opus_multistream_decoder_get_size(nil, streams, coupled)) {
		panic("opus multistream Go storage size mismatch")
	}
	return layout
}

func opusProjectionDecoderStorageType(channels, streams, coupled int32) reflect.Type {
	matrixSize := Opus_mapping_matrix_get_size(nil, channels, streams+coupled)
	matrixFields := []reflect.StructField{
		{Name: "Matrix", Type: reflect.TypeFor[OpusT_MappingMatrix]()},
		{Name: "HeaderPadding", Type: reflect.ArrayOf(4, reflect.TypeFor[byte]())},
	}
	if matrixSize > 16 {
		matrixFields = append(matrixFields, reflect.StructField{Name: "Coefficients", Type: reflect.ArrayOf(int((matrixSize-16)/2), reflect.TypeFor[int16]())})
	}
	matrix := reflect.StructOf(matrixFields)
	layout := reflect.StructOf([]reflect.StructField{
		{Name: "Decoder", Type: reflect.TypeFor[OpusT_OpusProjectionDecoder]()},
		{Name: "HeaderPadding", Type: reflect.ArrayOf(4, reflect.TypeFor[byte]())},
		{Name: "Demixing", Type: matrix},
		{Name: "Multistream", Type: opusMSDecoderStorageType(streams, coupled)},
	})
	mf, _ := layout.FieldByName("Demixing")
	msf, _ := layout.FieldByName("Multistream")
	if uint64(matrix.Size()) != uint64(matrixSize) || mf.Offset != 8 || uint64(msf.Offset) != uint64(8+matrixSize) || uint64(layout.Size()) != uint64(Opus_opus_projection_decoder_get_size(nil, channels, streams, coupled)) {
		panic("opus projection Go storage layout mismatch")
	}
	return layout
}

type opusProjectionStorageViews struct {
	Decoder           *OpusT_OpusProjectionDecoder
	Matrix            *OpusT_MappingMatrix
	Coefficients      *int16
	Multistream       opusMSStorageViews
	MultistreamOffset uint // actual reflection-field displacement
}

// Explicit flat ABI/image storage; normal projection creation uses Go fields.
func opusProjectionDecoderLegacyOwnedStorage(channels, streams, coupled int32) opusProjectionStorageViews {
	owner := opusDecoderAlignedStorage(opusProjectionDecoderStorageType(channels, streams, coupled))
	matrix := owner.FieldByName("Demixing")
	msField, _ := owner.Type().FieldByName("Multistream")
	views := opusProjectionStorageViews{
		Decoder:           owner.FieldByName("Decoder").Addr().Interface().(*OpusT_OpusProjectionDecoder),
		Matrix:            matrix.FieldByName("Matrix").Addr().Interface().(*OpusT_MappingMatrix),
		Multistream:       opusMSStorageViewsAt(owner.FieldByName("Multistream")),
		MultistreamOffset: uint(msField.Offset),
	}
	if coefficients := matrix.FieldByName("Coefficients"); coefficients.IsValid() {
		views.Coefficients = coefficients.Index(0).Addr().Interface().(*int16)
	}
	return views
}

func opusProjectionDecoderStorage(channels, streams, coupled int32) *OpusT_OpusProjectionDecoder {
	return opusProjectionDecoderLegacyOwnedStorage(channels, streams, coupled).Decoder
}

type opusMSStorageViews struct {
	Decoder  *OpusT_OpusMSDecoder
	Children []*OpusT_OpusDecoder
	Codecs   []opusDecoderStorageViews
	Offsets  []uint // numeric offsets from actual Go fields, not pointer subtraction
}

func opusMSStorageViewsAt(owner reflect.Value) opusMSStorageViews {
	views := opusMSStorageViews{Decoder: owner.FieldByName("Decoder").Addr().Interface().(*OpusT_OpusMSDecoder)}
	for _, name := range []string{"Coupled", "Mono"} {
		array := owner.FieldByName(name)
		if !array.IsValid() {
			continue
		}
		arrayField, _ := owner.Type().FieldByName(name)
		imageField, _ := array.Type().Elem().FieldByName("Image")
		for i := 0; i < array.Len(); i++ {
			image := array.Index(i).FieldByName("Image")
			child := opusDecoderStorageViewsAt(image)
			child.Clear = opusDecoderStorageClearAt(image)
			views.Children = append(views.Children, child.Decoder)
			views.Codecs = append(views.Codecs, child)
			views.Offsets = append(views.Offsets, uint(arrayField.Offset)+uint(i)*uint(array.Type().Elem().Size())+uint(imageField.Offset))
		}
	}
	return views
}

// Disabled DRED still has a concrete 12-byte scalar ABI header. Use its real
// Go type and the existing alignment policy, not a cast malloc byte image.
func opusDREDDecoderStorage() *OpusT_OpusDREDDecoder {
	owner := opusDecoderAlignedStorage(reflect.TypeFor[OpusT_OpusDREDDecoder]())
	return owner.Addr().Interface().(*OpusT_OpusDREDDecoder)
}

type opusDecoderStorageViews struct {
	Decoder *OpusT_OpusDecoder
	Silk    *OpusT_silk_decoder
	Celt    *OpusT_OpusCustomDecoder
	Clear   func()
}

func opusDecoderStorageViewsAt(owner reflect.Value) opusDecoderStorageViews {
	return opusDecoderStorageViews{
		Decoder: owner.FieldByName("Decoder").Addr().Interface().(*OpusT_OpusDecoder),
		Silk:    owner.FieldByName("Silk").Addr().Interface().(*OpusT_silk_decoder),
		Celt:    owner.FieldByName("Celt").Addr().Interface().(*OpusT_OpusCustomDecoder),
	}
}

// Clear declared fields, not the enclosing image: final Go padding can be
// outside the ABI extent (and inside a multistream child's aligned stride).
// SetZero uses each actual field type and its pointer write barriers.
func opusDecoderStorageClearAt(owner reflect.Value) func() {
	return func() {
		for i := 0; i < owner.NumField(); i++ {
			owner.Field(i).SetZero()
		}
	}
}

// Explicit flat-layout oracle for legacy/image interoperability tests.
func opusDecoderLegacyOwnedStorage(channels int32) opusDecoderStorageViews {
	owner := opusDecoderAlignedStorage(opusDecoderStorageType(channels))
	views := opusDecoderStorageViewsAt(owner)
	views.Clear = opusDecoderStorageClearAt(owner)
	return views
}

// Normal decoder allocation uses declared Go fields, never runtime reflection.
func opusDecoderOwnedStorage(channels int32) opusDecoderStorageViews {
	opusDecoderGoStorageCheck(channels)
	if channels == 2 {
		return opusGoStereoStorage(nil)
	}
	return opusGoMonoStorage(nil)
}

func opusDecoderStorage(channels int32) *OpusT_OpusDecoder {
	return opusDecoderOwnedStorage(channels).Decoder
}
