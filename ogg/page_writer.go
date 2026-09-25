package ogg

import "io"

// DefaultMaxPageGranules is PageWriter's default page span: one second of
// 48 kHz audio, matching opusenc. Page headers stay negligible (~0.2 kbps)
// while seeking, which lands on page boundaries, stays fine-grained.
const DefaultMaxPageGranules = 48000

// PageWriter writes a single Ogg stream, packing consecutive packets onto
// shared pages. PacketWriter puts every packet on a page of its own, which
// for 20 ms Opus packets costs ~11.6 kbps in page headers.
//
// A page is closed when the next packet wouldn't fit in its 255 lacing
// values, when its packets span MaxPageGranules, when a packet is written
// with eos set, or on FlushPage/Flush. Its granule position is that of the
// last packet completed on it.
type PageWriter struct {
	pw *PacketWriter

	// MaxPageGranules bounds how much audio (in granule units) one page
	// spans before it's closed. Zero means DefaultMaxPageGranules.
	MaxPageGranules uint64

	segs      []byte
	body      []byte
	granule   uint64 // granule position of the last packet queued
	pageStart uint64 // granule position the pending page started from
}

func NewPageWriter(w io.Writer, serial uint32) *PageWriter {
	return &PageWriter{pw: NewPacketWriter(w, serial)}
}

// WriteHeaderPacket writes packet alone on a page of its own, as RFC 7845
// section 3 requires for OpusHead and OpusTags, after closing any pending
// page. bos marks the stream's first page.
func (w *PageWriter) WriteHeaderPacket(packet []byte, bos bool) error {
	if err := w.FlushPage(); err != nil {
		return err
	}
	return w.pw.WritePacket(packet, 0, bos, false)
}

// WritePacket queues an audio packet onto the current page. granulePos is
// the stream position at the end of this packet (for Opus: total 48 kHz
// samples decodable through it, pre-skip included; the final packet's is
// end-trimmed to the real length). eos marks the stream's last packet and
// closes its page.
func (w *PageWriter) WritePacket(packet []byte, granulePos uint64, eos bool) error {
	lace, _ := oggLacing(packet)
	if len(lace) > 255 {
		// Too big to share a page: PacketWriter continues it across
		// pages of its own.
		if err := w.FlushPage(); err != nil {
			return err
		}
		w.granule, w.pageStart = granulePos, granulePos
		return w.pw.WritePacket(packet, granulePos, false, eos)
	}
	if len(w.segs)+len(lace) > 255 {
		if err := w.FlushPage(); err != nil {
			return err
		}
	}
	if len(w.segs) == 0 {
		w.pageStart = w.granule
	}
	w.segs = append(w.segs, lace...)
	w.body = append(w.body, packet...)
	w.granule = granulePos

	maxSpan := w.MaxPageGranules
	if maxSpan == 0 {
		maxSpan = DefaultMaxPageGranules
	}
	switch {
	case eos:
		return w.flushPage(0x04)
	case granulePos-w.pageStart >= maxSpan:
		return w.FlushPage()
	}
	return nil
}

// FlushPage closes the pending page, if any, without flushing the
// underlying writer.
func (w *PageWriter) FlushPage() error {
	return w.flushPage(0)
}

func (w *PageWriter) flushPage(headerType uint8) error {
	if len(w.segs) == 0 {
		return nil
	}
	if err := w.pw.writePage(headerType, w.granule, w.segs, w.body); err != nil {
		return err
	}
	w.pw.seq++
	w.segs = w.segs[:0]
	w.body = w.body[:0]
	return nil
}

// Flush closes the pending page and flushes the underlying writer.
func (w *PageWriter) Flush() error {
	if err := w.FlushPage(); err != nil {
		return err
	}
	return w.pw.Flush()
}
