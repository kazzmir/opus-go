// Package resample converts PCM between sample rates - e.g. 44.1 kHz WAV
// to the 48 kHz Opus encodes at, since libopus only accepts
// 8/12/16/24/48 kHz input.
//
// It is a polyphase windowed-sinc (Blackman) filter. The output/input
// rate ratio is rational, so one table of phases covers every output
// sample exactly: no per-sample trig and no drift, however long the
// stream.
package resample

import "math"

// halfTaps is the filter half-width in input samples at the cutoff;
// 32 taps keeps the passband flat and aliasing far below 16-bit noise.
const halfTaps = 16

// Resampler converts one interleaved stream from one rate to another,
// incrementally: feed input with Process, then call Flush once at the
// end. Across the whole stream it produces exactly
// ceil(inputFrames * to / from) frames.
type Resampler struct {
	channels int
	up, down int64 // output/input rate ratio, reduced
	half     int
	taps     int
	table    []float32 // [phase][tap]

	buf      []float32 // pending input, interleaved; buf[0] is input frame base
	base     int64
	inFrames int64 // input frames seen so far
	next     int64 // next output frame to produce
}

// New returns a Resampler for interleaved audio with channels channels,
// from Hz to to Hz.
func New(channels, from, to int) *Resampler {
	if channels <= 0 || from <= 0 || to <= 0 {
		panic("resample: channels and rates must be positive")
	}
	g := gcd(from, to)
	up, down := to/g, from/g
	// Cut off at the lower Nyquist, a little early for the window's
	// transition band.
	cutoff := 0.97 * math.Min(1, float64(to)/float64(from))
	half := int(math.Ceil(halfTaps / cutoff))
	taps := 2 * half

	// table[p][j] weighs input frame base+j-half+1 for output phase p,
	// whose position lies p/up of the way past input frame base.
	table := make([]float32, up*taps)
	for p := range up {
		row := table[p*taps : (p+1)*taps]
		var sum float64
		for j := range taps {
			d := float64(j-half+1) - float64(p)/float64(up)
			u := d / float64(half)
			if u <= -1 || u >= 1 {
				continue
			}
			win := 0.42 + 0.5*math.Cos(math.Pi*u) + 0.08*math.Cos(2*math.Pi*u)
			v := cutoff * sinc(cutoff*d) * win
			row[j] = float32(v)
			sum += v
		}
		for j := range row { // unity gain at DC for every phase
			row[j] = float32(float64(row[j]) / sum)
		}
	}
	return &Resampler{
		channels: channels,
		up:       int64(up),
		down:     int64(down),
		half:     half,
		taps:     taps,
		table:    table,
	}
}

// Process consumes interleaved input and returns every output frame whose
// filter window it now fully covers.
func (r *Resampler) Process(in []float32) []float32 {
	r.buf = append(r.buf, in...)
	r.inFrames += int64(len(in) / r.channels)
	return r.produce(false)
}

// Flush ends the stream (treating input past its end as silence) and
// returns the remaining output.
func (r *Resampler) Flush() []float32 {
	return r.produce(true)
}

// ProcessInt16 is Process for 16-bit PCM.
func (r *Resampler) ProcessInt16(in []int16) []int16 {
	f := make([]float32, len(in))
	for i, v := range in {
		f[i] = float32(v) / 32768
	}
	return toInt16(r.Process(f))
}

// FlushInt16 is Flush for 16-bit PCM.
func (r *Resampler) FlushInt16() []int16 {
	return toInt16(r.Flush())
}

func (r *Resampler) produce(final bool) []float32 {
	end := r.next // one past the last output frame to produce
	if final {
		end = (r.inFrames*r.up + r.down - 1) / r.down
	} else {
		// Output n needs input frames up to n*down/up + half.
		for ; ; end++ {
			if (end*r.down)/r.up+int64(r.half) >= r.inFrames {
				break
			}
		}
	}
	out := make([]float32, 0, int(end-r.next)*r.channels)
	for ; r.next < end; r.next++ {
		pos := r.next * r.down
		center := pos / r.up
		row := r.table[int(pos%r.up)*r.taps:][:r.taps]
		for c := range r.channels {
			var acc float32
			for j, w := range row {
				k := center + int64(j-r.half+1)
				if k < 0 || k >= r.inFrames {
					continue
				}
				acc += w * r.buf[int(k-r.base)*r.channels+c]
			}
			out = append(out, acc)
		}
	}
	// Drop input no future output frame reaches back to.
	keepFrom := (r.next*r.down)/r.up - int64(r.half) + 1
	if drop := keepFrom - r.base; drop > 0 {
		drop = min(drop, r.inFrames-r.base)
		r.buf = append(r.buf[:0], r.buf[int(drop)*r.channels:]...)
		r.base += drop
	}
	return out
}

// Int16 resamples a whole interleaved 16-bit buffer in one call.
func Int16(in []int16, channels, from, to int) []int16 {
	r := New(channels, from, to)
	return append(r.ProcessInt16(in), r.FlushInt16()...)
}

func toInt16(in []float32) []int16 {
	out := make([]int16, len(in))
	for i, v := range in {
		s := math.Round(float64(v) * 32768)
		out[i] = int16(max(-32768, min(32767, s)))
	}
	return out
}

func sinc(x float64) float64 {
	if x == 0 {
		return 1
	}
	return math.Sin(math.Pi*x) / (math.Pi * x)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
