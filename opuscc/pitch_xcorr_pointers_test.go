package opuscc

import (
	"math"
	"testing"
	"unsafe"
)

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
			Opus_celt_pitch_xcorr_c(nil, unsafe.SliceData(x), unsafe.SliceData(y), &out[1], int32(n), int32(lags), 0)
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
