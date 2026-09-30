package opuscc

import (
	"math"
	"testing"
)

func TestChannelShortPointers(t *testing.T) {
	opus_copy_channel_out_short(nil, nil, 0, 0, nil, 0, 0)
	src := [8]float32{-2, -.5, -.5 / 32768, .5 / 32768, 1.5 / 32768, 2.5 / 32768, 2, math.Float32frombits(0x7fc01234)}
	var dst [10]int16
	dst[0] = 77
	dst[9] = 88
	opus_copy_channel_out_short(nil, &dst[1], 1, 0, &src[0], 1, 8)
	if dst != [10]int16{77, -32768, -16384, 0, 0, 2, 2, 32767, -32768, 88} {
		t.Fatal(dst)
	}
	opus_copy_channel_out_short(nil, &dst[1], 2, 1, nil, 0, 4)
	if dst[2] != 0 || dst[4] != 0 || dst[6] != 0 || dst[8] != 0 || dst[0] != 77 || dst[9] != 88 {
		t.Fatal(dst)
	}
}
