package opuscc

import (
	"runtime"
	"testing"
)

func checkPulseICDFRow(t *testing.T, table *byte, symbols int32, decode func(*OpusT_ec_dec) int32) {
	t.Helper()
	buffer := make([]byte, 256)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, &buffer[0], 256)
	for i := int32(0); i < 48; i++ {
		Opus_ec_enc_icdf(nil, &encoder, i%symbols, table, 8)
	}
	Opus_ec_enc_done(nil, &encoder)
	var expected OpusT_ec_dec
	Opus_ec_dec_init(nil, &expected, &buffer[0], 256)
	got := expected
	entropyInitGrowStack(12)
	runtime.GC()
	for i := int32(0); i < 48; i++ {
		want := Opus_ec_dec_icdf(nil, &expected, table, 8)
		if symbol := decode(&got); symbol != want || symbol != i%symbols || got != expected {
			t.Fatal("pulse ICDF row/state", i, symbol, want)
		}
	}
}

func TestPulseDecodeRateRowsPointers(t *testing.T) {
	for row := int32(0); row < 2; row++ {
		checkPulseICDFRow(t, &Opus_silk_rate_levels_iCDF[row][0], 9, func(dec *OpusT_ec_dec) int32 { return silkDecodeRateICDF(nil, dec, row*2) })
	}
}

func TestPulseDecodePointers(t *testing.T) {
	for _, length := range []int32{16, 80, 120, 160, 240, 320} {
		for signal := int32(0); signal <= 2; signal++ {
			for offset := int32(0); offset <= 1; offset++ {
				n := ((length + 15) / 16) * 16
				out := make([]int16, n+2)
				for i := range out {
					out[i] = 123
				}
				dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 7}
				Opus_silk_decode_pulses(nil, &dec, &out[1], signal, offset, length)
				if out[0] != 123 || out[n+1] != 123 || dec.Fext != 123 || dec.Fend_window != 456 || dec.Fnend_bits != 7 {
					t.Fatal("guard or unrelated entropy fields changed")
				}
				for _, v := range out[1 : n+1] {
					if v != 0 {
						t.Fatalf("zero packet produced nonzero pulse: %d", v)
					}
				}
			}
		}
	}
}
