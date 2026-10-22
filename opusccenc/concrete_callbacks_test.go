package opusccenc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"testing"
	"weak"
)

type encoderCallbackOwner struct{ Values [32]int32 }
type encoderCallbackRoots struct {
	In      OpusT_opus_copy_channel_in_func
	Out     OpusT_opus_copy_channel_out_func
	Downmix OpusT_downmix_func
}

//go:noinline
func encoderCallbackOwners(t *testing.T) (encoderCallbackRoots, weak.Pointer[encoderCallbackOwner]) {
	owner := new(encoderCallbackOwner)
	owner.Values[31] = 71
	check := func() {
		runtime.GC()
		if owner.Values[31] != 71 {
			t.Fatal("lost concrete callback environment")
		}
	}
	return encoderCallbackRoots{
		In: func(tls *libc.TLS, dst uintptr, ds int32, src uintptr, ss, ch, n int32, user uintptr) {
			check()
			if tls != nil || dst != 0 || src != 0 || user != 0 || ds != 2 || ss != 4 || ch != 1 || n != 3 {
				t.Fatal("input callback arguments")
			}
			owner.Values[0]++
		},
		Out: func(tls *libc.TLS, dst uintptr, ds, ch int32, src uintptr, ss, n int32, user uintptr) {
			check()
			if tls != nil || dst != 0 || src != 0 || user != 0 || ds != 4 || ch != 2 || ss != 2 || n != 3 {
				t.Fatal("output callback arguments")
			}
			owner.Values[1]++
		},
		Downmix: func(tls *libc.TLS, src, dst uintptr, n, offset, c1, c2, channels int32) {
			check()
			if tls != nil || src != 0 || dst != 0 || n != 3 || offset != 5 || c1 != 1 || c2 != 2 || channels != 4 {
				t.Fatal("downmix callback arguments")
			}
			owner.Values[2]++
		},
	}, weak.Make(owner)
}

//go:noinline
func encoderCallbackGrowStack(n int) int {
	var pad [1024]int
	pad[n%len(pad)] = n
	if n > 0 {
		return pad[n%len(pad)] + encoderCallbackGrowStack(n-1)
	}
	return pad[0]
}

func TestEncoderConcreteCallbackRoots(t *testing.T) {
	original, owner := encoderCallbackOwners(t)
	callbacks := original
	original = encoderCallbackRoots{}
	encoderCallbackGrowStack(12)
	runtime.GC()
	if owner.Value() == nil {
		t.Fatal("callback copy must retain actual Go environment")
	}
	// Null pointer arguments only: these callbacks use captured typed Go storage,
	// not fabricated integer pointer sentinels or an erased function identity.
	callbacks.In(nil, 0, 2, 0, 4, 1, 3, 0)
	callbacks.Out(nil, 0, 4, 2, 0, 2, 3, 0)
	callbacks.Downmix(nil, 0, 0, 3, 5, 1, 2, 4)
	state := owner.Value()
	if state == nil || state.Values[0] != 1 || state.Values[1] != 1 || state.Values[2] != 1 {
		t.Fatal("typed callback dispatch/capture", state)
	}
	runtime.KeepAlive(callbacks)
}
