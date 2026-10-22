package opuscc

import (
	"math"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// From comb_mask_pointers_test.go.
func TestCombFilterTransitionPointers(t *testing.T) {
	Opus_comb_filter(nil, nil, nil, 0, 0, 0, 0, 0, -1, -1, nil, 0, 0)
	src := make([]float32, 64)
	for i := range src {
		src[i] = float32(i%17-8) * 0.125
	}
	before := slices.Clone(src)
	dst := [18]float32{}
	dst[0] = 77
	dst[17] = 88
	win := [8]float32{0, .1, .2, .3, .4, .5, .6, .7}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_comb_filter(nil, &dst[1], &src[32], 15, 20, 16, .25, .5, 0, 1, &win[0], 8, 0)
	if !slices.Equal(src, before) || dst[0] != 77 || dst[17] != 88 {
		t.Fatal("input/guards")
	}
	copyBuf := [5]float32{1, 2, 3, 4, 5}
	Opus_comb_filter(nil, &copyBuf[1], &copyBuf[0], 0, 0, 4, 0, 0, -1, -1, nil, 0, 0)
	if copyBuf != [5]float32{1, 1, 2, 3, 4} {
		t.Fatal("memmove")
	}
}

func TestCombFilterPointers(t *testing.T) {
	for _, period := range []int{15, 32} {
		for _, n := range []int{1, 17, 64} {
			for _, inPlace := range []bool{false, true} {
				offset := period + 2
				x := make([]float32, offset+n)
				for i := range x {
					x[i] = float32((i*7)%31-15) * 0.125
				}
				wantX := append([]float32(nil), x...)
				out := make([]float32, n+2)
				out[0], out[n+1] = 99, 88
				got := out[1 : n+1]
				if inPlace {
					got = x[offset:]
				}
				spanX, spanOut := slices.Clone(x), slices.Clone(out)
				span := spanOut[1 : n+1]
				if inPlace {
					span = spanX[offset:]
				}
				want := make([]float32, n)
				for i := range want {
					p := offset + i
					want[i] = wantX[p] + float32(0.5*wantX[p-period]) + float32(0.25*(wantX[p-period+1]+wantX[p-period-1])) + float32(-0.125*(wantX[p-period+2]+wantX[p-period-2]))
					if inPlace {
						wantX[p] = want[i]
					}
				}
				comb_filter_const_c(nil, &got[0], &x[0], int32(period), int32(n), 0.5, 0.25, -0.125)
				combFilterConstSpans(spanX, span, int32(period), .5, .25, -.125)
				if !slices.Equal(spanX, x) || !slices.Equal(spanOut, out) {
					t.Fatal("constant comb spans/whole buffers/guards")
				}
				for i := range want {
					if got[i] != want[i] {
						t.Fatalf("period=%d n=%d inplace=%t i=%d: got %g want %g", period, n, inPlace, i, got[i], want[i])
					}
				}
				if out[0] != 99 || out[n+1] != 88 {
					t.Fatal("output sentinel changed")
				}
				for i := 0; i < offset; i++ {
					if x[i] != wantX[i] {
						t.Fatal("history changed")
					}
				}
			}
		}
	}
	comb_filter_const_c(nil, nil, nil, 15, 0, 0.5, 0.25, 0.125)
	combFilterConstSpans(nil, nil, 15, .5, .25, .125)
}

func TestCollapseMaskPointers(t *testing.T) {
	for _, blocks := range []int{2, 4, 8, 16, 32} {
		for _, width := range []int{1, 3} {
			x := make([]int32, blocks*width)
			if got := extract_collapse_mask(nil, &x[0], int32(len(x)), int32(blocks)); got != 0 || extractCollapseMaskOwned(x, int32(len(x)), int32(blocks)) != got {
				t.Fatal("zero mask")
			}
			for b := 0; b < blocks; b++ {
				x[b*width+width-1] = -2147483648
				before := slices.Clone(x)
				entropyInitGrowStack(12)
				runtime.GC()
				if got := extract_collapse_mask(nil, &x[0], int32(len(x)), int32(blocks)); got != uint32(1)<<b || extractCollapseMaskOwned(x, int32(len(x)), int32(blocks)) != got {
					t.Fatalf("blocks=%d width=%d active=%d got=%x", blocks, width, b, got)
				}
				if !slices.Equal(x, before) {
					t.Fatal("collapse pulse source changed")
				}
				x[b*width+width-1] = 0
			}
		}
	}
	if got := extract_collapse_mask(nil, nil, 0, 1); got != 1 {
		t.Fatal("single block special case")
	}
	for _, N := range []int32{-2147483648, -1, 0, 1, 2147483647} {
		for _, B := range []int32{-2147483648, -1, 0, 1} {
			if extractCollapseMaskOwned(nil, N, B) != 1 || extract_collapse_mask(nil, nil, N, B) != 1 {
				t.Fatal("collapse unused input before guard", N, B)
			}
		}
	}
	// Integer division ignores a trailing remainder just as C does.
	x := [5]int32{1, -1, 0, 0, 99}
	if got := extract_collapse_mask(nil, &x[0], 5, 2); got != 1 || extractCollapseMaskOwned(x[:], 5, 2) != got {
		t.Fatalf("remainder/cancellation: %x", got)
	}
}

// From comb_window_owner_test.go.
func TestCombCanonicalWindowOwnerPointers(t *testing.T) {
	for _, n := range []int32{0, 1, 120} {
		window := combFilterKnownWindow(&window120[0], n)
		if !window.owned || len(window.owner) != 120 || &window.owner[0] != &window120[0] || len(window.span(n)) != int(n) {
			t.Fatal("canonical comb window actual owner", n)
		}
	}
	copied := window120
	for _, test := range []struct {
		pointer *float32
		overlap int32
	}{{&copied[0], 120}, {&window120[1], 119}, {&window120[0], 121}, {&window120[0], -1}, {nil, 0}} {
		window := combFilterKnownWindow(test.pointer, test.overlap)
		if window.owned || window.legacy != test.pointer {
			t.Fatal("noncanonical/unsupported window must retain raw adapter", test.overlap)
		}
	}
	if (combFilterWindow{legacy: &window120[0], owned: true}).span(0) != nil {
		t.Fatal("nil owner must not fall back")
	}
}

func TestCombCanonicalWindowLiveAliasPointers(t *testing.T) {
	saved := window120
	defer func() { window120 = saved }()
	for _, overlap := range []int32{1, 8, 16} {
		window120 = saved
		reference := saved
		entropyInitGrowStack(12)
		runtime.GC()
		// Copied table deliberately uses the raw adapter, same physical 120 words.
		Opus_comb_filter(nil, &reference[1], &reference[64], 15, 17, 16, .3, .7, 0, 1, &reference[0], overlap, 0)
		Opus_comb_filter(nil, &window120[1], &window120[64], 15, 17, 16, .3, .7, 0, 1, &window120[0], overlap, 0)
		for i := range reference {
			if math.Float32bits(reference[i]) != math.Float32bits(window120[i]) {
				t.Fatal("comb canonical window must stay live under sequential stores", overlap, i)
			}
		}
		window120 = saved
	}
}

// From postfilter_owned_test.go.
func TestCeltDecodePostfilterTailRequiredCellPointers(t *testing.T) {
	mode := newSynthesisTestMode()
	state := OpusT_OpusCustomDecoder{Fmode: mode}
	before := state
	samples := make([]float32, 120)
	var caught any
	func() {
		defer func() { caught = recover() }()
		celtDecodePostfilterTailWithHistory(nil, &state, mode, &samples[0], 120, 80, 0, 0, 120, samples, 0)
	}()
	if _, ok := caught.(runtime.Error); !ok || state != before {
		t.Fatal("tail argument requires the interior first cell even for zero gains", caught)
	}
	for _, sample := range samples {
		if sample != 0 {
			t.Fatal("tail argument check before stores")
		}
	}
}

// From deemphasis_pointers_test.go.
func TestDeemphasisOutputOriginPointers(t *testing.T) {
	if deemphasisOutputAt(nil, 0, 0) != nil {
		t.Fatal("zero output identity")
	}
	owner := [6]float32{77, .25, .5, 1, 2, 88}
	root := &owner[1]
	entropyInitGrowStack(12)
	runtime.GC()
	for _, coordinates := range [][2]int32{{0, 0}, {1, 2}, {2, 1}, {3, 0}, {2, -2}, {0, -1}} {
		channel, sample := coordinates[0], coordinates[1]
		old := celtNormConsumedAt(celtNormConsumedAt(root, channel), sample)
		if deemphasisOutputAt(root, channel, sample) != old {
			t.Fatal("output origin identity", coordinates)
		}
	}
	*deemphasisOutputAt(root, 1, 2) = 4
	if owner != [6]float32{77, .25, .5, 1, 4, 88} {
		t.Fatal("live output/guards", owner)
	}
}

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
					ownedOut := slices.Clone(out)
					ownedMemory := [2]float32{.25, -.5}
					historyOut := slices.Clone(out)
					historyMemory := ownedMemory
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
					deemphasisOwned(nil, input[:], &ownedOut[1], N, C, down, &coef, ownedMemory[:], accum)
					deemphasisWithHistory(nil, input[:], [2][]float32{left, right}, 0, &historyOut[1], N, C, down, &coef, historyMemory[:], accum)
					for i := range out {
						if math.Float32bits(ownedOut[i]) != math.Float32bits(out[i]) || math.Float32bits(historyOut[i]) != math.Float32bits(out[i]) {
							t.Fatal("owned input-list/whole PCM/guards", N, C, down, accum, i)
						}
						if math.Float32bits(out[i]) != math.Float32bits(want[i]) {
							t.Fatal("scratch output", N, C, down, accum, i, out[i], want[i])
						}
					}
					if memory != wm || ownedMemory != wm || historyMemory != wm {
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
	deemphasisOwned(nil, input[:], nil, 0, 2, 6, &coef, memory[:], 1)
	deemphasisOwned(nil, input[:], nil, 0, 2, 1, &coef, nil, 0)
	if memory != [2]float32{.25, -.5} {
		t.Fatal("zero length")
	}
	// Memory can alias output; downsampled PCM stores follow the memory write.
	left := []float32{32768, 0}
	input[0] = &left[0]
	out := []float32{.5, 77}
	ownedAlias := slices.Clone(out)
	coef = .5
	deemphasis(nil, &input[0], &out[0], 2, 1, 2, &coef, &out[0], 0)
	deemphasisOwned(nil, input[:], &ownedAlias[0], 2, 1, 2, &coef, ownedAlias[:1], 0)
	if !slices.Equal(out, ownedAlias) {
		t.Fatal("owned memory/PCM store ordering")
	}
	if out[0] != float32(float32(32768+.5)*float32(1.0/32768)) || out[1] != 77 {
		t.Fatal("memory/output store order", out)
	}
}

func TestDeemphasisHistoryPointers(t *testing.T) {
	left, right := []float32{77, 88, 32768, 0, -32768, 99}, []float32{11, 22, -32768, 16384, 0, 33}
	list := []*float32{&left[2], &right[2]}
	input := deemphasisInputs{pointers: celtSynthesisOutputs{owner: list, owned: true}, history: [2][]float32{left, right}, origin: 2, hasHistory: true}
	bound := input.channel(list[0], 0, 3)
	if !bound.owned || len(bound.samples(3)) != 3 || &bound.samples(3)[0] != &left[2] {
		t.Fatal("actual history/root/origin identity")
	}
	for _, tc := range []struct {
		pointer        *float32
		channel, count int32
	}{{list[1], 0, 3}, {list[0], 0, 0}, {list[0], 0, 99}, {list[0], 2, 3}} {
		fallback := input.channel(tc.pointer, tc.channel, tc.count)
		if fallback.owned || fallback.legacy != tc.pointer {
			t.Fatal("history mismatch/unused/unbound fallback")
		}
	}
	negative := input
	negative.origin = -1
	if negative.channel(list[0], 0, 3).owned {
		t.Fatal("signed origin stays legacy")
	}
	// Root metadata must not snapshot live pointer-list entries.
	list[0] = &right[2]
	if input.channel(input.span(2)[0], 0, 3).owned {
		t.Fatal("live input pointer reload")
	}
	list[0] = &left[2]
	// Go-only input/PCM overlap: owned samples stay live across stores.
	for _, accum := range []int32{0, 1} {
		a, b := slices.Clone(left), slices.Clone(left)
		ma, mb := []float32{.25}, []float32{.25}
		coef := float32(.5)
		deemphasis(nil, &[]*float32{&a[2]}[0], &a[3], 2, 1, 1, &coef, &ma[0], accum)
		deemphasisWithHistory(nil, []*float32{&b[2]}, [2][]float32{b}, 2, &b[3], 2, 1, 1, &coef, mb, accum)
		if !slices.Equal(a, b) || !slices.Equal(ma, mb) {
			t.Fatal("live owned input/PCM aliases", accum)
		}
	}
}

func TestDeemphasisPointers(t *testing.T) {
	left := [4]float32{32768, 0, 16384, -32768}
	right := [4]float32{-32768, 0, 0, 32768}
	originalL, originalR := left, right
	want := []float32{1, -1, 0.5, -0.5, 0.75, -0.25, -0.625, 0.875}
	output := [10]float32{111, 0, 0, 0, 0, 0, 0, 0, 0, 222}
	var state [2]float32
	spanOutput, spanState := output, state
	deemphasis_stereo_simple(nil, &left[0], &right[0], &output[1], 4, 0.5, &state)
	deemphasisStereoSpans(left[:], right[:], spanOutput[1:9], .5, &spanState)
	if spanOutput != output || spanState != state {
		t.Fatal("stereo spans/whole buffer/guards")
	}
	if !slices.Equal(output[1:9], want) || state != [2]float32{-10240, 14336} {
		t.Fatalf("PCM=%v state=%v", output, state)
	}
	if output[0] != 111 || output[9] != 222 || left != originalL || right != originalR {
		t.Fatal("input or sentinels changed")
	}
	before := state
	deemphasisStereoSpans(nil, nil, nil, .5, nil)
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

// From celt_fir5_pointers_test.go.
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
		got, want, numeric := input, input, input
		// Reference convolution over ORIGINAL input, in C's MAC order.
		for i := 0; i < n; i++ {
			for tap, coef := range coefs {
				if i > tap {
					want[i] += float32(coef * input[i-tap-1])
				}
			}
		}
		entropyInitGrowStack(12)
		runtime.GC()
		celt_fir5(nil, &got[0], &coefs[0], int32(n))
		celtFIR5Span(nil, numeric[:], coefs[:], int32(n))
		if got != want || numeric != want {
			t.Fatalf("length %d: got %v, want %v", n, got, want)
		}
	}
	celt_fir5(nil, nil, &coefs[0], 0)
	celtFIR5Span(nil, nil, coefs[:], 0)
	celtFIR5Span(nil, nil, coefs[:], -1)
	if !validationPanics(func() { celtFIR5Span(nil, nil, coefs[:4:4], 0) }) {
		t.Fatal("five coefficients required before zero-length guard")
	}
}

func TestCeltFIR5AliasedCoefficients(t *testing.T) {
	// The C implementation snapshots num[0:5] before modifying x.
	input := [7]float32{0.5, -0.25, 0.125, -0.0625, 0.03125, 1, -2}
	got, want, numeric := input, input, input
	for i := range want {
		for tap := 0; tap < 5 && tap < i; tap++ {
			want[i] += float32(input[tap] * input[i-tap-1])
		}
	}
	celt_fir5(nil, &got[0], &got[0], int32(len(got)))
	celtFIR5Span(nil, numeric[:], numeric[:], int32(len(numeric)))
	if got != want || numeric != want {
		t.Fatalf("aliased coefficients: got %v, want %v", got, want)
	}
}
