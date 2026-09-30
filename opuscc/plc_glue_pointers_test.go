package opuscc

import "testing"

func TestPLCGluePointers(t *testing.T) {
	var d OpusT_silk_decoder_state
	Opus_silk_PLC_glue_frames(nil, &d, nil, 0)
	var lost [16]int16
	for i := range lost {
		lost[i] = 100
	}
	d.FlossCnt = 1
	Opus_silk_PLC_glue_frames(nil, &d, &lost[0], 16)
	if d.FsPLC.Fconc_energy != 160000 || d.FsPLC.Flast_frame_lost != 1 {
		t.Fatal(d.FsPLC)
	}
	d.FlossCnt = 0
	good := [18]int16{0: 77, 17: 88}
	for i := 1; i < 17; i++ {
		good[i] = 1000
	}
	Opus_silk_PLC_glue_frames(nil, &d, &good[1], 16)
	if good[0] != 77 || good[17] != 88 || good[1] >= 1000 || good[16] != 1000 || d.FsPLC.Flast_frame_lost != 0 {
		t.Fatal(d.FsPLC, good)
	}
}
