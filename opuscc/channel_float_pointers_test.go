package opuscc

import (
	"math"
	"testing"
)

func TestChannelFloatPointers(t *testing.T) {
	opus_copy_channel_out_float(nil, nil, 0, 0, nil, 0, 0)
	b := [7]float32{77, 1, 2, 3, 4, 5, 88}
	opus_copy_channel_out_float(nil, &b[1], 1, 0, &b[1], 1, 5)
	opus_copy_channel_out_float(nil, &b[1], 2, 1, nil, 1, 2)
	if b != [7]float32{77, 1, 0, 3, 0, 5, 88} {
		t.Fatal(b)
	}
	v := math.Float32frombits(0x7fc01234)
	var out float32
	opus_copy_channel_out_float(nil, &out, 0, 0, &v, 0, 3)
	if math.Float32bits(out) != 0x7fc01234 {
		t.Fatal(out)
	}
	b = [7]float32{77, 1, 2, 3, 4, 5, 88}
	opus_copy_channel_out_float(nil, &b[2], 1, 0, &b[1], 1, 4)
	if b != [7]float32{77, 1, 1, 1, 1, 1, 88} {
		t.Fatal("forward overlap", b)
	}
}
