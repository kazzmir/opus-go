package ogg

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

// TestPageWriterRoundTrip writes two header packets and 3 s of 20 ms
// audio packets (plus one packet too big to share a page), then reads
// the stream back page by page and packet by packet.
func TestPageWriterRoundTrip(t *testing.T) {
	var buf bytes.Buffer
	w := NewPageWriter(&buf, 7)
	if err := w.WriteHeaderPacket([]byte("OpusHead-ish"), true); err != nil {
		t.Fatal(err)
	}
	if err := w.WriteHeaderPacket([]byte("OpusTags-ish"), false); err != nil {
		t.Fatal(err)
	}
	var want [][]byte
	var granules []uint64
	const packets = 150
	for i := range packets {
		size := 100 + i%7
		if i == 100 {
			size = 70000 // needs more than 255 lacing values: continued across pages
		}
		p := bytes.Repeat([]byte{byte(i)}, size)
		g := uint64(960 * (i + 1))
		if err := w.WritePacket(p, g, i == packets-1); err != nil {
			t.Fatal(err)
		}
		want = append(want, p)
		granules = append(granules, g)
	}
	if err := w.Flush(); err != nil {
		t.Fatal(err)
	}

	// Pages: 2 header pages, then ~1 s of audio per page, plus the big
	// packet's own pages - far fewer than one per packet.
	pr := NewPageReader(bytes.NewReader(buf.Bytes()))
	pr.VerifyCRC = true
	var pages []*Page
	for {
		p, err := pr.ReadPage()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if !p.CRCVerified {
			t.Fatalf("page %d: CRC not verified", len(pages))
		}
		if p.PageSequence != uint32(len(pages)) {
			t.Fatalf("page %d has sequence number %d", len(pages), p.PageSequence)
		}
		pages = append(pages, p)
	}
	if len(pages) > 10 {
		t.Errorf("%d pages for %d packets, want about one per second of audio", len(pages), packets)
	}
	if !pages[0].IsBOS() || pages[1].IsBOS() {
		t.Error("BOS must be set on the first page only")
	}
	for i, p := range pages {
		if p.IsEOS() != (i == len(pages)-1) {
			t.Errorf("page %d: EOS = %v", i, p.IsEOS())
		}
	}
	if last := pages[len(pages)-1]; last.GranulePosition != granules[packets-1] {
		t.Errorf("last page granule %d, want %d", last.GranulePosition, granules[packets-1])
	}

	r := NewPacketReader(bytes.NewReader(buf.Bytes()))
	for i := range 2 {
		if _, err := r.ReadPacket(); err != nil {
			t.Fatalf("header packet %d: %v", i, err)
		}
	}
	for i, w := range want {
		p, err := r.ReadPacket()
		if err != nil {
			t.Fatalf("packet %d: %v", i, err)
		}
		if !bytes.Equal(p.Data, w) {
			t.Fatalf("packet %d: %d bytes back, want %d", i, len(p.Data), len(w))
		}
		if p.GranuleValid && p.GranulePosition != granules[i] {
			t.Fatalf("packet %d: granule %d, want %d", i, p.GranulePosition, granules[i])
		}
	}
	if _, err := r.ReadPacket(); !errors.Is(err, io.EOF) {
		t.Fatalf("expected EOF after the last packet, got %v", err)
	}
}
