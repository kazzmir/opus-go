package opuscc

import (
	"testing"
	"unsafe"
)

func TestValidateLayoutFieldAccesses(t *testing.T) {
	valid := OpusT_ChannelLayout{
		Fnb_channels:        6,
		Fnb_streams:         3,
		Fnb_coupled_streams: 1,
		Fmapping:            [256]uint8{0, 1, 2, 3, 255, 0},
	}
	if got, want := Opus_validate_layout(nil, uintptr(unsafe.Pointer(&valid))), int32(1); got != want {
		t.Fatalf("valid layout: got %d, want %d", got, want)
	}

	invalidMapping := valid
	invalidMapping.Fmapping[3] = 4
	if got, want := Opus_validate_layout(nil, uintptr(unsafe.Pointer(&invalidMapping))), int32(0); got != want {
		t.Fatalf("invalid mapping: got %d, want %d", got, want)
	}

	tooManyStreams := valid
	tooManyStreams.Fnb_streams = 200
	tooManyStreams.Fnb_coupled_streams = 56
	if got, want := Opus_validate_layout(nil, uintptr(unsafe.Pointer(&tooManyStreams))), int32(0); got != want {
		t.Fatalf("too many streams: got %d, want %d", got, want)
	}
}
