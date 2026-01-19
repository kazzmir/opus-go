package opus

import (
	"errors"
	"fmt"
	"sync"

	libc "github.com/kazzmir/opus-go/libcshim"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opuscc"
)

var (
	ErrUnsupportedMapping = errors.New("opus: unsupported channel mapping")
	ErrBadPacket          = errors.New("opus: decode failed")
)

// Decoder is a pure-Go Opus decoder backed by ccgo-transpiled libopus.
//
// It supports both mapping family 0 (mono/stereo) and multistream mappings.
type Decoder struct {
	mu sync.Mutex

	tls *libc.TLS
	st  uintptr

	sampleRate int
	channels   int

	multistream bool
}

func NewDecoderFromHead(head ogg.OpusHead) (*Decoder, error) {
	// RFC 7845: Opus is decoded at 48 kHz.
	const fs = ogg.OpusSampleRateHz

	if head.ChannelMappingFamily == 0 {
		if head.Channels != 1 && head.Channels != 2 {
			return nil, fmt.Errorf("%w: mapping family 0 requires 1 or 2 channels, got %d", ErrUnsupportedMapping, head.Channels)
		}
		return NewDecoder(fs, int(head.Channels))
	}

	// Mapping family != 0 uses multistream.
	if head.StreamCount == 0 {
		return nil, fmt.Errorf("%w: missing stream count", ErrUnsupportedMapping)
	}
	if int(head.Channels) != len(head.ChannelMapping) {
		return nil, fmt.Errorf("%w: channel mapping length mismatch", ErrUnsupportedMapping)
	}

	return NewMultistreamDecoder(fs, int(head.Channels), int(head.StreamCount), int(head.CoupledStreamCount), head.ChannelMapping)
}

func NewDecoder(sampleRate, channels int) (*Decoder, error) {
	tls := libc.NewTLS()
	if tls == nil {
		return nil, errors.New("opus: failed to allocate TLS")
	}

	st, err := opuscc.Opus_opus_decoder_create(tls, opuscc.OpusT_opus_int32(sampleRate), int32(channels))
	if err != nil || st == 0 {
		if oe := (*opuscc.OpusError)(nil); errors.As(err, &oe) {
			msg := opusccErrorString(tls, oe.Code)
			tls.Close()
			return nil, fmt.Errorf("opus: decoder_create failed: %s (%d)", msg, oe.Code)
		}
		tls.Close()
		return nil, fmt.Errorf("opus: decoder_create failed: %w", err)
	}

	return &Decoder{tls: tls, st: st, sampleRate: sampleRate, channels: channels}, nil
}

func NewMultistreamDecoder(sampleRate, channels, streams, coupledStreams int, mapping []uint8) (*Decoder, error) {
	tls := libc.NewTLS()
	if tls == nil {
		return nil, errors.New("opus: failed to allocate TLS")
	}

	mappingPtr := libc.PtrUint8(mapping)

	st, err := opuscc.Opus_opus_multistream_decoder_create(
		tls,
		opuscc.OpusT_opus_int32(sampleRate),
		int32(channels),
		int32(streams),
		int32(coupledStreams),
		mappingPtr,
	)
	if err != nil || st == 0 {
		if oe := (*opuscc.OpusError)(nil); errors.As(err, &oe) {
			msg := opusccErrorString(tls, oe.Code)
			tls.Close()
			return nil, fmt.Errorf("opus: multistream_decoder_create failed: %s (%d)", msg, oe.Code)
		}
		tls.Close()
		return nil, fmt.Errorf("opus: multistream_decoder_create failed: %w", err)
	}

	return &Decoder{tls: tls, st: st, sampleRate: sampleRate, channels: channels, multistream: true}, nil
}

func (d *Decoder) Close() error {
	if d == nil {
		return nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.tls != nil {
		if d.st != 0 {
			if d.multistream {
				opuscc.Opus_opus_multistream_decoder_destroy(d.tls, d.st)
			} else {
				opuscc.Opus_opus_decoder_destroy(d.tls, d.st)
			}
			d.st = 0
		}
		opuscc.FreePseudostackTLS(d.tls)
		d.tls.Close()
		d.tls = nil
	}
	return nil
}

func (d *Decoder) SampleRate() int { return d.sampleRate }
func (d *Decoder) Channels() int   { return d.channels }

// Decode decodes a single Opus packet into interleaved signed 16-bit PCM.
//
// frameSize is the maximum number of samples per channel to decode.
// Use 5760 for the Opus max frame size at 48kHz (120 ms).
//
// Returns the number of samples per channel written into pcm.
func (d *Decoder) Decode(packet []byte, pcm []int16, frameSize int, decodeFEC bool) (int, error) {
	if d == nil {
		return 0, errors.New("opus: decoder closed")
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	if d == nil || d.tls == nil || d.st == 0 {
		return 0, errors.New("opus: decoder closed")
	}
	if frameSize <= 0 {
		return 0, errors.New("opus: invalid frameSize")
	}
	nNeeded := frameSize * d.channels
	if len(pcm) < nNeeded {
		return 0, fmt.Errorf("opus: pcm buffer too small: need %d samples, have %d", nNeeded, len(pcm))
	}

	dataPtr := libc.PtrByte(packet)
	dataLen := int32(len(packet))
	pcmPtr := libc.PtrInt16(pcm)
	fec := int32(0)
	if decodeFEC {
		fec = 1
	}

	var ret int32
	if d.multistream {
		ret = opuscc.Opus_opus_multistream_decode(d.tls, d.st, dataPtr, opuscc.OpusT_opus_int32(dataLen), pcmPtr, int32(frameSize), fec)
	} else {
		ret = opuscc.Opus_opus_decode(d.tls, d.st, dataPtr, opuscc.OpusT_opus_int32(dataLen), pcmPtr, int32(frameSize), fec)
	}

	if ret < 0 {
		return 0, fmt.Errorf("%w: %s (%d)", ErrBadPacket, opusccErrorString(d.tls, ret), ret)
	}
	return int(ret), nil
}

// convenience function to decode an Ogg OpusAudioPacket
func (decoder *Decoder) DecodePacket(packet *ogg.OpusAudioPacket, pcm []int16) ([]int16, int, error) {
    const maxMsPerFrame = 120
    maxSize := ogg.OpusSampleRateHz * maxMsPerFrame / 1000
    if len(pcm) < maxSize * decoder.channels {
        pcm = make([]int16, maxSize * decoder.channels)
    }

    n, err := decoder.Decode(packet.Data, pcm, maxSize, false)
    if err != nil {
        return nil, 0, err
    }

    return pcm[:n*decoder.channels], n, nil
}

func opusccErrorString(tls *libc.TLS, code int32) string {
	if tls == nil {
		return "(no tls)"
	}
	p := opuscc.Opus_opus_strerror(code)
	if p == 0 {
		return "(unknown)"
	}
	return libc.GoString(p)
}
