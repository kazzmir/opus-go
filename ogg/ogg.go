package ogg

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidCapturePattern = errors.New("ogg: invalid capture pattern")
	ErrUnsupportedVersion    = errors.New("ogg: unsupported version")
)

// Page represents a single Ogg page (RFC 3533).
//
// Note: CRC is verified by ReadPage; if you want to skip verification,
// use PageReader with VerifyCRC=false.
type Page struct {
	Version          uint8
	HeaderType       uint8
	GranulePosition  uint64
	BitstreamSerial  uint32
	PageSequence     uint32
	Checksum         uint32
	SegmentTable     []uint8
	SegmentData      []byte
	HeaderBytes      []byte // header bytes as read (including segment table)
	BodyBytes        []byte // body bytes as read
	CRCVerified      bool
	RawPageBytesSize int
}

func (p *Page) IsContinuedPacket() bool { return p.HeaderType&0x01 != 0 }
func (p *Page) IsBOS() bool             { return p.HeaderType&0x02 != 0 }
func (p *Page) IsEOS() bool             { return p.HeaderType&0x04 != 0 }

// PageReader reads Ogg pages from an io.Reader.
type PageReader struct {
	r         *bufio.Reader
	VerifyCRC bool
	crcTable  [256]uint32
}

func NewPageReader(r io.Reader) *PageReader {
	var tbl [256]uint32
	// Ogg uses the same CRC as Vorbis: polynomial 0x04C11DB7, MSB-first.
	// This is not the reflected IEEE CRC32 used by Go's hash/crc32.
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

	return &PageReader{
		r:         bufio.NewReaderSize(r, 64*1024),
		VerifyCRC: true,
		crcTable:  tbl,
	}
}

// ReadPage reads the next Ogg page.
func (pr *PageReader) ReadPage() (*Page, error) {
	// Fixed header is 27 bytes.
	header := make([]byte, 27)
	if _, err := io.ReadFull(pr.r, header); err != nil {
		return nil, err
	}
	if string(header[0:4]) != "OggS" {
		return nil, ErrInvalidCapturePattern
	}
	version := header[4]
	if version != 0 {
		return nil, fmt.Errorf("%w: %d", ErrUnsupportedVersion, version)
	}
	headerType := header[5]
	granule := binary.LittleEndian.Uint64(header[6:14])
	serial := binary.LittleEndian.Uint32(header[14:18])
	seq := binary.LittleEndian.Uint32(header[18:22])
	checksum := binary.LittleEndian.Uint32(header[22:26])
	pageSegments := int(header[26])

	segTable := make([]byte, pageSegments)
	if _, err := io.ReadFull(pr.r, segTable); err != nil {
		return nil, err
	}
	bodyLen := 0
	for _, v := range segTable {
		bodyLen += int(v)
	}
	body := make([]byte, bodyLen)
	if _, err := io.ReadFull(pr.r, body); err != nil {
		return nil, err
	}

	fullHeader := append(append([]byte{}, header...), segTable...)
	p := &Page{
		Version:          version,
		HeaderType:       headerType,
		GranulePosition:  granule,
		BitstreamSerial:  serial,
		PageSequence:     seq,
		Checksum:         checksum,
		SegmentTable:     bytesToU8(segTable),
		SegmentData:      body,
		HeaderBytes:      fullHeader,
		BodyBytes:        body,
		RawPageBytesSize: len(fullHeader) + len(body),
	}

	if pr.VerifyCRC {
		crcOk, err := pr.verifyCRC(fullHeader, body, checksum)
		if err != nil {
			return nil, err
		}
		p.CRCVerified = crcOk
		if !crcOk {
			return nil, fmt.Errorf("ogg: crc mismatch (serial=%d seq=%d)", serial, seq)
		}
	}

	return p, nil
}

func (pr *PageReader) verifyCRC(fullHeader []byte, body []byte, expected uint32) (bool, error) {
	// CRC is computed over the entire page with the checksum field set to 0.
	buf := make([]byte, 0, len(fullHeader)+len(body))
	buf = append(buf, fullHeader...)
	buf = append(buf, body...)
	if len(buf) < 27 {
		return false, errors.New("ogg: internal crc buffer too small")
	}
	buf[22] = 0
	buf[23] = 0
	buf[24] = 0
	buf[25] = 0
	got := oggCRC(buf, pr.crcTable)
	return got == expected, nil
}

func oggCRC(page []byte, table [256]uint32) uint32 {
	var crc uint32
	for _, b := range page {
		crc = (crc << 8) ^ table[byte(crc>>24)^b]
	}
	return crc
}

func bytesToU8(b []byte) []uint8 {
	u := make([]uint8, len(b))
	for i := range b {
		u[i] = b[i]
	}
	return u
}
