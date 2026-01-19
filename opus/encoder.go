package opus

import (
	"errors"
	"fmt"
	"sync"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"

	"github.com/kazzmir/opus-go/opusccenc"
)

var (
	ErrEncodeFailed = errors.New("opus: encode failed")
	ErrCtlFailed    = errors.New("opus: encoder ctl failed")
)

const (
	ApplicationVoIP               = int(opusccenc.OPUS_APPLICATION_VOIP)
	ApplicationAudio              = int(opusccenc.OPUS_APPLICATION_AUDIO)
	ApplicationRestrictedLowDelay = int(opusccenc.OPUS_APPLICATION_RESTRICTED_LOWDELAY)
	ApplicationRestrictedSilk     = int(opusccenc.OPUS_APPLICATION_RESTRICTED_SILK)
	ApplicationRestrictedCelt     = int(opusccenc.OPUS_APPLICATION_RESTRICTED_CELT)
)

// Encoder is an Opus encoder backed by ccgo-transpiled libopus.
//
// Note: This encoder uses the same stdlib-only runtime shim as the decoder.
type Encoder struct {
	mu sync.Mutex

	tls *libc.TLS
	st  uintptr

	sampleRate  int
	channels    int
	application int
}

func NewEncoder(sampleRate, channels, application int) (*Encoder, error) {
	tls := libc.NewTLS()
	if tls == nil {
		return nil, errors.New("opus: failed to allocate TLS")
	}

	errPtr := tls.Alloc(4)
	defer tls.Free(4)

	st := opusccenc.Opus_opus_encoder_create(tls, opusccenc.OpusT_opus_int32(sampleRate), int32(channels), int32(application), errPtr)
	errCode := *(*int32)(unsafe.Pointer(errPtr))
	if errCode != opusccenc.OPUS_OK || st == 0 {
		opusccenc.Opus_opus_encoder_destroy(tls, st)
		opusccenc.FreePseudostackTLS(tls)
		tls.Close()
		return nil, fmt.Errorf("opus: encoder_create failed: %s (%d)", opusccencErrorString(tls, errCode), errCode)
	}

	return &Encoder{tls: tls, st: st, sampleRate: sampleRate, channels: channels, application: application}, nil
}

func (e *Encoder) Close() error {
	if e == nil {
		return nil
	}
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.tls != nil {
		if e.st != 0 {
			opusccenc.Opus_opus_encoder_destroy(e.tls, e.st)
			e.st = 0
		}
		opusccenc.FreePseudostackTLS(e.tls)
		e.tls.Close()
		e.tls = nil
	}
	return nil
}

func (e *Encoder) SampleRate() int  { return e.sampleRate }
func (e *Encoder) Channels() int    { return e.channels }
func (e *Encoder) Application() int { return e.application }

func (e *Encoder) SetBitrate(bps int) error {
	return e.ctlInt32(int32(opusccenc.OPUS_SET_BITRATE_REQUEST), int32(bps))
}

func (e *Encoder) SetVBR(enabled bool) error {
	var v int32
	if enabled {
		v = 1
	}
	return e.ctlInt32(int32(opusccenc.OPUS_SET_VBR_REQUEST), v)
}

func (e *Encoder) SetComplexity(complexity int) error {
	return e.ctlInt32(int32(opusccenc.OPUS_SET_COMPLEXITY_REQUEST), int32(complexity))
}

// Lookahead returns the encoder lookahead in samples at 48 kHz.
//
// This is typically used as the OpusHead PreSkip value.
func (e *Encoder) Lookahead() (int, error) {
	if e == nil {
		return 0, errors.New("opus: encoder closed")
	}
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.tls == nil || e.st == 0 {
		return 0, errors.New("opus: encoder closed")
	}

	bp := e.tls.Alloc(32)
	defer e.tls.Free(32)
	// Store the output int32 in the second half to avoid overlap with VaList storage.
	outPtr := bp + 16
	*(*int32)(unsafe.Pointer(outPtr)) = 0

	ret := opusccenc.Opus_opus_encoder_ctl(
		e.tls,
		e.st,
		int32(opusccenc.OPUS_GET_LOOKAHEAD_REQUEST),
		libc.VaList(bp, uintptr(outPtr)),
	)
	if ret != opusccenc.OPUS_OK {
		return 0, fmt.Errorf("%w: %s (%d)", ErrCtlFailed, opusccencErrorString(e.tls, ret), ret)
	}
	return int(*(*int32)(unsafe.Pointer(outPtr))), nil
}

func (e *Encoder) ctlInt32(request int32, value int32) error {
	if e == nil {
		return errors.New("opus: encoder closed")
	}
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.tls == nil || e.st == 0 {
		return errors.New("opus: encoder closed")
	}

	bp := e.tls.Alloc(16)
	defer e.tls.Free(16)
	ret := opusccenc.Opus_opus_encoder_ctl(e.tls, e.st, request, libc.VaList(bp, value))
	if ret != opusccenc.OPUS_OK {
		return fmt.Errorf("%w: %s (%d)", ErrCtlFailed, opusccencErrorString(e.tls, ret), ret)
	}
	return nil
}

// Encode encodes interleaved signed 16-bit PCM into a single Opus packet.
//
// frameSize is the number of samples per channel.
func (e *Encoder) Encode(pcm []int16, frameSize int, packet []byte) (int, error) {
	if e == nil {
		return 0, errors.New("opus: encoder closed")
	}
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.tls == nil || e.st == 0 {
		return 0, errors.New("opus: encoder closed")
	}
	if frameSize <= 0 {
		return 0, errors.New("opus: invalid frameSize")
	}
	nNeeded := frameSize * e.channels
	if len(pcm) < nNeeded {
		return 0, fmt.Errorf("opus: pcm buffer too small: need %d samples, have %d", nNeeded, len(pcm))
	}
	if len(packet) == 0 {
		return 0, errors.New("opus: packet buffer is empty")
	}

	pcmPtr := uintptr(unsafe.Pointer(&pcm[0]))
	outPtr := uintptr(unsafe.Pointer(&packet[0]))

	ret := opusccenc.Opus_opus_encode(e.tls, e.st, pcmPtr, int32(frameSize), outPtr, opusccenc.OpusT_opus_int32(len(packet)))
	if ret < 0 {
		return 0, fmt.Errorf("%w: %s (%d)", ErrEncodeFailed, opusccencErrorString(e.tls, int32(ret)), ret)
	}
	return int(ret), nil
}

func opusccencErrorString(tls *libc.TLS, code int32) string {
	if tls == nil {
		return "(no tls)"
	}
	p := opusccenc.Opus_opus_strerror(tls, code)
	if p == 0 {
		return "(unknown)"
	}
	return libc.GoString(p)
}
