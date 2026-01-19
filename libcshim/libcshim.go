package libcshim

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"sync"
	"unsafe"
)

// This package implements the minimal C-runtime-like surface that the
// ccgo-generated Opus bindings rely on.
//
// It is stdlib-only and intentionally small; it is not a general-purpose libc.

type Tpthread_key_t = uint32

// TLS is a per-decoder state object used by ccgo-generated code.
//
// It provides:
// - a simple LIFO "stack" allocator (Alloc/Free) for temporary scratch
// - a very small malloc/free implementation for long-lived allocations
// - pthread-specific storage emulation for the ccgo TLS pseudostack
type TLS struct {
	stack []byte
	sp    int

	heapMu sync.Mutex
	heap   map[uintptr]*heapAlloc

	keysMu sync.Mutex
	keys   map[Tpthread_key_t]uintptr
}

type heapAlloc struct {
	buf []byte
}

func NewTLS() *TLS {
	return &TLS{
		stack: make([]byte, 64<<10),
		heap:  make(map[uintptr]*heapAlloc),
		keys:  make(map[Tpthread_key_t]uintptr),
	}
}

func (t *TLS) Close() {
	if t == nil {
		return
	}
	t.heapMu.Lock()
	t.heap = nil
	t.heapMu.Unlock()

	t.keysMu.Lock()
	t.keys = nil
	t.keysMu.Unlock()

	t.stack = nil
	t.sp = 0
}

func align16(n int) int { return (n + 15) &^ 15 }

func ptrAdd(p unsafe.Pointer, off uintptr) unsafe.Pointer { return unsafe.Add(p, off) }

// Alloc returns a pointer to at least n bytes of temporary memory.
// The memory is 16-byte aligned and valid until the matching Free.
func (t *TLS) Alloc(n int) uintptr {
	if t == nil {
		return 0
	}
	if n <= 0 {
		return 0
	}
	n = align16(n)
	need := t.sp + n
	if need > len(t.stack) {
		newCap := len(t.stack)
		if newCap == 0 {
			newCap = 64 << 10
		}
		for need > newCap {
			newCap *= 2
		}
		newStack := make([]byte, newCap)
		copy(newStack, t.stack)
		t.stack = newStack
	}
	base := unsafe.Pointer(unsafe.SliceData(t.stack))
	p := uintptr(ptrAdd(base, uintptr(t.sp)))
	t.sp += n
	return p
}

// Free releases the last Alloc(n) region.
func (t *TLS) Free(n int) {
	if t == nil {
		return
	}
	if n <= 0 {
		return
	}
	n = align16(n)
	if t.sp < n {
		panic("libcshim: TLS.Free underflow")
	}
	t.sp -= n
}

// ---- Varargs helpers (very small va_list model) ----

const ptrSize = int(unsafe.Sizeof(uintptr(0)))

// PtrSize is the size of a pointer in bytes for the current architecture.
const PtrSize = ptrSize

// VaList writes args into memory starting at p (as uintptr-sized slots) and
// returns p. This is sufficient for Opus' use of ctl varargs.
func VaList(p uintptr, args ...interface{}) uintptr {
	base := unsafe.Pointer(p)
	for i, a := range args {
		slot := ptrAdd(base, uintptr(i*ptrSize))
		switch v := a.(type) {
		case int32:
			*(*uintptr)(slot) = uintptr(uint32(v))
		case uint32:
			*(*uintptr)(slot) = uintptr(v)
		case uintptr:
			*(*uintptr)(slot) = v
		case int:
			*(*uintptr)(slot) = uintptr(v)
		default:
			panic(fmt.Sprintf("libcshim: unsupported VaList arg type %T", a))
		}
	}
	return p
}

// VaUintptr reads one uintptr argument and advances ap.
func VaUintptr(ap *uintptr) uintptr {
	v := *(*uintptr)(unsafe.Pointer(*ap))
	*ap += uintptr(ptrSize)
	return v
}

// VaInt32 reads one int32 argument (stored as a uintptr slot) and advances ap.
func VaInt32(ap *uintptr) int32 {
	v := *(*uintptr)(unsafe.Pointer(*ap))
	*ap += uintptr(ptrSize)
	return int32(uint32(v))
}

// ---- Memory allocation / libc-ish primitives ----

// Xmalloc allocates size bytes and returns a 16-byte aligned pointer.
func Xmalloc(tls *TLS, size uint64) uintptr {
	if tls == nil {
		return 0
	}
	if size == 0 {
		size = 1
	}
	if size > uint64(^uint(0)>>1) {
		panic("libcshim: Xmalloc too large")
	}
	buf := make([]byte, int(size)+16)
	base := uintptr(unsafe.Pointer(unsafe.SliceData(buf)))
	aligned := (base + 15) &^ 15

	tls.heapMu.Lock()
	if tls.heap == nil {
		tls.heap = make(map[uintptr]*heapAlloc)
	}
	tls.heap[aligned] = &heapAlloc{buf: buf}
	tls.heapMu.Unlock()
	return aligned
}

func Xfree(tls *TLS, p uintptr) {
	if tls == nil || p == 0 {
		return
	}
	tls.heapMu.Lock()
	if tls.heap != nil {
		delete(tls.heap, p)
	}
	tls.heapMu.Unlock()
}

func Xmemset(_ *TLS, s uintptr, c int32, n uint64) uintptr {
	if s == 0 || n == 0 {
		return s
	}
	if n > uint64(^uint(0)>>1) {
		panic("libcshim: Xmemset too large")
	}
	b := unsafe.Slice((*byte)(unsafe.Pointer(s)), int(n))
	fill := byte(c)
	for i := range b {
		b[i] = fill
	}
	return s
}

func Xmemcpy(_ *TLS, dst, src uintptr, n uint64) uintptr {
	if dst == 0 || src == 0 || n == 0 {
		return dst
	}
	if n > uint64(^uint(0)>>1) {
		panic("libcshim: Xmemcpy too large")
	}
	d := unsafe.Slice((*byte)(unsafe.Pointer(dst)), int(n))
	s := unsafe.Slice((*byte)(unsafe.Pointer(src)), int(n))
	copy(d, s)
	return dst
}

func Xmemmove(_ *TLS, dst, src uintptr, n uint64) uintptr {
	// Go's copy is memmove-safe.
	return Xmemcpy(nil, dst, src, n)
}

// ---- pthread TLS emulation (per *TLS instance) ----

func Xpthread_getspecific(tls *TLS, key uint32) uintptr {
	if tls == nil {
		return 0
	}
	tls.keysMu.Lock()
	defer tls.keysMu.Unlock()
	if tls.keys == nil {
		return 0
	}
	return tls.keys[Tpthread_key_t(key)]
}

func Xpthread_setspecific(tls *TLS, key uint32, value uintptr) int32 {
	if tls == nil {
		return -1
	}
	tls.keysMu.Lock()
	if tls.keys == nil {
		tls.keys = make(map[Tpthread_key_t]uintptr)
	}
	tls.keys[Tpthread_key_t(key)] = value
	tls.keysMu.Unlock()
	return 0
}

// ---- Fatal/assert/debug ----

var Xstderr uintptr = 0

func Xfprintf(_ *TLS, _ uintptr, format uintptr, ap uintptr) int32 {
	// Best-effort debug printing. Used by Opus_celt_fatal.
	_ = ap
	fmtStr := GoString(format)
	_, _ = fmt.Fprintf(os.Stderr, "%s", fmtStr)
	return 0
}

func Xabort(_ *TLS) { panic("libcshim: abort") }

func X__assert_fail(_ *TLS, assertion, file uintptr, line int32, _ uintptr) {
	panic(fmt.Sprintf("assertion failed: %s (%s:%d)", GoString(assertion), GoString(file), line))
}

// ---- Math/bit helpers ----

func Xabs(_ *TLS, x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func Xfabs(_ *TLS, x float64) float64   { return math.Abs(x) }
func Xfloor(_ *TLS, x float64) float64  { return math.Floor(x) }
func Xsqrt(_ *TLS, x float64) float64   { return math.Sqrt(x) }
func Xcos(_ *TLS, x float64) float64    { return math.Cos(x) }
func Xsin(_ *TLS, x float64) float64    { return math.Sin(x) }
func Xlog(_ *TLS, x float64) float64    { return math.Log(x) }
func Xlog10(_ *TLS, x float64) float64  { return math.Log10(x) }
func Xacos(_ *TLS, x float64) float64   { return math.Acos(x) }
func Xexp(_ *TLS, x float64) float64    { return math.Exp(x) }
func Xpow(_ *TLS, x, y float64) float64 { return math.Pow(x, y) }

// Xlrintf rounds to the nearest integer value, using ties-to-even.
func Xlrintf(_ *TLS, x float32) int64 { return int64(math.RoundToEven(float64(x))) }

// X__builtin_clz returns the number of leading 0-bits in x.
// C's __builtin_clz is undefined for 0; we return 32 for 0.
func X__builtin_clz(_ *TLS, x uint32) int32 {
	if x == 0 {
		return 32
	}
	return int32(bits.LeadingZeros32(x))
}

// ---- Conversions/helpers used by the generated bindings ----

func Bool(v bool) bool { return v }
func BoolInt32(v bool) int32 {
	if v {
		return 1
	}
	return 0
}
func BoolInt8(v bool) int8 {
	if v {
		return 1
	}
	return 0
}
func BoolInt64(v bool) int64 {
	if v {
		return 1
	}
	return 0
}
func BoolUint32(v bool) uint32 {
	if v {
		return 1
	}
	return 0
}
func BoolUint64(v bool) uint64 {
	if v {
		return 1
	}
	return 0
}
func BoolUintptr(v bool) uintptr {
	if v {
		return 1
	}
	return 0
}

func Float32FromFloat32(v float32) float32 { return v }
func Float32FromFloat64(v float64) float32 { return float32(v) }
func Float32FromInt32(v int32) float32     { return float32(v) }

func Float64FromFloat32(v float32) float64 { return float64(v) }
func Float64FromFloat64(v float64) float64 { return v }
func Float64FromInt32(v int32) float64     { return float64(v) }

func Int16FromInt32(v int32) int16 { return int16(v) }
func Int16FromUint8(v uint8) int16 { return int16(v) }

func Int32FromInt32(v int32) int32   { return v }
func Int32FromInt64(v int64) int32   { return int32(v) }
func Int32FromUint16(v uint16) int32 { return int32(v) }
func Int32FromUint32(v uint32) int32 { return int32(v) }
func Int32FromUint64(v uint64) int32 { return int32(v) }
func Int32FromUint8(v uint8) int32   { return int32(v) }
func Int64FromInt32(v int32) int64   { return int64(v) }

func Uint16FromInt32(v int32) uint16   { return uint16(v) }
func Uint16FromInt16(v int16) uint16   { return uint16(v) }
func Uint32FromInt16(v int16) uint32   { return uint32(uint16(v)) }
func Uint32FromInt32(v int32) uint32   { return uint32(v) }
func Uint32FromInt8(v int8) uint32     { return uint32(uint8(v)) }
func Uint32FromUint32(v uint32) uint32 { return v }
func Uint64FromInt16(v int16) uint64   { return uint64(uint16(v)) }
func Uint64FromInt32(v int32) uint64   { return uint64(uint32(v)) }
func Uint64FromInt64(v int64) uint64   { return uint64(v) }
func Uint64FromUint64(v uint64) uint64 { return v }
func Uint8FromInt16(v int16) uint8     { return uint8(uint16(v)) }
func Uint8FromInt32(v int32) uint8     { return uint8(v) }
func UintptrFromInt32(v int32) uintptr { return uintptr(uint32(v)) }

// GoString reads a NUL-terminated C string from p.
func GoString(p uintptr) string {
	if p == 0 {
		return ""
	}
	const max = 1 << 20
	buf := make([]byte, 0, 64)
	base := unsafe.Pointer(p)
	for i := 0; i < max; i++ {
		b := *(*byte)(ptrAdd(base, uintptr(i)))
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf)
}

// ---- Pointer helpers for higher-level packages (to reduce unsafe usage there) ----

func PtrUint8(s []uint8) uintptr {
	if len(s) == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(unsafe.SliceData(s)))
}

func PtrInt16(s []int16) uintptr {
	if len(s) == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(unsafe.SliceData(s)))
}

func PtrByte(s []byte) uintptr {
	if len(s) == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(unsafe.SliceData(s)))
}

func LoadInt32(p uintptr) int32 {
	if p == 0 {
		return 0
	}
	return *(*int32)(unsafe.Pointer(p))
}

func StoreInt32(p uintptr, v int32) {
	if p == 0 {
		return
	}
	*(*int32)(unsafe.Pointer(p)) = v
}

func LoadUintptr(p uintptr) uintptr {
	if p == 0 {
		return 0
	}
	return *(*uintptr)(unsafe.Pointer(p))
}

func StoreUintptr(p uintptr, v uintptr) {
	if p == 0 {
		return
	}
	*(*uintptr)(unsafe.Pointer(p)) = v
}

func LoadUintptrAt(base uintptr, off uintptr) uintptr {
	return LoadUintptr(base + off)
}

func StoreUintptrAt(base uintptr, off uintptr, v uintptr) {
	StoreUintptr(base+off, v)
}
