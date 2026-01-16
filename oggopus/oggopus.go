package oggopus

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"opusgo/ogg"
)

var (
	ErrNotOpusOgg     = errors.New("oggopus: not an Ogg Opus stream")
	ErrBadOpusHead    = errors.New("oggopus: invalid OpusHead")
	ErrBadOpusTags    = errors.New("oggopus: invalid OpusTags")
	ErrUnexpectedBOS  = errors.New("oggopus: unexpected BOS placement")
	ErrUnexpectedEOS  = errors.New("oggopus: unexpected EOS placement")
	ErrHeaderSequence = errors.New("oggopus: missing OpusHead/OpusTags")
)

var (
	opusHeadMagic = []byte("OpusHead")
	opusTagsMagic = []byte("OpusTags")
)

// OpusHead is the Opus identification header (RFC 7845).
type OpusHead struct {
	Version              uint8
	Channels             uint8
	PreSkip              uint16
	InputSampleRate      uint32
	OutputGainQ8         int16
	ChannelMappingFamily uint8

	// Present when ChannelMappingFamily != 0
	StreamCount        uint8
	CoupledStreamCount uint8
	ChannelMapping     []uint8
}

// OpusTags is the Opus comment header. This parser is minimal.
type OpusTags struct {
	Vendor   string
	Comments []string
}

type Packet struct {
	Data         []byte
	GranulePos   uint64
	GranuleValid bool
	EOS          bool
}

// Reader reads an Ogg Opus file/stream and yields audio packets.
type Reader struct {
	pr *ogg.PacketReader

	Head OpusHead
	Tags OpusTags

	headRead bool
	tagsRead bool
}

func NewReader(r io.Reader) (*Reader, error) {
	or := &Reader{pr: ogg.NewPacketReader(r)}
	if err := or.readHeaders(); err != nil {
		return nil, err
	}
	return or, nil
}

func (r *Reader) SetVerifyCRC(v bool) { r.pr.SetVerifyCRC(v) }

func (r *Reader) readHeaders() error {
	// Opus in Ogg is: BOS page contains OpusHead packet first, then OpusTags.
	pkt1, err := r.pr.ReadPacket()
	if err != nil {
		return err
	}
	if !pkt1.BOS {
		return fmt.Errorf("%w: first packet not BOS", ErrUnexpectedBOS)
	}
	head, err := parseOpusHead(pkt1.Data)
	if err != nil {
		return err
	}
	r.Head = head
	r.headRead = true

	pkt2, err := r.pr.ReadPacket()
	if err != nil {
		return err
	}
	tags, err := parseOpusTags(pkt2.Data)
	if err != nil {
		return err
	}
	r.Tags = tags
	r.tagsRead = true
	return nil
}

// ReadAudioPacket returns the next Opus audio packet (excluding OpusHead/Tags).
func (r *Reader) ReadAudioPacket() (*Packet, error) {
	if !r.headRead || !r.tagsRead {
		return nil, ErrHeaderSequence
	}
	pkt, err := r.pr.ReadPacket()
	if err != nil {
		return nil, err
	}
	if pkt.BOS {
		return nil, fmt.Errorf("%w: BOS after headers", ErrUnexpectedBOS)
	}
	if len(pkt.Data) >= 8 && bytes.Equal(pkt.Data[:8], opusHeadMagic) {
		return nil, ErrBadOpusHead
	}
	if len(pkt.Data) >= 8 && bytes.Equal(pkt.Data[:8], opusTagsMagic) {
		return nil, ErrBadOpusTags
	}
	return &Packet{
		Data:         pkt.Data,
		GranulePos:   pkt.GranulePosition,
		GranuleValid: pkt.GranuleValid,
		EOS:          pkt.EOS,
	}, nil
}

func parseOpusHead(b []byte) (OpusHead, error) {
	if len(b) < 19 {
		return OpusHead{}, fmt.Errorf("%w: too short", ErrBadOpusHead)
	}
	if string(b[:8]) != "OpusHead" {
		return OpusHead{}, ErrNotOpusOgg
	}
	h := OpusHead{}
	h.Version = b[8]
	h.Channels = b[9]
	h.PreSkip = binary.LittleEndian.Uint16(b[10:12])
	h.InputSampleRate = binary.LittleEndian.Uint32(b[12:16])
	h.OutputGainQ8 = int16(binary.LittleEndian.Uint16(b[16:18]))
	h.ChannelMappingFamily = b[18]

	if h.ChannelMappingFamily != 0 {
		if len(b) < 21 {
			return OpusHead{}, fmt.Errorf("%w: mapping too short", ErrBadOpusHead)
		}
		h.StreamCount = b[19]
		h.CoupledStreamCount = b[20]
		need := 21 + int(h.Channels)
		if len(b) < need {
			return OpusHead{}, fmt.Errorf("%w: mapping table too short", ErrBadOpusHead)
		}
		h.ChannelMapping = make([]uint8, h.Channels)
		copy(h.ChannelMapping, b[21:need])
	}
	return h, nil
}

func parseOpusTags(b []byte) (OpusTags, error) {
	if len(b) < 16 {
		return OpusTags{}, fmt.Errorf("%w: too short", ErrBadOpusTags)
	}
	if string(b[:8]) != "OpusTags" {
		return OpusTags{}, ErrHeaderSequence
	}
	off := 8
	vendorLen := int(binary.LittleEndian.Uint32(b[off : off+4]))
	off += 4
	if vendorLen < 0 || off+vendorLen > len(b) {
		return OpusTags{}, fmt.Errorf("%w: bad vendor length", ErrBadOpusTags)
	}
	vendor := string(b[off : off+vendorLen])
	off += vendorLen
	if off+4 > len(b) {
		return OpusTags{}, fmt.Errorf("%w: missing comment count", ErrBadOpusTags)
	}
	count := int(binary.LittleEndian.Uint32(b[off : off+4]))
	off += 4

	comments := make([]string, 0, count)
	for i := 0; i < count; i++ {
		if off+4 > len(b) {
			return OpusTags{}, fmt.Errorf("%w: truncated comment length", ErrBadOpusTags)
		}
		cl := int(binary.LittleEndian.Uint32(b[off : off+4]))
		off += 4
		if cl < 0 || off+cl > len(b) {
			return OpusTags{}, fmt.Errorf("%w: truncated comment", ErrBadOpusTags)
		}
		comments = append(comments, string(b[off:off+cl]))
		off += cl
	}
	return OpusTags{Vendor: vendor, Comments: comments}, nil
}
