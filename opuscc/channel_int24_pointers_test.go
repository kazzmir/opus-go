package opuscc

import "testing"

func TestChannelInt24Pointers(t *testing.T) {
	opus_copy_channel_out_int24(nil, nil, 0, 0, nil, 0, 0)
	src := [8]float32{-2, -.5, -.5 / 8388608, .5 / 8388608, 1.5 / 8388608, 2.5 / 8388608, 2, 255}
	var dst [10]int32
	dst[0] = 77
	dst[9] = 88
	opus_copy_channel_out_int24(nil, &dst[1], 1, 0, &src[0], 1, 8)
	if dst != [10]int32{77, -16777216, -4194304, 0, 0, 2, 2, 16777216, 2139095040, 88} {
		t.Fatal(dst)
	}
	opus_copy_channel_out_int24(nil, &dst[1], 2, 1, nil, 0, 4)
	if dst[2] != 0 || dst[4] != 0 || dst[6] != 0 || dst[8] != 0 || dst[0] != 77 || dst[9] != 88 {
		t.Fatal(dst)
	}
}
