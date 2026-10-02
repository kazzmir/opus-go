package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"slices"
	"testing"
	"unsafe"
	"weak"
)

func validationPanics(f func()) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	f()
	return false
}

func TestDecoderAllocationPointers(t *testing.T) {
	tls := libc.NewTLS()
	p := libc.XmallocPointer(tls, uint64(Opus_opus_decoder_get_size(nil, 2)))
	st := (*OpusT_OpusDecoder)(p)
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_decoder_init(nil, st, 48000, 2) != 0 || st.Fchannels != 2 || st.FFs != 48000 {
		t.Fatal("typed allocated initialization")
	}
	libc.XfreePointer(tls, p)
	tls.Close()
	entropyInitGrowStack(12)
	runtime.GC()
	if st.Fchannels != 2 || st.FFs != 48000 {
		t.Fatal("typed allocation ownership")
	}
}

func TestDecoderDestroyPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	release := func() weak.Pointer[OpusT_OpusDecoder] {
		st, err := Opus_opus_decoder_create_typed(tls, 48000, 2)
		if err != nil {
			t.Fatal(err)
		}
		w := weak.Make(st)
		entropyInitGrowStack(12)
		runtime.GC()
		Opus_opus_decoder_destroy_typed(tls, st)
		return w
	}
	w := release()
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	if w.Value() != nil {
		t.Fatal("destroy left allocation registered")
	}
	runtime.KeepAlive(tls)
	Opus_opus_decoder_destroy_typed(tls, nil)
	Opus_opus_decoder_destroy_typed(nil, nil)
	Opus_opus_decoder_destroy_typed(nil, &OpusT_OpusDecoder{})
}

func TestDecoderCreatePointers(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		for _, channels := range []int32{1, 2} {
			tls := libc.NewTLS()
			st, err := Opus_opus_decoder_create_typed(tls, rate, channels)
			if err != nil || st == nil || st.FFs != rate || st.Fchannels != channels {
				t.Fatal(rate, channels, st, err)
			}
			entropyInitGrowStack(12)
			runtime.GC()
			if st.FFs != rate || st.Fchannels != channels {
				t.Fatal("creation lifetime")
			}
			Opus_opus_decoder_destroy_typed(tls, st)
			tls.Close()
		}
	}
	for _, args := range [][2]int32{{44100, 1}, {48000, 0}, {48000, 3}, {48000, 1}} {
		st, err := Opus_opus_decoder_create_typed(nil, args[0], args[1])
		want := int32(-1)
		if args == [2]int32{48000, 1} {
			want = -7
		}
		if st != nil || err == nil || err.(*OpusError).Code != want {
			t.Fatal("error ordering", args, st, err)
		}
	}
}

func TestMSDecoderDestroyPointers(t *testing.T) {
	for _, legacy := range []bool{false, true} {
		tls := libc.NewTLS()
		release := func() weak.Pointer[OpusT_OpusMSDecoder] {
			mapping := [3]byte{0, 1, 2}
			st, err := Opus_opus_multistream_decoder_create_typed(tls, 48000, 3, 2, 1, &mapping[0])
			if err != nil {
				t.Fatal(err)
			}
			w := weak.Make(st)
			entropyInitGrowStack(12)
			runtime.GC()
			if legacy {
				Opus_opus_multistream_decoder_destroy(tls, uintptr(unsafe.Pointer(st)))
			} else {
				Opus_opus_multistream_decoder_destroy_typed(tls, st)
			}
			return w
		}
		w := release()
		for i := 0; i < 10; i++ {
			runtime.GC()
		}
		if w.Value() != nil {
			t.Fatal("multistream allocation retained", legacy)
		}
		runtime.KeepAlive(tls)
		tls.Close()
	}
	Opus_opus_multistream_decoder_destroy_typed(nil, nil)
	tls := libc.NewTLS()
	defer tls.Close()
	Opus_opus_multistream_decoder_destroy_typed(tls, nil)
}

func TestMSDecoderCreatePointers(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		tls := libc.NewTLS()
		mapping := [3]byte{0, 1, 2}
		entropyInitGrowStack(12)
		runtime.GC()
		st, err := Opus_opus_multistream_decoder_create_typed(tls, rate, 3, 2, 1, &mapping[0])
		if err != nil || st == nil || st.Flayout.Fmapping[2] != 2 || st.Flayout.Fnb_streams != 2 {
			t.Fatal(st, err)
		}
		runtime.GC()
		Opus_opus_multistream_decoder_destroy_typed(tls, st)
		tls.Close()
	}
	for _, test := range []struct{ rate, channels, streams, coupled, code int32 }{{48000, 0, 1, 0, -1}, {48000, 1, 0, 0, -1}, {48000, 1, 1, 2, -1}, {44100, 1, 1, 0, -7}, {48000, 1, 1, 0, -7}} {
		st, err := Opus_opus_multistream_decoder_create_typed(nil, test.rate, test.channels, test.streams, test.coupled, nil)
		if st != nil || err == nil || err.(*OpusError).Code != test.code {
			t.Fatal(test, st, err)
		}
	}
	tls := libc.NewTLS()
	defer tls.Close()
	mapping := [1]byte{1}
	st, err := Opus_opus_multistream_decoder_create_typed(tls, 48000, 1, 1, 0, &mapping[0])
	if st != nil || err == nil || err.(*OpusError).Code != -1 {
		t.Fatal("failed layout", st, err)
	}
}

func TestProjectionDecoderDestroyPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	release := func() weak.Pointer[OpusT_OpusProjectionDecoder] {
		matrix := [18]byte{}
		st, err := Opus_opus_projection_decoder_create_typed(tls, 48000, 3, 2, 1, &matrix[0], 18)
		if err != nil {
			t.Fatal(err)
		}
		w := weak.Make(st)
		if unsafe.Pointer(get_dec_demixing_matrix(nil, st)) == unsafe.Pointer(st) || unsafe.Pointer(get_multistream_decoder(nil, st)) == unsafe.Pointer(st) {
			t.Fatal("expected interiors")
		}
		entropyInitGrowStack(12)
		runtime.GC()
		Opus_opus_projection_decoder_destroy_typed(tls, st)
		return w
	}
	w := release()
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	if w.Value() != nil {
		t.Fatal("projection base retained")
	}
	runtime.KeepAlive(tls)
	Opus_opus_projection_decoder_destroy_typed(tls, nil)
	Opus_opus_projection_decoder_destroy_typed(nil, nil)
}

func TestProjectionDecoderCreatePointers(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		tls := libc.NewTLS()
		matrix := [18]byte{0, 0, 0xff, 0x7f, 0, 0x80, 0xff, 0xff, 1, 0, 0xfe, 0xff}
		before := matrix
		entropyInitGrowStack(12)
		runtime.GC()
		st, err := Opus_opus_projection_decoder_create_typed(tls, rate, 3, 2, 1, &matrix[0], 18)
		if err != nil || st == nil || matrix != before {
			t.Fatal(st, err)
		}
		runtime.GC()
		m := get_dec_demixing_matrix(nil, st)
		data := unsafe.Slice(Opus_mapping_matrix_get_data(nil, m), 9)
		if m.Frows != 3 || m.Fcols != 3 || data[1] != 32767 || data[2] != -32768 || get_multistream_decoder(nil, st).Flayout.Fmapping[2] != 2 {
			t.Fatal("created matrix/state")
		}
		Opus_opus_projection_decoder_destroy_typed(tls, st)
		tls.Close()
	}
	for _, args := range [][4]int32{{48000, 3, 2, 17}, {44100, 3, 2, 18}, {48000, 256, 1, 0}} {
		st, err := Opus_opus_projection_decoder_create_typed(nil, args[0], args[1], args[2], 0, nil, args[3])
		if st != nil || err == nil || err.(*OpusError).Code != -7 {
			t.Fatal("allocation precedes init arguments", args, st, err)
		}
	}
	tls := libc.NewTLS()
	defer tls.Close()
	st, err := Opus_opus_projection_decoder_create_typed(tls, 48000, 3, 2, 1, nil, 17)
	if st != nil || err == nil || err.(*OpusError).Code != -1 {
		t.Fatal("size failure", st, err)
	}
}

func TestCustomDecoderSizePointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{Foverlap: 120, FnbEBands: 21}
	before := mode
	entropyInitGrowStack(12)
	runtime.GC()
	mono := opus_custom_decoder_get_size(nil, &mode, 1)
	stereo := opus_custom_decoder_get_size(nil, &mode, 2)
	if mode != before {
		t.Fatal("mode changed")
	}
	if stereo-mono != (DEC_PITCH_BUF_SIZE+120+CELT_LPC_ORDER)*4 {
		t.Fatal(mono, stereo)
	}
	mode.FnbEBands++
	if opus_custom_decoder_get_size(nil, &mode, 1) != mono+32 {
		t.Fatal("band stride")
	}
	mode = before
	// Signed narrowing of the original generated size formula is retained.
	mode.Foverlap = 1 << 30
	got := opus_custom_decoder_get_size(nil, &mode, 2)
	want := int32(uint64(unsafe.Sizeof(OpusT_OpusCustomDecoder{})) + uint64(uint32(int32(2)*(DEC_PITCH_BUF_SIZE+mode.Foverlap)-1))*4 + uint64(uint32(mode.FnbEBands*8))*4 + uint64(uint32(int32(2)*CELT_LPC_ORDER))*4)
	if got != want {
		t.Fatal("wrap", got, want)
	}
}

func TestMSValidationPointers(t *testing.T) {
	st := OpusT_OpusMSDecoder{Flayout: OpusT_ChannelLayout{Fnb_channels: 2, Fnb_streams: 1, Fnb_coupled_streams: 1, Fmapping: [256]uint8{0, 1}}}
	before := st
	entropyInitGrowStack(12)
	runtime.GC()
	validate_ms_decoder(nil, &st)
	if st != before {
		t.Fatal("state changed")
	}
	// The original validator intentionally ignores an invalid layout return value.
	st.Flayout.Fmapping[0] = 7
	before = st
	validate_ms_decoder(nil, &st)
	if st != before || Opus_validate_layout(nil, &st.Flayout) != 0 {
		t.Fatal("invalid-layout behavior changed")
	}
}

type celtStateTestStorage struct {
	State OpusT_OpusCustomDecoder
	Tail  [5000]float32
}

func celtStateTestBuffer(mode *OpusT_OpusCustomMode, channels int32) (*celtStateTestStorage, []byte, int) {
	storage := new(celtStateTestStorage)
	size := int(opus_custom_decoder_get_size(nil, mode, channels))
	image := unsafe.Slice((*byte)(unsafe.Pointer(&storage.State)), size+16)
	// Never put arbitrary byte patterns in a GC-scanned pointer slot.
	for i := int(unsafe.Sizeof(storage.State.Fmode)); i < len(image); i++ {
		image[i] = 0xa5
	}
	storage.State.Fmode = mode
	storage.State.Fchannels = channels
	storage.State.Foverlap = mode.Foverlap
	return storage, image, size
}

func TestOpusDecoderInitPointers(t *testing.T) {
	for _, ch := range []int32{1, 2} {
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
			size := int(Opus_opus_decoder_get_size(nil, ch))
			backing := make([]uint64, (size+7)/8+3)
			backing[0] = 0x123456789abcdef0
			st := (*OpusT_OpusDecoder)(unsafe.Pointer(&backing[1]))
			image := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
			for i := size; i < len(image); i++ {
				image[i] = 0xa5
			}
			entropyInitGrowStack(12)
			runtime.GC()
			if Opus_opus_decoder_init(nil, st, rate, ch) != 0 {
				t.Fatal("init")
			}
			celt := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(st), st.Fcelt_dec_offset))
			if st.Fchannels != ch || st.FFs != rate || st.Fframe_size != rate/400 || st.FDecControl.FAPI_sampleRate != rate || celt.Fmode == nil || celt.Fsignalling != 0 || celt.Fdownsample != Opus_resampling_factor(nil, rate) {
				t.Fatal("initialized fields")
			}
			st.Fframe_size = 77
			celt.Frng = 88
			runtime.GC()
			if Opus_opus_decoder_init(nil, st, rate, ch) != 0 || celt.Frng != 0 || st.Fframe_size != rate/400 {
				t.Fatal("reinit")
			}
			if backing[0] != 0x123456789abcdef0 {
				t.Fatal("prefix guard")
			}
			for _, b := range image[size:] {
				if b != 0xa5 {
					t.Fatal("tail guard")
				}
			}
			before := slices.Clone(image)
			if Opus_opus_decoder_init(nil, st, 44100, ch) != OPUS_BAD_ARG || !slices.Equal(before, image) {
				t.Fatal("invalid rate changed state")
			}
		}
	}
	if Opus_opus_decoder_init(nil, nil, 44100, 2) != OPUS_BAD_ARG || Opus_opus_decoder_init(nil, nil, 48000, 0) != OPUS_BAD_ARG {
		t.Fatal("validation before access")
	}
}

func TestCeltDecoderInitPointers(t *testing.T) {
	mode, _ := Opus_opus_custom_mode_create(nil, 48000, 960)
	for _, rate := range []int32{48000, 24000, 16000, 12000, 8000} {
		storage, image, size := celtStateTestBuffer(mode, 2)
		entropyInitGrowStack(12)
		runtime.GC()
		code := Opus_celt_decoder_init(nil, &storage.State, rate, 2)
		factor := Opus_resampling_factor(nil, rate)
		want := int32(0)
		if factor == 0 {
			want = OPUS_BAD_ARG
		}
		if code != want || storage.State.Fmode != mode || storage.State.Fdownsample != factor || storage.State.Fskip_plc != 1 {
			t.Fatal(rate, code, &storage.State)
		}
		for _, b := range image[size:] {
			if b != 0xa5 {
				t.Fatal("guard")
			}
		}
	}
	storage, _, _ := celtStateTestBuffer(mode, 2)
	if !validationPanics(func() { Opus_celt_decoder_init(nil, &storage.State, 44100, 2) }) || storage.State.Fdownsample != 1 || storage.State.Fskip_plc != 1 {
		t.Fatal("rate assertion/initialized state")
	}
	if Opus_celt_decoder_init(nil, nil, 44100, 1) != -7 || Opus_celt_decoder_init(nil, nil, 48000, 3) != OPUS_BAD_ARG {
		t.Fatal("early errors")
	}
}

func TestCustomDecoderInitPointers(t *testing.T) {
	if opus_custom_decoder_init(nil, nil, nil, -1) != OPUS_BAD_ARG || opus_custom_decoder_init(nil, nil, nil, 1) != -7 {
		t.Fatal("argument order")
	}
	for _, channels := range []int32{0, 1, 2} {
		storage, image, size := func() (*celtStateTestStorage, []byte, int) {
			mode := &OpusT_OpusCustomMode{Foverlap: 120, FnbEBands: 21, FeffEBands: 19}
			s, b, n := celtStateTestBuffer(mode, channels)
			if opus_custom_decoder_init(nil, &s.State, mode, channels) != 0 {
				t.Fatal("init")
			}
			return s, b, n
		}()
		entropyInitGrowStack(12)
		runtime.GC()
		st := &storage.State
		wantInv := int32(0)
		if channels == 1 {
			wantInv = 1
		}
		if st.Fmode.FeffEBands != 19 || st.Fend != 19 || st.Fchannels != channels || st.Fstream_channels != channels || st.Fdownsample != 1 || st.Fsignalling != 1 || st.Fdisable_inv != wantInv || st.Frng != 0 || st.Fskip_plc != 1 {
			t.Fatal("state/mode ownership", st)
		}
		for _, b := range image[size:] {
			if b != 0xa5 {
				t.Fatal("guard")
			}
		}
		before := slices.Clone(image)
		if opus_custom_decoder_init(nil, st, nil, 3) != OPUS_BAD_ARG || !slices.Equal(before, image) {
			t.Fatal("failure modified state")
		}
	}
}

type opusCtlTestStorage struct {
	State OpusT_OpusDecoder
	Silk  OpusT_silk_decoder
	Celt  celtStateTestStorage
}

func newOpusCtlTestStorage() *opusCtlTestStorage {
	s := new(opusCtlTestStorage)
	s.State.Fsilk_dec_offset = int32(unsafe.Offsetof(s.Silk))
	s.State.Fcelt_dec_offset = int32(unsafe.Offsetof(s.Celt))
	s.State.FFs = 48000
	s.State.Fchannels = 1
	s.State.Fstream_channels = 1
	s.State.Fcomplexity = 10
	Opus_silk_InitDecoder(nil, &s.Silk)
	opus_custom_decoder_init(nil, &s.Celt.State, &mode48000_960_120, 1)
	return s
}
func TestOpusCtlPointers(t *testing.T) {
	s := newOpusCtlTestStorage()
	var out int32
	var rng uint32
	a := OpusDecoderCtlArgs{I32: &out, U32: &rng}
	entropyInitGrowStack(12)
	runtime.GC()
	for _, v := range []int32{-32769, -32768, 0, 32767, 32768} {
		r := Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_SET_GAIN_REQUEST, OpusDecoderCtlArgs{Value: v})
		if (r == 0) != (v >= -32768 && v <= 32767) {
			t.Fatal("gain", v, r)
		}
	}
	Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_GET_GAIN_REQUEST, a)
	if out != 32767 {
		t.Fatal("gain output")
	}
	s.State.Fprev_mode = MODE_CELT_ONLY
	s.Celt.State.Fpostfilter_period = 99
	Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_GET_PITCH_REQUEST, a)
	if out != 99 {
		t.Fatal("CELT pitch")
	}
	s.State.Fprev_mode = MODE_SILK_ONLY
	s.State.FDecControl.FprevPitchLag = 66
	Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_GET_PITCH_REQUEST, a)
	if out != 66 {
		t.Fatal("SILK pitch")
	}
	s.State.FrangeFinal = 0xfedcba98
	Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_GET_FINAL_RANGE_REQUEST, a)
	if rng != 0xfedcba98 {
		t.Fatal("range output")
	}
	if Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_GET_GAIN_REQUEST, OpusDecoderCtlArgs{}) != -1 {
		t.Fatal("nil output")
	}
	Opus_opus_decoder_ctl_typed(nil, &s.State, OPUS_RESET_STATE, a)
	if s.State.Fstream_channels != 1 || s.State.Fframe_size != 120 || s.State.FrangeFinal != 0 || s.Celt.State.Fskip_plc != 1 {
		t.Fatal("aggregate reset")
	}
	if Opus_opus_decoder_ctl_typed(nil, &s.State, 123456, OpusDecoderCtlArgs{}) != -5 {
		t.Fatal("unknown request")
	}
}

func TestCustomCtlPointers(t *testing.T) {
	mode := new(OpusT_OpusCustomMode)
	*mode = mode48000_960_120
	storage, _, _ := celtStateTestBuffer(mode, 1)
	st := &storage.State
	opus_custom_decoder_init(nil, st, mode, 1)
	var out int32
	var rng uint32
	var got *OpusT_OpusCustomMode
	args := OpusDecoderCtlArgs{I32: &out, U32: &rng, Mode: &got}
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_custom_decoder_ctl_typed(nil, st, CELT_GET_MODE_REQUEST, args) != 0 || got != mode {
		t.Fatal("mode output")
	}
	st.Fmode = nil
	mode = nil
	runtime.GC()
	if got.FnbEBands != 21 {
		t.Fatal("mode output owner")
	}
	st.Fmode = got
	for _, v := range []int32{-1, 0, 10, 11} {
		r := Opus_opus_custom_decoder_ctl_typed(nil, st, OPUS_SET_COMPLEXITY_REQUEST, OpusDecoderCtlArgs{Value: v})
		if (r == 0) != (v >= 0 && v <= 10) {
			t.Fatal("complexity", v, r)
		}
	}
	st.Ferror1 = 123
	if Opus_opus_custom_decoder_ctl_typed(nil, st, CELT_GET_AND_CLEAR_ERROR_REQUEST, OpusDecoderCtlArgs{I32: &st.Ferror1}) != 0 || st.Ferror1 != 0 {
		t.Fatal("clear output alias")
	}
	st.Frng = 0x89abcdef
	Opus_opus_custom_decoder_ctl_typed(nil, st, OPUS_GET_FINAL_RANGE_REQUEST, args)
	if rng != 0x89abcdef {
		t.Fatal("range output")
	}
	if Opus_opus_custom_decoder_ctl_typed(nil, nil, OPUS_GET_PITCH_REQUEST, OpusDecoderCtlArgs{}) != -1 || Opus_opus_custom_decoder_ctl_typed(nil, nil, 123456, OpusDecoderCtlArgs{}) != -5 {
		t.Fatal("nil/error order")
	}
	Opus_opus_custom_decoder_ctl_typed(nil, st, OPUS_RESET_STATE, args)
	if st.Fskip_plc != 1 || st.Fmode != got {
		t.Fatal("typed reset")
	}
}

func TestCeltResetPointers(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		mode := &OpusT_OpusCustomMode{Foverlap: 120, FnbEBands: 21}
		storage, image, size := celtStateTestBuffer(mode, channels)
		before := append([]byte(nil), image...)
		entropyInitGrowStack(12)
		runtime.GC()
		celt_decoder_reset(nil, &storage.State)
		start := int(unsafe.Offsetof(storage.State.Frng))
		if !slices.Equal(image[:start], before[:start]) || !slices.Equal(image[size:], before[size:]) || storage.State.Frng != 0 || storage.State.Fskip_plc != 1 || storage.State.Flast_frame_type != FRAME_NONE {
			t.Fatal(channels, "prefix/guards/reset")
		}
		oldBand := unsafe.Add(unsafe.Pointer(&storage.State.F_decode_mem[0]), int((DEC_PITCH_BUF_SIZE+mode.Foverlap)*channels)*4)
		logs := unsafe.Slice((*float32)(unsafe.Add(oldBand, int(2*mode.FnbEBands)*4)), 4*mode.FnbEBands)
		for _, v := range logs {
			if v != -28 {
				t.Fatal("log history")
			}
		}
	}
}

func TestCustomModePointers(t *testing.T) {
	var first *OpusT_OpusCustomMode
	for _, frame := range []int32{120, 240, 480, 960} {
		mode, err := Opus_opus_custom_mode_create(nil, 48000, frame)
		if err != nil || mode == nil {
			t.Fatal(frame, err)
		}
		if first == nil {
			first = mode
		}
		if mode != first {
			t.Fatal("mode identity")
		}
	}
	before := *first
	entropyInitGrowStack(12)
	runtime.GC()
	if *first != before {
		t.Fatal("mode changed")
	}
	for _, frame := range []int32{0, 60, 119, 121, 1920} {
		if mode, err := Opus_opus_custom_mode_create(nil, 48000, frame); mode != nil || err == nil {
			t.Fatal(frame, mode, err)
		}
	}
	if mode, err := Opus_opus_custom_mode_create(nil, 44100, 960); mode != nil || err == nil {
		t.Fatal("unsupported rate")
	}
	// Preserve the generated int32 shift's low bits, without using C's undefined signed overflow as an oracle.
	if mode, err := Opus_opus_custom_mode_create(nil, 48000, 0x20000078); err != nil || mode != first {
		t.Fatal("shift wrapping")
	}
}

func TestCeltValidationPointers(t *testing.T) {
	mode, _ := Opus_opus_custom_mode_create(nil, 48000, 960)
	st := OpusT_OpusCustomDecoder{Fmode: mode, Foverlap: 120, Fchannels: 2, Fstream_channels: 1, Fdownsample: 1, Fend: 21}
	before := st
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_validate_celt_decoder(nil, &st)
	if st != before {
		t.Fatal("state changed")
	}
	for _, pitch := range []int32{0, PLC_PITCH_LAG_MIN, PLC_PITCH_LAG_MAX} {
		st = before
		st.Flast_pitch_index = pitch
		Opus_validate_celt_decoder(nil, &st)
	}
	st = before
	st.Fpostfilter_period = MAX_PERIOD
	if !validationPanics(func() { Opus_validate_celt_decoder(nil, &st) }) {
		t.Fatal("period accepted")
	}
	st = before
	st.Fmode = nil
	if !validationPanics(func() { Opus_validate_celt_decoder(nil, &st) }) {
		t.Fatal("mode accepted")
	}
}

func TestOpusValidationPointers(t *testing.T) {
	st := OpusT_OpusDecoder{Fchannels: 2, FFs: 48000, Fstream_channels: 1}
	st.FDecControl.FAPI_sampleRate = 48000
	st.FDecControl.FnChannelsAPI = 2
	before := st
	entropyInitGrowStack(12)
	runtime.GC()
	validate_opus_decoder(nil, &st)
	if st != before {
		t.Fatal("state changed")
	}
	for _, bad := range []int32{0, 3, -1} {
		st = before
		st.Fchannels = bad
		if !validationPanics(func() { validate_opus_decoder(nil, &st) }) {
			t.Fatal("channel accepted", bad)
		}
	}
	st = before
	st.Farch = -1
	if !validationPanics(func() { validate_opus_decoder(nil, &st) }) {
		t.Fatal("arch accepted")
	}
}
