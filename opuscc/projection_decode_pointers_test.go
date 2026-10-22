package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"testing"
	"unsafe"
)

// From projection_float_pointers_test.go.
type projectionFloatTestOwner struct {
	Projection    OpusT_OpusProjectionDecoder
	Padding       [(8 - unsafe.Sizeof(OpusT_OpusProjectionDecoder{})%8) % 8]byte
	Matrix        OpusT_MappingMatrix
	MatrixPadding [(8 - unsafe.Sizeof(OpusT_MappingMatrix{})%8) % 8]byte
	Coefficients  [16]int16
	MS            msTwoStreamTestOwner
}

func newProjectionFloatOwner(t *testing.T) *projectionFloatTestOwner {
	t.Helper()
	owner := new(projectionFloatTestOwner)
	owner.MS = *newTwoStreamMSOwner(t)
	owner.MS.MS.Flayout.Fnb_channels = 4
	copy(owner.MS.MS.Flayout.Fmapping[:], []byte{0, 1, 2, 3})
	owner.Projection.Fdemixing_matrix_size_in_bytes = Opus_mapping_matrix_get_size(4, 4)
	if unsafe.Offsetof(owner.Matrix) != 8 || unsafe.Offsetof(owner.MS) != uintptr(owner.Projection.Fdemixing_matrix_size_in_bytes)+8 || get_multistream_decoder(nil, &owner.Projection) != &owner.MS.MS || Opus_mapping_matrix_get_data(nil, &owner.Matrix) != &owner.Coefficients[0] {
		t.Fatal("projection scanned geometry")
	}
	coeffs := [16]int16{16384, 4096, -8192, 2048, 8192, 16384, 2048, -4096, -4096, 8192, 16384, 4096, 2048, -2048, 8192, 16384}
	Opus_mapping_matrix_init(nil, &owner.Matrix, 4, 4, 0, &coeffs[0], 32)
	return owner
}

func TestProjectionFloatWrapperPointers(t *testing.T) {
	owner, baseline := newProjectionFloatOwner(t), newProjectionFloatOwner(t)
	packet := make([]byte, 129)
	packet[0] = 252
	for i := 1; i < len(packet); i++ {
		packet[i] = byte((i-1)*73 + 165)
	}
	combined := append([]byte{packet[0], 128}, packet[1:]...)
	combined = append(combined, packet...)
	out, want := make([]float32, 5760*4+2), make([]float32, 5760*4+2)
	out[0], out[len(out)-1] = 77, 88
	for _, step := range []int{0, 1, 2} {
		var data *byte
		length, fec := int32(0), int32(0)
		if step != 1 {
			data = &combined[0]
			length = int32(len(combined))
		}
		if step == 2 {
			fec = 1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := Opus_opus_projection_decode_float_typed(nil, &owner.Projection, data, length, &out[1], 5760, fec)
		copyOut := func(tls *libc.TLS, ds, dc int32, src *float32, ss, n int32) {
			opus_projection_copy_channel_out_float(tls, &want[1], ds, dc, src, ss, n, &baseline.Matrix)
		}
		expected := opusMSDecodeBound(nil, &baseline.MS.MS, data, length, copyOut, 5760, fec, 0)
		if got != expected || got <= 0 || owner.MS.Children[0].Decoder != baseline.MS.Children[0].Decoder || owner.MS.Children[1].Decoder != baseline.MS.Children[1].Decoder || out[0] != 77 || out[len(out)-1] != 88 {
			t.Fatal("projection float dispatch/state")
		}
		for i := int32(0); i < got*4; i++ {
			if out[i+1] != want[i+1] {
				t.Fatal("projection float PCM", i)
			}
		}
	}
	if opusProjectionDecodeFloat(nil, &owner.Projection, nil, 0, nil, 0, 0) != -1 {
		t.Fatal("projection minimum frame")
	}
	bad := []byte{3, 0}
	sample := float32(77)
	if opusProjectionDecodeFloat(nil, &owner.Projection, &bad[0], 2, &sample, 480, 0) != -4 || sample != 77 {
		t.Fatal("projection invalid packet writes")
	}
}

func TestProjectionFloatCallbackOwnerPointers(t *testing.T) {
	owner := newProjectionFloatOwner(t)
	dst := []float32{77, 9, 9, 9, 9, 9, 9, 9, 9, 88}
	copyOut := opusProjectionFloatCopy(&owner.Matrix, &dst[1])
	owner = nil
	entropyInitGrowStack(12)
	runtime.GC()
	src := []float32{1, 2}
	copyOut(nil, 4, 0, &src[0], 1, 2)
	if dst[0] != 77 || dst[9] != 88 || dst[1] != .5 || dst[2] != .125 || dst[3] != -.25 || dst[4] != .0625 || dst[5] != 1 || dst[6] != .25 || dst[7] != -.5 || dst[8] != .125 {
		t.Fatal("captured projection matrix", dst)
	}
	copyOut(nil, 4, 0, nil, 0, 2)
	for _, value := range dst[1:9] {
		if value != 0 {
			t.Fatal("projection nil source clear")
		}
	}
}

func TestProjectionFloatPointers(t *testing.T) {
	opus_projection_copy_channel_out_float(nil, nil, 0, 0, nil, 0, 0, nil)
	m := mappingTestBuffer{Header: OpusT_MappingMatrix{Frows: 2, Fcols: 2}, Data: [36]int16{16384, 8192, -8192, 16384}}
	src := [2]float32{1, 2}
	dst := [6]float32{77, 9, 9, 9, 9, 88}
	opus_projection_copy_channel_out_float(nil, &dst[1], 2, 0, &src[0], 1, 2, &m.Header)
	opus_projection_copy_channel_out_float(nil, &dst[1], 2, 1, &src[0], 1, 2, &m.Header)
	if dst != [6]float32{77, .25, .75, .5, 1.5, 88} {
		t.Fatal(dst)
	}
	opus_projection_copy_channel_out_float(nil, &dst[1], 2, 0, nil, 1, 2, nil)
	if dst != [6]float32{77, 0, 0, 0, 0, 88} {
		t.Fatal("clear", dst)
	}
}

// From projection_int24_pointers_test.go.
func TestProjectionInt24WrapperPointers(t *testing.T) {
	owner, baseline := newProjectionFloatOwner(t), newProjectionFloatOwner(t)
	packet := make([]byte, 129)
	packet[0] = 252
	for i := 1; i < len(packet); i++ {
		packet[i] = byte((i-1)*73 + 165)
	}
	combined := append([]byte{packet[0], 128}, packet[1:]...)
	combined = append(combined, packet...)
	out, want := make([]int32, 5760*4+2), make([]int32, 5760*4+2)
	out[0], out[len(out)-1] = 77, 88
	for _, step := range []int{0, 1, 2} {
		var data *byte
		length, fec := int32(0), int32(0)
		if step != 1 {
			data = &combined[0]
			length = int32(len(combined))
		}
		if step == 2 {
			fec = 1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := Opus_opus_projection_decode24_typed(nil, &owner.Projection, data, length, &out[1], 5760, fec)
		copyOut := func(tls *libc.TLS, ds, dc int32, src *float32, ss, n int32) {
			opus_projection_copy_channel_out_int24(tls, &want[1], ds, dc, src, ss, n, &baseline.Matrix)
		}
		expected := opusMSDecodeBound(nil, &baseline.MS.MS, data, length, copyOut, 5760, fec, 0)
		if got != expected || got <= 0 || owner.MS.Children[0].Decoder != baseline.MS.Children[0].Decoder || owner.MS.Children[1].Decoder != baseline.MS.Children[1].Decoder || out[0] != 77 || out[len(out)-1] != 88 {
			t.Fatal("projection int24 dispatch/state")
		}
		for i := int32(0); i < got*4; i++ {
			if out[i+1] != want[i+1] {
				t.Fatal("projection int24 PCM", i)
			}
		}
	}
	if opusProjectionDecodeInt24(nil, &owner.Projection, nil, 0, nil, 0, 0) != -1 {
		t.Fatal("projection int24 minimum frame")
	}
	bad := []byte{3, 0}
	sample := int32(77)
	if opusProjectionDecodeInt24(nil, &owner.Projection, &bad[0], 2, &sample, 480, 0) != -4 || sample != 77 {
		t.Fatal("projection int24 invalid packet writes")
	}
}

func TestProjectionInt24Pointers(t *testing.T) {
	opus_projection_copy_channel_out_int24(nil, nil, 0, 0, nil, 0, 0, nil)
	m := mappingTestBuffer{Header: OpusT_MappingMatrix{Frows: 2, Fcols: 2}, Data: [36]int16{16384, 8192, -8192, 16384}}
	src := [2]float32{1, 2}
	dst := [6]int32{77, 9, 9, 9, 9, 88}
	opus_projection_copy_channel_out_int24(nil, &dst[1], 2, 0, &src[0], 1, 2, &m.Header)
	opus_projection_copy_channel_out_int24(nil, &dst[1], 2, 1, &src[0], 1, 2, &m.Header)
	if dst != [6]int32{77, 2097152, 6291456, 4194304, 12582912, 88} {
		t.Fatal(dst)
	}
	opus_projection_copy_channel_out_int24(nil, &dst[1], 2, 1, nil, 1, 2, nil)
	if dst[1] != 2097152 {
		t.Fatal("nonfirst nil must not clear")
	}
	opus_projection_copy_channel_out_int24(nil, &dst[1], 2, 0, nil, 1, 2, nil)
	if dst != [6]int32{77, 0, 0, 0, 0, 88} {
		t.Fatal("clear", dst)
	}
}

// From projection_short_pointers_test.go.
func TestProjectionShortWrapperPointers(t *testing.T) {
	owner, baseline := newProjectionFloatOwner(t), newProjectionFloatOwner(t)
	packet := make([]byte, 129)
	packet[0] = 252
	for i := 1; i < len(packet); i++ {
		packet[i] = byte((i-1)*73 + 165)
	}
	combined := append([]byte{packet[0], 128}, packet[1:]...)
	combined = append(combined, packet...)
	for i := range owner.MS.Children {
		owner.MS.Children[i].Decoder.Fdecode_gain = 32767
		baseline.MS.Children[i].Decoder.Fdecode_gain = 32767
	}
	out, want := make([]int16, 5760*4+2), make([]int16, 5760*4+2)
	out[0], out[len(out)-1] = 77, 88
	for _, step := range []int{0, 1, 2} {
		var data *byte
		length, fec := int32(0), int32(0)
		if step != 1 {
			data = &combined[0]
			length = int32(len(combined))
		}
		if step == 2 {
			fec = 1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := Opus_opus_projection_decode_typed(nil, &owner.Projection, data, length, &out[1], 5760, fec)
		copyOut := func(tls *libc.TLS, ds, dc int32, src *float32, ss, n int32) {
			opus_projection_copy_channel_out_short(tls, &want[1], ds, dc, src, ss, n, &baseline.Matrix)
		}
		expected := opusMSDecodeBound(nil, &baseline.MS.MS, data, length, copyOut, 5760, fec, OPTIONAL_CLIP)
		if got != expected || got <= 0 || owner.MS.Children[0].Decoder != baseline.MS.Children[0].Decoder || owner.MS.Children[1].Decoder != baseline.MS.Children[1].Decoder || out[0] != 77 || out[len(out)-1] != 88 {
			t.Fatal("projection short dispatch/clipping/state")
		}
		for i := int32(0); i < got*4; i++ {
			if out[i+1] != want[i+1] {
				t.Fatal("projection short PCM", i)
			}
		}
	}
	if opusProjectionDecodeShort(nil, &owner.Projection, nil, 0, nil, 0, 0) != -1 {
		t.Fatal("projection short minimum frame")
	}
	bad := []byte{3, 0}
	sample := int16(77)
	if opusProjectionDecodeShort(nil, &owner.Projection, &bad[0], 2, &sample, 480, 0) != -4 || sample != 77 {
		t.Fatal("projection short invalid packet writes")
	}
}

func TestProjectionShortPointers(t *testing.T) {
	opus_projection_copy_channel_out_short(nil, nil, 0, 0, nil, 0, 0, nil)
	m := mappingTestBuffer{Header: OpusT_MappingMatrix{Frows: 2, Fcols: 2}, Data: [36]int16{16384, 8192, -8192, 16384}}
	src := [2]float32{.5, .25}
	dst := [6]int16{77, 9, 9, 9, 9, 88}
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 0, &src[0], 1, 2, &m.Header)
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 1, &src[0], 1, 2, &m.Header)
	if dst != [6]int16{77, 4096, 12288, 2048, 6144, 88} {
		t.Fatal(dst)
	}
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 1, nil, 1, 2, nil)
	if dst[1] != 4096 {
		t.Fatal("nonfirst nil must not clear")
	}
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 0, nil, 1, 2, nil)
	if dst != [6]int16{77, 0, 0, 0, 0, 88} {
		t.Fatal("clear", dst)
	}
}

// From projection_matrix_pointers_test.go.
func TestProjectionMatrixPointers(t *testing.T) {
	var matrix *OpusT_MappingMatrix
	func() {
		backing := make([]uint64, 64)
		header := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(&backing[0]))
		header.Fdemixing_matrix_size_in_bytes = 128
		matrix = get_dec_demixing_matrix(nil, header)
		if unsafe.Pointer(matrix) != unsafe.Add(unsafe.Pointer(header), 8) {
			t.Fatal("offset")
		}
		matrix.Frows = 3
		matrix.Fcols = 2
		matrix.Fgain = 77
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if matrix.Frows != 3 || matrix.Fcols != 2 || matrix.Fgain != 77 {
		t.Fatal(*matrix)
	}
	// The returned interior pointer retains coefficients in the same allocation.
	data := Opus_mapping_matrix_get_data(nil, matrix)
	*data = 1234
	runtime.GC()
	if *data != 1234 {
		t.Fatal("ownership")
	}
}

// From projection_multistream_pointers_test.go.
func TestProjectionNumericOffsetPointers(t *testing.T) {
	for _, tc := range []struct{ size, want int32 }{
		{0, 8}, {4, 8}, {8, 16}, {12, 16},
		{0x7ffffff0, 0x7ffffff8}, {0x7fffffff, -2147483640}, {-1, 8}, {-16, -8},
	} {
		state := OpusT_OpusProjectionDecoder{Fdemixing_matrix_size_in_bytes: tc.size}
		if got := opusProjectionMSOffset(&state); got != uint(tc.want) {
			t.Fatal("projection numeric narrowing/wrapping", tc.size, got, uint(tc.want))
		}
	}
}

func TestProjectionMultistreamPointers(t *testing.T) {
	var ms *OpusT_OpusMSDecoder
	func() {
		b := make([]uint64, 128)
		h := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(&b[0]))
		h.Fdemixing_matrix_size_in_bytes = 129
		ms = get_multistream_decoder(nil, h)
		if unsafe.Pointer(ms) != unsafe.Add(unsafe.Pointer(h), 136) {
			t.Fatal("alignment")
		}
		ms.Flayout.Fnb_channels = 7
		ms.Flayout.Fmapping[6] = 77
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if ms.Flayout.Fnb_channels != 7 || ms.Flayout.Fmapping[6] != 77 {
		t.Fatal(ms)
	}
}
