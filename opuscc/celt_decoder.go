// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"math"
	"math/bits"
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_validate_celt_decoder(tls *libc.TLS, st *OpusT_OpusCustomDecoder) {
	mode, _ := Opus_opus_custom_mode_create(tls, 48000, 960)
	if st.Fmode != mode {
		Opus_celt_fatal(tls, __ccgo_ts+3695, __ccgo_ts+3767, 147)
	}
	if st.Foverlap != 120 {
		Opus_celt_fatal(tls, __ccgo_ts+3790, __ccgo_ts+3767, 148)
	}
	if st.Fend > 21 {
		Opus_celt_fatal(tls, __ccgo_ts+3827, __ccgo_ts+3767, 149)
	}
	if !(st.Fchannels == 1 || st.Fchannels == 2) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+3767, 157)
	}
	if !(st.Fstream_channels == 1 || st.Fstream_channels == 2) {
		Opus_celt_fatal(tls, __ccgo_ts+925, __ccgo_ts+3767, 158)
	}
	if st.Fdownsample <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+3859, __ccgo_ts+3767, 159)
	}
	if !(st.Fstart == 0 || st.Fstart == 17) {
		Opus_celt_fatal(tls, __ccgo_ts+3896, __ccgo_ts+3767, 160)
	}
	if st.Fstart >= st.Fend {
		Opus_celt_fatal(tls, __ccgo_ts+3948, __ccgo_ts+3767, 161)
	}
	if st.Farch < 0 {
		Opus_celt_fatal(tls, __ccgo_ts+849, __ccgo_ts+3767, 163)
	}
	if st.Farch > OPUS_ARCHMASK {
		Opus_celt_fatal(tls, __ccgo_ts+881, __ccgo_ts+3767, 164)
	}
	if st.Flast_pitch_index > PLC_PITCH_LAG_MAX {
		Opus_celt_fatal(tls, __ccgo_ts+3986, __ccgo_ts+3767, 167)
	}
	if !(st.Flast_pitch_index >= PLC_PITCH_LAG_MIN || st.Flast_pitch_index == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4046, __ccgo_ts+3767, 168)
	}
	if st.Fpostfilter_period >= MAX_PERIOD {
		Opus_celt_fatal(tls, __ccgo_ts+4135, __ccgo_ts+3767, 170)
	}
	if !(st.Fpostfilter_period >= COMBFILTER_MINPERIOD || st.Fpostfilter_period == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4188, __ccgo_ts+3767, 171)
	}
	if st.Fpostfilter_period_old >= MAX_PERIOD {
		Opus_celt_fatal(tls, __ccgo_ts+4282, __ccgo_ts+3767, 172)
	}
	if !(st.Fpostfilter_period_old >= COMBFILTER_MINPERIOD || st.Fpostfilter_period_old == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4339, __ccgo_ts+3767, 173)
	}
	if st.Fpostfilter_tapset > 2 {
		Opus_celt_fatal(tls, __ccgo_ts+4441, __ccgo_ts+3767, 174)
	}
	if st.Fpostfilter_tapset < 0 {
		Opus_celt_fatal(tls, __ccgo_ts+4486, __ccgo_ts+3767, 175)
	}
	if st.Fpostfilter_tapset_old > 2 {
		Opus_celt_fatal(tls, __ccgo_ts+4531, __ccgo_ts+3767, 176)
	}
	if st.Fpostfilter_tapset_old < 0 {
		Opus_celt_fatal(tls, __ccgo_ts+4580, __ccgo_ts+3767, 177)
	}
}

func Opus_celt_decoder_get_size(tls *libc.TLS, channels int32) (r int32) {
	mode, _ := Opus_opus_custom_mode_create(tls, int32(48000), int32(960))
	return opus_custom_decoder_get_size(tls, mode, channels)
}

func opus_custom_decoder_get_size(tls *libc.TLS, mode *OpusT_OpusCustomMode, channels int32) int32 {
	// The generated uint64 expression narrowed to int32; preserve its low 32 bits.
	return int32(unsafe.Sizeof(OpusT_OpusCustomDecoder{})) + (channels*(DEC_PITCH_BUF_SIZE+mode.Foverlap)-1)*4 + mode.FnbEBands*32 + channels*CELT_LPC_ORDER*4
}

func Opus_celt_decoder_init(tls *libc.TLS, st *OpusT_OpusCustomDecoder, rate OpusT_opus_int32, channels int32) int32 {
	mode, _ := Opus_opus_custom_mode_create(tls, 48000, 960)
	if ret := opus_custom_decoder_init(tls, st, mode, channels); ret != OPUS_OK {
		return ret
	}
	// C initializes the complete state before rejecting an unsupported rate.
	st.Fdownsample = Opus_resampling_factor(tls, rate)
	if st.Fdownsample == 0 {
		return OPUS_BAD_ARG
	}
	return OPUS_OK
}

func opus_custom_decoder_init(tls *libc.TLS, st *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, channels int32) int32 {
	if channels < 0 || channels > 2 {
		return OPUS_BAD_ARG
	}
	if st == nil {
		return -7
	}
	size := opus_custom_decoder_get_size(tls, mode, channels)
	pointerBytes := unsafe.Sizeof(st.Fmode)
	// Clear the sole pointer through its typed slot (including the GC write barrier).
	st.Fmode = nil
	clear(unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(st), pointerBytes)), int(size)-int(pointerBytes)))
	st.Fmode = mode
	st.Foverlap = mode.Foverlap
	st.Fchannels = channels
	st.Fstream_channels = channels
	st.Fdownsample = 1
	st.Fstart = 0
	st.Fend = st.Fmode.FeffEBands
	st.Fsignalling = 1
	st.Fdisable_inv = libc.BoolInt32(channels == 1)
	st.Farch = 0
	celt_decoder_reset(tls, st)
	return OPUS_OK
}

// C documentation
//
//	/* Special case for stereo with no downsampling and no accumulation. This is
//	   quite common and we can make it faster by processing both channels in the
//	   same loop, reducing overhead due to the dependency loop in the IIR filter. */
func deemphasis_stereo_simple(tls *libc.TLS, left *OpusT_celt_sig, right *OpusT_celt_sig, pcm *OpusT_opus_res, N int32, coef0 OpusT_opus_val16, mem *[2]OpusT_celt_sig) {
	if N <= 0 {
		return
	}
	x0, x1 := unsafe.Slice(left, int(N)), unsafe.Slice(right, int(N))
	output := unsafe.Slice(pcm, 2*int(N))
	m0, m1 := mem[0], mem[1]
	for j := range x0 {
		// Add VERY_SMALL first, preserving the floating-point C operation order.
		tmp0 := x0[j] + float32(1e-30) + m0
		tmp1 := x1[j] + float32(1e-30) + m1
		m0 = OpusT_opus_val16(coef0 * tmp0)
		m1 = OpusT_opus_val16(coef0 * tmp1)
		output[2*j] = float32(float32(1) / float32(32768) * tmp0)
		output[2*j+1] = float32(float32(1) / float32(32768) * tmp1)
	}
	mem[0], mem[1] = m0, m1
}

func deemphasis_legacy(tls *libc.TLS, in, pcm uintptr, N, C, downsample int32, coef, mem uintptr, accum int32) {
	raw := unsafe.Slice((*uintptr)(unsafe.Pointer(in)), max(C, 1))
	channels := make([]*float32, len(raw))
	for i := range raw {
		channels[i] = (*float32)(unsafe.Pointer(raw[i]))
	}
	deemphasis(tls, unsafe.SliceData(channels), (*float32)(unsafe.Pointer(pcm)), N, C, downsample, (*float32)(unsafe.Pointer(coef)), (*float32)(unsafe.Pointer(mem)), accum)
}
func deemphasis(tls *libc.TLS, in **float32, pcm *float32, N int32, C int32, downsample int32, coef *float32, mem *float32, accum int32) {
	// Common stereo dispatch needs no TLS scratch, just as the C early return.
	if downsample == 1 && C == 2 && accum == 0 {
		channels := unsafe.Slice(in, 2)
		deemphasis_stereo_simple(tls, channels[0], channels[1], pcm, N, *coef, (*[2]float32)(unsafe.Pointer(mem)))
		return
	}
	var x, y *float32
	var Nd, apply_downsampling, c, j, v29 int32
	var coef0 OpusT_opus_val16
	var m, tmp, tmp1, tmp2 OpusT_celt_sig
	// Match the C temporary's extent and per-channel reuse without a TLS address.
	scratch := make([]float32, N)
	apply_downsampling = 0
	coef0 = *coef
	Nd = N / downsample
	c = 0
	for {
		m = unsafe.Slice(mem, max(C, 1))[c]
		x = unsafe.Slice(in, max(C, 1))[c]
		y = nil
		// If Nd==0, C's channel cursor is never consumed; avoid forming one-past.
		if N > 0 && (downsample <= 1 || Nd > 0) {
			y = celtNormAdd(pcm, c)
		}
		if downsample > int32(1) {
			/* Shortcut for the standard (non-custom modes) case */
			j = 0
			for {
				if !(j < N) {
					break
				}
				tmp = unsafe.Slice(x, N)[j] + float32(1e-30) + m
				m = OpusT_opus_val16(coef0 * tmp)
				scratch[j] = tmp
				j = j + 1
			}
			apply_downsampling = int32(1)
		} else {
			/* Shortcut for the standard (non-custom modes) case */
			if accum != 0 {
				j = 0
				for {
					if !(j < N) {
						break
					}
					tmp1 = unsafe.Slice(x, N)[j] + m + float32(1e-30)
					m = OpusT_opus_val16(coef0 * tmp1)
					*celtNormAdd(y, j*C) = *celtNormAdd(y, j*C) + float32(float32(1)/float32(32768)*tmp1)
					j = j + 1
				}
			} else {
				j = 0
				for {
					if !(j < N) {
						break
					}
					tmp2 = unsafe.Slice(x, N)[j] + float32(1e-30) + m
					m = OpusT_opus_val16(coef0 * tmp2)
					*celtNormAdd(y, j*C) = float32(float32(1) / float32(32768) * tmp2)
					j = j + 1
				}
			}
		}
		unsafe.Slice(mem, max(C, 1))[c] = m
		if apply_downsampling != 0 {
			/* Perform down-sampling */
			if accum != 0 {
				j = 0
				for {
					if !(j < Nd) {
						break
					}
					*celtNormAdd(y, j*C) = *celtNormAdd(y, j*C) + float32(float32(1)/float32(32768)*scratch[j*downsample])
					j = j + 1
				}
			} else {
				j = 0
				for {
					if !(j < Nd) {
						break
					}
					*celtNormAdd(y, j*C) = float32(float32(1) / float32(32768) * scratch[j*downsample])
					j = j + 1
				}
			}
		}
		c = c + 1
		v29 = c
		if !(v29 < C) {
			break
		}
	}
}

//go:uintptrescapes
func celt_synthesis_legacy(tls *libc.TLS, mode, X, out, energy uintptr, start, end, C, CC, transient, LM, downsample, silence, arch int32) {
	typedMode := (*OpusT_OpusCustomMode)(unsafe.Pointer(mode))
	typedX := (*float32)(unsafe.Pointer(X))
	typedEnergy := (*float32)(unsafe.Pointer(energy))
	raw := unsafe.Slice((*uintptr)(unsafe.Pointer(out)), max(CC, 1))
	// Retain normal decoder channel buffers before any allocation/stack growth.
	var pair [2]*float32
	for i := 0; i < min(len(raw), len(pair)); i++ {
		pair[i] = (*float32)(unsafe.Pointer(raw[i]))
	}
	outputs := pair[:min(len(raw), len(pair))]
	if len(raw) > len(pair) {
		outputs = make([]*float32, len(raw))
		copy(outputs, pair[:])
		for i := len(pair); i < len(raw); i++ {
			outputs[i] = (*float32)(unsafe.Pointer(raw[i]))
		}
	}
	celt_synthesis(tls, typedMode, typedX, unsafe.SliceData(outputs), typedEnergy, start, end, C, CC, transient, LM, downsample, silence, arch)
}
func celtSynthesisGeometry(mode *OpusT_OpusCustomMode, LM int32) (overlap, bands, N int32) {
	return mode.Foverlap, mode.FnbEBands, mode.FshortMdctSize << LM
}
func celtSynthesisIMDCT(tls *libc.TLS, mode *OpusT_OpusCustomMode, freq, out *float32, overlap, shift, stride, arch int32) {
	lookup := &mode.Fmdct
	st := lookup.Fkfft[shift]
	Opus_clt_mdct_backward_c(tls, lookup, st.Fbitrev, st.Ftwiddles, freq, out, mode.Fwindow, overlap, shift, stride, arch)
}
func celtSynthesisOutputCopy(destination, source *float32, N int32) {
	copy(unsafe.Slice(destination, N), unsafe.Slice(source, N))
}
func celt_synthesis(tls *libc.TLS, mode *OpusT_OpusCustomMode, X *float32, out_syn **float32, oldBandE *float32, start int32, effEnd int32, C int32, CC int32, isTransient int32, LM int32, downsample int32, silence int32, arch int32) {
	bandBoundaries := mode.FeBands
	overlap, nbEBands, N := celtSynthesisGeometry(mode, LM)
	// Reuse one N-sample interleaved spectrum, matching the C scratch extent.
	freq := make([]float32, N)
	freqHead := unsafe.SliceData(freq)
	M := int32(1) << LM
	B, NB, shift := int32(1), mode.FshortMdctSize<<LM, mode.FmaxLM-LM
	if isTransient != 0 {
		B, NB, shift = M, mode.FshortMdctSize, mode.FmaxLM
	}
	outputs := unsafe.Slice(out_syn, max(CC, 1))
	if CC == 2 && C == 1 {
		Opus_denormalise_bands(tls, bandBoundaries, mode.FshortMdctSize, X, freqHead, oldBandE, start, effEnd, M, downsample, silence)
		// Copy before either transform, and finish all left blocks before right blocks.
		temp := celtNormAdd(outputs[1], overlap/2)
		celtSynthesisOutputCopy(temp, freqHead, N)
		for b := int32(0); b < B; b++ {
			celtSynthesisIMDCT(tls, mode, celtNormAdd(temp, b), celtNormAdd(outputs[0], NB*b), overlap, shift, B, arch)
		}
		for b := int32(0); b < B; b++ {
			celtSynthesisIMDCT(tls, mode, celtNormAdd(freqHead, b), celtNormAdd(outputs[1], NB*b), overlap, shift, B, arch)
		}
	} else if CC == 1 && C == 2 {
		temp := celtNormAdd(outputs[0], overlap/2)
		Opus_denormalise_bands(tls, bandBoundaries, mode.FshortMdctSize, X, freqHead, oldBandE, start, effEnd, M, downsample, silence)
		Opus_denormalise_bands(tls, bandBoundaries, mode.FshortMdctSize, celtNormAdd(X, N), temp, celtNormAdd(oldBandE, nbEBands), start, effEnd, M, downsample, silence)
		other := unsafe.Slice(temp, N)
		// Round both halves independently (including on FMA-capable architectures).
		for i := range freq {
			freq[i] = float32(float32(.5)*freq[i]) + float32(float32(.5)*other[i])
		}
		for b := int32(0); b < B; b++ {
			celtSynthesisIMDCT(tls, mode, celtNormAdd(freqHead, b), celtNormAdd(outputs[0], NB*b), overlap, shift, B, arch)
		}
	} else {
		// The C do-while visits channel zero even when CC is zero.
		for c := int32(0); c < max(CC, 1); c++ {
			Opus_denormalise_bands(tls, bandBoundaries, mode.FshortMdctSize, celtNormAdd(X, c*N), freqHead, celtNormAdd(oldBandE, c*nbEBands), start, effEnd, M, downsample, silence)
			for b := int32(0); b < B; b++ {
				celtSynthesisIMDCT(tls, mode, celtNormAdd(freqHead, b), celtNormAdd(outputs[c], NB*b), overlap, shift, B, arch)
			}
		}
	}
	// SATURATE(x,SIG_SAT) is the identity in the floating-point C build.
	for c := int32(0); c < max(CC, 1); c++ {
		out := unsafe.Slice(outputs[c], N)
		for i := range out {
			out[i] = out[i]
		}
	}
}

func tf_decode(tls *libc.TLS, start, end, isTransient int32, tfRes *int32, LM int32, dec *OpusT_ec_dec) {
	budget := dec.Fstorage * 8
	tell := uint32(dec.Fnbits_total - int32(bits.Len32(dec.Frng)))
	logp := int32(4)
	if isTransient != 0 {
		logp = 2
	}
	reserved := int32(0)
	if LM > 0 && tell+uint32(logp)+1 <= budget {
		reserved = 1
	}
	budget -= uint32(reserved)
	out := unsafe.Slice(tfRes, end)
	curr, changed := int32(0), int32(0)
	for i := start; i < end; i++ {
		if tell+uint32(logp) <= budget {
			curr ^= Opus_ec_dec_bit_logp(tls, dec, uint32(logp))
			tell = uint32(dec.Fnbits_total - int32(bits.Len32(dec.Frng)))
			changed |= curr
		}
		out[i] = curr
		logp = 5
		if isTransient != 0 {
			logp = 4
		}
	}
	selectBit := int32(0)
	table := Opus_tf_select_table[LM]
	if reserved != 0 && table[4*isTransient+changed] != table[4*isTransient+2+changed] {
		selectBit = Opus_ec_dec_bit_logp(tls, dec, 1)
	}
	for i := start; i < end; i++ {
		out[i] = int32(table[4*isTransient+2*selectBit+out[i]])
	}
}

func celt_plc_pitch_search(tls *libc.TLS, st *OpusT_OpusCustomDecoder, left, right *float32, C, arch int32) int32 {
	// st is unused in the non-QEXT C implementation.
	var lowpass [DEC_PITCH_BUF_SIZE / 2]float32
	var pitch int32
	Opus_pitch_downsample(tls, left, right, &lowpass[0], DEC_PITCH_BUF_SIZE/2, C, 2, arch)
	Opus_pitch_search(tls, &lowpass[PLC_PITCH_LAG_MAX/2], &lowpass[0], DEC_PITCH_BUF_SIZE-PLC_PITCH_LAG_MAX, PLC_PITCH_LAG_MAX-PLC_PITCH_LAG_MIN, &pitch, arch)
	return PLC_PITCH_LAG_MAX - pitch
}

//go:uintptrescapes
func prefilter_and_fold_legacy(tls *libc.TLS, st uintptr, N int32) {
	prefilter_and_fold(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)), N)
}
func prefilterFoldState(st *OpusT_OpusCustomDecoder) (mode *OpusT_OpusCustomMode, overlap, channels int32) {
	return st.Fmode, st.Foverlap, st.Fchannels
}
func prefilterFoldHistory(st *OpusT_OpusCustomDecoder, overlap, channel int32) *float32 {
	return celtNormAdd(&st.F_decode_mem[0], channel*(DEC_PITCH_BUF_SIZE+overlap))
}
func prefilterFoldTDAC(mode *OpusT_OpusCustomMode, memory, filtered *float32, overlap int32) {
	for i := int32(0); i < overlap/2; i++ {
		// Keep the generated/C operand order and live window/sample reads.
		*celtNormAdd(memory, i) = float32(OpusT_celt_coef(*celtNormAdd(mode.Fwindow, i))**celtNormAdd(filtered, overlap-1-i)) + float32(OpusT_celt_coef(*celtNormAdd(mode.Fwindow, overlap-i-1))**celtNormAdd(filtered, i))
	}
}
func prefilter_and_fold(tls *libc.TLS, st1 *OpusT_OpusCustomDecoder, N int32) {
	mode, overlap, CC := prefilterFoldState(st1)
	scratch := make([]float32, overlap)
	etmp := unsafe.SliceData(scratch)
	var decodeMem [2]*float32
	// Match both C do-while loops, including channel zero for CC==0.
	for c := int32(0); c < max(CC, 1); c++ {
		decodeMem[c] = prefilterFoldHistory(st1, overlap, c)
	}
	// Neither the zero-length comb filter nor fold consumes an audio cursor.
	// Avoid an unused one-past pointer for exact-sized histories when N==0.
	if overlap == 0 {
		return
	}
	for c := int32(0); c < max(CC, 1); c++ {
		input := celtNormAdd(decodeMem[c], DEC_PITCH_BUF_SIZE-N)
		// Filter the overlap before writing any TDAC samples, then reuse scratch
		// for the next channel. State controls remain live reads per channel.
		Opus_comb_filter(tls, etmp, input, st1.Fpostfilter_period_old, st1.Fpostfilter_period, overlap, -st1.Fpostfilter_gain_old, -st1.Fpostfilter_gain, st1.Fpostfilter_tapset_old, st1.Fpostfilter_tapset, nil, 0, st1.Farch)
		prefilterFoldTDAC(mode, input, etmp, overlap)
	}
}

func celtPLCFIRStorage(length int32) []float32 { return make([]float32, length) }

func celtPLCExcitationStorage(period int32) []float32 { return make([]float32, period+CELT_LPC_ORDER) }

func celtPLCNoiseStorage(N, channels int32) []float32 { return make([]float32, N*channels) }

func celtPLCNoise(tls *libc.TLS, state *OpusT_OpusCustomDecoder, eBands *int16, spectrum *float32, N, start, end, LM, channels int32) {
	seed := state.Frng
	if channels > 0 && start < end {
		bands := unsafe.Slice(eBands, end+1)
		x := unsafe.Slice(spectrum, N*channels)
		for c := int32(0); c < channels; c++ {
			for i := start; i < end; i++ {
				offset := N*c + int32(bands[i])<<LM
				length := (int32(bands[i+1]) - int32(bands[i])) << LM
				for j := int32(0); j < length; j++ {
					seed = Opus_celt_lcg_rand(tls, seed)
					x[offset+j] = float32(int32(seed) >> 20)
				}
				if length > 0 {
					Opus_renormalise_vector(tls, &x[offset], length, 1, state.Farch)
				}
			}
		}
	}
	state.Frng = seed
}

func celtPLCFinish(state *OpusT_OpusCustomDecoder, loss, LM, frameType int32) {
	state.Floss_duration = min(int32(10000), loss+int32(1)<<LM)
	state.Fplc_duration = min(int32(10000), state.Fplc_duration+int32(1)<<LM)
	state.Flast_frame_type = frameType
}

func celtPLCLPCHistory(memory *[CELT_LPC_ORDER]float32, history *float32, size, N int32) {
	h := unsafe.Slice(history, size)
	for i := int32(0); i < CELT_LPC_ORDER; i++ {
		memory[i] = h[size-N-1-i]
	}
}

func celtPLCExtrapolate(history, exc *float32, size, period, N, overlap, pitch int32, fade, decay float32) float32 {
	length := N + overlap
	if length <= 0 {
		return 0
	}
	h, x := unsafe.Slice(history, size+overlap), unsafe.Slice(exc, period)
	offset := period - pitch
	attenuation := float32(fade * decay)
	energy := float32(0)
	j := int32(0)
	for i := int32(0); i < length; i++ {
		if j >= pitch {
			j -= pitch
			attenuation = float32(attenuation * decay)
		}
		h[size-N+i] = float32(attenuation * x[offset+j])
		sample := h[size-period-N+offset+j]
		energy += float32(sample * sample)
		j++
	}
	return energy
}

// Forward loads/stores deliberately preserve the original loop's alias order.
func celtPLCExcitationHistory(exc, history *float32, size, period int32) {
	x, h := unsafe.Slice(exc, period+CELT_LPC_ORDER), unsafe.Slice(history, size)
	for i := int32(0); i < period+CELT_LPC_ORDER; i++ {
		x[i] = h[size-period-CELT_LPC_ORDER+i]
	}
}

func celtPLCSynthesisAttenuate(output, window *float32, length, overlap int32, s1 float32) {
	if length <= 0 {
		return
	}
	x := unsafe.Slice(output, length)
	s2 := float32(0)
	for _, sample := range x {
		s2 += float32(sample * sample)
	}
	if !(s1 > float32(float32(.2)*s2)) {
		clear(x)
	} else if s1 < s2 {
		ratio := float32(math.Sqrt(float64((s1 + float32(1)) / (s2 + float32(1)))))
		w := unsafe.Slice(window, overlap)
		for i := int32(0); i < overlap; i++ {
			g := float32(1) - float32(w[i]*(float32(1)-ratio))
			x[i] = float32(g * x[i])
		}
		for i := overlap; i < length; i++ {
			x[i] = float32(ratio * x[i])
		}
	}
}

func celtPLCExcitationDecay(exc *float32, period, length int32) float32 {
	half := length >> 1
	if half <= 0 {
		return 1
	}
	x := unsafe.Slice(exc, period)
	e1, e2 := float32(1), float32(1)
	for i := int32(0); i < half; i++ {
		e := x[period-half+i]
		e1 += float32(e * e)
		e = x[period-2*half+i]
		e2 += float32(e * e)
	}
	if !(e1 < e2) {
		e1 = e2
	}
	return float32(math.Sqrt(float64(e1 / e2)))
}

func celtPLCLagWindow(ac *[CELT_LPC_ORDER + 1]float32) {
	ac[0] *= float32(1.0001)
	for i := int32(1); i <= CELT_LPC_ORDER; i++ {
		term := float32(float32(float32(ac[i]*float32(float32(.008)*float32(.008)))*float32(i)) * float32(i))
		ac[i] -= term
	}
}

// celtPLCDecay retains MAXG's ordered live loads (including NaN selection)
// and the C do/while channel count. The caller still owns legacy storage.
func celtPLCDecay(energy, background *float32, bands, start, end, channels, loss int32) {
	if start >= end {
		return
	}
	a, b := unsafe.Slice(energy, max(channels, 1)*bands), unsafe.Slice(background, max(channels, 1)*bands)
	decay := float32(0.5)
	if loss == 0 {
		decay = 1.5
	}
	for c := int32(0); c < max(channels, 1); c++ {
		for i := start; i < end; i++ {
			index := c*bands + i
			if b[index] > a[index]-decay {
				a[index] = b[index]
			} else {
				a[index] = a[index] - decay
			}
		}
	}
}

func celtPLCMode(state *OpusT_OpusCustomDecoder) (mode *OpusT_OpusCustomMode, bands, overlap int32, eBands *int16) {
	mode = state.Fmode
	bands, overlap, eBands = mode.FnbEBands, mode.Foverlap, mode.FeBands
	return
}

func celtPLCDispatch(state *OpusT_OpusCustomDecoder) (loss, start, frameType int32) {
	loss, start = state.Floss_duration, state.Fstart
	frameType = FRAME_PLC_PERIODIC
	if state.Fplc_duration >= 40 || start != 0 || state.Fskip_plc != 0 {
		frameType = FRAME_PLC_NOISE
	}
	return
}

func celt_decode_lost(tls *libc.TLS, st1 *OpusT_OpusCustomDecoder, N int32, LM int32) {
	base := uintptr(unsafe.Pointer(st1)) // Remaining legacy trailing-storage boundary.
	// All concealment scratch is Go-owned. Decoder/mode/history uintptr views
	// remain legacy; this is not yet a whole-path checkptr boundary.
	var C, c, curr_frame_type, curr_neural, decode_buffer_size, effEnd, end, exc_length, extrapolation_len, last_neural, loss_duration, max_period, nbEBands, overlap, pitch_index, start, v5, v7, v8 int32
	var S1 float32
	var X, _exc, exc, fir_tmp []float32
	var mode *OpusT_OpusCustomMode
	var eBands *int16
	var backgroundLogE, buf, lpc, oldBandE, oldLogE, oldLogE2, window uintptr
	var decay1, fade float32
	var ac [25]float32
	var decode_mem, out_syn [2]uintptr
	var lpc_mem [24]float32
	C = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	max_period = MAX_PERIOD
	mode, nbEBands, overlap, eBands = celtPLCMode(st1)
	c = 0
	for {
		decode_mem[c] = base + unsafe.Offsetof(OpusT_OpusCustomDecoder{}.F_decode_mem) + uintptr(c*(decode_buffer_size+overlap))*4
		out_syn[c] = decode_mem[c] + uintptr(decode_buffer_size)*4 - uintptr(N)*4
		c = c + 1
		v5 = c
		if !(v5 < C) {
			break
		}
	}
	oldBandE = base + unsafe.Offsetof(OpusT_OpusCustomDecoder{}.F_decode_mem) + uintptr((decode_buffer_size+overlap)*C)*4
	oldLogE = oldBandE + uintptr(int32(2)*nbEBands)*4
	oldLogE2 = oldLogE + uintptr(int32(2)*nbEBands)*4
	backgroundLogE = oldLogE2 + uintptr(int32(2)*nbEBands)*4
	lpc = backgroundLogE + uintptr(int32(2)*nbEBands)*4
	loss_duration, start, curr_frame_type = celtPLCDispatch(st1)
	if curr_frame_type == int32(FRAME_PLC_NOISE) {
		end = st1.Fend
		if end < mode.FeffEBands {
			v7 = end
		} else {
			v7 = mode.FeffEBands
		}
		if start > v7 {
			v5 = start
		} else {
			if end < mode.FeffEBands {
				v8 = end
			} else {
				v8 = mode.FeffEBands
			}
			v5 = v8
		}
		effEnd = v5
		X = celtPLCNoiseStorage(N, C)
		c = 0
		for {
			libc.Xmemmove(tls, decode_mem[c], decode_mem[c]+uintptr(N)*4, uint64(uint32(decode_buffer_size-N+overlap))*uint64(4)+uint64(0*((int64(decode_mem[c])-int64(decode_mem[c]+uintptr(N)*4))/4)))
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold != 0 {
			prefilter_and_fold(tls, st1, N)
		}
		celtPLCDecay((*float32)(unsafe.Pointer(oldBandE)), (*float32)(unsafe.Pointer(backgroundLogE)), nbEBands, start, end, C, loss_duration)
		celtPLCNoise(tls, st1, eBands, unsafe.SliceData(X), N, start, effEnd, LM, C)
		var outputs [2]*float32
		for channel := int32(0); channel < max(C, 1); channel++ {
			outputs[channel] = (*float32)(unsafe.Pointer(out_syn[channel]))
		}
		celt_synthesis(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), unsafe.SliceData(X), &outputs[0], (*float32)(unsafe.Pointer(oldBandE)), start, effEnd, C, C, 0, LM, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		/* Run the postfilter with the last parameters. */
		c = 0
		for {
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period > int32(COMBFILTER_MINPERIOD) {
				v7 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
			} else {
				v7 = int32(COMBFILTER_MINPERIOD)
			}
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period = v7
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old > int32(COMBFILTER_MINPERIOD) {
				v5 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old
			} else {
				v5 = int32(COMBFILTER_MINPERIOD)
			}
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = v5
			comb_filter_legacy(tls, out_syn[c], out_syn[c], (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			if LM != 0 {
				comb_filter_legacy(tls, out_syn[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, out_syn[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, N-(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			}
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold = 0
		/* Skip regular PLC until we get two consecutive packets. */
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fskip_plc = int32(1)
	} else {
		fade = float32(1)
		curr_neural = libc.BoolInt32(curr_frame_type == int32(FRAME_PLC_NEURAL) || curr_frame_type == int32(FRAME_DRED))
		last_neural = libc.BoolInt32((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type == int32(FRAME_PLC_NEURAL) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type == int32(FRAME_DRED))
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type != int32(FRAME_PLC_PERIODIC) && !(last_neural != 0 && curr_neural != 0) {
			v5 = celt_plc_pitch_search(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*float32)(unsafe.Pointer(decode_mem[0])), (*float32)(unsafe.Pointer(decode_mem[1])), C, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			pitch_index = v5
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_pitch_index = v5
		} else {
			pitch_index = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_pitch_index
			fade = float32(0.8)
		}
		/* We want the excitation for 2 pitch periods in order to look for a
		   decaying signal, but we can't get more than MAX_PERIOD. */
		if int32(2)*pitch_index < max_period {
			v5 = int32(2) * pitch_index
		} else {
			v5 = max_period
		}
		exc_length = v5
		_exc = celtPLCExcitationStorage(max_period)
		fir_tmp = celtPLCFIRStorage(exc_length)
		exc = _exc[CELT_LPC_ORDER:]
		window = uintptr(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow))
		c = 0
		for {
			S1 = float32(0)
			buf = decode_mem[c]
			celtPLCExcitationHistory(unsafe.SliceData(_exc), (*float32)(unsafe.Pointer(buf)), decode_buffer_size, max_period)
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type != int32(FRAME_PLC_PERIODIC) && !(last_neural != 0 && curr_neural != 0) {
				/* Compute LPC coefficients for the last MAX_PERIOD samples before
				   the first loss so we can work in the excitation-filter domain. */
				Opus__celt_autocorr(tls, unsafe.SliceData(exc), &ac[0], (*float32)(unsafe.Pointer(window)), overlap, CELT_LPC_ORDER, max_period, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
				// Noise floor followed by rounded lag-window products.
				celtPLCLagWindow(&ac)
				Opus__celt_lpc(tls, (*OpusT_opus_val16)(unsafe.Pointer(lpc+uintptr(c*int32(CELT_LPC_ORDER))*4)), &ac[0], int32(CELT_LPC_ORDER))
			}
			/* Initialize the LPC history with the samples just before the start
			   of the region for which we're computing the excitation. */
			/* Compute the excitation for exc_length samples before the loss. We need the copy
			   because celt_fir() cannot filter in-place. */
			Opus_celt_fir_c(tls, unsafe.SliceData(exc[max_period-exc_length:]), (*float32)(unsafe.Pointer(lpc+uintptr(c*CELT_LPC_ORDER)*4)), unsafe.SliceData(fir_tmp), exc_length, CELT_LPC_ORDER, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			copy(exc[max_period-exc_length:], fir_tmp)
			/* Check if the waveform is decaying, and if so how fast.
			   We do this to avoid adding energy when concealing in a segment
			   with decaying energy. */
			decay1 = celtPLCExcitationDecay(unsafe.SliceData(exc), max_period, exc_length)
			/* Move the decoder memory one frame to the left to give us room to
			   add the data for the new frame. We ignore the overlap that extends
			   past the end of the buffer, because we aren't going to use it. */
			libc.Xmemmove(tls, buf, buf+uintptr(N)*4, uint64(uint32(decode_buffer_size-N))*uint64(4)+uint64(0*((int64(buf)-int64(buf+uintptr(N)*4))/4)))
			/* Extrapolate from the end of the excitation with a period of
			   "pitch_index", scaling down each period by an additional factor of
			   "decay". */
			extrapolation_len = N + overlap
			S1 = celtPLCExtrapolate((*float32)(unsafe.Pointer(buf)), unsafe.SliceData(exc), decode_buffer_size, max_period, N, overlap, pitch_index, fade, decay1)
			/* Copy the last decoded samples (prior to the overlap region) to
			   synthesis filter memory so we can have a continuous signal. */
			celtPLCLPCHistory(&lpc_mem, (*float32)(unsafe.Pointer(buf)), decode_buffer_size, N)
			/* Apply the synthesis filter to convert the excitation back into
			   the signal domain. */
			Opus_celt_iir(tls, (*float32)(unsafe.Pointer(buf+uintptr(decode_buffer_size)*4-uintptr(N)*4)), (*float32)(unsafe.Pointer(lpc+uintptr(c*int32(CELT_LPC_ORDER))*4)), (*float32)(unsafe.Pointer(buf+uintptr(decode_buffer_size)*4-uintptr(N)*4)), extrapolation_len, int32(CELT_LPC_ORDER), &lpc_mem[0], (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			/* Check if the synthesis energy is higher than expected, which can
			   happen with the signal changes during our window. If so,
			   attenuate. */
			celtPLCSynthesisAttenuate((*float32)(unsafe.Pointer(buf+uintptr(decode_buffer_size-N)*4)), (*float32)(unsafe.Pointer(window)), extrapolation_len, overlap, S1)
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold = int32(1)
	}
	/* Saturate duration counters, then commit the frame type. */
	celtPLCFinish((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), loss_duration, LM, curr_frame_type)
}

func Opus_celt_decode_with_ec_dred(tls *libc.TLS, st1 uintptr, data uintptr, len1 int32, pcm uintptr, frame_size int32, dec uintptr, accum int32) (r int32) {
	var C, CC, LM, M, N, alloc_trim, anti_collapse_on, anti_collapse_rsv, boost, c, codedBands, decode_buffer_size, dynalloc_logp, dynalloc_loop_logp, effEnd, end, flag, i, intra_ener, isTransient, missing, nbEBands, octave, overlap, postfilter_pitch, postfilter_tapset, qg, quanta, shortBlocks, silence, spread_decision, start, width, v28, v37, v40 int32
	var E0, E1, E2, slope, v57 OpusT_opus_val32
	var X, _saved_stack, backgroundLogE, cap1, collapse_masks, eBands, fine_priority, fine_quant, mode, offsets, oldBandE, oldLogE, oldLogE2, pulses, st, tf_res, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var bits, tell, total_bits OpusT_opus_int32
	var decode_mem [2]uintptr
	var max_background_increase, safety, v35, v56, v61 OpusT_celt_glog
	var postfilter_gain OpusT_opus_val16
	var v60 float32
	var _dec OpusT_ec_dec
	var balance OpusT_opus_int32
	var dual_stereo int32
	var intensity int32
	var out_syn [2]uintptr
	CC = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
	intensity = 0
	dual_stereo = 0
	anti_collapse_on = 0
	C = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstream_channels
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v8 = st
	if (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v8)).Fglobal_stack == uintptr(0) {
		v13 = libc.Xmalloc(tls, uint64(GLOBAL_STACK_SIZE))
		v11 = v13
		v10 = v11
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v15 = libc.Xmalloc(tls, uint64(16))
			st = v15
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v17 = st
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fscratch_ptr = v10
		v5 = v10
	} else {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v19 = libc.Xmalloc(tls, uint64(16))
			st = v19
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v21 = st
		v5 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack
	}
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = v5
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	Opus_validate_celt_decoder(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)))
	mode = uintptr(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fmode))
	nbEBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands
	overlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	eBands = uintptr(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeBands))
	start = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstart
	end = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fend
	frame_size = frame_size * (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
	oldBandE = st1 + unsafe.Offsetof(OpusT_OpusCustomDecoder{}.F_decode_mem) + uintptr((decode_buffer_size+overlap)*CC)*4
	oldLogE = oldBandE + uintptr(int32(2)*nbEBands)*4
	oldLogE2 = oldLogE + uintptr(int32(2)*nbEBands)*4
	backgroundLogE = oldLogE2 + uintptr(int32(2)*nbEBands)*4
	LM = 0
	for {
		if !(LM <= (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM) {
			break
		}
		if (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize<<LM == frame_size {
			break
		}
		LM = LM + 1
	}
	if LM > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM {
		return -int32(1)
	}
	M = int32(1) << LM
	if len1 < 0 || len1 > int32(1275) || pcm == uintptr(uint32(0)) {
		return -int32(1)
	}
	N = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize
	c = 0
	for {
		decode_mem[c] = st1 + unsafe.Offsetof(OpusT_OpusCustomDecoder{}.F_decode_mem) + uintptr(c*(decode_buffer_size+overlap))*4
		out_syn[c] = decode_mem[c] + uintptr(decode_buffer_size)*4 - uintptr(N)*4
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	effEnd = end
	if effEnd > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands {
		effEnd = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands
	}
	if data == uintptr(uint32(0)) || len1 <= int32(1) {
		celt_decode_lost(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), N, LM)
		deemphasis_legacy(tls, uintptr(unsafe.Pointer(&out_syn[0])), pcm, N, CC, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, mode+16, st1+unsafe.Offsetof(OpusT_OpusCustomDecoder{}.Fpreemph_memD), accum)
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return frame_size / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
	}
	/* Check if there are at least two packets received consecutively before
	 * turning on the pitch-based PLC */
	if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration == 0 {
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fskip_plc = 0
	}
	if dec == uintptr(uint32(0)) {
		Opus_ec_dec_init(tls, &_dec, (*byte)(unsafe.Pointer(data)), uint32(len1))
		dec = uintptr(unsafe.Pointer(&_dec))
	}
	if C == int32(1) {
		i = 0
		for {
			if !(i < nbEBands) {
				break
			}
			if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) > *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(nbEBands+i)*4)) {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4))
			} else {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(nbEBands+i)*4))
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) = v35
			i = i + 1
		}
	}
	total_bits = len1 * int32(8)
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	tell = v28
	if tell >= total_bits {
		silence = int32(1)
	} else {
		if tell == int32(1) {
			silence = Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(15))
		} else {
			silence = 0
		}
	}
	if silence != 0 {
		/* Pretend we've read all the remaining bits */
		tell = len1 * int32(8)
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		(*OpusT_ec_ctx)(unsafe.Pointer(dec)).Fnbits_total += tell - v28
	}
	postfilter_gain = float32(0)
	postfilter_pitch = 0
	postfilter_tapset = 0
	if start == 0 && tell+int32(16) <= total_bits {
		if Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(1)) != 0 {
			octave = int32(Opus_ec_dec_uint(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(6)))
			postfilter_pitch = int32(uint32(int32(16)<<octave) + Opus_ec_dec_bits(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(int32(4)+octave)) - uint32(1))
			qg = int32(Opus_ec_dec_bits(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(3)))
			v1 = dec
			v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
			if v28+int32(2) <= total_bits {
				postfilter_tapset = Opus_ec_dec_icdf(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), &tapset_icdf9[0], uint32(2))
			}
			postfilter_gain = OpusT_opus_val16(float32(0.09375) * float32(qg+int32(1)))
		}
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		tell = v28
	}
	if LM > 0 && tell+int32(3) <= total_bits {
		isTransient = Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(3))
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		tell = v28
	} else {
		isTransient = 0
	}
	if isTransient != 0 {
		shortBlocks = M
	} else {
		shortBlocks = 0
	}
	/* Decode the global flags (first symbols in the stream) */
	if tell+int32(3) <= total_bits {
		v28 = Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(3))
	} else {
		v28 = 0
	}
	intra_ener = v28
	/* If recovering from packet loss, make sure we make the energy prediction safe to reduce the
	   risk of getting loud artifacts. */
	if !(intra_ener != 0) && (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration != 0 {
		c = 0
		for {
			safety = float32(0)
			if int32(10) < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration>>LM {
				v37 = int32(10)
			} else {
				v37 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration >> LM
			}
			missing = v37
			if LM == 0 {
				safety = float32(1.5)
			} else {
				if LM == int32(1) {
					safety = float32(0.5)
				}
			}
			i = start
			for {
				if !(i < end) {
					break
				}
				if *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) > *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) {
					v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
				} else {
					v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4))
				}
				if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) < v35 {
					E0 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))
					E1 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
					E2 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4))
					if E1-E0 > float32(float32(0.5)*(E2-E0)) {
						v57 = E1 - E0
					} else {
						v57 = float32(float32(0.5) * (E2 - E0))
					}
					slope = v57
					if slope < float32(2) {
						v57 = slope
					} else {
						v57 = float32(2)
					}
					slope = v57
					if float32(int32(0)) > OpusT_opus_val32(float32(int32(1)+missing)*slope) {
						v57 = float32(int32(0))
					} else {
						v57 = OpusT_opus_val32(float32(int32(1)+missing) * slope)
					}
					E0 = E0 - v57
					if -float32(20) > E0 {
						v60 = -float32(20)
					} else {
						v60 = E0
					}
					*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = v60
				} else {
					/* Otherwise take the min of the last frames. */
					if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) {
						v56 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))
					} else {
						v56 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
					}
					if v56 < *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) {
						if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) {
							v61 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))
						} else {
							v61 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
						}
						v35 = v61
					} else {
						v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4))
					}
					*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = v35
				}
				/* Shorter frames have more natural fluctuations -- play it safe. */
				*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) -= safety
				i = i + 1
			}
			c = c + 1
			v28 = c
			if !(v28 < int32(2)) {
				break
			}
		}
	}
	/* Get band energies */
	Opus_unquant_coarse_energy(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), start, end, (*OpusT_celt_glog)(unsafe.Pointer(oldBandE)), intra_ener, (*OpusT_ec_dec)(unsafe.Pointer(dec)), C, LM)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1395))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	tf_res = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	tf_decode(tls, start, end, isTransient, (*int32)(unsafe.Pointer(tf_res)), LM, (*OpusT_ec_dec)(unsafe.Pointer(dec)))
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	tell = v28
	spread_decision = int32(SPREAD_NORMAL)
	if tell+int32(4) <= total_bits {
		spread_decision = Opus_ec_dec_icdf(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), &spread_icdf9[0], uint32(5))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1403))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	cap1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	capMode := (*OpusT_OpusCustomMode)(unsafe.Pointer(mode))
	Opus_init_caps(tls, capMode.FeBands, capMode.Fcache.Fcaps, (*int32)(unsafe.Pointer(cap1)), capMode.FnbEBands, LM, C)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1407))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	offsets = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	dynalloc_logp = int32(6)
	total_bits = total_bits << int32(BITRES)
	tell = int32(Opus_ec_tell_frac(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec))))
	i = start
	for {
		if !(i < end) {
			break
		}
		width = C * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))) << LM
		/* quanta is 6 bits, but no more than 1 bit/sample
		   and no less than 1/8 bit/sample */
		if int32(6)<<int32(BITRES) > width {
			v37 = int32(6) << int32(BITRES)
		} else {
			v37 = width
		}
		if width<<int32(BITRES) < v37 {
			v28 = width << int32(BITRES)
		} else {
			if int32(6)<<int32(BITRES) > width {
				v40 = int32(6) << int32(BITRES)
			} else {
				v40 = width
			}
			v28 = v40
		}
		quanta = v28
		dynalloc_loop_logp = dynalloc_logp
		boost = 0
		for tell+dynalloc_loop_logp<<int32(BITRES) < total_bits && boost < *(*int32)(unsafe.Pointer(cap1 + uintptr(i)*4)) {
			flag = Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(dynalloc_loop_logp))
			tell = int32(Opus_ec_tell_frac(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec))))
			if !(flag != 0) {
				break
			}
			boost = boost + quanta
			total_bits = total_bits - quanta
			dynalloc_loop_logp = int32(1)
		}
		*(*int32)(unsafe.Pointer(offsets + uintptr(i)*4)) = boost
		/* Making dynalloc more likely */
		if boost > 0 {
			if int32(2) > dynalloc_logp-int32(1) {
				v28 = int32(2)
			} else {
				v28 = dynalloc_logp - int32(1)
			}
			dynalloc_logp = v28
		}
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1440))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	fine_quant = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	if tell+int32(6)<<int32(BITRES) <= total_bits {
		v28 = Opus_ec_dec_icdf(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), &trim_icdf9[0], uint32(7))
	} else {
		v28 = int32(5)
	}
	alloc_trim = v28
	bits = len1*int32(8)<<int32(BITRES) - int32(Opus_ec_tell_frac(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)))) - int32(1)
	if isTransient != 0 && LM >= int32(2) && bits >= (LM+int32(2))<<int32(BITRES) {
		v28 = int32(1) << int32(BITRES)
	} else {
		v28 = 0
	}
	anti_collapse_rsv = v28
	bits = bits - anti_collapse_rsv
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1448))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	pulses = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1449))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	fine_priority = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	codedBands = Opus_clt_compute_allocation(tls, mode, start, end, offsets, cap1, alloc_trim, uintptr(unsafe.Pointer(&intensity)), uintptr(unsafe.Pointer(&dual_stereo)), bits, uintptr(unsafe.Pointer(&balance)), pulses, fine_quant, fine_priority, C, LM, dec, 0, 0, 0)
	Opus_unquant_fine_energy(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), start, end, (*OpusT_celt_glog)(unsafe.Pointer(oldBandE)), nil, (*int32)(unsafe.Pointer(fine_quant)), (*OpusT_ec_dec)(unsafe.Pointer(dec)), C)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(C*N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1457))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(C*N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(C*N))*(uint64(4)/uint64(1))) /**< Interleaved normalised MDCTs */
	c = 0
	for {
		libc.Xmemmove(tls, decode_mem[c], decode_mem[c]+uintptr(N)*4, uint64(uint32(decode_buffer_size-N+overlap))*uint64(4)+uint64(0*((int64(decode_mem[c])-int64(decode_mem[c]+uintptr(N)*4))/4)))
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	/* Decode fixed codebook */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(1)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(1)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	if !(int64(int32(uint64(uint32(C*nbEBands))*(uint64(1)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1487))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack += uintptr(uint64(uint32(C*nbEBands)) * (uint64(1) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	collapse_masks = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(C*nbEBands))*(uint64(1)/uint64(1)))
	if C == int32(2) {
		v1 = X + uintptr(N)*4
	} else {
		v1 = uintptr(uint32(0))
	}
	Opus_quant_all_bands(tls, 0, mode, start, end, X, v1, collapse_masks, uintptr(uint32(0)), pulses, shortBlocks, spread_decision, dual_stereo, intensity, tf_res, len1*(int32(8)<<int32(BITRES))-anti_collapse_rsv, balance, dec, LM, codedBands, st1+unsafe.Offsetof(OpusT_OpusCustomDecoder{}.Frng), 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdisable_inv)
	if anti_collapse_rsv > 0 {
		anti_collapse_on = int32(Opus_ec_dec_bits(tls, (*OpusT_ec_dec)(unsafe.Pointer(dec)), uint32(1)))
	}
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	Opus_unquant_energy_finalise(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), start, end, (*OpusT_celt_glog)(unsafe.Pointer(oldBandE)), (*int32)(unsafe.Pointer(fine_quant)), (*int32)(unsafe.Pointer(fine_priority)), len1*int32(8)-v28, (*OpusT_ec_dec)(unsafe.Pointer(dec)), C)
	if anti_collapse_on != 0 {
		anti_collapse_legacy(tls, mode, X, collapse_masks, LM, C, N, start, end, oldBandE, oldLogE, oldLogE2, pulses, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
	}
	if silence != 0 {
		i = 0
		for {
			if !(i < C*nbEBands) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) = -float32(28)
			i = i + 1
		}
	}
	if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold != 0 {
		prefilter_and_fold_legacy(tls, st1, N)
	}
	celt_synthesis_legacy(tls, mode, X, uintptr(unsafe.Pointer(&out_syn[0])), oldBandE, start, effEnd, C, CC, isTransient, LM, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, silence, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
	c = 0
	for {
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period > int32(COMBFILTER_MINPERIOD) {
			v37 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
		} else {
			v37 = int32(COMBFILTER_MINPERIOD)
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period = v37
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old > int32(COMBFILTER_MINPERIOD) {
			v28 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old
		} else {
			v28 = int32(COMBFILTER_MINPERIOD)
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = v28
		comb_filter_legacy(tls, out_syn[c], out_syn[c], (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		if LM != 0 {
			comb_filter_legacy(tls, out_syn[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, out_syn[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, postfilter_pitch, N-(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, postfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, postfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		}
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period = postfilter_pitch
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain = postfilter_gain
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset = postfilter_tapset
	if LM != 0 {
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset
	}
	if C == int32(1) {
		libc.Xmemcpy(tls, oldBandE+uintptr(nbEBands)*4, oldBandE, uint64(uint32(nbEBands))*uint64(4)+uint64(0*((OpusT___predefined_ptrdiff_t(oldBandE+uintptr(nbEBands)*4)-int64(oldBandE))/4)))
	}
	if !(isTransient != 0) {
		libc.Xmemcpy(tls, oldLogE2, oldLogE, uint64(uint32(int32(2)*nbEBands))*uint64(4)+uint64(0*((int64(oldLogE2)-int64(oldLogE))/4)))
		libc.Xmemcpy(tls, oldLogE, oldBandE, uint64(uint32(int32(2)*nbEBands))*uint64(4)+uint64(0*((int64(oldLogE)-int64(oldBandE))/4)))
	} else {
		i = 0
		for {
			if !(i < int32(2)*nbEBands) {
				break
			}
			if *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4))
			} else {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4))
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4)) = v35
			i = i + 1
		}
	}
	/* In normal circumstances, we only allow the noise floor to increase by
	   up to 2.4 dB/second, but when we're in DTX we give the weight of
	   all missing packets to the update packet. */
	if int32(160) < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration+M {
		v28 = int32(160)
	} else {
		v28 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration + M
	}
	max_background_increase = OpusT_celt_glog(float32(v28) * float32(0.001))
	i = 0
	for {
		if !(i < int32(2)*nbEBands) {
			break
		}
		if *(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(i)*4))+max_background_increase < *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) {
			v35 = *(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(i)*4)) + max_background_increase
		} else {
			v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4))
		}
		*(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(i)*4)) = v35
		i = i + 1
	}
	/* In case start or end were to change */
	c = 0
	for {
		i = 0
		for {
			if !(i < start) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = float32(0)
			v35 = -float32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) = v35
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) = v35
			i = i + 1
		}
		i = end
		for {
			if !(i < nbEBands) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = float32(0)
			v35 = -float32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) = v35
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) = v35
			i = i + 1
		}
		c = c + 1
		v28 = c
		if !(v28 < int32(2)) {
			break
		}
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng = (*OpusT_ec_dec)(unsafe.Pointer(dec)).Frng
	deemphasis_legacy(tls, uintptr(unsafe.Pointer(&out_syn[0])), pcm, N, CC, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, mode+16, st1+unsafe.Offsetof(OpusT_OpusCustomDecoder{}.Fpreemph_memD), accum)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration = 0
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fplc_duration = 0
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type = int32(FRAME_NORMAL)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold = 0
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	if v28 > int32(8)*len1 {
		return -int32(3)
	}
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(dec)).Ferror1
	if v28 != 0 {
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Ferror1 = int32(1)
	}
	return frame_size / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
}

func Opus_celt_decode_with_ec(tls *libc.TLS, st uintptr, data uintptr, len1 int32, pcm uintptr, frame_size int32, dec uintptr, accum int32) (r int32) {
	return Opus_celt_decode_with_ec_dred(tls, st, data, len1, pcm, frame_size, dec, accum)
}
