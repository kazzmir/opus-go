package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestPitchXCorrOwnedAliasesPointers(t *testing.T) {
	for _, n := range []int32{3, 4, 5, 7, 8, 9, 16, 19} {
		for _, lags := range []int32{1, 2, 3, 4, 5, 8, 9, 13} {
			for _, origin := range []int32{1, 2, n + 3, n + 4} {
				a := make([]float32, 3*n+3*lags+8)
				for i := range a {
					a[i] = float32(i%11-5) / 8
				}
				a[0], a[len(a)-1] = 77, 88
				b := slices.Clone(a)
				entropyInitGrowStack(12)
				runtime.GC()
				Opus_celt_pitch_xcorr_c(nil, &a[1], &a[n+3], &a[origin], n, lags, 0)
				pitchXCorrOwned(nil, b[1:1+n], b[n+3:n+3+n+lags-1], b[origin:origin+lags], n, lags, 0)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("xcorr full-buffer live aliases/guards", n, lags, origin, i)
					}
				}
			}
		}
	}
}

func TestPitchXCorrValidationPointers(t *testing.T) {
	want := "Fatal (internal) error in " + opusDiagnosticString(4659) + ", line 265: " + opusDiagnosticString(4629) + "\n"
	checkDiagnosticFatal(t, func() { Opus_celt_pitch_xcorr_c(nil, nil, nil, nil, math.MaxInt32, 0, 0) }, want)
	checkDiagnosticFatal(t, func() { pitchXCorrOwned(nil, nil, nil, nil, math.MaxInt32, 0, 0) }, want)
	x, y, out := [2]float32{}, [5]float32{}, [6]float32{77, 1, 2, 3, 4, 88}
	before := out
	want = "Fatal (internal) error in " + opusDiagnosticString(3374) + ", line 69: " + opusDiagnosticString(3349) + "\n"
	checkDiagnosticFatal(t, func() { pitchXCorrOwned(nil, x[:], y[:], out[1:5], 2, 4, 0) }, want)
	if out != before {
		t.Fatal("unrolled assertion before any output store")
	}
	short := [1]float32{}
	var caught any
	func() {
		defer func() { caught = recover() }()
		pitchXCorrOwned(nil, short[:], y[:], out[1:5], 2, 4, 0)
	}()
	if _, ok := caught.(runtime.Error); !ok || out != before {
		t.Fatal("input view precedes unrolled assertion")
	}
}

func TestPitchXCorrHistoryExtentPointers(t *testing.T) {
	x, history := [3]float32{1, 2, 3}, [6]float32{1, 2, 3, 4, 5, 6}
	a, b := [6]float32{77, 0, 0, 0, 0, 88}, [6]float32{77, 0, 0, 0, 0, 88}
	Opus_celt_pitch_xcorr_c(nil, &x[0], &history[0], &a[1], 3, 4, 0)
	pitchXCorrOwned(nil, x[:], history[:], b[1:5], 3, 4, 0)
	if a != b || b != [6]float32{77, 14, 20, 26, 32, 88} {
		t.Fatal("exact final-lag history extent/guards", a, b)
	}
	before := b
	var caught any
	func() {
		defer func() { caught = recover() }()
		pitchXCorrOwned(nil, x[:], history[:5:5], b[1:5], 3, 4, 0)
	}()
	if _, ok := caught.(runtime.Error); !ok || b != before {
		t.Fatal("short history rejected before stores/unrolled work")
	}
	want := "Fatal (internal) error in " + opusDiagnosticString(3374) + ", line 69: " + opusDiagnosticString(3349) + "\n"
	checkDiagnosticFatal(t, func() { pitchXCorrOwned(nil, nil, nil, b[1:5], 0, 4, 0) }, want)
	if b != before {
		t.Fatal("zero-length unrolled assertion has no stores or history reads")
	}
}

func TestPitchXCorrPointers(t *testing.T) {
	for n := 0; n <= 19; n++ {
		for lags := 1; lags <= 13; lags++ {
			if lags > 3 && n < 3 {
				continue
			}
			x := make([]float32, n)
			y := make([]float32, 0)
			if n > 0 {
				y = make([]float32, n+lags-1)
			}
			for i := range x {
				x[i] = float32(i-7) * 0.125
			}
			for i := range y {
				y[i] = float32(i%11-5) * 0.75
			}
			out := make([]float32, lags+2)
			out[0] = 123
			out[lags+1] = 456
			span := append([]float32(nil), out...)
			owned := slices.Clone(out)
			Opus_celt_pitch_xcorr_c(nil, unsafe.SliceData(x), unsafe.SliceData(y), &out[1], int32(n), int32(lags), 0)
			pitchXCorrSpans(nil, x, y, span[1:lags+1], int32(n), int32(lags))
			pitchXCorrOwned(nil, x, y, owned[1:lags+1], int32(n), int32(lags), 0)
			for i := range out {
				if math.Float32bits(out[i]) != math.Float32bits(span[i]) || math.Float32bits(out[i]) != math.Float32bits(owned[i]) {
					t.Fatal("pitch xcorr kernel/guards", n, lags, i)
				}
			}
			for lag := 0; lag < lags; lag++ {
				var want float32
				for j := range x {
					want += float32(x[j] * y[lag+j])
				}
				if math.Float32bits(out[lag+1]) != math.Float32bits(want) {
					t.Fatalf("n=%d lags=%d lag=%d", n, lags, lag)
				}
			}
			if out[0] != 123 || out[lags+1] != 456 {
				t.Fatal("guards")
			}
		}
	}
}
