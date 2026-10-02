package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestQuantBandN1Pointers(t *testing.T) {
	buffer := [16]byte{}
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, &buffer[0], 16)
	ctx := band_ctx{Fencode: 1, Fresynth: 1, Fremaining_bits: 8}
	x := float32(-.5)
	low := float32(77)
	entropyInitGrowStack(12)
	runtime.GC()
	if quant_band_n1(nil, &ctx, &encoder, &x, &x, &low) != 1 || x != 1 || low != 1 || ctx.Fremaining_bits != 0 {
		t.Fatal("alias/budget/store order", x, low, ctx.Fremaining_bits)
	}
	// Entropy output can legally alias the integer encode field through bytes.
	ctx = band_ctx{Fencode: 1, Fresynth: 1, Fremaining_bits: 16}
	Opus_ec_enc_init(nil, &encoder, (*byte)(unsafe.Pointer(&ctx.Fencode)), 4)
	encoder.Fnend_bits = 32
	encoder.Fend_window = 0
	x = -.5
	y := float32(-.25)
	quant_band_n1(nil, &ctx, &encoder, &x, &y, nil)
	if ctx.Fencode != 0 || encoder.Fnend_bits != 2 || encoder.Fend_window != 3 || ctx.Fremaining_bits != 0 {
		t.Fatal("cached encode across aliasing byte stores", ctx, encoder)
	}
	ctx = band_ctx{Fencode: 1, Fremaining_bits: 0}
	if quant_band_n1(nil, &ctx, nil, nil, nil, nil) != 1 {
		t.Fatal("unused nil inputs")
	}
}

func TestBandContextEntropyPointers(t *testing.T) {
	makeEncoder := func(capacity int) *band_ctx {
		buffer := make([]byte, capacity)
		ec := new(OpusT_ec_ctx)
		Opus_ec_enc_init(nil, ec, unsafe.SliceData(buffer), uint32(capacity))
		return &band_ctx{Fencode: 1, Fresynth: 1, Fremaining_bits: 4096, Fec: ec}
	}
	compare := func(a, b *band_ctx) {
		t.Helper()
		g, c := *a.Fec, *b.Fec
		if !slices.Equal(unsafe.Slice(g.Fbuf, int(g.Fstorage)), unsafe.Slice(c.Fbuf, int(c.Fstorage))) {
			t.Fatal("owned packet bytes")
		}
		g.Fbuf = nil
		c.Fbuf = nil
		if g != c || a.Fremaining_bits != b.Fremaining_bits {
			t.Fatal("owned entropy/context state", g, c)
		}
	}
	var encoded []byte
	for _, capacity := range []int{0, 1, 32} {
		ctx, ref := makeEncoder(capacity), makeEncoder(capacity)
		saved := *ctx
		ctx.Fec = nil
		entropyInitGrowStack(12)
		runtime.GC()
		*ctx = saved
		for i := 0; i < 20; i++ {
			x, y, low := float32(-.5), float32(.25), float32(77)
			if i%2 != 0 {
				x, y = y, x
			}
			rx, ry, rl := x, y, low
			gm := quant_band_n1(nil, ctx, ctx.Fec, &x, &y, &low)
			cm := quant_band_n1(nil, ref, ref.Fec, &rx, &ry, &rl)
			if gm != cm || x != rx || y != ry || low != rl {
				t.Fatal("owned encoder sign/store")
			}
			compare(ctx, ref)
		}
		Opus_ec_enc_done(nil, ctx.Fec)
		Opus_ec_enc_done(nil, ref.Fec)
		compare(ctx, ref)
		if capacity == 32 {
			encoded = slices.Clone(unsafe.Slice(ctx.Fec.Fbuf, capacity))
		}
	}
	makeDecoder := func(input []byte) *band_ctx {
		owned := slices.Clone(input)
		ec := new(OpusT_ec_ctx)
		Opus_ec_dec_init(nil, ec, unsafe.SliceData(owned), uint32(len(owned)))
		return &band_ctx{Fresynth: 1, Fremaining_bits: 4096, Fec: ec}
	}
	ctx, ref := makeDecoder(encoded), makeDecoder(encoded)
	saved := *ctx
	ctx.Fec = nil
	entropyInitGrowStack(12)
	runtime.GC()
	*ctx = saved
	for i := 0; i < 20; i++ {
		x, y, low := float32(77), float32(88), float32(99)
		rx, ry, rl := x, y, low
		gm := quant_band_n1(nil, ctx, ctx.Fec, &x, &y, &low)
		cm := quant_band_n1(nil, ref, ref.Fec, &rx, &ry, &rl)
		wantX, wantY := float32(-1), float32(1)
		if i%2 != 0 {
			wantX, wantY = wantY, wantX
		}
		if gm != cm || x != rx || y != ry || low != rl || x != wantX || y != wantY || low != x {
			t.Fatal("owned decoder sign/store", i, x, y, low)
		}
		compare(ctx, ref)
	}
	runtime.KeepAlive(ctx)
}

func TestThetaOutputPointers(t *testing.T) {
	for alias := 0; alias < 4; alias++ {
		data := [8]byte{17, 255, 88, 1, 192, 0, 77, 43}
		var ec OpusT_ec_ctx
		Opus_ec_dec_init(nil, &ec, &data[0], 8)
		log := [1]int16{24}
		m := OpusT_OpusCustomMode{FnbEBands: 1, FlogN: &log[0]}
		ctx := band_ctx{Fm: &m, Fec: &ec, Fintensity: 0, Fremaining_bits: 300}
		var split split_ctx
		b, fill := int32(16), int32(15)
		bp, fp := &b, &fill
		switch alias {
		case 1:
			fp = bp
		case 2:
			ctx.Fremaining_bits = b
			bp = &ctx.Fremaining_bits
		case 3:
			split.Fitheta = b
			split.Fqalloc = fill
			bp = &split.Fitheta
			fp = &split.Fqalloc
		}
		entropyInitGrowStack(12)
		runtime.GC()
		compute_theta(nil, &ctx, &split, nil, nil, 2, bp, 1, 1, 0, 1, fp)
		if split.Fimid != 32767 || split.Fiside != 0 || split.Fdelta != -16384 || split.Fitheta != 0 || split.Fqalloc != 0 || split.Finv != 0 {
			t.Fatal(alias, split)
		}
		wantB := [4]int32{16, 0, 16, 0}
		wantFill := [4]int32{1, 0, 1, 0}
		if *bp != wantB[alias] || *fp != wantFill[alias] {
			t.Fatal("aliased output order", alias, *bp, *fp)
		}
	}
}

func TestQuantBandContextPointers(t *testing.T) {
	ctx := new(band_ctx)
	saved := *ctx
	entropyInitGrowStack(12)
	runtime.GC()
	if mask := quant_band(nil, ctx, nil, 1, 0, 1, nil, 0, nil, 1, nil, 1); mask != 1 || *ctx != saved {
		t.Fatal("typed mono context / unused nil spectra", mask, ctx)
	}
}

func TestQuantStereoContextPointers(t *testing.T) {
	ctx := new(band_ctx)
	saved := *ctx
	entropyInitGrowStack(12)
	runtime.GC()
	if mask := quant_band_stereo(nil, ctx, 0, 0, 1, 0, 1, 0, 0, 0, 0, 3); mask != 1 || *ctx != saved {
		t.Fatal("typed stereo context / unused nil spectra", mask, ctx)
	}
}

func TestThetaSpectrumPointers(t *testing.T) {
	buffer := make([]byte, 32)
	var ec OpusT_ec_ctx
	Opus_ec_enc_init(nil, &ec, &buffer[0], 32)
	log := [1]int16{24}
	m := OpusT_OpusCustomMode{FnbEBands: 1, FlogN: &log[0]}
	energy := [2]float32{.8, 1.2}
	ctx := band_ctx{Fm: &m, Fec: &ec, FbandE: &energy[0], Fencode: 1, Fremaining_bits: 300}
	x, y := []float32{77, .3, -.4, 88}, []float32{99, .2, .1, 66}
	rx, ry := slices.Clone(x), slices.Clone(y)
	intensity_stereo(nil, &m, &rx[1], &ry[1], &energy[0], 0, 2)
	var split split_ctx
	b, fill := int32(0), int32(3)
	entropyInitGrowStack(12)
	runtime.GC()
	compute_theta(nil, &ctx, &split, &x[1], &y[1], 2, &b, 1, 1, 0, 1, &fill)
	if !slices.Equal(x, rx) || !slices.Equal(y, ry) || split.Fitheta != 0 || b != 0 || fill != 1 {
		t.Fatal("typed theta spectra", split, x, y)
	}
}

func TestPartitionSpectrumPointers(t *testing.T) {
	for LM := int32(1); LM <= 3; LM++ {
		N := int32(1) << LM
		for _, budget := range []int32{0, 24, 400} {
			for _, fold := range []bool{false, true} {
				data := []byte{17, 255, 88, 1, 192, 0, 77, 43}
				var ec OpusT_ec_ctx
				Opus_ec_dec_init(nil, &ec, &data[0], 8)
				ctx := band_ctx{Fm: &mode48000_960_120, Fec: &ec, Fresynth: 1, Fremaining_bits: 600, Fseed: 123456}
				x := make([]float32, N+2)
				x[0] = 77
				x[N+1] = 88
				low := make([]float32, N+2)
				for i := range low {
					low[i] = float32(i+1) / 17
				}
				var lp *float32
				if fold {
					lp = &low[1]
				}
				before := slices.Clone(low)
				entropyInitGrowStack(12)
				runtime.GC()
				mask := quant_partition(nil, &ctx, &x[1], N, budget, 1, lp, LM, 1, 1)
				if x[0] != 77 || x[N+1] != 88 || !slices.Equal(low, before) || mask > 1 {
					t.Fatal("partition guards/mask", LM, budget, mask)
				}
			}
		}
	}
	// Sequential fold stores must retain overlapping-read behavior.
	data := []byte{0}
	var ec OpusT_ec_ctx
	Opus_ec_dec_init(nil, &ec, &data[0], 1)
	ctx := band_ctx{Fm: &mode48000_960_120, Fec: &ec, Fresynth: 1, Fseed: 123456}
	shared := []float32{.2, .3, .4, .5, .6}
	quant_partition(nil, &ctx, &shared[1], 4, 0, 1, &shared[0], 2, 1, 1)
	if shared[0] != .2 {
		t.Fatal("overlapping fold guard")
	}
}

func TestMonoSpectrumPointers(t *testing.T) {
	for LM := int32(1); LM <= 3; LM++ {
		N := int32(1) << LM
		for _, B := range []int32{1, N} {
			for _, tf := range []int32{-1, 0, 1} {
				if tf > 0 && B < 2 {
					continue
				}
				for _, budget := range []int32{0, 24, 400} {
					data := []byte{17, 255, 88, 1, 192, 0, 77, 43}
					var ec OpusT_ec_ctx
					Opus_ec_dec_init(nil, &ec, &data[0], 8)
					ctx := band_ctx{Fm: &mode48000_960_120, Fec: &ec, Fresynth: 1, Fremaining_bits: 600, Fseed: 123456, Ftf_change: tf}
					x, low, out, scratch := make([]float32, N+2), make([]float32, N+2), make([]float32, N+2), make([]float32, N+2)
					for i := range x {
						x[i] = float32(i+1) / 17
						low[i] = float32(i+2) / 19
						out[i] = 77
						scratch[i] = 88
					}
					entropyInitGrowStack(12)
					runtime.GC()
					quant_band(nil, &ctx, &x[1], N, budget, B, &low[1], LM, &out[1], 1, &scratch[1], (1<<B)-1)
					if out[0] != 77 || out[N+1] != 77 || scratch[0] != 88 || scratch[N+1] != 88 {
						t.Fatal("mono buffer guards", LM, B, tf, budget)
					}
				}
			}
		}
	}
}

func TestSpreadingPointers(t *testing.T) {
	bands := [3]int16{0, 1, 10}
	x := [12]float32{77}
	x[11] = 88
	weights := [2]int32{1, 1}
	state := [3]int32{256, 20, 1}
	entropyInitGrowStack(12)
	runtime.GC()
	decision := Opus_spreading_decision(nil, &bands[0], 2, 10, &x[1], &state[0], 2, &state[1], &state[2], 1, 2, 1, 1, &weights[0])
	if decision < SPREAD_NONE || decision > SPREAD_AGGRESSIVE || x[0] != 77 || x[11] != 88 {
		t.Fatal(decision, state, x)
	}
	short := [2]int16{0, 8}
	if Opus_spreading_decision(nil, &short[0], 1, 8, nil, nil, 0, nil, nil, 1, 1, 1, 1, nil) != SPREAD_NONE {
		t.Fatal("early exit")
	}
	if !validationPanics(func() { Opus_spreading_decision(nil, nil, 0, 0, nil, nil, 0, nil, nil, 0, 0, 1, 1, nil) }) {
		t.Fatal("end assertion")
	}
}

func TestQuantBandN1FieldAccesses(t *testing.T) {
	buffer := make([]byte, 16)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, unsafe.SliceData(buffer), uint32(len(buffer)))

	context := band_ctx{
		Fencode:         1,
		Fresynth:        1,
		Fec:             &encoder,
		Fremaining_bits: 24,
	}
	x := OpusT_celt_norm(-0.375)
	y := OpusT_celt_norm(0.625)
	lowband := OpusT_celt_norm(0)

	if got, want := quant_band_n1(nil, &context, &encoder, &x, &y, &lowband), uint32(1); got != want {
		t.Fatalf("coded dimensions: got %d, want %d", got, want)
	}
	Opus_ec_enc_done(nil, &encoder)

	if got, want := x, OpusT_celt_norm(-1); got != want {
		t.Fatalf("resynthesized X: got %v, want %v", got, want)
	}
	if got, want := y, OpusT_celt_norm(1); got != want {
		t.Fatalf("resynthesized Y: got %v, want %v", got, want)
	}
	if got, want := lowband, OpusT_celt_norm(-1); got != want {
		t.Fatalf("lowband output: got %v, want %v", got, want)
	}
	if got, want := context.Fremaining_bits, OpusT_opus_int32(8); got != want {
		t.Fatalf("remaining bits: got %d, want %d", got, want)
	}
	if got, want := encoder.Fnbits_total, int32(35); got != want {
		t.Fatalf("encoded sign bit count: got %d, want %d", got, want)
	}
}
