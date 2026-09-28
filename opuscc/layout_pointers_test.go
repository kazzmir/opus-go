package opuscc

import "testing"

func TestLayoutChannelPointers(t *testing.T) {
	layout := OpusT_ChannelLayout{Fnb_channels: 9, Fnb_streams: 4, Fnb_coupled_streams: 2,
		Fmapping: [256]uint8{0, 1, 2, 3, 4, 5, 0, 1, 255}}
	before := layout
	for _, tc := range []struct {
		kind               string
		stream, prev, want int32
	}{
		{"left", 0, -9, 0}, {"left", 0, 0, 6}, {"left", 0, 6, -1},
		{"right", 0, -1, 1}, {"right", 0, 1, 7}, {"right", 0, 7, -1},
		{"left", 1, -1, 2}, {"right", 1, -1, 3},
		{"mono", 2, -1, 4}, {"mono", 3, -1, 5}, {"mono", 3, 5, -1},
		{"left", 0, 9, -1}, {"right", 0, 9, -1}, {"mono", 2, 9, -1},
	} {
		var got int32
		switch tc.kind {
		case "left":
			got = Opus_get_left_channel(nil, &layout, tc.stream, tc.prev)
		case "right":
			got = Opus_get_right_channel(nil, &layout, tc.stream, tc.prev)
		case "mono":
			got = Opus_get_mono_channel(nil, &layout, tc.stream, tc.prev)
		}
		if got != tc.want {
			t.Fatalf("%s stream=%d prev=%d: got %d want %d", tc.kind, tc.stream, tc.prev, got, tc.want)
		}
	}
	if Opus_validate_layout(nil, &layout) != 1 || layout != before {
		t.Fatal("valid layout rejected or modified")
	}
	layout.Fnb_streams = 254
	layout.Fnb_coupled_streams = 1
	layout.Fmapping[0] = 254
	if Opus_validate_layout(nil, &layout) != 1 {
		t.Fatal("255 coded channels should be valid")
	}
	layout.Fnb_streams++
	if Opus_validate_layout(nil, &layout) != 0 {
		t.Fatal("256 coded channels should be rejected")
	}
	layout = OpusT_ChannelLayout{}
	if Opus_get_left_channel(nil, &layout, 0, -1) != -1 || Opus_get_right_channel(nil, &layout, 0, -1) != -1 || Opus_get_mono_channel(nil, &layout, 0, -1) != -1 {
		t.Fatal("empty layout should have no channels")
	}
}
