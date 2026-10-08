package opuscc

import (
	"math/bits"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// Generates excitation for CNG LPC synthesis.
func silk_CNG_exc(tls *libc.TLS, exc_Q14 *OpusT_opus_int32, exc_buf_Q14 *OpusT_opus_int32, length int32, rand_seed *OpusT_opus_int32) {
	if length == 0 {
		return
	}
	out := unsafe.Slice(exc_Q14, int(length))
	exc_mask := int32(CNG_BUF_MASK_MAX)
	for exc_mask > length {
		exc_mask >>= 1
	}
	// The mask is inclusive: even length == mask needs mask+1 input slots.
	buffer := unsafe.Slice(exc_buf_Q14, int(exc_mask)+1)
	seed := *rand_seed
	for i := range out {
		seed = int32(uint32(RAND_INCREMENT) + uint32(seed)*uint32(RAND_MULTIPLIER))
		idx := (seed >> 24) & exc_mask
		out[i] = buffer[idx]
	}
	*rand_seed = seed
}

func Opus_silk_CNG_Reset(tls *libc.TLS, dec *OpusT_silk_decoder_state) {
	var NLSF_acc_Q15, NLSF_step_Q15, i int32
	_, _, _ = NLSF_acc_Q15, NLSF_step_Q15, i
	cng := &dec.FsCNG
	NLSF_step_Q15 = int32(silk_int16_MAX1) / (dec.FLPC_order + int32(1))
	NLSF_acc_Q15 = 0
	i = 0
	for {
		if !(i < dec.FLPC_order) {
			break
		}
		NLSF_acc_Q15 = NLSF_acc_Q15 + NLSF_step_Q15
		cng.FCNG_smth_NLSF_Q15[i] = int16(NLSF_acc_Q15)
		i = i + 1
	}
	cng.FCNG_smth_Gain_Q16 = 0
	cng.Frand_seed = int32(3176576)
}

// cngSqrtApprox mirrors silk_SQRT_APPROX/CLZ_FRAC in silk/Inlines.h,
// including the signed 16-bit multiplier used by SMLAWB.
func cngSqrtApprox(x int32) int32 {
	if x <= 0 {
		return 0
	}
	lz := bits.LeadingZeros32(uint32(x))
	frac := int32(bits.RotateLeft32(uint32(x), lz-24) & 0x7f)
	y := int32(46214)
	if lz&1 != 0 {
		y = 32768
	}
	y >>= uint(lz >> 1)
	return int32(int64(y) + (int64(y) * int64(int16(213*frac)) >> 16))
}

// Updates CNG estimates and adds comfort noise on lost packets. The excitation
// and full 16-word synthesis history are Go-owned, not TLS pseudostack addresses.
func Opus_silk_CNG(tls *libc.TLS, dec *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, frame *int16, length int32) {
	cng := &dec.FsCNG
	if dec.Ffs_kHz != cng.Ffs_kHz {
		Opus_silk_CNG_Reset(tls, dec)
		cng.Ffs_kHz = dec.Ffs_kHz
	}
	if dec.FlossCnt == 0 && dec.FprevSignalType == TYPE_NO_VOICE_ACTIVITY {
		for i := int32(0); i < dec.FLPC_order; i++ {
			cng.FCNG_smth_NLSF_Q15[i] = int16(int32(cng.FCNG_smth_NLSF_Q15[i]) + int32(int64(int32(dec.FprevNLSF_Q15[i])-int32(cng.FCNG_smth_NLSF_Q15[i]))*int64(int16(CNG_NLSF_SMTH_Q16))>>16))
		}
		maxGain, subfr := int32(0), int32(0)
		for i := int32(0); i < dec.Fnb_subfr; i++ {
			if control.FGains_Q16[i] > maxGain {
				maxGain = control.FGains_Q16[i]
				subfr = i
			}
		}
		// Shift first, then copy, then read gains again: control may alias excitation.
		copy(cng.FCNG_exc_buf_Q14[dec.Fsubfr_length:dec.Fnb_subfr*dec.Fsubfr_length], cng.FCNG_exc_buf_Q14[:(dec.Fnb_subfr-1)*dec.Fsubfr_length])
		copy(cng.FCNG_exc_buf_Q14[:dec.Fsubfr_length], dec.Fexc_Q14[subfr*dec.Fsubfr_length:(subfr+1)*dec.Fsubfr_length])
		for i := int32(0); i < dec.Fnb_subfr; i++ {
			cng.FCNG_smth_Gain_Q16 += int32(int64(control.FGains_Q16[i]-cng.FCNG_smth_Gain_Q16) * int64(int16(CNG_GAIN_SMTH_Q16)) >> 16)
			if int32(int64(cng.FCNG_smth_Gain_Q16)*int64(CNG_GAIN_SMTH_THRESHOLD_Q16)>>16) > control.FGains_Q16[i] {
				cng.FCNG_smth_Gain_Q16 = control.FGains_Q16[i]
			}
		}
	}
	if dec.FlossCnt != 0 {
		signal := make([]int32, int(length)+MAX_LPC_ORDER)
		gain := int32(int64(dec.FsPLC.FrandScale_Q14) * int64(dec.FsPLC.FprevGain_Q16[1]) >> 16)
		if gain >= 1<<21 || cng.FCNG_smth_Gain_Q16 > 1<<23 {
			gain = (gain >> 16) * (gain >> 16)
			gain = (cng.FCNG_smth_Gain_Q16>>16)*(cng.FCNG_smth_Gain_Q16>>16) - int32(uint32(gain)<<5)
			gain = int32(uint32(cngSqrtApprox(gain)) << 16)
		} else {
			gain = int32(int64(gain) * int64(gain) >> 16)
			gain = int32(int64(cng.FCNG_smth_Gain_Q16)*int64(cng.FCNG_smth_Gain_Q16)>>16) - int32(uint32(gain)<<5)
			gain = int32(uint32(cngSqrtApprox(gain)) << 8)
		}
		gain >>= 6
		silk_CNG_exc(tls, unsafe.SliceData(signal[MAX_LPC_ORDER:]), &cng.FCNG_exc_buf_Q14[0], length, &cng.Frand_seed)
		var coefficients [MAX_LPC_ORDER]int16
		Opus_silk_NLSF2A(tls, &coefficients[0], &cng.FCNG_smth_NLSF_Q15[0], dec.FLPC_order, dec.Farch)
		copy(signal[:MAX_LPC_ORDER], cng.FCNG_synth_state[:])
		if !(dec.FLPC_order == 10 || dec.FLPC_order == 16) {
			opusCeltFatal(tls, opusDiagnosticString(5777), opusDiagnosticString(5763), 153)
		}
		pcm := unsafe.Slice(frame, length)
		for i := int32(0); i < length; i++ {
			// SMLAWB narrows after every MAC and rounds products towards -infinity.
			pred := dec.FLPC_order >> 1
			for j := int32(0); j < 10; j++ {
				pred = int32(int64(pred) + (int64(signal[MAX_LPC_ORDER+i-j-1]) * int64(coefficients[j]) >> 16))
			}
			if dec.FLPC_order == 16 {
				for j := int32(10); j < 16; j++ {
					pred = int32(int64(pred) + (int64(signal[MAX_LPC_ORDER+i-j-1]) * int64(coefficients[j]) >> 16))
				}
			}
			// LSHIFT_SAT32 followed by ADD_SAT32, without signed-overflow assumptions.
			if pred > 2147483647>>4 {
				pred = 2147483647 >> 4
			} else if pred < -2147483648>>4 {
				pred = -2147483648 >> 4
			}
			sample := int64(signal[MAX_LPC_ORDER+i]) + int64(int32(uint32(pred)<<4))
			if sample > 2147483647 {
				sample = 2147483647
			} else if sample < -2147483648 {
				sample = -2147483648
			}
			signal[MAX_LPC_ORDER+i] = int32(sample)
			// SMULWW narrows before RSHIFT_ROUND; both additions then saturate to int16.
			scaled := int32(int64(signal[MAX_LPC_ORDER+i]) * int64(gain) >> 16)
			scaled = ((scaled >> 7) + 1) >> 1
			if scaled > 32767 {
				scaled = 32767
			} else if scaled < -32768 {
				scaled = -32768
			}
			sum := int32(pcm[i]) + scaled
			if sum > 32767 {
				sum = 32767
			} else if sum < -32768 {
				sum = -32768
			}
			pcm[i] = int16(sum)
		}
		// This store follows every PCM write, including overlapping history/PCM views.
		copy(cng.FCNG_synth_state[:], signal[length:length+MAX_LPC_ORDER])
	} else {
		clear(cng.FCNG_synth_state[:dec.FLPC_order])
	}
}

const silk_int16_MAX2 = 0x7FFF
