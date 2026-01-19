package ogg

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// BuildOpusHeadPacket builds an OpusHead identification header packet (RFC 7845).
func BuildOpusHeadPacket(h OpusHead) ([]byte, error) {
	if h.Version == 0 {
		h.Version = 1
	}
	if h.Channels == 0 {
		return nil, fmt.Errorf("%w: channels=0", ErrBadOpusHead)
	}
	if h.ChannelMappingFamily == 0 {
		b := make([]byte, 19)
		copy(b[0:8], opusHeadMagic)
		b[8] = h.Version
		b[9] = h.Channels
		binary.LittleEndian.PutUint16(b[10:12], h.PreSkip)
		binary.LittleEndian.PutUint32(b[12:16], h.InputSampleRate)
		binary.LittleEndian.PutUint16(b[16:18], uint16(h.OutputGainQ8))
		b[18] = 0
		return b, nil
	}

	if int(h.Channels) != len(h.ChannelMapping) {
		return nil, fmt.Errorf("%w: mapping len mismatch", ErrBadOpusHead)
	}
	b := make([]byte, 21+len(h.ChannelMapping))
	copy(b[0:8], opusHeadMagic)
	b[8] = h.Version
	b[9] = h.Channels
	binary.LittleEndian.PutUint16(b[10:12], h.PreSkip)
	binary.LittleEndian.PutUint32(b[12:16], h.InputSampleRate)
	binary.LittleEndian.PutUint16(b[16:18], uint16(h.OutputGainQ8))
	b[18] = h.ChannelMappingFamily
	b[19] = h.StreamCount
	b[20] = h.CoupledStreamCount
	copy(b[21:], h.ChannelMapping)
	return b, nil
}

// BuildOpusTagsPacket builds an OpusTags comment header packet (RFC 7845).
func BuildOpusTagsPacket(t OpusTags) ([]byte, error) {
	vendor := t.Vendor
	if vendor == "" {
		vendor = "opusgo"
	}

	var buf bytes.Buffer
	buf.Write(opusTagsMagic)
	_ = binary.Write(&buf, binary.LittleEndian, uint32(len(vendor)))
	buf.WriteString(vendor)
	_ = binary.Write(&buf, binary.LittleEndian, uint32(len(t.Comments)))
	for _, c := range t.Comments {
		_ = binary.Write(&buf, binary.LittleEndian, uint32(len(c)))
		buf.WriteString(c)
	}
	return buf.Bytes(), nil
}
