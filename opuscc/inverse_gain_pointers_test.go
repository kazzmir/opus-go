package opuscc

import "testing"

func TestInverseGainPointers(t *testing.T) {
	for _, tc := range []struct {
		a    int16
		gain int32
	}{{0, 1 << 30}, {1024, 1006632960}, {-1024, 1006632960}, {4096, 0}, {-4096, 0}, {32767, 0}, {-32768, 0}} {
		a := [3]int16{123, tc.a, 456}
		before := a
		if got := Opus_silk_LPC_inverse_pred_gain_c(nil, &a[1], 1); got != tc.gain || a != before {
			t.Fatalf("a=%d gain=%d want=%d changed=%v", tc.a, got, tc.gain, a != before)
		}
	}
	for _, order := range []int32{2, 10, 16, 24} {
		a := make([]int32, order+2)
		a[0] = 123
		a[order+1] = 456
		a[order] = 1 << 22
		if got := LPC_inverse_pred_gain_QA_c(nil, &a[1], order); got != 1006632960 || a[0] != 123 || a[order+1] != 456 {
			t.Fatalf("order=%d gain=%d state=%v", order, got, a)
		}
		a[order] = 16773023
		if got := LPC_inverse_pred_gain_QA_c(nil, &a[1], order); got != 0 {
			t.Fatal("unstable coefficient accepted")
		}
	}
}
