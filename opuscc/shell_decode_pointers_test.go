package opuscc

import "testing"

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
