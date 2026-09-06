package opuscc

import (
	"io"
	"os"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

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
			cstring := func(s string) uintptr {
				p := libc.Xmalloc(tls, uint64(len(s)+1))
				t.Cleanup(func() { libc.Xfree(tls, p) })
				copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(s)+1), append([]byte(s), 0))
				return p
			}
			message, file := cstring(tc.message), cstring(tc.file)
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
