package opuscc

import (
	"runtime"
	"testing"
)

func checkPulseICDFRow(t *testing.T, table []byte, decode func(*OpusT_ec_dec) int32) {
	t.Helper()
	var symbols []int32
	previous := uint16(256)
	for i, v := range table {
		if uint16(v) < previous {
			symbols = append(symbols, int32(i))
		}
		previous = uint16(v)
	}
	if len(symbols) == 0 || previous != 0 {
		t.Fatal("invalid ICDF fixture")
	}
	buffer := make([]byte, 256)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, &buffer[0], 256)
	for i := int32(0); i < 48; i++ {
		Opus_ec_enc_icdf(nil, &encoder, symbols[int(i)%len(symbols)], &table[0], 8)
	}
	Opus_ec_enc_done(nil, &encoder)
	var expected OpusT_ec_dec
	Opus_ec_dec_init(nil, &expected, &buffer[0], 256)
	got := expected
	entropyInitGrowStack(12)
	runtime.GC()
	for i := int32(0); i < 48; i++ {
		want := Opus_ec_dec_icdf(nil, &expected, &table[0], 8)
		if symbol := decode(&got); symbol != want || symbol != symbols[int(i)%len(symbols)] || got != expected {
			t.Fatal("pulse ICDF row/state", i, symbol, want)
		}
	}
}

func TestPulseDecodeRateRowsPointers(t *testing.T) {
	for row := int32(0); row < 2; row++ {
		checkPulseICDFRow(t, Opus_silk_rate_levels_iCDF[row][:], func(dec *OpusT_ec_dec) int32 { return silkDecodeRateICDF(nil, dec, row*2) })
	}
}

func TestPulseDecodeCountRowsPointers(t *testing.T) {
	for row := int32(0); row < 10; row++ {
		checkPulseICDFRow(t, Opus_silk_pulses_per_block_iCDF[row][:], func(dec *OpusT_ec_dec) int32 { return silkDecodePulseICDF(nil, dec, row, 0) })
	}
}

func TestPulseDecodeEscapeRowsPointers(t *testing.T) {
	for start := int32(0); start < 2; start++ {
		checkPulseICDFRow(t, Opus_silk_pulses_per_block_iCDF[N_RATE_LEVELS-1][start:], func(dec *OpusT_ec_dec) int32 { return silkDecodePulseICDF(nil, dec, N_RATE_LEVELS-1, start) })
	}
	buffer := make([]byte, 256)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, &buffer[0], 256)
	Opus_ec_enc_icdf(nil, &encoder, 0, &Opus_silk_rate_levels_iCDF[0][0], 8)
	Opus_ec_enc_icdf(nil, &encoder, SILK_MAX_PULSES+1, &Opus_silk_pulses_per_block_iCDF[0][0], 8)
	for shift := 1; shift <= 10; shift++ {
		start, symbol := 0, int32(SILK_MAX_PULSES+1)
		if shift == 10 {
			start, symbol = 1, 0
		}
		Opus_ec_enc_icdf(nil, &encoder, symbol, &Opus_silk_pulses_per_block_iCDF[N_RATE_LEVELS-1][start], 8)
	}
	for i := 0; i < 160; i++ {
		Opus_ec_enc_icdf(nil, &encoder, 0, &Opus_silk_lsb_iCDF[0], 8)
	}
	Opus_ec_enc_done(nil, &encoder)
	var expected OpusT_ec_dec
	Opus_ec_dec_init(nil, &expected, &buffer[0], 256)
	got := expected
	rate := ec_dec_icdf(nil, &expected, &Opus_silk_rate_levels_iCDF[0][0], 8)
	sum := ec_dec_icdf(nil, &expected, &Opus_silk_pulses_per_block_iCDF[rate][0], 8)
	shifts := 0
	for sum == SILK_MAX_PULSES+1 {
		shifts++
		start := 0
		if shifts == 10 {
			start = 1
		}
		sum = ec_dec_icdf(nil, &expected, &Opus_silk_pulses_per_block_iCDF[N_RATE_LEVELS-1][start], 8)
	}
	if shifts != 10 || sum != 0 {
		t.Fatal("ten-shift fixture", shifts, sum)
	}
	for i := 0; i < 160; i++ {
		if ec_dec_icdf(nil, &expected, &Opus_silk_lsb_iCDF[0], 8) != 0 {
			t.Fatal("LSB fixture")
		}
	}
	out := [18]int16{77}
	out[17] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_silk_decode_pulses(nil, &got, &out[1], 0, 0, 16)
	if got != expected || out[0] != 77 || out[17] != 88 {
		t.Fatal("ten-shift context/guards")
	}
	for _, pulse := range out[1:17] {
		if pulse != 0 {
			t.Fatal("ten-shift pulse", pulse)
		}
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
