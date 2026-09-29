package opuscc

import "testing"

func TestPacketDurationPointers(t *testing.T) {
	for _, n := range []int32{-2, -1, 0} {
		if got := Opus_opus_packet_get_nb_samples(nil, nil, n, 48000); got != -1 {
			t.Fatal(n, got)
		}
	}
	toc := byte(3)
	if got := Opus_opus_packet_get_nb_samples(nil, &toc, 1, 48000); got != -4 {
		t.Fatal(got)
	}
	packet := [2]byte{0x83, 48}
	if got := Opus_opus_packet_get_nb_samples(nil, &packet[0], 2, 48000); got != 5760 {
		t.Fatal(got)
	}
	packet[1] = 49
	if got := Opus_opus_packet_get_nb_samples(nil, &packet[0], 2, 48000); got != -4 {
		t.Fatal(got)
	}
	packet[1] = 0
	if got := Opus_opus_packet_get_nb_samples(nil, &packet[0], 2, 48000); got != 0 {
		t.Fatal(got)
	}
}
