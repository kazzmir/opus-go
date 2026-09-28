package opuscc

import (
	"math"
	"math/rand"
	"slices"
	"testing"
)

func TestDeemphasisPointers(t *testing.T) {
	left := [4]float32{32768, 0, 16384, -32768}
	right := [4]float32{-32768, 0, 0, 32768}
	originalL, originalR := left, right
	want := []float32{1, -1, 0.5, -0.5, 0.75, -0.25, -0.625, 0.875}
	output := [10]float32{111, 0, 0, 0, 0, 0, 0, 0, 0, 222}
	var state [2]float32
	deemphasis_stereo_simple(nil, &left[0], &right[0], &output[1], 4, 0.5, &state)
	if !slices.Equal(output[1:9], want) || state != [2]float32{-10240, 14336} {
		t.Fatalf("PCM=%v state=%v", output, state)
	}
	if output[0] != 111 || output[9] != 222 || left != originalL || right != originalR {
		t.Fatal("input or sentinels changed")
	}
	before := state
	deemphasis_stereo_simple(nil, nil, nil, nil, 0, 0.5, &state)
	if state != before {
		t.Fatal("empty frame changed state")
	}
	var zero float32
	var silence [2]float32
	state = [2]float32{}
	deemphasis_stereo_simple(nil, &zero, &zero, &silence[0], 1, 0.5, &state)
	tiny := float32(1e-30)
	if silence != [2]float32{tiny / 32768, tiny / 32768} || state != [2]float32{tiny * 0.5, tiny * 0.5} {
		t.Fatal("VERY_SMALL injection changed")
	}
}

func TestDeemphasisChunkedPointers(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{2, 17, 120, 960} {
		for _, coef := range []float32{0, 0.5, 0.85, 1} {
			left, right := make([]float32, n), make([]float32, n)
			for i := range left {
				left[i] = float32(rng.NormFloat64() * 32768)
				right[i] = float32(rng.NormFloat64() * 32768)
			}
			whole, chunked := make([]float32, 2*n), make([]float32, 2*n)
			state := [2]float32{1234, -5678}
			chunks := state
			deemphasis_stereo_simple(nil, &left[0], &right[0], &whole[0], int32(n), coef, &state)
			split := n / 2
			deemphasis_stereo_simple(nil, &left[0], &right[0], &chunked[0], int32(split), coef, &chunks)
			deemphasis_stereo_simple(nil, &left[split], &right[split], &chunked[2*split], int32(n-split), coef, &chunks)
			if state != chunks {
				t.Fatal("chunked state differs")
			}
			for i := range whole {
				if math.Float32bits(whole[i]) != math.Float32bits(chunked[i]) {
					t.Fatalf("n=%d coef=%g i=%d: chunked PCM differs", n, coef, i)
				}
			}
		}
	}
}
