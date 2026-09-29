package opuscc

import "testing"

func TestDownInterpolPointers(t *testing.T) {
	for _, order := range []int32{18, 24, 36} {
		if silk_resampler_private_down_FIR_INTERPOL(nil, nil, nil, nil, order, 1, 0, 1) != 0 {
			t.Fatal("empty")
		}
		in := make([]int32, order)
		coefs := make([]int16, order/2)
		out := [3]int16{123, 0, 456}
		in[0] = 2147483647
		in[order-1] = 1
		coefs[0] = 1
		n := silk_resampler_private_down_FIR_INTERPOL(nil, &out[1], &in[0], &coefs[0], order, 1, 65536, 65536)
		want := int16(-512)
		if order == 18 {
			want = 512
		}
		if n != 1 || out != [3]int16{123, want, 456} {
			t.Fatalf("order %d: %v (%d)", order, out, n)
		}
	}
}
