package opuscc

import (
	"runtime"
	"slices"
	"testing"
)

// From shell_encode_pointers_test.go.
func TestShellEncodePointers(t *testing.T) {
	var b [1024]byte
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[0], 1024)
	var empty [16]int32
	before := e
	Opus_silk_shell_encoder(nil, &e, &empty)
	if e != before {
		t.Fatal("zero tree")
	}
	for total := int32(1); total <= 16; total++ {
		for pos := 0; pos < 16; pos++ {
			var p [16]int32
			p[pos] = total
			Opus_silk_shell_encoder(nil, &e, &p)
		}
	}
	Opus_ec_enc_done(nil, &e)
	if e.Ferror1 != 0 {
		t.Fatal(e)
	}
	var d OpusT_ec_dec
	Opus_ec_dec_init(nil, &d, &b[0], 1024)
	for total := int32(1); total <= 16; total++ {
		for pos := 0; pos < 16; pos++ {
			var got [16]int16
			Opus_silk_shell_decoder(nil, &got, &d, total)
			for j, v := range got {
				want := int16(0)
				if j == pos {
					want = int16(total)
				}
				if v != want {
					t.Fatal(total, pos, j, v)
				}
			}
		}
	}
}

// From shell_decode_pointers_test.go.
func TestShellDecodeSplitRowsPointers(t *testing.T) {
	for _, table := range [][]byte{Opus_silk_shell_code_table0[:], Opus_silk_shell_code_table1[:], Opus_silk_shell_code_table2[:], Opus_silk_shell_code_table3[:]} {
		for p := int32(1); p <= 16; p++ {
			offset := int32(Opus_silk_shell_code_table_offsets[p])
			checkPulseICDFRow(t, table[offset:offset+p+1], func(dec *OpusT_ec_dec) int32 {
				out := [4]int16{77, -1, -1, 88}
				decode_split(nil, &out[1], &out[2], dec, p, table)
				if out[0] != 77 || out[3] != 88 || int32(out[2]) != p-int32(out[1]) {
					t.Fatal("split sequential stores/guards", p, out)
				}
				return int32(out[1])
			})
			state := OpusT_ec_dec{Frng: 1 << 31, Fval: uint32(table[offset]) * (1 << 23), Fnbits_total: 33}
			expected := state
			left := ec_dec_icdf(nil, &expected, &table[offset], 8)
			shared := int16(-1)
			decode_split(nil, &shared, &shared, &state, p, table)
			if shared != int16(p-left) || state != expected {
				t.Fatal("aliased positive split stores", p)
			}
		}
	}
	for _, p := range []int32{1, 17} {
		state := OpusT_ec_dec{Frng: 123, Fval: 45}
		before := state
		out := [2]int16{77, 88}
		panicked := false
		func() {
			defer func() { panicked = recover() != nil }()
			decode_split(nil, &out[0], &out[1], &state, p, nil)
		}()
		if !panicked || state != before || out != [2]int16{77, 88} {
			t.Fatal("split validation before mutation", p)
		}
	}
}

func TestShellDecodePointers(t *testing.T) {
	for total := int32(0); total <= 16; total++ {
		state := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
		var output [18]int16
		output[0], output[17] = 111, 222
		for i := 1; i <= 16; i++ {
			output[i] = -1
		}
		Opus_silk_shell_decoder(nil, (*[16]int16)(output[1:17]), &state, total)
		for i := 1; i <= 16; i++ {
			want := int16(0)
			if i == 16 {
				want = int16(total)
			}
			if output[i] != want {
				t.Fatalf("total=%d output=%v", total, output)
			}
		}
		if output[0] != 111 || output[17] != 222 {
			t.Fatal("sentinels changed")
		}
		if total == 0 && state != (OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}) {
			t.Fatal("zero pulses changed entropy state")
		}
	}
	var child int16 = 123
	decode_split(nil, &child, &child, nil, 0, nil)
	if child != 0 {
		t.Fatal("zero split did not clear aliased output")
	}
}

func TestICDFCorePointers(t *testing.T) {
	table := []uint8{250, 180, 100, 0}
	for symbol := int32(0); symbol < 4; symbol++ {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: uint32(table[symbol]) * (1 << 23), Fnbits_total: 33}
		if got := ec_dec_icdf(nil, &dec, &table[0], 8); got != symbol || dec.Fval != 0 {
			t.Fatalf("symbol=%d got=%d state=%+v", symbol, got, dec)
		}
	}
}

// From sign_encode_pointers_test.go.
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

// From signs_pointers_test.go.
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

// From encode_pulses_pointers_test.go.
func TestEncodePulsesPointers(t *testing.T) {
	for _, pair := range [][2]int32{{2, 128}, {3, 32}, {8, 8}, {176, 1}} {
		n, k := pair[0], pair[1]
		var b [512]byte
		var e OpusT_ec_enc
		Opus_ec_enc_init(nil, &e, &b[0], 512)
		var ownedBuffer [512]byte
		var ownedEnc OpusT_ec_enc
		Opus_ec_enc_init(nil, &ownedEnc, &ownedBuffer[0], 512)
		vectors := make([][]int32, 20)
		for i := range vectors {
			y := make([]int32, n)
			y[i%int(n)] = k
			if i%2 != 0 {
				y[i%int(n)] = -k
			}
			vectors[i] = y
			before := slices.Clone(y)
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_encode_pulses(nil, &y[0], n, k, &e)
			encodePulsesOwned(nil, y, n, k, &ownedEnc)
			probe := ownedEnc
			probe.Fbuf = e.Fbuf // Only the independent buffer's pointer identity differs.
			if probe != e || ownedBuffer != b || !slices.Equal(y, before) {
				t.Fatal("owned pulse encode complete entropy/bytes/input", n, k, i)
			}
		}
		Opus_ec_enc_done(nil, &e)
		Opus_ec_enc_done(nil, &ownedEnc)
		probe := ownedEnc
		probe.Fbuf = e.Fbuf
		if probe != e || ownedBuffer != b {
			t.Fatal("owned pulse encode finalized entropy/full bytes", n, k)
		}
		if e.Ferror1 != 0 {
			t.Fatal(e)
		}
		var d OpusT_ec_dec
		Opus_ec_dec_init(nil, &d, &b[0], 512)
		for _, want := range vectors {
			got := make([]int32, n)
			Opus_decode_pulses(nil, &got[0], n, k, &d)
			if !slices.Equal(got, want) {
				t.Fatal(n, k, got, want)
			}
		}
	}
}

// From pulse_decode_pointers_test.go.
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
