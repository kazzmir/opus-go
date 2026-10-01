package opuscc

import (
	"runtime"
	"slices"
	"testing"
)

func TestPVQSearchPointers(t *testing.T) {
	x := [4]float32{77, -.25, .5, 88}
	p := [4]int32{77, 0, 0, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	yy := Opus_op_pvq_search_c(nil, &x[1], &p[1], 5, 2, 0)
	if p != [4]int32{77, -2, 3, 88} || x != [4]float32{77, .25, .5, 88} || yy != 13 {
		t.Fatal("pulse search", x, p, yy)
	}
	zero := [2]float32{}
	pulses := [2]int32{}
	if got := Opus_op_pvq_search_c(nil, &zero[0], &pulses[0], 16, 2, 0); got != 256 || pulses != [2]int32{16, 0} || zero != [2]float32{1, 0} {
		t.Fatal("silence fallback", got, pulses, zero)
	}
}

func TestAlgUnquantPointers(t *testing.T) {
	for _, n := range []int32{2, 8, 32, 128} {
		for _, blocks := range []int32{1, 2} {
			for spread := int32(0); spread <= 3; spread++ {
				dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
				ref := dec
				pulses := make([]int32, n)
				want := make([]float32, n)
				out := make([]float32, n+2)
				out[0] = 123
				out[n+1] = 456
				energy := Opus_decode_pulses(nil, &pulses[0], n, 1, &ref)
				normalise_residual(nil, &pulses[0], &want[0], n, energy, 0.75, 0)
				Opus_exp_rotation(nil, &want[0], n, -1, blocks, 1, spread)
				mask := Opus_alg_unquant(nil, &out[1], n, 1, spread, blocks, &dec, 0.75)
				if !slices.Equal(out[1:n+1], want) || mask != extract_collapse_mask(nil, &pulses[0], n, blocks) || dec != ref || out[0] != 123 || out[n+1] != 456 {
					t.Fatalf("n=%d B=%d spread=%d", n, blocks, spread)
				}
			}
		}
	}
}
