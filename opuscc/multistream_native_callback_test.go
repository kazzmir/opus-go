package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"slices"
	"testing"
	"weak"
)

type msNativeMarker struct{ Value int }
type msNativePayload struct {
	Marker *msNativeMarker
	Order  []int32
	Muted  int
}

//go:noinline
func msNativeCallbackOwners(t *testing.T, output []float32) (OpusMSChannelCopy, weak.Pointer[msNativeMarker], weak.Pointer[msNativeMarker], weak.Pointer[msNativePayload]) {
	payload := &msNativePayload{Marker: &msNativeMarker{71}}
	captured := &msNativeMarker{93}
	callback := func(_ *libc.TLS, ds, dc int32, src *float32, ss, n int32) {
		entropyInitGrowStack(12)
		runtime.GC()
		if captured.Value != 93 || payload.Marker.Value != 71 {
			t.Fatal("typed closure/payload Go ownership")
		}
		payload.Order = append(payload.Order, dc)
		if src == nil {
			payload.Muted++
		} else if ss != 2 {
			t.Fatal("stereo source stride", ss)
		}
		dst := output[:ds*n]
		for i := int32(0); i < n; i++ {
			dst[i*ds+dc] = float32(dc + 1)
		}
	}
	return callback, weak.Make(payload.Marker), weak.Make(captured), weak.Make(payload)
}

func TestMultistreamCallbackAliasRootsPointers(t *testing.T) {
	pcm := []float32{77, 0, 88}
	callback, marker, captured, _ := msNativeCallbackOwners(t, pcm[1:2])
	var bound OpusT___ccgo_fp__Xopus_multistream_decode_native_4 = callback
	callback = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if marker.Value() == nil || captured.Value() == nil {
		t.Fatal("typed closure/payload roots")
	}
	src := [2]float32{.25, -.25}
	bound(nil, 1, 0, &src[0], 2, 1)
	if pcm[0] != 77 || pcm[1] != 1 || pcm[2] != 88 {
		t.Fatal("typed destination/guards", pcm)
	}
	runtime.KeepAlive(bound)
}

func TestMultistreamNativeCallbackPointers(t *testing.T) {
	mapping := [4]byte{0, 1, 0, 255}
	st, err := Opus_opus_multistream_decoder_create_typed(48000, 4, 1, 1, &mapping[0])
	if err != nil {
		t.Fatal(err)
	}
	pcm := make([]float32, 4*120+2)
	pcm[0], pcm[len(pcm)-1] = 77, 88
	callback, marker, captured, payloadOwner := msNativeCallbackOwners(t, pcm[1:len(pcm)-1])
	runtime.GC()
	if marker.Value() == nil || captured.Value() == nil {
		t.Fatal("callback roots")
	}
	if got := Opus_opus_multistream_decode_native(nil, st, nil, 0, callback, 120, 0, 0); got != 120 {
		t.Fatal(got)
	}
	payload := payloadOwner.Value()
	if payload == nil || !slices.Equal(payload.Order, []int32{0, 2, 1, 3}) || payload.Muted != 1 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
		t.Fatal("callback order/muted/guards", payload)
	}
	for i, v := range pcm[1 : len(pcm)-1] {
		if v != float32(i%4+1) {
			t.Fatal("callback destination", i, v)
		}
	}
	if got := Opus_opus_multistream_decode_native(nil, st, nil, 0, nil, 0, 0, 0); got != -1 {
		t.Fatal("validation must not invoke nil callback", got)
	}
	runtime.KeepAlive(callback)
}
