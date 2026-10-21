package opuscc

import (
	"runtime"
	"testing"
)

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
