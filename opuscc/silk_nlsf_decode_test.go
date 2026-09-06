package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestNLSFDecodeLocalArrays(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	indices := make([]OpusT_opus_int8, 17)
	var decoded [16]OpusT_opus_int16
	Opus_silk_NLSF_decode(tls, uintptr(unsafe.Pointer(&decoded[0])), uintptr(unsafe.Pointer(&indices[0])), uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_WB)))

	want := [16]OpusT_opus_int16{896, 2944, 4864, 6912, 8832, 10880, 12800, 14848, 16768, 18816, 20736, 22784, 24704, 26624, 28544, 30592}
	if decoded != want {
		t.Fatalf("decoded NLSFs: got %v, want %v", decoded, want)
	}
}

func TestNLSFDecodeNegativeResidualIndices(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	/* expected values from the C reference implementation (silk/NLSF_decode.c);
	   the residual dequantizer must sign-extend negative opus_int8 indices */
	indices1 := []OpusT_opus_int8{5, 0, -2, 1, 0, 0, -1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	want1 := []int16{633, 1144, 7947, 10055, 13266, 16136, 21156, 23040, 26240, 29184}

	indices2 := []OpusT_opus_int8{3, -3, -2, -1, -3, -2, -1, -2, -1, -1, 0, 0, 0, 0, 0, 0, 0}
	want2 := []int16{250, 417, 3103, 3106, 8055, 12125, 15239, 20354, 24398, 28544}

	indices3 := []OpusT_opus_int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	want3 := []int16{1536, 4480, 7680, 10624, 13824, 16896, 20096, 23040, 26368, 29184}

	for _, tc := range []struct {
		indices []OpusT_opus_int8
		want    []int16
	}{
		{indices1, want1},
		{indices2, want2},
		{indices3, want3},
	} {
		var decoded [16]OpusT_opus_int16
		Opus_silk_NLSF_decode(tls, uintptr(unsafe.Pointer(&decoded[0])), uintptr(unsafe.Pointer(&tc.indices[0])), uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_NB_MB)))
		if got := decoded[:10]; !equalInt16s(got, tc.want) {
			t.Fatalf("decoded NLSFs for indices %v: got %v, want %v", tc.indices, got, tc.want)
		}
	}
}
