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

func opusFrameCeltSilence(tls *libc.TLS, state *OpusT_OpusCustomDecoder, silence *[2]byte, pcm *float32, frame, accum int32) int32 {
	return celt_decode_with_ec_dred(tls, state, &silence[0], 2, pcm, frame, nil, accum)
}

func opusFrameRedundantPacket(data *byte, offset, length int32) *byte {
	if length <= 1 {
		return nil
	}
	return unsafe.SliceData(unsafe.Slice(data, offset+length)[offset : offset+length])
}
func opusFrameCeltRedundant(tls *libc.TLS, state *OpusT_OpusCustomDecoder, data *byte, offset, length int32, pcm *float32, frame int32) int32 {
	return celt_decode_with_ec_dred(tls, state, opusFrameRedundantPacket(data, offset, length), length, pcm, frame, nil, 0)
}

func opusFrameCelt(tls *libc.TLS, state *OpusT_OpusCustomDecoder, data *byte, length int32, pcm *float32, frame int32, ec *OpusT_ec_ctx, fec, accum int32) int32 {
	if fec != 0 {
		data = nil
	}
	return celt_decode_with_ec_dred(tls, state, data, length, pcm, frame, ec, accum)
}

func celtDecodeFrameLM(mode *OpusT_OpusCustomMode, frameSize int32) int32 {
	LM := int32(0)
	for LM <= mode.FmaxLM {
		if mode.FshortMdctSize<<LM == frameSize {
			break
		}
		LM++
	}
	if LM > mode.FmaxLM {
		return -1
	}
	return LM
}

func celtDecodePacketArguments(pcm *float32, length int32) bool {
	return length >= 0 && length <= 1275 && pcm != nil
}

func celtDecodePacketLost(data *byte, length int32) bool { return data == nil || length <= 1 }

func celtDecodePacketStart(state *OpusT_OpusCustomDecoder) {
	if state.Floss_duration == 0 {
		state.Fskip_plc = 0
	}
}

func celtDecodeHistoryViews(state *OpusT_OpusCustomDecoder, overlap, channels, N int32) (history [2][]float32, output [2]*float32) {
	count := max(int32(1), channels)
	stride := DEC_PITCH_BUF_SIZE + overlap
	memory := unsafe.Slice(&state.F_decode_mem[0], count*stride)
	for c := int32(0); c < count; c++ {
		history[c] = memory[c*stride : (c+1)*stride]
		if N > 0 {
			output[c] = &history[c][DEC_PITCH_BUF_SIZE-N]
		}
	}
	return
}

func celtDecodeEnergyViews(state *OpusT_OpusCustomDecoder, bands, overlap, channels int32) (energy, log, previous, background *float32) {
	if bands == 0 {
		return
	}
	offset := (DEC_PITCH_BUF_SIZE + overlap) * channels
	memory := unsafe.Slice(&state.F_decode_mem[0], offset+8*bands)
	energy = &memory[offset]
	log = &memory[offset+2*bands]
	previous = &memory[offset+4*bands]
	background = &memory[offset+6*bands]
	return
}

func celtDecodeEntropy(tls *libc.TLS, provided, local *OpusT_ec_ctx, data *byte, length int32) *OpusT_ec_ctx {
	if provided != nil {
		return provided
	}
	Opus_ec_dec_init(tls, local, data, uint32(length))
	return local
}

func celtDecodeMode(state *OpusT_OpusCustomDecoder) (mode *OpusT_OpusCustomMode, bands, overlap int32, boundaries *int16) {
	mode = state.Fmode
	bands = mode.FnbEBands
	overlap = mode.Foverlap
	boundaries = mode.FeBands
	return
}

func celtDecodeAntiCollapseBit(tls *libc.TLS, ec *OpusT_ec_ctx, reserved int32) int32 {
	if reserved > 0 {
		return int32(Opus_ec_dec_bits(tls, ec, 1))
	}
	return 0
}
func celtDecodeAntiCollapse(tls *libc.TLS, state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, spectrum *float32, masks *byte, pulses *int32, energy, log, previous *float32, N, LM, channels, start, end, on int32) {
	if on != 0 {
		Opus_anti_collapse(tls, mode.FeBands, mode.FnbEBands, spectrum, masks, LM, channels, N, start, end, energy, log, previous, pulses, state.Frng, 0, state.Farch)
	}
}

func celtDecodeFinalEnergy(tls *libc.TLS, mode *OpusT_OpusCustomMode, energy *float32, fine, priority *int32, start, end, length, channels int32, ec *OpusT_ec_ctx) {
	tell := ec.Fnbits_total - int32(bits.Len32(ec.Frng))
	Opus_unquant_energy_finalise(tls, mode, start, end, energy, fine, priority, length*8-tell, ec, channels)
}

func celtDecodeAllocationBudget(tls *libc.TLS, ec *OpusT_ec_ctx, length, transient, LM int32) (budget, reserved int32) {
	budget = (length * 8) << BITRES
	budget -= int32(Opus_ec_tell_frac(tls, ec))
	budget -= 1
	if transient != 0 && LM >= 2 && budget >= (LM+2)<<BITRES {
		reserved = 1 << BITRES
	}
	budget -= reserved
	return
}

func celtDecodeTrim(tls *libc.TLS, ec *OpusT_ec_ctx, tell, total int32) int32 {
	if tell+(int32(6)<<BITRES) <= total {
		return Opus_ec_dec_icdf(tls, ec, &trim_icdf9[0], 7)
	}
	return 5
}

func celtDecodeSpread(tls *libc.TLS, ec *OpusT_ec_ctx, total int32) (spread, tell int32) {
	tell = ec.Fnbits_total - int32(bits.Len32(ec.Frng))
	spread = SPREAD_NORMAL
	if tell+4 <= total {
		spread = Opus_ec_dec_icdf(tls, ec, &spread_icdf9[0], 5)
	}
	return
}

func celtDecodeGlobalFlags(tls *libc.TLS, ec *OpusT_ec_ctx, LM, M, total, tell int32) (transient, short, intra, updatedTell int32) {
	updatedTell = tell
	if LM > 0 && tell+3 <= total {
		transient = Opus_ec_dec_bit_logp(tls, ec, 3)
		updatedTell = ec.Fnbits_total - int32(bits.Len32(ec.Frng))
	}
	if transient != 0 {
		short = M
	}
	if updatedTell+3 <= total {
		intra = Opus_ec_dec_bit_logp(tls, ec, 3)
	}
	return
}

func celtDecodePostfilterHeader(tls *libc.TLS, ec *OpusT_ec_ctx, start, total, tell int32) (pitch int32, gain float32, tapset, updatedTell int32) {
	updatedTell = tell
	if start == 0 && tell+16 <= total {
		if Opus_ec_dec_bit_logp(tls, ec, 1) != 0 {
			octave := int32(Opus_ec_dec_uint(tls, ec, 6))
			pitch = int32(uint32(int32(16)<<octave) + Opus_ec_dec_bits(tls, ec, uint32(4+octave)) - 1)
			qg := int32(Opus_ec_dec_bits(tls, ec, 3))
			if ec.Fnbits_total-int32(bits.Len32(ec.Frng))+2 <= total {
				tapset = Opus_ec_dec_icdf(tls, ec, &tapset_icdf9[0], 2)
			}
			gain = float32(.09375 * float32(qg+1))
		}
		updatedTell = ec.Fnbits_total - int32(bits.Len32(ec.Frng))
	}
	return
}

func celtDecodeSilence(tls *libc.TLS, ec *OpusT_ec_ctx, total int32) (silence, tell int32) {
	tell = ec.Fnbits_total - int32(bits.Len32(ec.Frng))
	if tell >= total {
		silence = 1
	} else if tell == 1 {
		silence = Opus_ec_dec_bit_logp(tls, ec, 15)
	}
	if silence != 0 {
		tell = total
		ec.Fnbits_total += tell - (ec.Fnbits_total - int32(bits.Len32(ec.Frng)))
	}
	return
}

func celtDecodePacketError(state *OpusT_OpusCustomDecoder, ec *OpusT_ec_ctx, length int32) int32 {
	tell := ec.Fnbits_total - int32(bits.Len32(ec.Frng))
	if tell > 8*length {
		return -3
	}
	if ec.Ferror1 != 0 {
		state.Ferror1 = 1
	}
	return 0
}

func celtDecodeDeemphasis(tls *libc.TLS, state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, outputs **float32, pcm *float32, N, channels, accum int32) {
	deemphasis(tls, outputs, pcm, N, channels, state.Fdownsample, &mode.Fpreemph[0], &state.Fpreemph_memD[0], accum)
}

func celtDecodePrefilter(tls *libc.TLS, state *OpusT_OpusCustomDecoder, N int32) {
	if state.Fprefilter_and_fold != 0 {
		prefilter_and_fold(tls, state, N)
	}
}

func celtDecodeEnergyMergeMono(energy *float32, bands int32) {
	if bands <= 0 {
		return
	}
	e := unsafe.Slice(energy, 2*bands)
	for i := int32(0); i < bands; i++ {
		value := e[bands+i]
		if e[i] > value {
			value = e[i]
		}
		e[i] = value
	}
}

func celtDecodePostfilter(tls *libc.TLS, state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, outputs **float32, channels, N, LM, period int32, gain float32, tapset, overlap int32) {
	out := unsafe.Slice(outputs, max(int32(1), channels))
	for c := int32(0); ; c++ {
		celtDecodePostfilterClamp(state)
		celtDecodePostfilterFirst(tls, state, mode, out[c], overlap)
		if LM != 0 {
			celtDecodePostfilterTail(tls, state, mode, out[c], N, period, gain, tapset, overlap)
		}
		if c+1 >= channels {
			break
		}
	}
}

func celtDecodePostfilterTail(tls *libc.TLS, state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, output *float32, N, period int32, gain float32, tapset, overlap int32) {
	tail := &unsafe.Slice(output, N)[mode.FshortMdctSize]
	Opus_comb_filter(tls, tail, tail, state.Fpostfilter_period, period, N-mode.FshortMdctSize, state.Fpostfilter_gain, gain, state.Fpostfilter_tapset, tapset, mode.Fwindow, overlap, state.Farch)
}

func celtDecodePostfilterFirst(tls *libc.TLS, state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, output *float32, overlap int32) {
	Opus_comb_filter(tls, output, output, state.Fpostfilter_period_old, state.Fpostfilter_period, mode.FshortMdctSize, state.Fpostfilter_gain_old, state.Fpostfilter_gain, state.Fpostfilter_tapset_old, state.Fpostfilter_tapset, mode.Fwindow, overlap, state.Farch)
}

func celtDecodePostfilterClamp(state *OpusT_OpusCustomDecoder) {
	state.Fpostfilter_period = max(int32(COMBFILTER_MINPERIOD), state.Fpostfilter_period)
	state.Fpostfilter_period_old = max(int32(COMBFILTER_MINPERIOD), state.Fpostfilter_period_old)
}

func celtDecodePacketFinish(state *OpusT_OpusCustomDecoder) {
	state.Floss_duration = 0
	state.Fplc_duration = 0
	state.Flast_frame_type = FRAME_NORMAL
	state.Fprefilter_and_fold = 0
}

func celtDecodeRecoverEnergy(state *OpusT_OpusCustomDecoder, energy, log, previous *float32, bands, start, end, LM, intra int32) {
	if intra != 0 || state.Floss_duration == 0 {
		return
	}
	var e, l, p []float32
	if start < end {
		length := bands + end
		e = unsafe.Slice(energy, length)
		l = unsafe.Slice(log, length)
		p = unsafe.Slice(previous, length)
	}
	for c := int32(0); c < 2; c++ {
		missing, safety := celtDecodeRecoverySafety(state, LM)
		for i := start; i < end; i++ {
			index := c*bands + i
			celtDecodeRecoveryBand(&e[index], &l[index], &p[index], missing, safety)
		}
	}
}

func celtDecodeRecoveryBand(energy, log, previous *float32, missing int32, safety float32) {
	highest := *previous
	if *log > highest {
		highest = *log
	}
	if *energy < highest {
		E0, E1, E2 := *energy, *log, *previous
		delta := float32(E1 - E0)
		half := float32(.5 * float32(E2-E0))
		slope := half
		if delta > half {
			slope = delta
		}
		if !(slope < 2) {
			slope = 2
		}
		decrease := float32(float32(1+missing) * slope)
		if 0 > decrease {
			decrease = 0
		}
		E0 = float32(E0 - decrease)
		if -20 > E0 {
			E0 = -20
		}
		*energy = E0
	} else {
		lowest := *log
		if *energy < lowest {
			lowest = *energy
		}
		if !(lowest < *previous) {
			lowest = *previous
		}
		*energy = lowest
	}
	*energy = float32(*energy - safety)
}

func celtDecodeRecoverySafety(state *OpusT_OpusCustomDecoder, LM int32) (missing int32, safety float32) {
	missing = min(int32(10), state.Floss_duration>>LM)
	if LM == 0 {
		safety = 1.5
	} else if LM == 1 {
		safety = .5
	}
	return
}

func celtDecodePostfilterFinish(state *OpusT_OpusCustomDecoder, period int32, gain float32, tapset, LM int32) {
	state.Fpostfilter_period_old = state.Fpostfilter_period
	state.Fpostfilter_gain_old = state.Fpostfilter_gain
	state.Fpostfilter_tapset_old = state.Fpostfilter_tapset
	state.Fpostfilter_period = period
	state.Fpostfilter_gain = gain
	state.Fpostfilter_tapset = tapset
	if LM != 0 {
		state.Fpostfilter_period_old = state.Fpostfilter_period
		state.Fpostfilter_gain_old = state.Fpostfilter_gain
		state.Fpostfilter_tapset_old = state.Fpostfilter_tapset
	}
}

func celtDecodeBoosts(tls *libc.TLS, bands *int16, cap, offsets *int32, start, end, C, LM, total int32, ec *OpusT_ec_ctx) (remaining, tell int32) {
	remaining = total << BITRES
	tell = int32(Opus_ec_tell_frac(tls, ec))
	if start >= end {
		return
	}
	e := unsafe.Slice(bands, end+1)
	caps := unsafe.Slice(cap, end)
	out := unsafe.Slice(offsets, end)
	logp := int32(6)
	for i := start; i < end; i++ {
		width := C * (int32(e[i+1]) - int32(e[i])) << LM
		quanta := min(width<<BITRES, max(int32(6)<<BITRES, width))
		loopLogp := logp
		boost := int32(0)
		for tell+(loopLogp<<BITRES) < remaining && boost < caps[i] {
			flag := Opus_ec_dec_bit_logp(tls, ec, uint32(loopLogp))
			tell = int32(Opus_ec_tell_frac(tls, ec))
			if flag == 0 {
				break
			}
			boost += quanta
			remaining -= quanta
			loopLogp = 1
		}
		out[i] = boost
		if boost > 0 {
			logp = max(int32(2), logp-1)
		}
	}
	return
}

func celtDecodeSilenceEnergy(energy *float32, bands, channels int32) {
	length := bands * channels
	if length <= 0 {
		return
	}
	e := unsafe.Slice(energy, length)
	for i := range e {
		e[i] = -28
	}
}

func celtDecodeHistoryMove(history *float32, N, length int32) {
	if length == 0 {
		return
	}
	h := unsafe.Slice(history, N+length)
	copy(h[:length], h[N:])
}

func celtDecodeMaskStorage(bands, channels int32) []byte { return make([]byte, channels*bands) }

func celtDecodeSpectrumStorage(N, channels int32) []float32 { return make([]float32, channels*N) }
func celtDecodeSpectrumChannels(spectrum []float32, N, channels int32) (x, y *float32) {
	if N == 0 {
		return
	}
	x = unsafe.SliceData(spectrum)
	if channels == 2 {
		y = &spectrum[N]
	}
	return
}

func celtDecodePriorityStorage(bands int32) []int32 { return make([]int32, bands) }

func celtDecodePulseStorage(bands int32) []int32 { return make([]int32, bands) }

func celtDecodeFineStorage(bands int32) []int32 { return make([]int32, bands) }

func celtDecodeOffsetsStorage(bands int32) []int32 { return make([]int32, bands) }

func celtDecodeCapsStorage(tls *libc.TLS, mode *OpusT_OpusCustomMode, bands, LM, channels int32) []int32 {
	caps := make([]int32, bands)
	Opus_init_caps(tls, mode.FeBands, mode.Fcache.Fcaps, unsafe.SliceData(caps), mode.FnbEBands, LM, channels)
	return caps
}

func celtDecodeTFStorage(tls *libc.TLS, bands, start, end, transient, LM int32, ec *OpusT_ec_ctx) []int32 {
	flags := make([]int32, bands)
	tf_decode(tls, start, end, transient, unsafe.SliceData(flags), LM, ec)
	return flags
}

func celtDecodeEnergyClear(energy, log, previous *float32, bands, start, end int32) {
	if bands <= 0 || start <= 0 && end >= bands {
		return
	}
	e, l, p := unsafe.Slice(energy, 2*bands), unsafe.Slice(log, 2*bands), unsafe.Slice(previous, 2*bands)
	for c := int32(0); c < 2; c++ {
		for i := int32(0); i < start; i++ {
			index := c*bands + i
			e[index] = 0
			p[index] = -28
			l[index] = -28
		}
		for i := end; i < bands; i++ {
			index := c*bands + i
			e[index] = 0
			p[index] = -28
			l[index] = -28
		}
	}
}

func celtDecodeEnergyBackground(state *OpusT_OpusCustomDecoder, background, energy *float32, bands, M int32) {
	increase := float32(min(int32(160), state.Floss_duration+M)) * float32(.001)
	if bands <= 0 {
		return
	}
	b, e := unsafe.Slice(background, 2*bands), unsafe.Slice(energy, 2*bands)
	for i := range b {
		if b[i]+increase < e[i] {
			b[i] = b[i] + increase
		} else {
			b[i] = e[i]
		}
	}
}

func celtDecodeEnergyLogs(energy, log, previous *float32, bands, transient int32) {
	if bands <= 0 {
		return
	}
	e, l := unsafe.Slice(energy, 2*bands), unsafe.Slice(log, 2*bands)
	if transient == 0 {
		p := unsafe.Slice(previous, 2*bands)
		copy(p, l)
		copy(l, e)
	} else {
		for i := range l {
			if l[i] < e[i] {
				l[i] = l[i]
			} else {
				l[i] = e[i]
			}
		}
	}
}

func celtDecodeEnergyMono(energy *float32, bands int32) {
	if bands <= 0 {
		return
	}
	e := unsafe.Slice(energy, 2*bands)
	copy(e[bands:], e[:bands])
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

func celtPLCHistoryViews(state *OpusT_OpusCustomDecoder, overlap, bands, channels, N int32) (history [2][]float32, output [2]*float32, energy, background, coefficients *float32) {
	stride := int32(DEC_PITCH_BUF_SIZE) + overlap
	energyOffset := stride * channels
	total := max(stride*max(channels, 1), energyOffset+8*bands+max(channels, 1)*CELT_LPC_ORDER)
	storage := unsafe.Slice(&state.F_decode_mem[0], total)
	for c := int32(0); c < max(channels, 1); c++ {
		history[c] = storage[c*stride : (c+1)*stride]
		if N > 0 || overlap > 0 {
			output[c] = &history[c][DEC_PITCH_BUF_SIZE-N]
		}
	}
	energy = &storage[energyOffset]
	background = &storage[energyOffset+6*bands]
	coefficients = &storage[energyOffset+8*bands]
	return
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
	// Fully typed active concealment with Go-owned scratch. The outer decoder
	// ABI remains legacy; byte-backed decoder allocations still do not scan pointers.
	var C, c, curr_frame_type, curr_neural, decode_buffer_size, effEnd, end, exc_length, extrapolation_len, last_neural, loss_duration, max_period, nbEBands, overlap, pitch_index, start, v5, v7, v8 int32
	var S1 float32
	var X, _exc, exc, fir_tmp []float32
	var mode *OpusT_OpusCustomMode
	var eBands *int16
	var buf []float32
	var window *float32
	var decay1, fade float32
	var ac [25]float32
	var lpc_mem [24]float32
	C = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	max_period = MAX_PERIOD
	mode, nbEBands, overlap, eBands = celtPLCMode(st1)
	decode_mem, out_syn, oldBandE, backgroundLogE, coefficients := celtPLCHistoryViews(st1, overlap, nbEBands, C, N)
	lpc := unsafe.Slice(coefficients, max(C, 1)*CELT_LPC_ORDER)
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
			copy(decode_mem[c], decode_mem[c][N:])
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold != 0 {
			prefilter_and_fold(tls, st1, N)
		}
		celtPLCDecay(oldBandE, backgroundLogE, nbEBands, start, end, C, loss_duration)
		celtPLCNoise(tls, st1, eBands, unsafe.SliceData(X), N, start, effEnd, LM, C)
		celt_synthesis(tls, mode, unsafe.SliceData(X), &out_syn[0], oldBandE, start, effEnd, C, C, 0, LM, st1.Fdownsample, 0, st1.Farch)
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
			st1.Fpostfilter_period_old = v5
			Opus_comb_filter(tls, out_syn[c], out_syn[c], st1.Fpostfilter_period_old, st1.Fpostfilter_period, mode.FshortMdctSize, st1.Fpostfilter_gain_old, st1.Fpostfilter_gain, st1.Fpostfilter_tapset_old, st1.Fpostfilter_tapset, mode.Fwindow, overlap, st1.Farch)
			if LM != 0 {
				tail := &decode_mem[c][decode_buffer_size-N+mode.FshortMdctSize]
				Opus_comb_filter(tls, tail, tail, st1.Fpostfilter_period, st1.Fpostfilter_period, N-mode.FshortMdctSize, st1.Fpostfilter_gain, st1.Fpostfilter_gain, st1.Fpostfilter_tapset, st1.Fpostfilter_tapset, mode.Fwindow, overlap, st1.Farch)
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
			v5 = celt_plc_pitch_search(tls, st1, unsafe.SliceData(decode_mem[0]), unsafe.SliceData(decode_mem[1]), C, st1.Farch)
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
		window = mode.Fwindow
		c = 0
		for {
			S1 = float32(0)
			buf = decode_mem[c]
			celtPLCExcitationHistory(unsafe.SliceData(_exc), unsafe.SliceData(buf), decode_buffer_size, max_period)
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type != int32(FRAME_PLC_PERIODIC) && !(last_neural != 0 && curr_neural != 0) {
				/* Compute LPC coefficients for the last MAX_PERIOD samples before
				   the first loss so we can work in the excitation-filter domain. */
				Opus__celt_autocorr(tls, unsafe.SliceData(exc), &ac[0], (*float32)(unsafe.Pointer(window)), overlap, CELT_LPC_ORDER, max_period, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
				// Noise floor followed by rounded lag-window products.
				celtPLCLagWindow(&ac)
				Opus__celt_lpc(tls, &lpc[c*CELT_LPC_ORDER], &ac[0], CELT_LPC_ORDER)
			}
			/* Initialize the LPC history with the samples just before the start
			   of the region for which we're computing the excitation. */
			/* Compute the excitation for exc_length samples before the loss. We need the copy
			   because celt_fir() cannot filter in-place. */
			Opus_celt_fir_c(tls, unsafe.SliceData(exc[max_period-exc_length:]), &lpc[c*CELT_LPC_ORDER], unsafe.SliceData(fir_tmp), exc_length, CELT_LPC_ORDER, st1.Farch)
			copy(exc[max_period-exc_length:], fir_tmp)
			/* Check if the waveform is decaying, and if so how fast.
			   We do this to avoid adding energy when concealing in a segment
			   with decaying energy. */
			decay1 = celtPLCExcitationDecay(unsafe.SliceData(exc), max_period, exc_length)
			/* Move the decoder memory one frame to the left to give us room to
			   add the data for the new frame. We ignore the overlap that extends
			   past the end of the buffer, because we aren't going to use it. */
			copy(buf[:decode_buffer_size-N], buf[N:decode_buffer_size])
			/* Extrapolate from the end of the excitation with a period of
			   "pitch_index", scaling down each period by an additional factor of
			   "decay". */
			extrapolation_len = N + overlap
			S1 = celtPLCExtrapolate(unsafe.SliceData(buf), unsafe.SliceData(exc), decode_buffer_size, max_period, N, overlap, pitch_index, fade, decay1)
			/* Copy the last decoded samples (prior to the overlap region) to
			   synthesis filter memory so we can have a continuous signal. */
			celtPLCLPCHistory(&lpc_mem, unsafe.SliceData(buf), decode_buffer_size, N)
			/* Apply the synthesis filter to convert the excitation back into
			   the signal domain. */
			Opus_celt_iir(tls, &buf[decode_buffer_size-N], &lpc[c*CELT_LPC_ORDER], &buf[decode_buffer_size-N], extrapolation_len, CELT_LPC_ORDER, &lpc_mem[0], st1.Farch)
			/* Check if the synthesis energy is higher than expected, which can
			   happen with the signal changes during our window. If so,
			   attenuate. */
			celtPLCSynthesisAttenuate(&buf[decode_buffer_size-N], window, extrapolation_len, overlap, S1)
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

//go:uintptrescapes
func Opus_celt_decode_with_ec_dred(tls *libc.TLS, st1, data uintptr, len1 int32, pcm uintptr, frame_size int32, decAddress uintptr, accum int32) int32 {
	return celt_decode_with_ec_dred(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*byte)(unsafe.Pointer(data)), len1, (*float32)(unsafe.Pointer(pcm)), frame_size, (*OpusT_ec_ctx)(unsafe.Pointer(decAddress)), accum)
}

func celt_decode_with_ec_dred(tls *libc.TLS, st1 *OpusT_OpusCustomDecoder, data *byte, len1 int32, pcm *float32, frame_size int32, dec *OpusT_ec_ctx, accum int32) (r int32) {
	var C, CC, LM, M, N, alloc_trim, anti_collapse_on, anti_collapse_rsv, c, codedBands, decode_buffer_size, effEnd, end, intra_ener, isTransient, nbEBands, overlap, postfilter_pitch, postfilter_tapset, shortBlocks, silence, spread_decision, start, v28 int32
	var tf_res, cap1, offsets, fine_quant, pulses, fine_priority []int32
	var X []float32
	var collapse_masks []byte
	var eBands *int16
	var mode *OpusT_OpusCustomMode
	var backgroundLogE, oldBandE, oldLogE, oldLogE2 *float32
	var bits, tell, total_bits OpusT_opus_int32
	var decode_mem [2][]float32
	var postfilter_gain OpusT_opus_val16
	var _dec OpusT_ec_dec
	var balance OpusT_opus_int32
	var dual_stereo int32
	var intensity int32
	var out_syn [2]*float32
	CC = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
	intensity = 0
	dual_stereo = 0
	anti_collapse_on = 0
	C = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstream_channels
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	Opus_validate_celt_decoder(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)))
	mode, nbEBands, overlap, eBands = celtDecodeMode((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)))
	start = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstart
	end = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fend
	frame_size = frame_size * (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
	oldBandE, oldLogE, oldLogE2, backgroundLogE = celtDecodeEnergyViews((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), nbEBands, overlap, CC)
	LM = celtDecodeFrameLM(mode, frame_size)
	if LM < 0 {
		return -int32(1)
	}
	M = int32(1) << LM
	if !celtDecodePacketArguments(pcm, len1) {
		return -int32(1)
	}
	N = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize
	decode_mem, out_syn = celtDecodeHistoryViews((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), overlap, CC, N)
	effEnd = end
	if effEnd > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands {
		effEnd = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands
	}
	if celtDecodePacketLost(data, len1) {
		celt_decode_lost(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), N, LM)
		celtDecodeDeemphasis(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), &out_syn[0], (*float32)(unsafe.Pointer(pcm)), N, CC, accum)
		return frame_size / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
	}
	/* Check if there are at least two packets received consecutively before
	 * turning on the pitch-based PLC */
	celtDecodePacketStart(st1)
	dec = celtDecodeEntropy(tls, dec, &_dec, (*byte)(unsafe.Pointer(data)), len1)
	if C == 1 {
		celtDecodeEnergyMergeMono((*float32)(unsafe.Pointer(oldBandE)), nbEBands)
	}
	total_bits = len1 * int32(8)
	silence, tell = celtDecodeSilence(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), total_bits)
	postfilter_pitch, postfilter_gain, postfilter_tapset, tell = celtDecodePostfilterHeader(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), start, total_bits, tell)
	isTransient, shortBlocks, intra_ener, tell = celtDecodeGlobalFlags(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), LM, M, total_bits, tell)
	/* If recovering from packet loss, make sure we make the energy prediction safe to reduce the
	   risk of getting loud artifacts. */
	celtDecodeRecoverEnergy((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*float32)(unsafe.Pointer(oldBandE)), (*float32)(unsafe.Pointer(oldLogE)), (*float32)(unsafe.Pointer(oldLogE2)), nbEBands, start, end, LM, intra_ener)
	/* Get band energies */
	Opus_unquant_coarse_energy(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), start, end, (*float32)(unsafe.Pointer(oldBandE)), intra_ener, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), C, LM)
	tf_res = celtDecodeTFStorage(tls, nbEBands, start, end, isTransient, LM, (*OpusT_ec_ctx)(unsafe.Pointer(dec)))
	spread_decision, tell = celtDecodeSpread(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), total_bits)
	cap1 = celtDecodeCapsStorage(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), nbEBands, LM, C)
	offsets = celtDecodeOffsetsStorage(nbEBands)
	total_bits, tell = celtDecodeBoosts(tls, eBands, unsafe.SliceData(cap1), unsafe.SliceData(offsets), start, end, C, LM, total_bits, (*OpusT_ec_ctx)(unsafe.Pointer(dec)))
	fine_quant = celtDecodeFineStorage(nbEBands)
	alloc_trim = celtDecodeTrim(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), tell, total_bits)
	bits, anti_collapse_rsv = celtDecodeAllocationBudget(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), len1, isTransient, LM)
	pulses = celtDecodePulseStorage(nbEBands)
	fine_priority = celtDecodePriorityStorage(nbEBands)
	codedBands = clt_compute_allocation(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), start, end, unsafe.SliceData(offsets), unsafe.SliceData(cap1), alloc_trim, &intensity, &dual_stereo, bits, &balance, unsafe.SliceData(pulses), unsafe.SliceData(fine_quant), unsafe.SliceData(fine_priority), C, LM, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), 0, 0, 0)
	Opus_unquant_fine_energy(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), start, end, (*OpusT_celt_glog)(unsafe.Pointer(oldBandE)), nil, unsafe.SliceData(fine_quant), (*OpusT_ec_dec)(unsafe.Pointer(dec)), C)
	X = celtDecodeSpectrumStorage(N, C) // Contiguous per-channel normalized MDCT spectra.
	c = 0
	for {
		celtDecodeHistoryMove(unsafe.SliceData(decode_mem[c]), N, decode_buffer_size-N+overlap)
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	// Decode fixed codebook using Go-owned per-channel collapse flags.
	collapse_masks = celtDecodeMaskStorage(nbEBands, C)
	spectralX, spectralY := celtDecodeSpectrumChannels(X, N, C)
	quant_all_bands(tls, 0, mode, start, end, spectralX, spectralY, unsafe.SliceData(collapse_masks), nil, unsafe.SliceData(pulses), shortBlocks, spread_decision, dual_stereo, intensity, unsafe.SliceData(tf_res), len1*(int32(8)<<int32(BITRES))-anti_collapse_rsv, balance, dec, LM, codedBands, &st1.Frng, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdisable_inv)
	anti_collapse_on = celtDecodeAntiCollapseBit(tls, (*OpusT_ec_ctx)(unsafe.Pointer(dec)), anti_collapse_rsv)
	celtDecodeFinalEnergy(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), (*float32)(unsafe.Pointer(oldBandE)), unsafe.SliceData(fine_quant), unsafe.SliceData(fine_priority), start, end, len1, C, (*OpusT_ec_ctx)(unsafe.Pointer(dec)))
	celtDecodeAntiCollapse(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), unsafe.SliceData(X), unsafe.SliceData(collapse_masks), unsafe.SliceData(pulses), (*float32)(unsafe.Pointer(oldBandE)), (*float32)(unsafe.Pointer(oldLogE)), (*float32)(unsafe.Pointer(oldLogE2)), N, LM, C, start, end, anti_collapse_on)
	if silence != 0 {
		celtDecodeSilenceEnergy((*float32)(unsafe.Pointer(oldBandE)), nbEBands, C)
	}
	celtDecodePrefilter(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), N)
	celt_synthesis(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), unsafe.SliceData(X), &out_syn[0], (*float32)(unsafe.Pointer(oldBandE)), start, effEnd, C, CC, isTransient, LM, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, silence, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
	celtDecodePostfilter(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), &out_syn[0], CC, N, LM, postfilter_pitch, postfilter_gain, postfilter_tapset, overlap)
	celtDecodePostfilterFinish((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), postfilter_pitch, postfilter_gain, postfilter_tapset, LM)
	if C == 1 {
		celtDecodeEnergyMono((*float32)(unsafe.Pointer(oldBandE)), nbEBands)
	}
	celtDecodeEnergyLogs((*float32)(unsafe.Pointer(oldBandE)), (*float32)(unsafe.Pointer(oldLogE)), (*float32)(unsafe.Pointer(oldLogE2)), nbEBands, isTransient)
	/* In normal circumstances, we only allow the noise floor to increase by
	   up to 2.4 dB/second, but when we're in DTX we give the weight of
	   all missing packets to the update packet. */
	celtDecodeEnergyBackground((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*float32)(unsafe.Pointer(backgroundLogE)), (*float32)(unsafe.Pointer(oldBandE)), nbEBands, M)
	/* In case start or end were to change: energy, previous, log store order. */
	celtDecodeEnergyClear((*float32)(unsafe.Pointer(oldBandE)), (*float32)(unsafe.Pointer(oldLogE)), (*float32)(unsafe.Pointer(oldLogE2)), nbEBands, start, end)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng = (*OpusT_ec_dec)(unsafe.Pointer(dec)).Frng
	celtDecodeDeemphasis(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)), &out_syn[0], (*float32)(unsafe.Pointer(pcm)), N, CC, accum)
	celtDecodePacketFinish((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)))
	if errorCode := celtDecodePacketError((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)), (*OpusT_ec_ctx)(unsafe.Pointer(dec)), len1); errorCode != 0 {
		return errorCode
	}
	return frame_size / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
}

//go:uintptrescapes
func Opus_celt_decode_with_ec(tls *libc.TLS, st uintptr, data uintptr, len1 int32, pcm uintptr, frame_size int32, dec uintptr, accum int32) (r int32) {
	return Opus_celt_decode_with_ec_dred(tls, st, data, len1, pcm, frame_size, dec, accum)
}
