package libcshim

import (
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

func TestGoStringPointer(t *testing.T) {
	if GoStringPointer(nil) != "" {
		t.Fatal("nil C string")
	}
	for _, tc := range []struct {
		bytes  []byte
		offset int
		want   string
	}{
		{[]byte{0}, 0, ""},
		{[]byte("file%name.c\x00ignored"), 0, "file%name.c"},
		{[]byte("prefix:message\x00tail"), 7, "message"},
		{[]byte{0xff, 0x80, 'x', 0}, 0, string([]byte{0xff, 0x80, 'x'})},
	} {
		pointer := func() *byte {
			storage := make([]byte, len(tc.bytes))
			copy(storage, tc.bytes)
			return &storage[tc.offset]
		}()
		runtime.GC()
		if got := GoStringPointer(pointer); got != tc.want {
			t.Fatalf("got %q want %q", got, tc.want)
		}
	}
	const limit = 1 << 20
	storage := make([]byte, limit+1)
	for i := range storage {
		storage[i] = 'x'
	}
	if got := GoStringPointer(unsafe.SliceData(storage)); got != strings.Repeat("x", limit) {
		t.Fatal("C string scan limit")
	}
}
