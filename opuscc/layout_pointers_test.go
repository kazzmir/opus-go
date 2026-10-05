package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestFFTArchitectureOwnerPointers(t *testing.T) {
	cfg := new(OpusT_kiss_fft_state)
	architecture := new(OpusT_arch_fft_state)
	payload := []byte{77, 88}
	architecture.Fis_supported = 1
	architecture.Fpriv = unsafe.Pointer(&payload[0])
	cfg.Farch_fft = architecture
	architecture = nil
	payload = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if cfg.Farch_fft == nil || cfg.Farch_fft.Fis_supported != 1 || *(*byte)(cfg.Farch_fft.Fpriv) != 77 {
		t.Fatal("FFT architecture ownership")
	}
	cfg.Farch_fft.Fpriv = nil
	cfg.Farch_fft = nil
	runtime.GC()
}

func TestRepacketizerAliasPointers(t *testing.T) {
	holder := new(struct{ State OpusRepacketizer })
	frame, padding := []byte{17, 19}, []byte{23, 29}
	holder.State.Fframes[47] = &frame[0]
	holder.State.Fpaddings[47] = &padding[0]
	frame, padding = nil, nil
	var canonical *OpusT_OpusRepacketizer = &holder.State
	entropyInitGrowStack(12)
	runtime.GC()
	if *canonical.Fframes[47] != 17 || *canonical.Fpaddings[47] != 23 {
		t.Fatal("canonical alias owners")
	}
}

func TestRepacketizerPaddingPointers(t *testing.T) {
	owner := new(OpusT_OpusRepacketizer)
	for i := range owner.Fpaddings {
		padding := []byte{byte(255 - i), byte(i)}
		owner.Fpaddings[i] = &padding[0]
		owner.Fpadding_len[i] = 2
		owner.Fpadding_nb_frames[i] = 1
	}
	entropyInitGrowStack(12)
	runtime.GC()
	for i, p := range owner.Fpaddings {
		if p == nil || *p != byte(255-i) || owner.Fpadding_len[i] != 2 || owner.Fpadding_nb_frames[i] != 1 {
			t.Fatal("repacketizer padding owner", i)
		}
	}
}

func TestRepacketizerFramePointers(t *testing.T) {
	owner := new(OpusT_OpusRepacketizer)
	for i := range owner.Fframes {
		packet := []byte{byte(i), byte(i + 1)}
		owner.Fframes[i] = &packet[0]
		owner.Flen1[i] = 2
	}
	entropyInitGrowStack(12)
	runtime.GC()
	for i, p := range owner.Fframes {
		if p == nil || *p != byte(i) || owner.Flen1[i] != 2 {
			t.Fatal("repacketizer frame owner", i)
		}
	}
}

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
