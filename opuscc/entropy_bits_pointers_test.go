package opuscc

import "testing"

func TestEntropyBitsPointers(t *testing.T) {
	for available := int32(0); available <= 32; available++ {
		for width := uint32(0); width <= 25; width++ {
			window := uint32(0xd6b59a73) & ((uint32(1) << available) - 1)
			dec := OpusT_ec_dec{Frng: 1 << 31, Fval: 123, Fext: 456, Frem: 78, Ferror1: 9, Fend_window: window, Fnend_bits: available, Fnbits_total: 33}
			want := dec
			if uint32(want.Fnend_bits) < width {
				for {
					want.Fnend_bits += 8
					if want.Fnend_bits > 24 {
						break
					}
				}
			}
			want.Fnend_bits -= int32(width)
			want.Fnbits_total += int32(width)
			want.Fend_window >>= width
			expected := window & ((uint32(1) << width) - 1)
			if got := Opus_ec_dec_bits(nil, &dec, width); got != expected || dec != want {
				t.Fatalf("available=%d width=%d got=%x want=%x state=%+v want=%+v", available, width, got, expected, dec, want)
			}
		}
	}
}
