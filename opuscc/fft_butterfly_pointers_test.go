package opuscc

import (
	"slices"
	"testing"
)

func TestFFTButterfly2Pointers(t *testing.T) {
	kf_bfly2(nil, nil, 4, 0)
	data := make([]OpusT_kiss_fft_cpx, 18)
	for i := range data {
		data[i] = OpusT_kiss_fft_cpx{Fr: float32(i), Fi: float32(-i)}
	}
	before := slices.Clone(data)
	kf_bfly2(nil, &data[1], 4, 2)
	if data[0] != before[0] || data[17] != before[17] || data[1].Fr != 6 || data[5].Fr != -4 {
		t.Fatal(data)
	}
}
