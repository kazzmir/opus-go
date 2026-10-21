package opuscc

import (
	"io"
	"os"
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// Nil and embedded-NUL cases preserve the Go shim's existing behavior.
func TestCeltFatalTypedPointers(t *testing.T) {
	checkDiagnosticFatal(t, func() { Opus_celt_fatal(nil, nil, nil, 0) }, "Fatal (internal) error in , line 0: \n")
	message := []byte("first\x00ignored\x00")
	file := []byte("file\x00ignored\x00")
	checkDiagnosticFatal(t, func() { Opus_celt_fatal(nil, &message[0], &file[0], -7) }, "Fatal (internal) error in file, line -7: first\n")
}

// Diagnostics from /tmp/opencode/celt_fatal_ref.c linked with
// ../opus/.libs/libopus.a. C terminates with SIGABRT; the shim panics instead.
func TestCeltFatalCReference(t *testing.T) {
	for _, tc := range []struct {
		message, file string
		line          int32
		want          string
	}{
		{"assertion failed: count > 0", "celt/test.c", 123, "Fatal (internal) error in celt/test.c, line 123: assertion failed: count > 0\n"},
		{"100% invalid\nsecond line", "file%name.c", -7, "Fatal (internal) error in file%name.c, line -7: 100% invalid\nsecond line\n"},
	} {
		t.Run(tc.file, func(t *testing.T) {
			tls := libc.NewTLS()
			t.Cleanup(func() { tls.Close() })
			cstring := func(s string) *byte {
				storage := make([]byte, len(s)+1)
				copy(storage, s)
				return unsafe.SliceData(storage)
			}
			message, file := cstring(tc.message), cstring(tc.file)
			entropyInitGrowStack(12)
			runtime.GC()
			output, err := os.CreateTemp(t.TempDir(), "stderr")
			if err != nil {
				t.Fatal(err)
			}
			defer output.Close()
			var caught interface{}
			func() {
				old := os.Stderr
				os.Stderr = output
				defer func() { os.Stderr = old; caught = recover() }()
				Opus_celt_fatal(tls, message, file, tc.line)
			}()
			if caught != "libcshim: abort" {
				t.Fatalf("panic = %v", caught)
			}
			if _, err := output.Seek(0, 0); err != nil {
				t.Fatal(err)
			}
			got, err := io.ReadAll(output)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tc.want {
				t.Fatalf("stderr = %q, want %q", got, tc.want)
			}
		})
	}
}
