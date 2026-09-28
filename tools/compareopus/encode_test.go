//go:build compareopus && cgo

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kazzmir/opus-go/wav"
)

// Explicitly opt in with: go test -tags compareopus ./tools/compareopus
func TestEncodeSilence(t *testing.T) {
	for _, channels := range []int{1, 2} {
		for _, frames := range []int{1, 648, 960, 961, 1920} {
			for _, limit := range []int{0, 1} {
				t.Run(fmt.Sprintf("channels%d/frames%d/limit%d", channels, frames, limit), func(t *testing.T) {
					path := filepath.Join(t.TempDir(), "silence.wav")
					f, err := os.Create(path)
					if err != nil {
						t.Fatal(err)
					}
					defer f.Close()
					w, err := wav.NewWriter(f, 48000, channels)
					if err != nil {
						t.Fatal(err)
					}
					if err := w.WriteInt16PCM(make([]int16, frames*channels)); err != nil {
						t.Fatal(err)
					}
					if err := w.Close(); err != nil {
						t.Fatal(err)
					}
					if err := f.Close(); err != nil {
						t.Fatal(err)
					}
					opt := encodeOptions{bitrate: 64000, complexity: 10, vbr: true, exact: true}
					if err := compareEncode(path, opt, limit); err != nil {
						t.Fatal(err)
					}
				})
			}
		}
	}
}
