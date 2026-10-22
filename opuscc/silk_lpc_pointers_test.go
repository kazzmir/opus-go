package opuscc

import (
	"slices"
	"testing"
)

// From lpc_analysis_pointers_test.go.
func TestLPCAnalysisPointers(t *testing.T) {
	for _, d := range []int{6, 10, 16, 24} {
		for _, extra := range []int{0, 1, 19} {
			n := d + extra
			in := make([]int16, n)
			b := make([]int16, d)
			for i := range in {
				in[i] = int16(32767)
				if i%2 == 0 {
					in[i] = -32768
				}
			}
			for i := range b {
				b[i] = -32768
			}
			before := slices.Clone(in)
			out := make([]int16, n+2)
			out[0] = 123
			out[n+1] = 456
			Opus_silk_LPC_analysis_filter(nil, &out[1], &in[0], &b[0], int32(n), int32(d), 0)
			if out[0] != 123 || out[n+1] != 456 || !slices.Equal(in, before) {
				t.Fatal("buffer boundary changed")
			}
			for i := 0; i < n; i++ {
				var want int16
				if i >= d {
					var sum int64
					for j := range b {
						sum += int64(in[i-j-1]) * int64(b[j])
					}
					r := int32((int64(in[i]) << 12) - sum)
					rounded := (int64(r) + 2048) >> 12
					want = int16(max(-32768, min(32767, rounded)))
				}
				if out[i+1] != want {
					t.Fatalf("d=%d n=%d i=%d got=%d want=%d", d, n, i, out[i+1], want)
				}
			}
		}
	}
}

// From lpc_fit_pointers_test.go.
func TestLPCFitPointers(t *testing.T) {
	input := []int32{111, -7, -5, -3, -1, 1, 3, 5, 7, 222}
	before := slices.Clone(input)
	output := [10]int16{123, 0, 0, 0, 0, 0, 0, 0, 0, 234}
	Opus_silk_LPC_fit(nil, &output[1], &input[1], 12, 13, 8)
	if !slices.Equal(output[:], []int16{123, -3, -2, -1, 0, 1, 2, 3, 4, 234}) || !slices.Equal(input, before) {
		t.Fatal("rounding or no-expansion input differs")
	}
	// A large late coefficient reaches C's ten-iteration clipping fallback.
	for _, sign := range []int32{-1, 1} {
		var high [26]int32
		high[0], high[25] = 111, 222
		high[24] = sign * 2000000000
		var out [26]int16
		out[0], out[25] = 123, 234
		Opus_silk_LPC_fit(nil, &out[1], &high[1], 12, 13, 24)
		expected := int16(32767)
		if sign < 0 {
			expected = -32768
		}
		if out[24] != expected || high[24] != int32(expected)<<1 || high[0] != 111 || high[25] != 222 || out[0] != 123 || out[25] != 234 {
			t.Fatalf("sign=%d out=%v input=%v", sign, out, high)
		}
	}
	Opus_silk_LPC_fit(nil, nil, nil, 12, 13, 0)
}

// From inverse_gain_pointers_test.go.
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

// From bwexpander_pointers_test.go.
func TestBWExpanderPointers(t *testing.T) {
	for _, n := range []int{1, 10, 16} {
		for _, chirp := range []int32{0, 32768, 64881, 65535, 65536} {
			a16, want16 := make([]int16, n+2), make([]int16, n+2)
			a32, want32 := make([]int32, n+2), make([]int32, n+2)
			for i := range a16 {
				a16[i] = int16(32767 - i*5000)
				a32[i] = int32(2147483647 - int64(i)*300000000)
			}
			copy(want16, a16)
			copy(want32, a32)
			c := chirp
			for i := 1; i <= n; i++ {
				product := int32(int64(c) * int64(want16[i]))
				want16[i] = int16((int64(product) + 32768) >> 16)
				want32[i] = int32((int64(c) * int64(want32[i])) >> 16)
				product = int32(int64(c) * int64(chirp-65536))
				c += int32((int64(product) + 32768) >> 16)
			}
			Opus_silk_bwexpander(nil, &a16[1], int32(n), chirp)
			Opus_silk_bwexpander_32(nil, &a32[1], int32(n), chirp)
			for i := range a16 {
				if a16[i] != want16[i] || a32[i] != want32[i] {
					t.Fatalf("n=%d chirp=%d index=%d: got (%d,%d), want (%d,%d)", n, chirp, i, a16[i], a32[i], want16[i], want32[i])
				}
			}
		}
	}
}
