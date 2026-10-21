package opuscc

import (
	"math"
	"runtime"
	"testing"
)

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
