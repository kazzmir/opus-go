package wav

import (
	"encoding/binary"
	"io"
)

// Writer writes 16-bit PCM WAV files.
//
// It requires an io.WriteSeeker so it can patch sizes on Close.
type Writer struct {
	w          io.WriteSeeker
	sampleRate uint32
	channels   uint16

	dataBytes uint32
	closed    bool
}

func NewWriter(w io.WriteSeeker, sampleRate int, channels int) (*Writer, error) {
	wr := &Writer{w: w, sampleRate: uint32(sampleRate), channels: uint16(channels)}
	if err := wr.writeHeaderPlaceholder(); err != nil {
		return nil, err
	}
	return wr, nil
}

func (wr *Writer) WriteInt16PCM(pcm []int16) error {
	if wr.closed {
		return io.ErrClosedPipe
	}
	var buf [2]byte
	for _, s := range pcm {
		binary.LittleEndian.PutUint16(buf[:], uint16(s))
		if _, err := wr.w.Write(buf[:]); err != nil {
			return err
		}
		wr.dataBytes += 2
	}
	return nil
}

func (wr *Writer) Close() error {
	if wr.closed {
		return nil
	}
	wr.closed = true

	// Patch RIFF chunk size and data chunk size.
	// RIFF size = 4 (WAVE) + (8+fmt) + (8+data)
	// We wrote: 12 + (8+16) + 8 + data
	riffSize := uint32(4 + (8 + 16) + (8 + wr.dataBytes))

	if _, err := wr.w.Seek(4, io.SeekStart); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, riffSize); err != nil {
		return err
	}
	if _, err := wr.w.Seek(40, io.SeekStart); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, wr.dataBytes); err != nil {
		return err
	}

	_, err := wr.w.Seek(0, io.SeekEnd)
	return err
}

func (wr *Writer) writeHeaderPlaceholder() error {
	// Minimal PCM WAV header.
	// RIFF header
	if _, err := wr.w.Write([]byte("RIFF")); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, uint32(0)); err != nil { // placeholder
		return err
	}
	if _, err := wr.w.Write([]byte("WAVE")); err != nil {
		return err
	}

	// fmt chunk
	if _, err := wr.w.Write([]byte("fmt ")); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, uint32(16)); err != nil { // PCM
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, uint16(1)); err != nil { // audio format PCM
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, wr.channels); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, wr.sampleRate); err != nil {
		return err
	}
	byteRate := wr.sampleRate * uint32(wr.channels) * 2
	blockAlign := wr.channels * 2
	if err := binary.Write(wr.w, binary.LittleEndian, byteRate); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, blockAlign); err != nil {
		return err
	}
	if err := binary.Write(wr.w, binary.LittleEndian, uint16(16)); err != nil { // bits per sample
		return err
	}

	// data chunk
	if _, err := wr.w.Write([]byte("data")); err != nil {
		return err
	}
	return binary.Write(wr.w, binary.LittleEndian, uint32(0)) // placeholder
}
