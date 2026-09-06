package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestLPVariableCutoffLocalTaps(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	state := OpusT_silk_LP_state{
		FIn_LP_State:         [2]OpusT_opus_int32{18000, -26000},
		Ftransition_frame_no: 17,
		Fmode:                1,
	}
	frame := []OpusT_opus_int16{1200, -2300, 3400, -4500, 5600, -6700, 7800, -8900}

	Opus_silk_LP_variable_cutoff(tls, uintptr(unsafe.Pointer(&state)), uintptr(unsafe.Pointer(&frame[0])), int32(len(frame)))

	wantFrame := []OpusT_opus_int16{454, -96, -78, 45, 7, -12, 3, 2}
	for i, value := range frame {
		if value != wantFrame[i] {
			t.Fatalf("frame[%d]: got %d, want %d", i, value, wantFrame[i])
		}
	}
	if got, want := state.FIn_LP_State, [2]OpusT_opus_int32{-15339370, -13652716}; got != want {
		t.Fatalf("filter state: got %v, want %v", got, want)
	}
	if got, want := state.Ftransition_frame_no, int32(18); got != want {
		t.Fatalf("transition frame: got %d, want %d", got, want)
	}
}
