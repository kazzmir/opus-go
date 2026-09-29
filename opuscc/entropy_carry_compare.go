//go:build compareopus

package opuscc

func CompareEntropyCarry(enc *OpusT_ec_enc, c int32) { ec_enc_carry_out(nil, enc, c) }
