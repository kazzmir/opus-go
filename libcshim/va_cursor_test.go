package libcshim

import (
	"encoding/binary"
	"runtime"
	"testing"
	"unsafe"
)

func TestVaCursorPointers(t *testing.T) {
	for _, want := range []int32{-2147483648, -7, 0, 2147483647} {
		word := int(want)
		cursor := NewVaCursor(unsafe.Pointer(&word))
		runtime.GC()
		if got := cursor.Int32(); got != want {
			t.Fatal("native scalar word", got, want)
		}
		if cursor.base != unsafe.Pointer(&word) || cursor.offset != uint(PtrSize) {
			t.Fatal("advance offset without forming one-past pointer")
		}
	}
	storage := make([]byte, PtrSize)
	if PtrSize == 8 {
		binary.NativeEndian.PutUint64(storage, 0x9abcdef080000007)
	} else {
		binary.NativeEndian.PutUint32(storage, 0x80000007)
	}
	cursor := NewVaCursor(unsafe.Pointer(unsafe.SliceData(storage)))
	if cursor.Int32() != -2147483641 {
		t.Fatal("low 32 bits of native word")
	}
	value := int32(123)
	slots := [2]*int32{&value, nil}
	cursor = NewVaCursor(unsafe.Pointer(&slots[0]))
	runtime.GC()
	pointer := (*int32)(cursor.Pointer())
	if pointer != &value || cursor.Pointer() != nil {
		t.Fatal("actual pointer slots")
	}
	runtime.GC()
	if *pointer != 123 {
		t.Fatal("typed pointer retains pointee")
	}
}
