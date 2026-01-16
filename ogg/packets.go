package ogg

import (
	"errors"
	"fmt"
	"io"
	"math"
)

var (
	ErrSerialMismatch   = errors.New("ogg: bitstream serial mismatch")
	ErrBadContinuedPage = errors.New("ogg: page marked continued but no prior packet")
)

// Packet is a reassembled Ogg packet.
type Packet struct {
	Data            []byte
	Serial          uint32
	PageSequenceEnd uint32
	GranulePosition uint64
	GranuleValid    bool
	BOS             bool
	EOS             bool
}

// PacketReader converts Ogg pages into packets by applying lacing rules.
//
// It assumes a single logical bitstream; if multiple serial numbers
// are present it will return ErrSerialMismatch.
type PacketReader struct {
	pr          *PageReader
	serial      *uint32
	pending     []byte
	havePending bool
	queue       []*Packet
}

func NewPacketReader(r io.Reader) *PacketReader {
	return &PacketReader{pr: NewPageReader(r)}
}

func (r *PacketReader) SetVerifyCRC(v bool) { r.pr.VerifyCRC = v }

func (r *PacketReader) ReadPacket() (*Packet, error) {
	if len(r.queue) > 0 {
		pkt := r.queue[0]
		r.queue[0] = nil
		r.queue = r.queue[1:]
		return pkt, nil
	}

	for {
		page, err := r.pr.ReadPage()
		if err != nil {
			return nil, err
		}

		if r.serial == nil {
			s := page.BitstreamSerial
			r.serial = &s
		} else if page.BitstreamSerial != *r.serial {
			return nil, fmt.Errorf("%w: got=%d want=%d", ErrSerialMismatch, page.BitstreamSerial, *r.serial)
		}

		if page.IsContinuedPacket() && !r.havePending {
			return nil, ErrBadContinuedPage
		}
		if !page.IsContinuedPacket() {
			// If the page does not continue a previous packet, any prior pending data
			// is invalid per spec; drop it to recover.
			r.pending = r.pending[:0]
			r.havePending = false
		}

		// Walk lacing values to assemble all completed packets on this page.
		body := page.SegmentData
		off := 0
		completed := make([][]byte, 0, 4)
		for _, segLenU8 := range page.SegmentTable {
			segLen := int(segLenU8)
			if off+segLen > len(body) {
				return nil, fmt.Errorf("ogg: segment data underrun (seq=%d)", page.PageSequence)
			}
			r.pending = append(r.pending, body[off:off+segLen]...)
			r.havePending = true
			off += segLen

			if segLenU8 < 255 {
				completed = append(completed, append([]byte{}, r.pending...))
				r.pending = r.pending[:0]
				r.havePending = false
			}
		}

		if len(completed) == 0 {
			// Page ended with an incomplete packet; keep pending and read next page.
			continue
		}

		granuleValid := page.GranulePosition != math.MaxUint64
		for i := range completed {
			pkt := &Packet{
				Data:            completed[i],
				Serial:          page.BitstreamSerial,
				PageSequenceEnd: page.PageSequence,
				GranulePosition: page.GranulePosition,
				GranuleValid:    granuleValid && i == len(completed)-1,
				BOS:             page.IsBOS() && i == 0,
				EOS:             page.IsEOS() && i == len(completed)-1,
			}
			r.queue = append(r.queue, pkt)
		}

		pkt := r.queue[0]
		r.queue[0] = nil
		r.queue = r.queue[1:]
		return pkt, nil
	}
}
