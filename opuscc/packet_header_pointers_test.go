package opuscc

import "testing"

func TestPacketHeaderPointers(t *testing.T) {
	// RFC 6716's 32 TOC configurations, expressed independently of bit masks.
	bandwidths := [32]int32{
		OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND,
		OPUS_BANDWIDTH_MEDIUMBAND, OPUS_BANDWIDTH_MEDIUMBAND, OPUS_BANDWIDTH_MEDIUMBAND, OPUS_BANDWIDTH_MEDIUMBAND,
		OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND,
		OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND,
		OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND,
		OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND,
		OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND,
		OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND,
	}
	// Frame durations in units of 2.5 ms.
	durations := [32]int32{4, 8, 16, 24, 4, 8, 16, 24, 4, 8, 16, 24, 4, 8, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8}
	for config := 0; config < 32; config++ {
		for stereo := 0; stereo < 2; stereo++ {
			for code := 0; code < 4; code++ {
				toc := byte(config*8 + stereo*4 + code)
				if got := Opus_opus_packet_get_bandwidth(nil, &toc); got != bandwidths[config] {
					t.Fatalf("TOC %#x bandwidth: %d", toc, got)
				}
				if got := Opus_opus_packet_get_nb_channels(nil, &toc); got != int32(stereo+1) {
					t.Fatalf("TOC %#x channels: %d", toc, got)
				}
				for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
					want := rate * durations[config] / 400
					if got := Opus_opus_packet_get_samples_per_frame(nil, &toc, rate); got != want {
						t.Fatalf("TOC %#x rate %d: got %d samples, want %d", toc, rate, got, want)
					}
				}
				wantFrames := [4]int32{1, 2, 2, -4}[code]
				if got := Opus_opus_packet_get_nb_frames(nil, &toc, 1); got != wantFrames {
					t.Fatalf("TOC %#x one-byte frames: %d", toc, got)
				}
				for second := 0; second < 256; second++ {
					packet := [2]byte{toc, byte(second)}
					want := wantFrames
					if code == 3 {
						want = int32(second % 64)
					}
					if got := Opus_opus_packet_get_nb_frames(nil, &packet[0], 2); got != want {
						t.Fatalf("packet %x: got %d frames, want %d", packet, got, want)
					}
				}
			}
		}
	}
	for _, length := range []int32{-1, 0} {
		if got := Opus_opus_packet_get_nb_frames(nil, nil, length); got != -1 {
			t.Fatalf("length %d: got %d, want OPUS_BAD_ARG", length, got)
		}
	}
}
