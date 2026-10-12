package libcshim

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"
)

func TestDiagnosticAssertFail(t *testing.T) {
	assertion, file := []byte("test expression"), []byte("test-file.c")
	runtime.GC()
	var caught any
	func() { defer func() { caught = recover() }(); AssertFail(nil, string(assertion), string(file), 17) }()
	if caught != "assertion failed: test expression (test-file.c:17)" {
		t.Fatal("typed assert diagnostic", caught)
	}
}

// Raw GoString reconstruction remains an ordinary ABI compatibility fixture;
// it is not a scoped checkptr claim for legacy integer string addresses.
func TestAssertFailLegacyABI(t *testing.T) {
	assertion, file := []byte("test expression\x00"), []byte("test-file.c\x00")
	runtime.GC()
	var caught any
	func() {
		defer func() { caught = recover() }()
		X__assert_fail(nil, uintptr(unsafe.Pointer(&assertion[0])), uintptr(unsafe.Pointer(&file[0])), 17, 0)
	}()
	if caught != "assertion failed: test expression (test-file.c:17)" {
		t.Fatal("raw assert diagnostic", caught)
	}
	runtime.KeepAlive(assertion)
	runtime.KeepAlive(file)
}

//go:noinline
func registeredGoInterior(tls *TLS) (uintptr, weak.Pointer[[64]byte]) {
	parent := new(struct {
		Header int64
		Child  *[64]byte
	})
	parent.Child = new([64]byte)
	parent.Child[0] = 77
	p := RegisterGoAllocation(tls, &parent.Header)
	return uintptr(unsafe.Pointer(p)), weak.Make(parent.Child)
}

type registryTestViews struct{ Child *[64]byte }

//go:noinline
func registerMetadataRoot(tls *TLS) (*uint64, weak.Pointer[[64]byte]) {
	p := RegisterGoAllocation(tls, new(uint64))
	child := new([64]byte)
	child[0] = 77
	if !SetGoAllocationMetadata(tls, p, registryTestViews{Child: child}) {
		panic("metadata set")
	}
	return p, weak.Make(child)
}

func TestRegisteredGoAllocationMetadata(t *testing.T) {
	tls := NewTLS()
	defer tls.Close()
	p, child := registerMetadataRoot(tls)
	runtime.GC()
	view, ok := GoAllocationMetadata[uint64, registryTestViews](tls, p)
	if !ok || child.Value() == nil || view.Child != child.Value() || view.Child[0] != 77 {
		t.Fatal("metadata actual Go owner")
	}
	if _, ok := GoAllocationMetadata[uint64, int](tls, p); ok {
		t.Fatal("metadata type mismatch")
	}
	wrong := (*int64)(unsafe.Pointer(p))
	if SetGoAllocationMetadata(tls, wrong, 1) || RemoveGoAllocation(tls, wrong) {
		t.Fatal("root type mismatch")
	}
	bytes := XmallocPointer(tls, 8)
	if SetGoAllocationMetadata(tls, (*uint64)(bytes), view) || RemoveGoAllocation(tls, (*uint64)(bytes)) {
		t.Fatal("opaque bytes promoted")
	}
	XfreePointer(tls, bytes)
	if !RemoveGoAllocation(tls, p) || RemoveGoAllocation(tls, p) {
		t.Fatal("exact-root cleanup/idempotence")
	}
	if _, ok := GoAllocationMetadata[uint64, registryTestViews](tls, p); ok {
		t.Fatal("freed metadata")
	}
	if SetGoAllocationMetadata[uint64](nil, p, view) || RemoveGoAllocation[uint64](tls, nil) {
		t.Fatal("nil registry")
	}
}

func TestRegisteredGoAllocationLookup(t *testing.T) {
	tls := NewTLS()
	defer tls.Close()
	root := RegisterGoAllocation(tls, new(int32))
	handle := uintptr(unsafe.Pointer(root))
	if RegisteredGoAllocation[int32](tls, handle) != root || RegisteredGoAllocation[int64](tls, handle) != nil || RegisteredGoAllocation[int32](nil, handle) != nil || RegisteredGoAllocation[int32](tls, 0) != nil {
		t.Fatal("typed allocation lookup")
	}
	bytes := Xmalloc(tls, 32)
	if RegisteredGoAllocation[int32](tls, bytes) != nil {
		t.Fatal("opaque byte allocation promoted to Go object")
	}
	Xfree(tls, handle)
	if RegisteredGoAllocation[int32](tls, handle) != nil {
		t.Fatal("freed registered root")
	}
}

func TestRegisterGoAllocation(t *testing.T) {
	tls := NewTLS()
	defer tls.Close()
	address, child := registeredGoInterior(tls)
	for i := 0; i < 3; i++ {
		runtime.GC()
		if child.Value() == nil || child.Value()[0] != 77 {
			t.Fatal("interior Go root must scan its actual parent")
		}
	}
	if len(tls.heap) != 1 {
		t.Fatal("Go registry insertion")
	}
	Xfree(tls, address)
	if len(tls.heap) != 0 {
		t.Fatal("Go registry release")
	}
	if RegisterGoAllocation[int64](nil, new(int64)) != nil || RegisterGoAllocation[int64](tls, nil) != nil {
		t.Fatal("nil registration")
	}
}

func TestMallocPointer(t *testing.T) {
	if XmallocPointer(nil, ^uint64(0)) != nil || Xmalloc(nil, 1) != 0 {
		t.Fatal("nil TLS")
	}
	for _, size := range []uint64{0, 1, 15, 16, 17, 255, 4096} {
		tls := NewTLS()
		p := XmallocPointer(tls, size)
		if p == nil || uintptr(p)&15 != 0 || len(tls.heap) != 1 || tls.heap[uintptr(p)] == nil {
			t.Fatal("allocation/alignment/registry", size)
		}
		n := int(size)
		if n == 0 {
			n = 1
		}
		view := unsafe.Slice((*byte)(p), n)
		for _, b := range view {
			if b != 0 {
				t.Fatal("zeroed storage")
			}
		}
		view[n-1] = 77
		runtime.GC()
		if view[n-1] != 77 {
			t.Fatal("lifetime")
		}
		XfreePointer(tls, p)
		if len(tls.heap) != 0 {
			t.Fatal("unregister")
		}
		tls.Close()
		runtime.GC()
		if *(*byte)(unsafe.Add(p, n-1)) != 77 {
			t.Fatal("typed backing ownership")
		}
	}
	tls := NewTLS()
	defer tls.Close()
	a, b := XmallocPointer(tls, 32), XmallocPointer(tls, 32)
	if a == b || len(tls.heap) != 2 {
		t.Fatal("independent allocations")
	}
	Xfree(tls, uintptr(a))
	XfreePointer(tls, b)
	if len(tls.heap) != 0 {
		t.Fatal("mixed free boundaries")
	}
	panicked := false
	func() { defer func() { panicked = recover() != nil }(); XmallocPointer(tls, ^uint64(0)) }()
	if !panicked || len(tls.heap) != 0 {
		t.Fatal("oversized allocation")
	}
}

// C converts a signed integer to a wider unsigned type by value modulo
// 2^N - i.e. it sign-extends first. Transpiled expressions like
// (opus_uint32)idx << 10 on a negative opus_int8 depend on it.
func TestSignedToUnsignedSignExtends(t *testing.T) {
	cases := []struct {
		name      string
		got, want uint64
	}{
		{"Uint32FromInt8(-4)", uint64(Uint32FromInt8(-4)), 0xFFFFFFFC},
		{"Uint32FromInt16(-2)", uint64(Uint32FromInt16(-2)), 0xFFFFFFFE},
		{"Uint64FromInt16(-3)", Uint64FromInt16(-3), 0xFFFFFFFFFFFFFFFD},
		{"Uint64FromInt32(-1)", Uint64FromInt32(-1), 0xFFFFFFFFFFFFFFFF},
		{"Uint32FromInt8(5)", uint64(Uint32FromInt8(5)), 5},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %#x, want %#x", c.name, c.got, c.want)
		}
	}
	// (uintptr)(int32)-1 is all ones on every architecture, so p + that
	// offset steps backwards instead of ~4 GiB forwards on 64-bit.
	if got := UintptrFromInt32(-1); got != ^uintptr(0) {
		t.Errorf("UintptrFromInt32(-1) = %#x, want %#x", got, ^uintptr(0))
	}
	if got := Int32FromUint32(Uint32FromInt8(-4) << 10); got != -4096 {
		t.Errorf("(int32)((uint32)(int8)-4 << 10) = %d, want -4096", got)
	}
}
