// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type

func Opus__celt_lpc(tls *libc.TLS, _lpc *OpusT_opus_val16, ac *OpusT_opus_val32, p int32) {
	lpc := unsafe.Slice(_lpc, int(p))
	correlation := unsafe.Slice(ac, int(p)+1)
	error1 := correlation[0]
	clear(lpc)
	if correlation[0] > float32(1e-10) {
		for i := range lpc {
			var rr OpusT_opus_val32
			for j := 0; j < i; j++ {
				rr = rr + float32(lpc[j]*correlation[i-j])
			}
			rr = rr + correlation[i+1]
			r := -(rr / error1)
			lpc[i] = r
			for j := 0; j < (i+1)>>1; j++ {
				tmp1, tmp2 := lpc[j], lpc[i-1-j]
				lpc[j] = tmp1 + OpusT_opus_val32(r*tmp2)
				lpc[i-1-j] = tmp2 + OpusT_opus_val32(r*tmp1)
			}
			error1 = error1 - OpusT_opus_val32(OpusT_opus_val32(r*r)*error1)
			// Bail out once we get 30 dB gain, leaving the remaining taps zero.
			if error1 <= OpusT_opus_val32(float32(0.001)*correlation[0]) {
				break
			}
		}
	}
}

func Opus_celt_fir_c(tls *libc.TLS, x, num, y *float32, N, ord, arch int32) {
	celtFIRWithHistory(tls, x, num, y, N, ord, arch, nil)
}

// history, when supplied, starts ord samples before x and includes N samples.
func celtFIRWithHistory(tls *libc.TLS, x, num, y *float32, N, ord, arch int32, history []float32) {
	if x == y {
		opusCeltFatal(tls, opusDiagnosticString(3305), opusDiagnosticString(3330), 157)
	}
	coeff := unsafe.Slice(num, ord)
	reversed := make([]float32, ord)
	for i := range reversed {
		reversed[i] = coeff[len(coeff)-1-i]
	}
	// x points just past ord history samples, all within the same allocation.
	var input []float32
	if history != nil {
		input = history[:N+ord]
	} else if ord == 0 {
		input = unsafe.Slice(x, N)
	} else {
		input = unsafe.Slice((*float32)(unsafe.Add(unsafe.Pointer(x), -int(ord)*4)), N+ord)
	}
	output := unsafe.Slice(y, N)
	i := int32(0)
	for ; i < N-3; i += 4 {
		if ord < 3 {
			opusCeltFatal(tls, opusDiagnosticString(3349), opusDiagnosticString(3374), 69)
		}
		sum := [4]float32{input[ord+i], input[ord+i+1], input[ord+i+2], input[ord+i+3]}
		celtCorrelation4(reversed, input[i:], &sum)
		copy(output[i:i+4], sum[:])
	}
	for ; i < N; i++ {
		sum := input[ord+i]
		for j := int32(0); j < ord; j++ {
			sum += float32(reversed[j] * input[i+j])
		}
		output[i] = sum
	}
}

// Each lane accumulates in coefficient order; explicit products prevent ARM64 FMA.
func celtCorrelation4(coeff, input []float32, sum *[4]float32) {
	for j, c := range coeff {
		for lane := 0; lane < 4; lane++ {
			sum[lane] += float32(c * input[j+lane])
		}
	}
}

func Opus_celt_iir(tls *libc.TLS, x, den, out *float32, N, ord int32, mem *float32, arch int32) {
	if ord&3 != 0 {
		opusCeltFatal(tls, opusDiagnosticString(3390), opusDiagnosticString(3330), 225)
	}
	input, coeff, output, memory := unsafe.Slice(x, N), unsafe.Slice(den, ord), unsafe.Slice(out, N), unsafe.Slice(mem, ord)
	reversed := make([]float32, ord)
	history := make([]float32, N+ord)
	for i := int32(0); i < ord; i++ {
		reversed[i] = coeff[ord-i-1]
	}
	for i := int32(0); i < ord; i++ {
		history[i] = -memory[ord-i-1]
	}
	i := int32(0)
	for ; i < N-3; i += 4 {
		if ord < 3 {
			opusCeltFatal(tls, opusDiagnosticString(3349), opusDiagnosticString(3374), 69)
		}
		sum := [4]float32{input[i], input[i+1], input[i+2], input[i+3]}
		celtCorrelation4(reversed, history[i:], &sum)
		// Preserve patch order and output stores, including in-place input/output.
		for lane := int32(0); lane < 4; lane++ {
			for j := int32(0); j < lane; j++ {
				sum[lane] += float32(history[i+ord+lane-j-1] * coeff[j])
			}
			history[i+ord+lane] = -sum[lane]
			output[i+lane] = sum[lane]
		}
	}
	for ; i < N; i++ {
		sum := input[i]
		for j := int32(0); j < ord; j++ {
			sum -= float32(reversed[j] * history[i+j])
		}
		// The C tail stores the positive sum, unlike the four-sample block.
		history[i+ord] = sum
		output[i] = sum
	}
	var tail []float32
	if N >= ord {
		tail = output[N-ord:]
	} else {
		// For N<ord, C reads caller-owned samples preceding out as well.
		tail = unsafe.Slice((*float32)(unsafe.Add(unsafe.Pointer(out), int(N-ord)*4)), ord)
	}
	for i := int32(0); i < ord; i++ {
		memory[i] = tail[ord-i-1]
	}
}

func Opus__celt_autocorr(tls *libc.TLS, x, ac, window *float32, overlap, lag, n, arch int32) int32 {
	if n <= 0 {
		opusCeltFatal(tls, opusDiagnosticString(3419), opusDiagnosticString(3330), 302)
	}
	if overlap < 0 {
		opusCeltFatal(tls, opusDiagnosticString(3441), opusDiagnosticString(3330), 303)
	}
	input := unsafe.Slice(x, n)
	signal := input
	if overlap != 0 {
		signal = make([]float32, n)
		copy(signal, input)
		weights := unsafe.Slice(window, overlap)
		for i := int32(0); i < overlap; i++ {
			// Always multiply the original input, even when the two ends overlap.
			signal[i] = float32(input[i] * weights[i])
			signal[n-i-1] = float32(input[n-i-1] * weights[i])
		}
	}
	fastN := n - lag
	Opus_celt_pitch_xcorr_c(tls, unsafe.SliceData(signal), unsafe.SliceData(signal), ac, fastN, lag+1, arch)
	out := unsafe.Slice(ac, lag+1)
	for k := int32(0); k <= lag; k++ {
		d := float32(0)
		for i := k + fastN; i < n; i++ {
			d += float32(signal[i] * signal[i-k])
		}
		out[k] += d
	}
	return 0
}

var log2_x_norm_coeff5 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff5 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

/* The guts header contains all the multiplication and addition macros that are defined for
   complex numbers.  It also declares the kf_ internal functions.
*/
