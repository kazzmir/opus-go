package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

// This port has no DRED model/SIMD initialization; test its existing disabled
// header behavior, not an enabled upstream DRED decoder or model lifetime.
func TestDREDInitPointers(t *testing.T) {
	var alias OpusDREDDecoder
	var canonical *OpusT_OpusDREDDecoder = &alias
	if canonical != &alias || unsafe.Sizeof(alias) != 12 {
		t.Fatal("canonical disabled DRED alias/layout")
	}
	owner := &struct {
		Before  uint32
		Decoder OpusT_OpusDREDDecoder
		After   uint32
	}{Before: 0x12345678, After: 0x87654321}
	owner.Decoder = OpusT_OpusDREDDecoder{Floaded: -1, Farch: -1, Fmagic: 0xffffffff}
	entropyInitGrowStack(12)
	runtime.GC()
	if ret := opusDREDDecoderInit(nil, &owner.Decoder); ret != OPUS_OK {
		t.Fatal("disabled DRED init", ret)
	}
	if owner.Decoder != (OpusT_OpusDREDDecoder{Fmagic: 0xD8EDDEC0}) || owner.Before != 0x12345678 || owner.After != 0x87654321 {
		t.Fatal("disabled DRED fields/guards", owner)
	}
}

// The function-value slot is a Go-only opaque-owner fixture, not a C function
// address and not permission to invoke Go closures through legacy integer ABIs.
func TestCleanupRecordPointers(t *testing.T) {
	makeRecord := func() *__ptcb {
		payload := []byte{41}
		counter := new(int32)
		closure := func(arg unsafe.Pointer) { *counter += int32(*(*byte)(arg)) }
		tail := &__ptcb{F__x: unsafe.Pointer(counter)}
		return &__ptcb{F__f: unsafe.Pointer(&closure), F__x: unsafe.Pointer(&payload[0]), F__next: tail}
	}
	record := makeRecord()
	entropyInitGrowStack(12)
	runtime.GC()
	callback := *(*func(unsafe.Pointer))(record.F__f)
	callback(record.F__x)
	if record.F__next == nil || *(*int32)(record.F__next.F__x) != 41 {
		t.Fatal("cleanup chain owners")
	}
	type oldLayout struct{ Function, Argument, Next uintptr }
	var old oldLayout
	if unsafe.Sizeof(*record) != unsafe.Sizeof(old) || unsafe.Offsetof(record.F__x) != unsafe.Offsetof(old.Argument) || unsafe.Offsetof(record.F__next) != unsafe.Offsetof(old.Next) {
		t.Fatal("cleanup layout changed")
	}
	record.F__f = nil
	record.F__x = nil
	record.F__next = nil
	runtime.GC()
}

func TestTimezoneOwnerPointers(t *testing.T) {
	type oldLayout struct {
		Fields [9]int32
		Offset int64
		Zone   uintptr
	}
	holder := new(tm)
	zone := []byte{'U', 'T', 'C', 0}
	holder.Ftm_zone = &zone[0]
	holder.Ftm_year = 123
	zone = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if got := unsafe.Slice(holder.Ftm_zone, 4); got[0] != 'U' || got[1] != 'T' || got[2] != 'C' || got[3] != 0 || holder.Ftm_year != 123 {
		t.Fatal("timezone owner", got)
	}
	var old oldLayout
	if unsafe.Sizeof(*holder) != unsafe.Sizeof(old) || unsafe.Offsetof(holder.Ftm_zone) != unsafe.Offsetof(old.Zone) {
		t.Fatal("existing timezone layout changed")
	}
}

func TestTimerHandlePointers(t *testing.T) {
	payload := []byte{29, 31}
	holder := new(struct{ Timer OpusT_timer_t })
	holder.Timer = unsafe.Pointer(&payload[0])
	payload = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if *(*byte)(holder.Timer) != 29 {
		t.Fatal("opaque timer owner")
	}
	if unsafe.Sizeof(holder.Timer) != unsafe.Sizeof(unsafe.Pointer(nil)) {
		t.Fatal("timer handle width")
	}
	holder.Timer = nil
	runtime.GC()
}

func TestLocaleHandlePointers(t *testing.T) {
	payload := []byte{19, 23}
	holder := new(struct{ Locale OpusT_locale_t })
	holder.Locale = unsafe.Pointer(&payload[0])
	payload = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if *(*byte)(holder.Locale) != 19 {
		t.Fatal("opaque locale owner")
	}
	if unsafe.Sizeof(holder.Locale) != unsafe.Sizeof(unsafe.Pointer(nil)) {
		t.Fatal("locale handle width")
	}
	holder.Locale = nil
	runtime.GC()
}

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
