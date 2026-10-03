//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestCeltPLCExtrapolateAgainstC(t *testing.T) {
	for _, pitch := range []int32{40, 100, 511, 1024} {
		for _, N := range []int32{120, 240, 960} {
			for _, decay := range []float32{0, .923, 1} {
				a := make([]float32, 2170)
				x := make([]float32, 1024)
				a[0], a[len(a)-1] = 77, 88
				for i := 1; i < len(a)-1; i++ {
					a[i] = float32((i*37)%79-39) / 13
				}
				for i := range x {
					x[i] = float32((i*43)%89-44) / 17
				}
				b := slices.Clone(a)
				g := opuscc.CompareCeltPLCExtrapolate(&a[1], &x[0], 2048, 1024, N, 120, pitch, .8, decay)
				c := nativeCeltPLCExtrapolate(&b[1], &x[0], 2048, 1024, N, 120, pitch, .8, decay)
				if math.Float32bits(g) != math.Float32bits(c) {
					t.Fatal("extrapolation energy", pitch, N, decay, g, c)
				}
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("extrapolation", pitch, N, decay, i)
					}
				}
			}
		}
	}
}
func TestCeltPLCExcitationHistoryAgainstC(t *testing.T) {
	for _, period := range []int32{0, 1, 512, 1024} {
		h := make([]float32, 2050)
		for i := range h {
			h[i] = math.Float32frombits(uint32(i) * uint32(7919))
		}
		before := slices.Clone(h)
		a := make([]float32, period+26)
		a[0], a[len(a)-1] = 77, 88
		b := slices.Clone(a)
		opuscc.CompareCeltPLCExcitationHistory(&a[1], &h[1], 2048, period)
		nativeCeltPLCExcitationHistory(&b[1], &h[1], 2048, period)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("excitation history", period, i)
			}
		}
		if !slices.Equal(h, before) {
			t.Fatal("history input changed")
		}
	}
}
func TestCeltPLCSynthesisAttenuateAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2453))
	for _, length := range []int32{1, 120, 240, 1080} {
		for _, overlap := range []int32{0, 1, min(120, length)} {
			for _, factor := range []float32{0, .1, .2, .21, .5, 1, 2, float32(math.NaN())} {
				a := make([]float32, length+2)
				a[0], a[len(a)-1] = 77, 88
				w := make([]float32, overlap)
				for i := int32(0); i < length; i++ {
					a[i+1] = float32(rng.NormFloat64() * 7)
				}
				for i := range w {
					w[i] = float32(i+1) / float32(len(w)+1)
				}
				b := slices.Clone(a)
				s2 := float32(0)
				for _, v := range a[1 : len(a)-1] {
					s2 += float32(v * v)
				}
				s1 := float32(factor * s2)
				opuscc.CompareCeltPLCSynthesisAttenuate(&a[1], unsafe.SliceData(w), length, overlap, s1)
				nativeCeltPLCSynthesisAttenuate(&b[1], unsafe.SliceData(w), length, overlap, s1)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("attenuation", length, overlap, factor, i, a[i], b[i])
					}
				}
			}
		}
	}
}
func TestCeltPLCExcitationDecayAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2452))
	for _, length := range []int32{0, 1, 2, 31, 120, 512, 1024} {
		for trial := 0; trial < 80; trial++ {
			a := make([]float32, 1024)
			for i := range a {
				a[i] = float32(rng.NormFloat64() * 1e4)
			}
			g := opuscc.CompareCeltPLCExcitationDecay(&a[0], 1024, length)
			c := nativeCeltPLCExcitationDecay(&a[0], 1024, length)
			if math.Float32bits(g) != math.Float32bits(c) {
				t.Fatal("excitation decay", length, trial, g, c)
			}
		}
	}
}
func TestCeltPLCLagWindowAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2451))
	for trial := 0; trial < 1000; trial++ {
		var a [25]float32
		for i := range a {
			a[i] = float32(rng.NormFloat64() * 1e10)
		}
		b := a
		opuscc.CompareCeltPLCLagWindow(&a)
		nativeCeltPLCLagWindow(&b)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("lag window", trial, i, a[i], b[i])
			}
		}
	}
}
func TestCeltPLCDecayAgainstC(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, loss := range []int32{0, 1, 99} {
			for _, alias := range []int{0, 1, 2} {
				a := make([]float32, 66)
				b := make([]float32, 66)
				for i := range a {
					a[i] = float32(i) - 22
					b[i] = float32(i%5) - 12
				}
				a[0], a[65] = 77, 88
				want := slices.Clone(a)
				wb := slices.Clone(b)
				gp, cp := &b[1], &wb[1]
				if alias == 1 {
					gp, cp = &a[1], &want[1]
				}
				if alias == 2 {
					gp, cp = &a[2], &want[2]
				}
				opuscc.CompareCeltPLCDecay(&a[1], gp, 21, 1, 20, channels, loss)
				nativeCeltPLCDecay(&want[1], cp, 21, 1, 20, channels, loss)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
						t.Fatal("PLC decay", channels, loss, alias, i)
					}
				}
			}
		}
	}
}
func TestPLCPitchSearchAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1316))
	for _, channels := range []int32{0, 1, 2, 3} {
		for trial := 0; trial < 80; trial++ {
			left := make([]float32, opuscc.DEC_PITCH_BUF_SIZE)
			right := make([]float32, opuscc.DEC_PITCH_BUF_SIZE)
			for i := range left {
				left[i] = float32(rng.NormFloat64())
				right[i] = float32(rng.NormFloat64())
			}
			if trial == 0 {
				clear(left)
				clear(right)
			}
			if trial%3 == 1 {
				for i := range left {
					left[i] = float32(math.Sin(float64(i) * 0.17))
					right[i] = float32(math.Sin(float64(i) * 0.13))
				}
			}
			beforeL := slices.Clone(left)
			beforeR := slices.Clone(right)
			var rp *float32
			if channels == 2 {
				rp = &right[0]
			}
			g := opuscc.ComparePLCPitchSearch(&left[0], rp, channels)
			c := nativePLCPitchSearch(left, right, channels)
			if g != c || !sameFloatBits(left, beforeL) || !sameFloatBits(right, beforeR) {
				t.Fatal(channels, trial, g, c)
			}
		}
	}
}

func TestRemoveDoublingAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1215))
	for _, maxPeriod := range []int{16, 31, 64, 128} {
		for _, minPeriod := range []int{4, 5, 8} {
			for _, n := range []int{8, 17, 64, 240} {
				for trial := 0; trial < 100; trial++ {
					x := make([]float32, maxPeriod/2+n/2)
					for i := range x {
						x[i] = float32(rng.NormFloat64())
					}
					if trial == 0 {
						clear(x)
					}
					if trial%3 == 1 {
						for i := range x {
							x[i] = float32(math.Sin(float64(i) * 0.3))
						}
					}
					before := append([]float32(nil), x...)
					initial := minPeriod + rng.Intn(maxPeriod-minPeriod+8)
					previous := rng.Intn(maxPeriod)
					previousGain := float32(rng.Float64())
					g := [3]int32{77, int32(initial), 88}
					c := g
					gg := opuscc.Opus_remove_doubling(nil, &x[0], int32(maxPeriod), int32(minPeriod), int32(n), &g[1], int32(previous), previousGain, 0)
					cg := nativeRemoveDoubling(x, int32(maxPeriod), int32(minPeriod), int32(n), &c[1], int32(previous), previousGain)
					if g != c || math.Float32bits(gg) != math.Float32bits(cg) || !sameFloatBits(x, before) {
						t.Fatal(maxPeriod, minPeriod, n, trial, initial, previous, g, c, gg, cg)
					}
				}
			}
		}
	}
}

func TestPitchSearchAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1114))
	for _, n := range []int{4, 8, 12, 17, 32, 128, 240} {
		for _, maxPitch := range []int{4, 8, 12, 16, 31, 128} {
			if maxPitch>>2 > 3 && n>>2 < 3 {
				continue
			}
			for trial := 0; trial < 60; trial++ {
				x := make([]float32, n/2)
				y := make([]float32, (n+maxPitch)/2)
				for i := range x {
					x[i] = float32(rng.NormFloat64())
				}
				for i := range y {
					y[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					clear(x)
					clear(y)
				}
				if trial == 1 {
					for i := range x {
						x[i] = float32(math.Sin(float64(i) * 0.7))
					}
					for i := range y {
						y[i] = float32(math.Sin(float64(i) * 0.7))
					}
				}
				x0 := append([]float32(nil), x...)
				y0 := append([]float32(nil), y...)
				g := [3]int32{77, -1, 88}
				c := g
				opuscc.Opus_pitch_search(nil, &x[0], &y[0], int32(n), int32(maxPitch), &g[1], 0)
				nativePitchSearch(x, y, int32(n), int32(maxPitch), &c[1])
				if g != c || !sameFloatBits(x, x0) || !sameFloatBits(y, y0) {
					t.Fatal(n, maxPitch, trial, g, c)
				}
			}
		}
	}
}

func TestPitchDownsampleAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1013))
	for _, n := range []int{7, 8, 17, 64, 240} {
		for _, factor := range []int{1, 2, 3, 4} {
			for _, channels := range []int{0, 1, 2, 3} {
				for trial := 0; trial < 12; trial++ {
					left := make([]float32, n*factor)
					right := make([]float32, n*factor)
					for i := range left {
						left[i] = float32(rng.NormFloat64())
						right[i] = float32(rng.NormFloat64())
					}
					if trial == 0 {
						clear(left)
						clear(right)
					}
					l0 := append([]float32(nil), left...)
					r0 := append([]float32(nil), right...)
					g := make([]float32, n+2)
					for i := range g {
						g[i] = 77
					}
					c := append([]float32(nil), g...)
					var rp *float32
					if channels == 2 {
						rp = &right[0]
					}
					opuscc.Opus_pitch_downsample(nil, &left[0], rp, &g[1], int32(n), int32(channels), int32(factor), 0)
					nativePitchDownsample(left, right, c[1:], int32(n), int32(channels), int32(factor))
					if !sameFloatBits(g, c) || !sameFloatBits(left, l0) || !sameFloatBits(right, r0) {
						t.Fatal(n, factor, channels, trial, g, c)
					}
				}
			}
		}
	}
}

func TestAutocorrAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(912))
	for _, n := range []int{1, 4, 5, 8, 17, 64, 128} {
		for _, lag := range []int{0, 1, 3, 4, 8, 24} {
			// The four-lag scalar C kernel requires at least three samples.
			if lag >= n || (lag >= 3 && n-lag < 3) {
				continue
			}
			for _, overlap := range []int{0, 1, n / 2, n} {
				for trial := 0; trial < 12; trial++ {
					input := make([]float32, n)
					window := make([]float32, overlap)
					for i := range input {
						input[i] = float32(rng.NormFloat64() * 3)
					}
					for i := range window {
						window[i] = float32(rng.Float64())
					}
					if trial == 0 {
						for i := range input {
							input[i] = math.Float32frombits(uint32(i%2) << 31)
						}
					}
					if trial == 1 {
						for i := range input {
							input[i] = math.Float32frombits(uint32(i + 1))
						}
					}
					before := append([]float32(nil), input...)
					weights := append([]float32(nil), window...)
					goOut := make([]float32, lag+3)
					for i := range goOut {
						goOut[i] = 77
					}
					cOut := append([]float32(nil), goOut...)
					r := opuscc.Opus__celt_autocorr(nil, &input[0], &goOut[1], unsafe.SliceData(window), int32(overlap), int32(lag), int32(n), 0)
					c := nativeAutocorr(input, cOut[1:], window, int32(overlap), int32(lag), int32(n))
					if r != c || !sameFloatBits(goOut, cOut) || !sameFloatBits(input, before) || !sameFloatBits(window, weights) {
						t.Fatal(n, lag, overlap, trial, goOut, cOut)
					}
				}
			}
		}
	}
	// Input/output aliasing has C's correlation-then-tail update order.
	for _, overlap := range []int{0, 4} {
		g := make([]float32, 32)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		window := []float32{0.2, 0.4, 0.6, 0.8}
		opuscc.Opus__celt_autocorr(nil, &g[0], &g[1], &window[0], int32(overlap), 4, 32, 0)
		nativeAutocorr(c, c[1:], window, int32(overlap), 4, 32)
		if !sameFloatBits(g, c) {
			t.Fatal("alias", overlap, g, c)
		}
	}
}

func TestIIRAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(811))
	for _, ord := range []int{4, 8, 24} {
		for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 31, 64} {
			for trial := 0; trial < 12; trial++ {
				input := make([]float32, n+2)
				coeff := make([]float32, ord)
				mem := make([]float32, ord+2)
				for i := range input {
					input[i] = float32(rng.NormFloat64() * 3)
				}
				for i := range coeff {
					coeff[i] = float32(rng.NormFloat64() * 0.02)
				}
				for i := range mem {
					mem[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i%2) << 31)
					}
				}
				if trial == 1 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i + 1))
					}
				}
				before := append([]float32(nil), input...)
				goOut := make([]float32, ord+n+2)
				for i := range goOut {
					goOut[i] = 77
				}
				cOut := append([]float32(nil), goOut...)
				cMem := append([]float32(nil), mem...)
				opuscc.Opus_celt_iir(nil, &input[0], &coeff[0], &goOut[ord], int32(n), int32(ord), &mem[1], 0)
				nativeIIR(input, coeff, cOut[ord:], cMem[1:], int32(n), int32(ord))
				if !sameFloatBits(goOut, cOut) || !sameFloatBits(mem, cMem) || !sameFloatBits(input, before) {
					t.Fatal(ord, n, trial, goOut, cOut, mem, cMem)
				}
			}
		}
	}
	// In-place and partial overlaps preserve block read-ahead and store order.
	for _, offset := range []int{3, 4, 5} {
		g := make([]float32, 40)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		coeff := []float32{0.1, 0.2, 0.3, 0.4}
		gm := []float32{1, 2, 3, 4}
		cm := append([]float32(nil), gm...)
		opuscc.Opus_celt_iir(nil, &g[4], &coeff[0], &g[offset], 24, 4, &gm[0], 0)
		nativeIIR(c[4:], coeff, c[offset:], cm, 24, 4)
		if !sameFloatBits(g, c) || !sameFloatBits(gm, cm) {
			t.Fatal("overlap", offset, g, c)
		}
	}
}

func TestFIRAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(710))
	for _, ord := range []int{3, 4, 5, 7, 8, 24} {
		for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 31, 64} {
			for trial := 0; trial < 12; trial++ {
				input := make([]float32, ord+n+2)
				coeff := make([]float32, ord)
				for i := range input {
					input[i] = float32(rng.NormFloat64() * 3)
				}
				for i := range coeff {
					coeff[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i%2) << 31)
					}
				}
				if trial == 1 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i + 1))
					}
				}
				before := append([]float32(nil), input...)
				goOut := make([]float32, n+2)
				for i := range goOut {
					goOut[i] = 77
				}
				cOut := append([]float32(nil), goOut...)
				opuscc.Opus_celt_fir_c(nil, &input[ord], &coeff[0], &goOut[1], int32(n), int32(ord), 0)
				nativeFIR(input, coeff, cOut[1:], int32(n), int32(ord))
				if !sameFloatBits(goOut, cOut) || !sameFloatBits(input, before) {
					t.Fatal(ord, n, trial, goOut, cOut)
				}
			}
		}
	}
	// Partial overlaps are permitted (only exact x==y is asserted against).
	for _, offset := range []int{3, 5} {
		g := make([]float32, 40)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		coeff := []float32{0.1, 0.2, 0.3, 0.4}
		opuscc.Opus_celt_fir_c(nil, &g[4], &coeff[0], &g[offset], 24, 4, 0)
		nativeFIR(c, coeff, c[offset:], 24, 4)
		if !sameFloatBits(g, c) {
			t.Fatal("overlap", offset, g, c)
		}
	}
}

func TestLPCAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, p := range []int{1, 4, 16, 24} {
		for trial := 0; trial < 100; trial++ {
			signal := make([]float32, 128)
			for i := range signal {
				signal[i] = float32(rng.NormFloat64())
			}
			ac := make([]float32, p+1)
			for lag := range ac {
				for i := lag; i < len(signal); i++ {
					ac[lag] += float32(signal[i] * signal[i-lag])
				}
			}
			if trial == 0 {
				clear(ac)
			}
			if trial == 1 {
				for i := range ac {
					ac[i] = 1
				}
			}
			g, c := make([]float32, p), make([]float32, p)
			opuscc.Opus__celt_lpc(nil, &g[0], &ac[0], int32(p))
			nativeLPC(c, ac)
			for i := range g {
				if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
					t.Fatalf("p=%d trial=%d tap=%d Go=%g C=%g", p, trial, i, g[i], c[i])
				}
			}
		}
	}
}
