package opuscc

import (
	"math"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestDeemphasisDriverPointers(t *testing.T) {
	channels := func() []*float32 {
		l, r := make([]float32, 8), make([]float32, 8)
		for i := range l {
			l[i] = float32(i*17 - 31)
			r[i] = float32(i*23 - 37)
		}
		return []*float32{&l[0], &r[0]}
	}()
	coef := float32(.85)
	memory := [2]float32{.25, -.5}
	out := make([]float32, 18)
	out[0] = 77
	out[17] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	deemphasis(nil, &channels[0], &out[1], 8, 2, 1, &coef, &memory[0], 0)
	if out[0] != 77 || out[17] != 88 {
		t.Fatal("driver guards")
	}
	var mem [2]float32
	deemphasis(nil, &channels[0], nil, 0, 2, 1, &coef, &mem[0], 0)
	for _, N := range []int32{0, -1} {
		// The stereo leaf consumes neither PCM nor memory for N<=0.
		deemphasis(nil, &channels[0], nil, N, 2, 1, &coef, nil, 0)
	}
	t.Run("zero-still-loads-coefficient", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatal("zero work must still load the coefficient")
			}
		}()
		deemphasis(nil, &channels[0], nil, 0, 2, 1, nil, nil, 0)
	})
}

func TestDeemphasisScratchPointers(t *testing.T) {
	for _, N := range []int32{0, 1, 8, 120} {
		for _, C := range []int32{0, 1, 2} {
			for _, down := range []int32{1, 2, 3, 6} {
				for _, accum := range []int32{0, 1} {
					left, right := make([]float32, N), make([]float32, N)
					for i := range left {
						left[i] = float32(i%17-8) * 1700
						right[i] = float32(i%13-6) * 2300
					}
					input := [2]*float32{unsafe.SliceData(left), unsafe.SliceData(right)}
					out := make([]float32, (N/down)*max(C, 1)+2)
					for i := range out {
						out[i] = float32(i+77) / 31
					}
					want := append([]float32(nil), out...)
					memory, wm := [2]float32{.25, -.5}, [2]float32{.25, -.5}
					coef := float32(.85)
					channels := [][]float32{left, right}
					for c := int32(0); c < max(C, 1); c++ {
						m := wm[c]
						scratch := make([]float32, N)
						for j := int32(0); j < N; j++ {
							var tmp float32
							if down > 1 || accum == 0 {
								tmp = float32(float32(channels[c][j]+float32(1e-30)) + m)
							} else {
								tmp = float32(float32(channels[c][j]+m) + float32(1e-30))
							}
							m = float32(coef * tmp)
							if down > 1 {
								scratch[j] = tmp
							} else {
								scaled := float32(tmp * float32(1.0/32768))
								if accum != 0 {
									want[1+j*C+c] += scaled
								} else {
									want[1+j*C+c] = scaled
								}
							}
						}
						wm[c] = m
						if down > 1 {
							for j := int32(0); j < N/down; j++ {
								scaled := float32(scratch[j*down] * float32(1.0/32768))
								if accum != 0 {
									want[1+j*C+c] += scaled
								} else {
									want[1+j*C+c] = scaled
								}
							}
						}
					}
					entropyInitGrowStack(12)
					runtime.GC()
					deemphasis(nil, &input[0], &out[1], N, C, down, &coef, &memory[0], accum)
					for i := range out {
						if math.Float32bits(out[i]) != math.Float32bits(want[i]) {
							t.Fatal("scratch output", N, C, down, accum, i, out[i], want[i])
						}
					}
					if memory != wm {
						t.Fatal("scratch memory", N, C, down, accum, memory, wm)
					}
				}
			}
		}
	}
	// N=0 ignores spectral/PCM buffers while preserving generic channel histories.
	var input [2]*float32
	coef := float32(.85)
	memory := [2]float32{.25, -.5}
	deemphasis(nil, &input[0], nil, 0, 2, 6, &coef, &memory[0], 1)
	if memory != [2]float32{.25, -.5} {
		t.Fatal("zero length")
	}
	// Memory can alias output; downsampled PCM stores follow the memory write.
	left := []float32{32768, 0}
	input[0] = &left[0]
	out := []float32{.5, 77}
	coef = .5
	deemphasis(nil, &input[0], &out[0], 2, 1, 2, &coef, &out[0], 0)
	if out[0] != float32(float32(32768+.5)*float32(1.0/32768)) || out[1] != 77 {
		t.Fatal("memory/output store order", out)
	}
}

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
