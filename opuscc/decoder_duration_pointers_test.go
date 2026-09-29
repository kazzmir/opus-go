package opuscc

import "testing"

func TestDecoderDurationPointers(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		dec := OpusT_OpusDecoder{FFs: rate}
		before := dec
		for _, n := range []int32{-1, 0} {
			if got := Opus_opus_decoder_get_nb_samples(nil, &dec, nil, n); got != -1 {
				t.Fatal(got)
			}
		}
		p := [2]byte{0x83, 48}
		if got := Opus_opus_decoder_get_nb_samples(nil, &dec, &p[0], 2); got != rate*120/1000 {
			t.Fatal(rate, got)
		}
		p[1] = 49
		if got := Opus_opus_decoder_get_nb_samples(nil, &dec, &p[0], 2); got != -4 {
			t.Fatal(rate, got)
		}
		if dec != before {
			t.Fatal("state modified")
		}
	}
}
