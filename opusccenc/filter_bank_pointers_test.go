package opusccenc

import (
	"slices"
	"testing"
)

func TestFilterBankPointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, 20000, -30000, 0, 0, 1234}
	before := slices.Clone(input)
	wantL := []int16{-7653, 6051, 3611, 385, -6167}
	wantH := []int16{12966, 25842, -8844, -4362, -23179}
	initial := [2]int32{12345, -98765}
	wantState := [2]int32{-29581159, 5921226}
	for _, n := range []int32{10, 11} {
		state := initial
		low, high := [7]int16{111, 0, 0, 0, 0, 0, 222}, [7]int16{333, 0, 0, 0, 0, 0, 444}
		Opus_silk_ana_filt_bank_1(nil, &input[0], &state, &low[1], &high[1], n)
		if !slices.Equal(low[1:6], wantL) || !slices.Equal(high[1:6], wantH) || state != wantState {
			t.Fatalf("n=%d low=%v high=%v state=%v", n, low, high, state)
		}
		if low[0] != 111 || low[6] != 222 || high[0] != 333 || high[6] != 444 || !slices.Equal(input, before) {
			t.Fatal("sentinels or input changed")
		}
	}
	state := initial
	low := slices.Clone(input)
	high := make([]int16, 5)
	Opus_silk_ana_filt_bank_1(nil, &low[0], &state, &low[0], &high[0], 4)
	Opus_silk_ana_filt_bank_1(nil, &low[4], &state, &low[2], &high[2], 6)
	if !slices.Equal(low[:5], wantL) || !slices.Equal(high, wantH) || state != wantState {
		t.Fatal("in-place/chunked result differs")
	}
	for _, n := range []int32{0, 1} {
		Opus_silk_ana_filt_bank_1(nil, nil, &state, nil, nil, n)
		if state != wantState {
			t.Fatal("incomplete pair changed state")
		}
	}
	// Output order follows C even when the two output bands alias.
	shared := make([]int16, 5)
	state = initial
	Opus_silk_ana_filt_bank_1(nil, &input[0], &state, &shared[0], &shared[0], 10)
	if !slices.Equal(shared, wantH) {
		t.Fatal("output store order changed")
	}
}
