package opuscc

import "testing"

func TestProjectionShortPointers(t *testing.T) {
	opus_projection_copy_channel_out_short(nil, nil, 0, 0, nil, 0, 0, nil)
	m := mappingTestBuffer{Header: OpusT_MappingMatrix{Frows: 2, Fcols: 2}, Data: [36]int16{16384, 8192, -8192, 16384}}
	src := [2]float32{.5, .25}
	dst := [6]int16{77, 9, 9, 9, 9, 88}
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 0, &src[0], 1, 2, &m.Header)
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 1, &src[0], 1, 2, &m.Header)
	if dst != [6]int16{77, 4096, 12288, 2048, 6144, 88} {
		t.Fatal(dst)
	}
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 1, nil, 1, 2, nil)
	if dst[1] != 4096 {
		t.Fatal("nonfirst nil must not clear")
	}
	opus_projection_copy_channel_out_short(nil, &dst[1], 2, 0, nil, 1, 2, nil)
	if dst != [6]int16{77, 0, 0, 0, 0, 88} {
		t.Fatal("clear", dst)
	}
}
