package opus

import "testing"

// Lookahead is at the encoder's own rate; PreSkip is always 48 kHz.
func TestPreSkipIs48kHz(t *testing.T) {
	for _, rate := range []int{8000, 12000, 16000, 24000, 48000} {
		enc, err := NewEncoder(rate, 1, ApplicationAudio)
		if err != nil {
			t.Fatal(err)
		}
		lookahead, err := enc.Lookahead()
		if err != nil {
			t.Fatal(err)
		}
		preSkip, err := enc.PreSkip()
		if err != nil {
			t.Fatal(err)
		}
		enc.Close()
		// Fs/400 + Fs/250 in libopus: 6.5 ms.
		if want := rate * 13 / 2000; lookahead != want {
			t.Errorf("%d Hz: Lookahead %d, want %d", rate, lookahead, want)
		}
		if preSkip != 312 {
			t.Errorf("%d Hz: PreSkip %d, want 312", rate, preSkip)
		}
	}
}
