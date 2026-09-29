package opuscc

import "testing"

func TestSignEncodePointers(t *testing.T) {
	Opus_silk_encode_signs(nil, nil, nil, 7, 0, 0, nil)
	for signal := int32(0); signal < 3; signal++ {
		for quant := int32(0); quant < 2; quant++ {
			var pulses [128]int8
			var sums [8]int32
			for i := range sums {
				sums[i] = int32(i + 1)
			}
			for i := range pulses {
				pulses[i] = int8(i)
				if i%3 == 0 {
					pulses[i] = 0
				} else if i%2 == 0 {
					pulses[i] = -pulses[i]
				}
			}
			var b [64]byte
			var e OpusT_ec_enc
			Opus_ec_enc_init(nil, &e, &b[0], 64)
			Opus_silk_encode_signs(nil, &e, &pulses[0], 120, signal, quant, &sums[0])
			Opus_ec_enc_done(nil, &e)
			if e.Ferror1 != 0 {
				t.Fatal(e)
			}
			var d OpusT_ec_dec
			Opus_ec_dec_init(nil, &d, &b[0], 64)
			var magnitudes [128]int16
			for i, v := range pulses {
				magnitudes[i] = int16(v)
				if magnitudes[i] < 0 {
					magnitudes[i] = -magnitudes[i]
				}
			}
			Opus_silk_decode_signs(nil, &d, &magnitudes[0], 120, signal, quant, &sums[0])
			for i, v := range magnitudes {
				if v != int16(pulses[i]) {
					t.Fatal(signal, quant, i, v, pulses[i])
				}
			}
		}
	}
	var zeroPulses [16]int8
	var skipped [1]int32
	e := OpusT_ec_enc{Frng: 1 << 31, Frem: -1}
	before := e
	Opus_silk_encode_signs(nil, &e, &zeroPulses[0], 16, 0, 0, &skipped[0])
	if e != before {
		t.Fatal("skipped")
	}
}
