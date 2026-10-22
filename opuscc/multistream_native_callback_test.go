package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"slices"
	"testing"
	"unsafe"
	"weak"
)

type msNativeMarker struct{ Value int }
type msNativePayload struct {
	Marker *msNativeMarker
	Order  []int32
	Muted  int
}

//go:noinline
func msNativeCallbackOwners(t *testing.T) (unsafe.Pointer, weak.Pointer[msNativeMarker], OpusMSChannelCopy, weak.Pointer[msNativeMarker]) {
	payload := &msNativePayload{Marker: &msNativeMarker{71}}
	captured := &msNativeMarker{93}
	callback := func(_ *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32, user unsafe.Pointer) {
		entropyInitGrowStack(12)
		runtime.GC()
		p := (*msNativePayload)(user)
		if captured.Value != 93 || p.Marker.Value != 71 {
			t.Fatal("callback/payload Go ownership")
		}
		p.Order = append(p.Order, dc)
		if src == nil {
			p.Muted++
		} else if ss != 2 {
			t.Fatal("stereo source stride", ss)
		}
		output := unsafe.Slice((*float32)(dst), ds*n)
		for i := int32(0); i < n; i++ {
			output[i*ds+dc] = float32(dc + 1)
		}
	}
	return unsafe.Pointer(payload), weak.Make(payload.Marker), callback, weak.Make(captured)
}

func TestMultistreamCallbackAliasRootsPointers(t *testing.T) {
	user, marker, callback, captured := msNativeCallbackOwners(t)
	pcm := []float32{77, 0, 88}
	bound := opusMSBindBoundCopy(callback, user, unsafe.Pointer(&pcm[1]))
	callback, user = nil, nil
	entropyInitGrowStack(12)
	runtime.GC()
	if marker.Value() == nil || captured.Value() == nil {
		t.Fatal("bound typed callback/payload roots")
	}
	src := [2]float32{.25, -.25}
	bound(nil, 1, 0, &src[0], 2, 1)
	if pcm[0] != 77 || pcm[1] != 1 || pcm[2] != 88 {
		t.Fatal("bound destination/guards", pcm)
	}
	runtime.KeepAlive(bound)
}

func TestMultistreamNativeCallbackPointers(t *testing.T) {
	mapping := [4]byte{0, 1, 0, 255}
	st, err := Opus_opus_multistream_decoder_create_typed(48000, 4, 1, 1, &mapping[0])
	if err != nil {
		t.Fatal(err)
	}
	user, marker, callback, captured := msNativeCallbackOwners(t)
	pcm := make([]float32, 4*120+2)
	pcm[0], pcm[len(pcm)-1] = 77, 88
	runtime.GC()
	if marker.Value() == nil || captured.Value() == nil {
		t.Fatal("callback roots")
	}
	if got := Opus_opus_multistream_decode_native(nil, st, nil, 0, unsafe.Pointer(&pcm[1]), callback, 120, 0, 0, user); got != 120 {
		t.Fatal(got)
	}
	payload := (*msNativePayload)(user)
	if !slices.Equal(payload.Order, []int32{0, 2, 1, 3}) || payload.Muted != 1 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
		t.Fatal("callback order/muted/guards", payload)
	}
	for i, v := range pcm[1 : len(pcm)-1] {
		if v != float32(i%4+1) {
			t.Fatal("callback destination", i, v)
		}
	}
	if got := Opus_opus_multistream_decode_native(nil, st, nil, 0, nil, nil, 0, 0, 0, nil); got != -1 {
		t.Fatal("validation must not invoke nil callback", got)
	}
}
