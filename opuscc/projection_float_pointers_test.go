package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"testing"
	"unsafe"
)

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
	owner.Projection.Fdemixing_matrix_size_in_bytes = Opus_mapping_matrix_get_size(nil, 4, 4)
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
		got := opusProjectionDecodeFloat(nil, &owner.Projection, data, length, &out[1], 5760, fec)
		copyOut := func(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
			opus_projection_copy_channel_out_float(tls, (*float32)(dst), ds, dc, src, ss, n, &baseline.Matrix)
		}
		expected := opusMSDecodeNative(nil, &baseline.MS.MS, data, length, unsafe.Pointer(&want[1]), copyOut, 5760, fec, 0)
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
	copyOut := opusProjectionFloatCopy(&owner.Matrix)
	owner = nil
	entropyInitGrowStack(12)
	runtime.GC()
	src := []float32{1, 2}
	dst := []float32{77, 9, 9, 9, 9, 9, 9, 9, 9, 88}
	copyOut(nil, unsafe.Pointer(&dst[1]), 4, 0, &src[0], 1, 2)
	if dst[0] != 77 || dst[9] != 88 || dst[1] != .5 || dst[2] != .125 || dst[3] != -.25 || dst[4] != .0625 || dst[5] != 1 || dst[6] != .25 || dst[7] != -.5 || dst[8] != .125 {
		t.Fatal("captured projection matrix", dst)
	}
	copyOut(nil, unsafe.Pointer(&dst[1]), 4, 0, nil, 0, 2)
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
