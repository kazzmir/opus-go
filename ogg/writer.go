package ogg

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
)

// PacketWriter writes Ogg packets as pages using the lacing scheme (RFC 3533).
//
// This is a minimal single-stream writer.
type PacketWriter struct {
	bw *bufio.Writer

	serial uint32
	seq    uint32

	crcTable [256]uint32
}

func NewPacketWriter(w io.Writer, serial uint32) *PacketWriter {
	var tbl [256]uint32
	const poly uint32 = 0x04C11DB7
	for i := 0; i < 256; i++ {
		c := uint32(i) << 24
		for j := 0; j < 8; j++ {
			if c&0x80000000 != 0 {
				c = (c << 1) ^ poly
			} else {
				c <<= 1
			}
		}
		tbl[i] = c
	}

	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriterSize(w, 1<<20)
	}

	return &PacketWriter{
		bw:       bw,
		serial:   serial,
		seq:      0,
		crcTable: tbl,
	}
}

func (pw *PacketWriter) Flush() error {
	if pw == nil || pw.bw == nil {
		return nil
	}
	return pw.bw.Flush()
}

// WritePacket writes a single logical Ogg packet.
//
// If the packet is too large for one page, it will be continued across pages.
// For continued packets, only the final page carries the provided granulePos;
// earlier pages use granulePos = -1 (0xFFFFFFFFFFFFFFFF).
func (pw *PacketWriter) WritePacket(packet []byte, granulePos uint64, bos bool, eos bool) error {
	lace, laceDataLens := oggLacing(packet)
	dataOff := 0

	pageIndex := 0
	for len(lace) > 0 {
		nSeg := len(lace)
		if nSeg > 255 {
			nSeg = 255
		}
		segTable := lace[:nSeg]
		dataLen := 0
		for i := 0; i < nSeg; i++ {
			dataLen += laceDataLens[i]
		}

		// Determine header type.
		headerType := uint8(0)
		if pageIndex > 0 {
			headerType |= 0x01 // continued packet
		}
		if bos && pageIndex == 0 {
			headerType |= 0x02
		}
		isLastPage := nSeg == len(lace)
		if eos && isLastPage {
			headerType |= 0x04
		}

		gp := granulePos
		if !isLastPage {
			gp = math.MaxUint64
		}

		pageData := packet[dataOff : dataOff+dataLen]
		if err := pw.writePage(headerType, gp, segTable, pageData); err != nil {
			return err
		}

		pw.seq++
		pageIndex++
		dataOff += dataLen
		lace = lace[nSeg:]
		laceDataLens = laceDataLens[nSeg:]
	}

	return nil
}

func oggLacing(packet []byte) (segTable []byte, segDataLens []int) {
	// Generate lacing values (255-byte segments). Packet ends when a segment < 255.
	remaining := len(packet)
	for remaining >= 255 {
		segTable = append(segTable, 255)
		segDataLens = append(segDataLens, 255)
		remaining -= 255
	}
	segTable = append(segTable, byte(remaining))
	segDataLens = append(segDataLens, remaining)
	return segTable, segDataLens
}

func (pw *PacketWriter) writePage(headerType uint8, granulePos uint64, segTable []byte, segData []byte) error {
	var header [27]byte
	copy(header[0:4], []byte("OggS"))
	header[4] = 0
	header[5] = headerType
	binary.LittleEndian.PutUint64(header[6:14], granulePos)
	binary.LittleEndian.PutUint32(header[14:18], pw.serial)
	binary.LittleEndian.PutUint32(header[18:22], pw.seq)
	binary.LittleEndian.PutUint32(header[22:26], 0) // checksum placeholder
	header[26] = byte(len(segTable))

	crc := oggCRC3(header[:], segTable, segData, pw.crcTable)
	binary.LittleEndian.PutUint32(header[22:26], crc)

	if _, err := pw.bw.Write(header[:]); err != nil {
		return err
	}
	if _, err := pw.bw.Write(segTable); err != nil {
		return err
	}
	if len(segData) > 0 {
		if _, err := pw.bw.Write(segData); err != nil {
			return err
		}
	}
	return nil
}
