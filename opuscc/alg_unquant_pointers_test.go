package opuscc

import (
	"slices"
	"testing"
)

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
