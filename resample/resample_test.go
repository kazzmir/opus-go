package resample

import (
	"math"
	"slices"
	"testing"
)

func sine(rate, channels, frames int, hz float64) []float32 {
	s := make([]float32, frames*channels)
	for i := range frames {
		v := float32(0.5 * math.Sin(2*math.Pi*hz*float64(i)/float64(rate)))
		for c := range channels {
			s[i*channels+c] = v
		}
	}
	return s
}

// pitch estimates channel 0's frequency from zero crossings, ignoring the
// filter's edge transients.
func pitch(s []float32, channels, rate int) float64 {
	frames := len(s) / channels
	lo, hi := frames/10, frames*9/10
	n := 0
	for i := lo + 1; i < hi; i++ {
		if (s[(i-1)*channels] < 0) != (s[i*channels] < 0) {
			n++
		}
	}
	return float64(n) / 2 / (float64(hi-lo) / float64(rate))
}

func TestRates(t *testing.T) {
	const hz = 440
	for _, tc := range []struct{ from, to, channels, frames int }{
		{44100, 48000, 2, 44100},
		{22050, 48000, 1, 22050 + 17},
		{48000, 16000, 1, 48000},
		{32000, 48000, 2, 1},
		{44100, 48000, 1, 0},
	} {
		in := sine(tc.from, tc.channels, tc.frames, hz)
		r := New(tc.channels, tc.from, tc.to)
		out := append(r.Process(in), r.Flush()...)
		wantFrames := (tc.frames*tc.to + tc.from - 1) / tc.from
		if got := len(out) / tc.channels; got != wantFrames {
			t.Errorf("%+v: %d output frames, want %d", tc, got, wantFrames)
		}
		if tc.frames < tc.from/2 {
			continue
		}
		if got := pitch(out, tc.channels, tc.to); math.Abs(got-hz) > 2 {
			t.Errorf("%+v: pitch %.1f Hz, want %d", tc, got, hz)
		}
		// Steady-state amplitude is preserved (unity passband gain).
		var peak float32
		for _, v := range out[len(out)/4 : len(out)*3/4] {
			peak = max(peak, v)
		}
		if math.Abs(float64(peak)-0.5) > 0.01 {
			t.Errorf("%+v: peak %.3f, want 0.5", tc, peak)
		}
	}
}

// TestStreamingMatchesOneShot feeds input in uneven chunks and checks the
// output is identical to resampling it all at once.
func TestStreamingMatchesOneShot(t *testing.T) {
	in := sine(44100, 2, 30000, 1000)
	r := New(2, 44100, 48000)
	want := append(r.Process(in), r.Flush()...)

	r = New(2, 44100, 48000)
	var got []float32
	for off, step := 0, 1; off < len(in); step = step*3%997 + 1 {
		end := min(off+step*2, len(in))
		got = append(got, r.Process(in[off:end])...)
		off = end
	}
	got = append(got, r.Flush()...)
	if !slices.Equal(got, want) {
		t.Fatalf("streamed output (%d samples) differs from one-shot (%d)", len(got), len(want))
	}
}

func TestInt16(t *testing.T) {
	in := make([]int16, 4410)
	for i := range in {
		in[i] = int16(10000 * math.Sin(2*math.Pi*300*float64(i)/44100))
	}
	out := Int16(in, 1, 44100, 48000)
	if len(out) != 4800 {
		t.Fatalf("%d samples, want 4800", len(out))
	}
}
