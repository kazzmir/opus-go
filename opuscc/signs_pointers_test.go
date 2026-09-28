package opuscc

import "testing"

func TestSignsPointers(t *testing.T) {
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	before := dec
	Opus_silk_decode_signs(nil, &dec, nil, 0, 0, 0, nil)
	if dec != before {
		t.Fatal("empty input changed state")
	}
	sums := [2]int32{0, 32}
	var pulses [34]int16
	pulses[0] = 123
	pulses[33] = 456
	for i := 1; i <= 32; i++ {
		pulses[i] = int16(i)
	}
	pulses[17] = 0
	pulses[18] = -32768
	pulses[19] = 32767
	want := pulses
	for i := 17; i <= 32; i++ {
		if want[i] > 0 {
			want[i] = -want[i]
		}
	}
	// length=24 rounds to two full shell blocks, including padded samples.
	Opus_silk_decode_signs(nil, &dec, &pulses[1], 24, 2, 1, &sums[0])
	if pulses != want {
		t.Fatalf("got=%v want=%v", pulses, want)
	}
	sums = [2]int32{0, -1}
	before = dec
	Opus_silk_decode_signs(nil, &dec, &pulses[1], 32, 2, 1, &sums[0])
	if pulses != want || dec != before {
		t.Fatal("empty shell blocks consumed signs")
	}
}
