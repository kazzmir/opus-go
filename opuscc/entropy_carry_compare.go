//go:build compareopus

package opuscc

func CompareEntropyNormalize(enc *OpusT_ec_enc) { ec_enc_normalize(nil, enc) }

func CompareEntropyCarry(enc *OpusT_ec_enc, c int32) { ec_enc_carry_out(nil, enc, c) }
