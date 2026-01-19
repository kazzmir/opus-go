package wav

import (
	"bufio"
	"encoding/binary"
	"io"
)

// Writer writes 16-bit PCM WAV files.
//
// It requires an io.WriteSeeker so it can patch sizes on Close.
type Writer struct {
	w          io.WriteSeeker
	bw         *bufio.Writer
	sampleRate uint32
	channels   uint16

	buf []byte

	dataBytes uint32
	closed    bool
}

func NewWriter(w io.WriteSeeker, sampleRate int, channels int) (*Writer, error) {
	wr := &Writer{
		w:          w,
		bw:         bufio.NewWriterSize(w, 1<<20),
		sampleRate: uint32(sampleRate),
		channels:   uint16(channels),
	}
	if err := wr.writeHeaderPlaceholder(); err != nil {
		return nil, err
	}
	// Ensure header is on disk before we later Seek() to patch sizes.
	if err := wr.bw.Flush(); err != nil {
		return nil, err
	}
	return wr, nil
}

func (wr *Writer) WriteInt16PCM(pcm []int16) error {
	if wr.closed {
		return io.ErrClosedPipe
	}
	if len(pcm) == 0 {
		return nil
	}

	n := len(pcm) * 2
	if cap(wr.buf) < n {
		wr.buf = make([]byte, n)
	}
	buf := wr.buf[:n]
	for i, s := range pcm {
		binary.LittleEndian.PutUint16(buf[i*2:], uint16(s))
	}
	if _, err := wr.bw.Write(buf); err != nil {
		return err
	}
	wr.dataBytes += uint32(n)
	return nil
}

func (wr *Writer) Close() error {
	if wr.closed {
		return nil
	}
	wr.closed = true
	if wr.bw != nil {
		if err := wr.bw.Flush(); err != nil {
			return err
		}
	}

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
	if _, err := wr.bw.Write([]byte("RIFF")); err != nil {
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, uint32(0)); err != nil { // placeholder
		return err
	}
	if _, err := wr.bw.Write([]byte("WAVE")); err != nil {
		return err
	}

	// fmt chunk
	if _, err := wr.bw.Write([]byte("fmt ")); err != nil {
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, uint32(16)); err != nil { // PCM
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, uint16(1)); err != nil { // audio format PCM
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, wr.channels); err != nil {
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, wr.sampleRate); err != nil {
		return err
	}
	byteRate := wr.sampleRate * uint32(wr.channels) * 2
	blockAlign := wr.channels * 2
	if err := binary.Write(wr.bw, binary.LittleEndian, byteRate); err != nil {
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, blockAlign); err != nil {
		return err
	}
	if err := binary.Write(wr.bw, binary.LittleEndian, uint16(16)); err != nil { // bits per sample
		return err
	}

	// data chunk
	if _, err := wr.bw.Write([]byte("data")); err != nil {
		return err
	}
	return binary.Write(wr.bw, binary.LittleEndian, uint32(0)) // placeholder
}
