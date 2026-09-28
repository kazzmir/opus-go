package opuscc

import "testing"

func TestCeltFIR5Pointers(t *testing.T) {
	coefs := [5]float32{0.5, -0.25, 0.125, -0.0625, 0.03125}
	// C's zero-initialized history gives an impulse followed by all five taps.
	impulse := [9]float32{99, 1, 0, 0, 0, 0, 0, 0, 88}
	celt_fir5(nil, &impulse[1], &coefs[0], 7)
	want := [9]float32{99, 1, 0.5, -0.25, 0.125, -0.0625, 0.03125, 0, 88}
	if impulse != want {
		t.Fatalf("impulse: got %v, want %v", impulse, want)
	}

	input := [8]float32{1.25, -2.5, 0.75, 4, -0.125, 100.25, -70, 0.1}
	for n := 0; n <= len(input); n++ {
		got, want := input, input
		// Reference convolution over ORIGINAL input, in C's MAC order.
		for i := 0; i < n; i++ {
			for tap, coef := range coefs {
				if i > tap {
					want[i] += float32(coef * input[i-tap-1])
				}
			}
		}
		celt_fir5(nil, &got[0], &coefs[0], int32(n))
		if got != want {
			t.Fatalf("length %d: got %v, want %v", n, got, want)
		}
	}
	celt_fir5(nil, nil, &coefs[0], 0)
}

func TestCeltFIR5AliasedCoefficients(t *testing.T) {
	// The C implementation snapshots num[0:5] before modifying x.
	input := [7]float32{0.5, -0.25, 0.125, -0.0625, 0.03125, 1, -2}
	got, want := input, input
	for i := range want {
		for tap := 0; tap < 5 && tap < i; tap++ {
			want[i] += float32(input[tap] * input[i-tap-1])
		}
	}
	celt_fir5(nil, &got[0], &got[0], int32(len(got)))
	if got != want {
		t.Fatalf("aliased coefficients: got %v, want %v", got, want)
	}
}
