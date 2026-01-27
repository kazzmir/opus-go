package ogg

import (
	"errors"
	"fmt"
	"io"
	"math"
    "bufio"
    "bytes"
    "slices"
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
    reader      io.Reader // the original reader for seeking
	serial      *uint32
	pending     []byte
	havePending bool
	queue       []*Packet
}

func NewPacketReader(r io.Reader) *PacketReader {
	return &PacketReader{
        pr: NewPageReader(r),
        reader: r,
    }
}

func (r *PacketReader) reset() {
    r.pending = nil
    r.queue = nil
    r.havePending = false
}

// keep reading bytes until we find the start of an ogg page
// return the byte position of the start of the page
func (r *PacketReader) findNextPage(seeker io.ReadSeeker, position int64, seekBuffer *bytes.Buffer) (int64, error) {

    seekBuffer.Reset()

    // these bytes are the start of an ogg page
    scanBytes := []byte{'O', 'g', 'g', 'S', 0}

    _, err := seeker.Seek(position, io.SeekStart)
    if err != nil {
        return 0, err
    }

    buffered := bufio.NewReader(seeker)

    // rounds := 0
    for seekBuffer.Len() < 65536 {
        // rounds += 1

        reader := io.LimitReader(buffered, 4096)

        n, err := seekBuffer.ReadFrom(reader)
        if err != nil {
            return 0, err
        }
        if n == 0 {
            return -1, io.EOF
        }

        index := bytes.Index(seekBuffer.Bytes(), scanBytes)

        if index != -1 {
            // fmt.Printf("found page at position %d after %d rounds\n", position+int64(index), rounds)
            return position + int64(index), nil
        }
    }

    return -1, io.EOF
}

// seek to the first page that contains the granule position
// note that the granule position on the page is the position of the last complete packet on that page
// For example, if granulePos == 40000 then we find the page that has the closest granule position below
// 40000, which may be GranulePosition=38000. Then we keep reading packets until we see the last valid packet
// with GranulePosition=38000. The packet we care about must be the next one
func (r *PacketReader) SeekToPage(granulePos uint64) (uint64, error) {
    seeker, ok := r.reader.(io.ReadSeeker)
    if !ok {
        return 0, fmt.Errorf("ogg: underlying reader is not seekable")
    }

    // seek to beginning
    if granulePos == 0 {
        r.reset()
        _, err := seeker.Seek(0, io.SeekStart)
        if err != nil {
            return 0, err
        }
        r.pr = NewPageReader(seeker)
        // skip headers
        for range 2 {
            _, err := r.ReadPacket()
            if err != nil {
                return 0, err
            }
        }

        return 0, nil
    }

    total, err := seeker.Seek(0, io.SeekEnd)
    if err != nil {
        return 0, err
    }

    start, err := seeker.Seek(0, io.SeekStart)
    if err != nil {
        return 0, err
    }

    // map of granule positions to file positions that starts the page where the granule is
    // TODO: possibly consider persisting this so that the granule positions can be reused
    granulePositions := make(map[uint64]int64)

    highestGranule := uint64(0)
    highestPage := uint32(0)
    lowestGranule := uint64(0)

    var seekBuffer bytes.Buffer

    pages := 0

    for start < total {
        position := (total + start) / 2

        pagePosition, err := r.findNextPage(seeker, position, &seekBuffer)

        if err != nil {
            if errors.Is(err, io.EOF) {
                total = position
                continue
            }

            return 0, err
        }

        if pagePosition == -1 {
            // if we don't find a page then back up
            total = position
        } else {
            // found a page at index, seek there and read the page
            // if the granule position is less than the target, move start up
            // if the granule position is greater than the target then this might be the page we want
            // but we still have to check the previous page

            seeker.Seek(pagePosition, io.SeekStart)
            r.reset()
            r.pr = NewPageReader(seeker)
            page, err := r.ReadPacket()
            if err != nil {
                // this wasn't a page after all?
                total = position
                continue
            }

            pages += 1

            last, ok := granulePositions[page.GranulePosition]
            if ok && last == pagePosition {
                // we already scanned this page. total must be too close to start
                // force scanning at start
                total = start + 1

                // tie break
                if position == start {
                    break
                }

                continue
            }

            granulePositions[page.GranulePosition] = pagePosition

            // the highest granule position that is above the target but closest to the target
            if page.GranulePosition > granulePos && (highestGranule == 0 || page.GranulePosition <= highestGranule) {
                highestGranule = page.GranulePosition
                highestPage = page.PageSequenceEnd
                // search down for the next highest page
                total = position

                // this is the earliest page that can have data on it
                if highestPage == 2 {
                    break
                }
            }

            // this page is too low
            if page.GranulePosition < granulePos {
                if page.GranulePosition > lowestGranule {
                    lowestGranule = page.GranulePosition
                }

                // this is one byte after the start of the page we just found, but we
                // have no way to know how many bytes this page takes up, so just jump ahead a bit
                // there must be another page that comes after this one
                start = pagePosition + 1
            }
        }
    }

    if highestGranule == 0 {
        // never found the page, or the target granule position is beyond the end of the stream
        return 0, fmt.Errorf("ogg: could not find page with granule position %d", granulePos)
    }

    position, ok := granulePositions[lowestGranule]
    if !ok {
        // just start at the beginning
        position = 0
    }

    seeker.Seek(position, io.SeekStart)
    r.reset()
    r.pr = NewPageReader(seeker)

    // fmt.Printf("start sequential scan at position %d pages %d\n", position, pages)
    last := uint64(0)
    for {
        page, err := r.ReadPacket()
        if err != nil {
            return 0, err
        }
        pages += 1
        if page.GranulePosition >= granulePos {
            // put back in the queue
            r.queue = slices.Insert(r.queue, 0, page)
            // fmt.Printf("ogg: SeekToPage scanned %d pages to find granule position %d\n", pages, granulePos)
            return last, nil
        } else {
            last = page.GranulePosition
        }
    }
}

// return the last granule position in the stream
// this is a destructive operation that will position the stream at the end
// so no more packets can be read after this, unless the stream is seek'd somewhere
func (r *PacketReader) LastPageGranule() (int64, error) {
    seeker, ok := r.reader.(io.ReadSeeker)
    if ok {
        total, err := seeker.Seek(0, io.SeekEnd)
        if err != nil {
            return 0, err
        }

        // FIXME: we could probably read less than 64k, but its ok for now
        backup := max(0, total-65536)
        position, err := seeker.Seek(backup, io.SeekStart)
        if err != nil {
            return 0, err
        }

        index, err := r.findNextPage(seeker, position, &bytes.Buffer{})
        if err != nil {
            return 0, err
        }
        if index == -1 {
            return 0, fmt.Errorf("ogg: could not find last page granule")
        }
        seeker.Seek(index, io.SeekStart)

        r.pr = NewPageReader(seeker)
        r.reset()
    }

    for {
        page, err := r.ReadPacket()
        if err == nil && page.EOS && page.GranuleValid {
            return int64(page.GranulePosition), nil
        }
        if err != nil {
            return 0, fmt.Errorf("ogg: could not find last page granule: %v", err)
        }
    }
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
		pagePackets := make([]*Packet, 0, 4)
		pktIndex := 0
		for _, segLenU8 := range page.SegmentTable {
			segLen := int(segLenU8)
			if off+segLen > len(body) {
				return nil, fmt.Errorf("ogg: segment data underrun (seq=%d)", page.PageSequence)
			}
			r.pending = append(r.pending, body[off:off+segLen]...)
			r.havePending = true
			off += segLen

			if segLenU8 < 255 {
				data := r.pending
				// Detach the backing array so future appends won't overwrite returned data.
				r.pending = nil
				r.havePending = false
				pagePackets = append(pagePackets, &Packet{
					Data:            data,
					Serial:          page.BitstreamSerial,
					PageSequenceEnd: page.PageSequence,
					GranulePosition: page.GranulePosition,
					GranuleValid:    false, // set below for last packet only
					BOS:             page.IsBOS() && pktIndex == 0,
					EOS:             false, // set below for last packet only
				})
				pktIndex++
			}
		}

		if len(pagePackets) == 0 {
			// Page ended with an incomplete packet; keep pending and read next page.
			continue
		}

		granuleValid := page.GranulePosition != math.MaxUint64
		last := pagePackets[len(pagePackets)-1]
		last.GranuleValid = granuleValid
		last.EOS = page.IsEOS()
		r.queue = append(r.queue, pagePackets...)

		pkt := r.queue[0]
		r.queue[0] = nil
		r.queue = r.queue[1:]
		return pkt, nil
	}
}
