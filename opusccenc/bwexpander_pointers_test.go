package opusccenc

import (
	"slices"
	"testing"
)

func TestBWExpanderPointers(t *testing.T) {
	for _, chirp := range []int32{0, 32768, 65536} {
		a16 := []int16{111, 1, 3, -3, -32768, 222}
		a32 := []int32{111, 1, 3, -3, -32768, 222}
		w16 := slices.Clone(a16)
		w32 := slices.Clone(a32)
		switch chirp {
		case 0:
			clear(w16[1:5])
			clear(w32[1:5])
		case 32768:
			copy(w16[1:5], []int16{1, 1, 0, -2048})
			copy(w32[1:5], []int32{0, 0, -1, -2048})
		}
		Opus_silk_bwexpander(nil, &a16[1], 4, chirp)
		Opus_silk_bwexpander_32(nil, &a32[1], 4, chirp)
		if !slices.Equal(a16, w16) || !slices.Equal(a32, w32) {
			t.Fatalf("chirp=%d 16=%v 32=%v", chirp, a16, a32)
		}
	}
	for _, value := range []int16{-32768, -3, -1, 0, 1, 3, 32767} {
		got := value
		Opus_silk_bwexpander(nil, &got, 1, 32768)
		want := int16((int32(value) + 1) >> 1)
		if got != want {
			t.Fatalf("value=%d got=%d want=%d", value, got, want)
		}
	}
}
