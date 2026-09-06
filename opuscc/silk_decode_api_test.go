package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestSilkDecodeLostFrameState(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	var size int32
	Opus_silk_Get_Decoder_Size(tls, uintptr(unsafe.Pointer(&size)))
	memory := libc.Xmalloc(tls, uint64(size))
	if got := Opus_silk_InitDecoder(tls, memory); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_silk_decoder)(unsafe.Pointer(memory))
	decoder.FnChannelsAPI = 1
	decoder.FnChannelsInternal = 1
	decoder.Fchannel_state[0].Ffs_API_hz = 8000
	decoder.Fchannel_state[0].Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder.Fchannel_state[0])), 8, 8000); got != OPUS_OK {
		t.Fatalf("sample rate setup: got %d", got)
	}
	Opus_silk_PLC_Reset(tls, uintptr(unsafe.Pointer(&decoder.Fchannel_state[0])))
	decoder.Fchannel_state[0].FsPLC.FprevGain_Q16 = [2]OpusT_opus_int32{65536, 65536}
	decoder.Fchannel_state[0].FsPLC.FpitchL_Q8 = 20 << 8
	decoder.Fchannel_state[0].FsPLC.Frand_seed = 12345
	var control OpusT_silk_DecControlStruct
	control.FAPI_sampleRate = 8000
	control.FnChannelsAPI = 1
	control.FnChannelsInternal = 1
	output := make([]int16, 320)
	var samples int32
	if got := Opus_silk_Decode(tls, memory, uintptr(unsafe.Pointer(&control)), 1, 1, 0, uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&samples)), 0); got != OPUS_OK {
		t.Fatalf("decode result: got %d", got)
	}
	t.Logf("samples=%d output=%v losses=%d", samples, output[:8], decoder.Fchannel_state[0].FlossCnt)
}
