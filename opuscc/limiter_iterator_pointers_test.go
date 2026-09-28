package opuscc

import (
	"math"
	"testing"
)

func TestLimiterPointers(t *testing.T) {
	bits := []uint32{0xff800000, 0xc0400000, 0xc0000000, 0xbf800000, 0x80000000, 0, 0x3f800000, 0x40000000, 0x40400000, 0x7f800000, 0x7fc01234}
	want := []uint32{0xc0000000, 0xc0000000, 0xc0000000, 0xbf800000, 0x80000000, 0, 0x3f800000, 0x40000000, 0x40000000, 0x40000000, 0x7fc01234}
	pcm := make([]float32, len(bits)+2)
	pcm[0], pcm[len(pcm)-1] = 99, 88
	for i, b := range bits {
		pcm[i+1] = math.Float32frombits(b)
	}
	if ret := Opus_opus_limit2_checkwithin1_c(nil, &pcm[1], int32(len(bits))); ret != 0 {
		t.Fatalf("nonempty return: %d", ret)
	}
	for i, b := range want {
		if got := math.Float32bits(pcm[i+1]); got != b {
			t.Fatalf("index %d: bits=%08x want=%08x", i, got, b)
		}
	}
	if pcm[0] != 99 || pcm[len(pcm)-1] != 88 {
		t.Fatal("sentinel modified")
	}
	for _, n := range []int32{0, -1} {
		if ret := Opus_opus_limit2_checkwithin1_c(nil, nil, n); ret != 1 {
			t.Fatalf("count=%d return=%d", n, ret)
		}
	}
	within := float32(0.5)
	if ret := Opus_opus_limit2_checkwithin1_c(nil, &within, 1); ret != 0 {
		t.Fatal("scalar limiter must not promise the +/-1 hint")
	}
}

func TestIteratorPointers(t *testing.T) {
	// Integer address sentinels are never dereferenced; internal uintptr
	// fields are deliberately outside this argument-only migration.
	iter := OpusT_OpusExtensionIterator{
		Fdata: 100, Fcurr_data: 101, Frepeat_data: 102, Flast_long: 103, Fsrc_data: 104,
		Flen1: 20, Fcurr_len: 19, Frepeat_len: 18, Fsrc_len: 17, Ftrailing_short_len: 16,
		Fnb_frames: 8, Fframe_max: 7, Fcurr_frame: 6, Frepeat_frame: 5, Frepeat_l: 1,
	}
	want := iter
	want.Fcurr_data, want.Frepeat_data = 100, 100
	want.Flast_long = 0
	want.Fcurr_len = 20
	want.Fcurr_frame, want.Frepeat_frame, want.Ftrailing_short_len = 0, 0, 0
	Opus_opus_extension_iterator_reset(nil, &iter)
	if iter != want {
		t.Fatalf("reset: got %+v want %+v", iter, want)
	}
	Opus_opus_extension_iterator_reset(nil, &iter)
	if iter != want {
		t.Fatal("reset not idempotent")
	}
	for _, frameMax := range []int32{0, 1, 48} {
		want.Fframe_max = frameMax
		Opus_opus_extension_iterator_set_frame_max(nil, &iter, frameMax)
		if iter != want {
			t.Fatalf("set max %d changed unexpected fields", frameMax)
		}
	}
}
