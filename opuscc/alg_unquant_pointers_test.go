package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

func TestAlgQuantPointers(t *testing.T) {
	for spread := int32(0); spread < 4; spread++ {
		buffer := [32]byte{}
		var enc OpusT_ec_enc
		Opus_ec_enc_init(nil, &enc, &buffer[0], 32)
		x := [18]float32{}
		x[0] = 77
		x[17] = 88
		for i := 1; i < 17; i++ {
			x[i] = float32(i-8) * .1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		mask := Opus_alg_quant(nil, &x[1], 16, 4, spread, 2, &enc, .75, 1, 0)
		Opus_ec_enc_done(nil, &enc)
		if x[0] != 77 || x[17] != 88 {
			t.Fatal("guards")
		}
		var dec OpusT_ec_dec
		Opus_ec_dec_init(nil, &dec, &buffer[0], 32)
		var out [16]float32
		dm := Opus_alg_unquant(nil, &out[0], 16, 4, spread, 2, &dec, .75)
		if mask != dm {
			t.Fatal("collapse mask", mask, dm)
		}
		for i := range out {
			if math.Float32bits(out[i]) != math.Float32bits(x[i+1]) {
				t.Fatal("round trip", spread, i, out[i], x[i+1])
			}
		}
	}
}

func TestPVQSearchPointers(t *testing.T) {
	x := [4]float32{77, -.25, .5, 88}
	p := [4]int32{77, 0, 0, 88}
	ownedX, ownedP := x, p
	entropyInitGrowStack(12)
	runtime.GC()
	yy := Opus_op_pvq_search_c(nil, &x[1], &p[1], 5, 2, 0)
	ownedYY := pvqSearchPulsesOwned(&ownedX[1], ownedP[1:3], 5, 2)
	if ownedX != x || ownedP != p || ownedYY != yy {
		t.Fatal("owned PVQ producer full buffers/energy")
	}
	if p != [4]int32{77, -2, 3, 88} || x != [4]float32{77, .25, .5, 88} || yy != 13 {
		t.Fatal("pulse search", x, p, yy)
	}
	zero := [2]float32{}
	pulses := [2]int32{}
	ownedZero, ownedSilence := zero, pulses
	if got := pvqSearchPulsesOwned(&ownedZero[0], ownedSilence[:], 16, 2); got != 256 || ownedSilence != [2]int32{16, 0} || ownedZero != [2]float32{1, 0} {
		t.Fatal("owned silence fallback", got, ownedSilence, ownedZero)
	}
	if got := Opus_op_pvq_search_c(nil, &zero[0], &pulses[0], 16, 2, 0); got != 256 || pulses != [2]int32{16, 0} || zero != [2]float32{1, 0} {
		t.Fatal("silence fallback", got, pulses, zero)
	}
}

func TestPVQSearchOwnedVariantsPointers(t *testing.T) {
	for _, N := range []int32{2, 4, 8, 16, 32, 176} {
		for _, K := range []int32{1, N / 2, N/2 + 1, 2 * N, 127} {
			x, p := make([]float32, N+2), make([]int32, N+2)
			x[0], x[N+1], p[0], p[N+1] = 77, 88, 77, 88
			for j := int32(0); j < N; j++ {
				x[j+1] = float32(j%7-3) / 17
			}
			ox, op := slices.Clone(x), slices.Clone(p)
			entropyInitGrowStack(12)
			runtime.GC()
			yy := Opus_op_pvq_search_c(nil, &x[1], &p[1], K, N, 0)
			oy := pvqSearchPulsesOwned(&ox[1], op[1:N+1], K, N)
			if !slices.Equal(p, op) || math.Float32bits(yy) != math.Float32bits(oy) {
				t.Fatal("owned PVQ pulses/energy", N, K)
			}
			for j := range x {
				if math.Float32bits(x[j]) != math.Float32bits(ox[j]) {
					t.Fatal("owned PVQ full spectrum bits", N, K, j)
				}
			}
			if x[0] != 77 || x[N+1] != 88 || p[0] != 77 || p[N+1] != 88 {
				t.Fatal("owned PVQ guards")
			}
		}
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
