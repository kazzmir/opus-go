package wav

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	ErrNotWAV         = errors.New("wav: not a RIFF/WAVE file")
	ErrUnsupportedWAV = errors.New("wav: unsupported WAV format")
)

// Reader reads 16-bit PCM little-endian WAV data.
//
// It supports a minimal subset sufficient for decoding audio for Opus encoding.
// Currently supported: PCM (format 1), 16-bit, any channel count.
//
// Data is returned as interleaved signed 16-bit PCM.
type Reader struct {
	br *bufio.Reader

	sampleRate int
	channels   int

	dataRemaining uint32
}

func NewReader(r io.Reader) (*Reader, error) {
	wr := &Reader{br: bufio.NewReaderSize(r, 1<<20)}
	if err := wr.readHeader(); err != nil {
		return nil, err
	}
	return wr, nil
}

func (r *Reader) SampleRate() int { return r.sampleRate }
func (r *Reader) Channels() int   { return r.channels }

// ReadInt16PCM reads up to len(dst) samples (not frames) into dst.
// Returns number of samples read.
func (r *Reader) ReadInt16PCM(dst []int16) (int, error) {
	if len(dst) == 0 {
		return 0, nil
	}
	if r.dataRemaining == 0 {
		return 0, io.EOF
	}

	maxSamples := int(r.dataRemaining / 2)
	if maxSamples <= 0 {
		r.dataRemaining = 0
		return 0, io.EOF
	}
	n := len(dst)
	if n > maxSamples {
		n = maxSamples
	}

	buf := make([]byte, n*2)
	if _, err := io.ReadFull(r.br, buf); err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		dst[i] = int16(binary.LittleEndian.Uint16(buf[i*2:]))
	}
	r.dataRemaining -= uint32(n * 2)
	return n, nil
}

func (r *Reader) readHeader() error {
	var riff [12]byte
	if _, err := io.ReadFull(r.br, riff[:]); err != nil {
		return err
	}
	if string(riff[0:4]) != "RIFF" || string(riff[8:12]) != "WAVE" {
		return ErrNotWAV
	}

	var (
		haveFmt  bool
		haveData bool
	)

	for {
		var chdr [8]byte
		if _, err := io.ReadFull(r.br, chdr[:]); err != nil {
			return err
		}
		id := string(chdr[0:4])
		sz := binary.LittleEndian.Uint32(chdr[4:8])

		switch id {
		case "fmt ":
			if sz < 16 {
				return fmt.Errorf("%w: fmt chunk too short", ErrUnsupportedWAV)
			}
			buf := make([]byte, sz)
			if _, err := io.ReadFull(r.br, buf); err != nil {
				return err
			}
			audioFormat := binary.LittleEndian.Uint16(buf[0:2])
			channels := binary.LittleEndian.Uint16(buf[2:4])
			sampleRate := binary.LittleEndian.Uint32(buf[4:8])
			bitsPerSample := binary.LittleEndian.Uint16(buf[14:16])

			if audioFormat != 1 {
				return fmt.Errorf("%w: audio format=%d", ErrUnsupportedWAV, audioFormat)
			}
			if bitsPerSample != 16 {
				return fmt.Errorf("%w: bits per sample=%d", ErrUnsupportedWAV, bitsPerSample)
			}
			if channels == 0 {
				return fmt.Errorf("%w: channels=0", ErrUnsupportedWAV)
			}
			r.sampleRate = int(sampleRate)
			r.channels = int(channels)
			haveFmt = true

		case "data":
			if !haveFmt {
				return fmt.Errorf("%w: data before fmt", ErrUnsupportedWAV)
			}
			r.dataRemaining = sz
			haveData = true
			return nil

		default:
			// Skip unknown chunk (plus padding byte if needed).
			if _, err := io.CopyN(io.Discard, r.br, int64(sz)); err != nil {
				return err
			}
			if sz%2 == 1 {
				if _, err := r.br.ReadByte(); err != nil {
					return err
				}
			}
		}

		if haveData {
			return nil
		}
	}
}
