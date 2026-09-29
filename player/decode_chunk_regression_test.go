package player

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// Different read paths reach the decoder at different stack depths. A frame's
// entropy context must survive stack growth even across legacy uintptr APIs.
func TestDecodeChunkSizeRegression(t *testing.T) {
	const length = 32768
	open := func(t *testing.T) *OpusPlayer[int16] {
		t.Helper()
		p, err := NewPlayerFromFile(testFilePath, true)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { p.decoder.Close() })
		return p
	}
	var reference bytes.Buffer
	if _, err := io.CopyN(&reference, open(t), length); err != nil {
		t.Fatal(err)
	}
	for _, size := range []int{4, 512, 8192} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			p := open(t)
			got := make([]byte, length)
			for offset := 0; offset < len(got); offset += size {
				if _, err := io.ReadFull(p, got[offset:min(offset+size, len(got))]); err != nil {
					t.Fatal(err)
				}
			}
			if !bytes.Equal(got, reference.Bytes()) {
				for i := range got {
					if got[i] != reference.Bytes()[i] {
						t.Fatalf("PCM differs at byte %d: got %02x want %02x", i, got[i], reference.Bytes()[i])
					}
				}
			}
		})
	}
}
