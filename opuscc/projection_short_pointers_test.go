package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"testing"
	"unsafe"
)

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
		got := opusProjectionDecodeShort(nil, &owner.Projection, data, length, &out[1], 5760, fec)
		copyOut := func(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
			opus_projection_copy_channel_out_short(tls, (*int16)(dst), ds, dc, src, ss, n, &baseline.Matrix)
		}
		expected := opusMSDecodeNative(nil, &baseline.MS.MS, data, length, unsafe.Pointer(&want[1]), copyOut, 5760, fec, OPTIONAL_CLIP)
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
