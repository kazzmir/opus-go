//go:build compareopus

package opuscc

import (
	"reflect"
	"runtime"
	"slices"
	"testing"
	"weak"
)

func compareImageMustPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("invalid serialized image accepted")
		}
	}()
	f()
}

func compareInitializedOpusImage(t *testing.T, channels int32) []byte {
	t.Helper()
	data := make([]byte, Opus_opus_decoder_get_size(nil, channels)+16)
	image := compareDecoderImageStorage(opusDecoderStorageType(channels), data)
	state := image.owner.Field(0).FieldByName("Decoder").Addr().Interface().(*OpusT_OpusDecoder)
	if r := Opus_opus_decoder_init(nil, state, 48000, channels); r != 0 {
		t.Fatal("Opus image init", r)
	}
	celt := image.owner.Field(0).FieldByName("Celt").Addr().Interface().(*OpusT_OpusCustomDecoder)
	celt.Fmode = nil
	state.Fdecode_gain = -99
	image.save(data)
	return data
}

func TestCompareOpusCtlImagePointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		data := compareInitializedOpusImage(t, channels)
		field, _ := reflect.TypeFor[OpusT_OpusDecoder]().FieldByName("Fdecode_gain")
		r, out := CompareOpusCtl(data, OPUS_GET_SAMPLE_RATE_REQUEST, 0, int32(field.Offset))
		if r != 0 || out != 77 {
			t.Fatal("Opus scalar alias result", r, out)
		}
		st, image := compareOpusDecoderImage(data)
		if st.Fdecode_gain != 48000 {
			t.Fatal("live sample-rate/gain alias")
		}
		celt := image.owner.Field(0).FieldByName("Celt").Addr().Interface().(*OpusT_OpusCustomDecoder)
		silk := image.owner.Field(0).FieldByName("Silk").Addr().Interface().(*OpusT_silk_decoder)
		mode := new(OpusT_OpusCustomMode)
		book := new(OpusT_silk_NLSF_CB_struct)
		book.FCB1_NLSF_Q8 = new(byte)
		*book.FCB1_NLSF_Q8 = 77
		modeRoot, bookRoot := weak.Make(mode), weak.Make(book)
		celt.Fmode = mode
		silk.Fchannel_state[0].FpsNLSF_CB = book
		st, celt, silk, mode, book = nil, nil, nil, nil, nil
		entropyInitGrowStack(12)
		runtime.GC()
		if modeRoot.Value() == nil || bookRoot.Value() == nil || *bookRoot.Value().FCB1_NLSF_Q8 != 77 {
			t.Fatal("composite scanned CELT/SILK owners")
		}
		compareImageMustPanic(t, func() { image.save(data) })
		celt = image.owner.Field(0).FieldByName("Celt").Addr().Interface().(*OpusT_OpusCustomDecoder)
		silk = image.owner.Field(0).FieldByName("Silk").Addr().Interface().(*OpusT_silk_decoder)
		celt.Fmode = nil
		silk.Fchannel_state[0].FpsNLSF_CB = nil
		image.save(data)
		r, out = CompareOpusCtl(data, OPUS_GET_GAIN_REQUEST, 0, -1)
		if r != 0 || out != 48000 {
			t.Fatal("Opus scalar alias serialized bits")
		}
		runtime.KeepAlive(image)
	}
}

func TestComparePrefilterImagePointers(t *testing.T) {
	for _, active := range []int32{0, 1} {
		for _, n := range []int32{0, 120, 960} {
			size := opus_custom_decoder_get_size(nil, &mode48000_960_120, 1)
			data := make([]byte, size+16)
			st, initial := compareCustomDecoderImage(data)
			if r := opus_custom_decoder_init(nil, st, &mode48000_960_120, 1); r != 0 {
				t.Fatal("custom fixture init", r)
			}
			st.Fprefilter_and_fold = active
			st.Fpostfilter_period_old, st.Fpostfilter_period = 20, 40
			st.Fpostfilter_gain_old, st.Fpostfilter_gain = .25, .5
			history, _ := celtDecodeHistoryViews(st, st.Foverlap, 1, n)
			for i := range history[0] {
				history[0][i] = float32(i%29-14) / 37
			}
			st.Fmode = nil
			initial.save(data)
			want := slices.Clone(data)
			reference, expected := compareCustomDecoderImage(want)
			reference.Fmode = &mode48000_960_120
			if active != 0 {
				prefilter_and_fold(nil, reference, n)
			}
			reference.Fmode = nil
			expected.save(want)
			entropyInitGrowStack(12)
			runtime.GC()
			CompareCeltDecodePrefilterImage(data, n)
			if !slices.Equal(data, want) {
				t.Fatal("prefilter dispatch/entire image/guards", active, n)
			}
		}
	}
}

func TestCompareCustomCtlImagePointers(t *testing.T) {
	data := make([]byte, int(reflect.TypeFor[OpusT_OpusCustomDecoder]().Size())+64)
	st, initial := compareCustomDecoderImage(data)
	st.Ferror1, st.Fcomplexity, st.Frng = -123, 5, 0x89abcdef
	initial.save(data)
	field, _ := reflect.TypeFor[OpusT_OpusCustomDecoder]().FieldByName("Fcomplexity")
	entropyInitGrowStack(12)
	runtime.GC()
	r, out := CompareCustomCtl(data, CELT_GET_AND_CLEAR_ERROR_REQUEST, 0, int32(field.Offset))
	if r != 0 || out != 77 {
		t.Fatal("custom CTL alias result", r, out)
	}
	st, result := compareCustomDecoderImage(data)
	if st.Ferror1 != 0 || st.Fcomplexity != -123 || st.Fmode != nil {
		t.Fatal("custom CTL live scalar alias/cleared mode")
	}
	result.save(data)
	r, out = CompareCustomCtl(data, OPUS_GET_FINAL_RANGE_REQUEST, 0, -1)
	if r != 0 || out != 0x89abcdef {
		t.Fatal("custom range bits")
	}
	r, out = CompareCustomCtl(data, CELT_GET_MODE_REQUEST, 0, -1)
	if r != 0 || out != 1 {
		t.Fatal("custom mode identity")
	}
}

func TestCompareCustomDecoderImageOwnerPointers(t *testing.T) {
	size := int(reflect.TypeFor[OpusT_OpusCustomDecoder]().Size()) + 64
	data := make([]byte, size)
	for i := size - 64; i < size; i++ {
		data[i] = byte(i*37 + 19)
	}
	before := slices.Clone(data)
	state, image := compareCustomDecoderImage(data)
	if image.owner.Field(0).Type() != reflect.TypeFor[OpusT_OpusCustomDecoder]() || int(image.owner.Type().Size()) != size || len(image.bytes) != size {
		t.Fatal("actual scanned header/exact extent")
	}
	mode := new(OpusT_OpusCustomMode)
	mode.Fwindow = new(float32)
	*mode.Fwindow = .375
	root := weak.Make(mode)
	state.Fmode = mode
	state, mode = nil, nil
	entropyInitGrowStack(12)
	runtime.GC()
	if root.Value() == nil || *root.Value().Fwindow != .375 {
		t.Fatal("scanned mode/window owner lost")
	}
	compareImageMustPanic(t, func() { image.save(data) })
	state = image.owner.Field(0).Addr().Interface().(*OpusT_OpusCustomDecoder)
	state.Fmode = nil
	image.save(data)
	if !slices.Equal(data, before) {
		t.Fatal("scalar payload/guards round trip")
	}
	if len(image.slots) != 1 {
		t.Fatal("custom mode pointer slot")
	}
	bad := slices.Clone(before)
	bad[image.slots[0][0]] = 1
	compareImageMustPanic(t, func() { compareCustomDecoderImage(bad) })
	compareImageMustPanic(t, func() { compareCustomDecoderImage(nil) })
	compareImageMustPanic(t, func() { compareCustomDecoderImage(before[:size-1]) })
	// Variable C tails need not be pointer-aligned. Keep the exact original
	// view even when the containing Go struct requires trailing ABI padding.
	short := slices.Clone(before[:size-4])
	_, padded := compareCustomDecoderImage(short)
	if len(padded.bytes) != len(short) || cap(padded.bytes) != len(short) {
		t.Fatal("padding exposed as fixture storage")
	}
	padded.save(short)
	if !slices.Equal(short, before[:size-4]) {
		t.Fatal("unaligned logical extent round trip")
	}
	runtime.KeepAlive(image)
}
