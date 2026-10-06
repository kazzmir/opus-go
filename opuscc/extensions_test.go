package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestExtensionIteratorInitPointers(t *testing.T) {
	owned := func() OpusT_OpusExtensionIterator {
		packet := [3]byte{7, 99, 0}
		var st OpusT_OpusExtensionIterator
		Opus_opus_extension_iterator_init(nil, &st, &packet[0], 3, 48)
		return st
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *owned.Fdata != 7 || owned.Fcurr_data != owned.Fdata || owned.Frepeat_data != owned.Fdata || owned.Fsrc_data != nil || owned.Flast_long != nil || owned.Fcurr_len != 3 || owned.Fframe_max != 48 {
		t.Fatal("state/ownership")
	}
	guarded := struct {
		before uint64
		state  OpusT_OpusExtensionIterator
		after  uint64
	}{before: 77, after: 88}
	Opus_opus_extension_iterator_init(nil, &guarded.state, nil, 0, 0)
	if guarded.before != 77 || guarded.after != 88 || guarded.state != (OpusT_OpusExtensionIterator{}) {
		t.Fatal("empty/guards")
	}
	for _, args := range [][2]int32{{-1, 1}, {1, 1}, {0, -1}, {0, 49}} {
		before := owned
		panicked := false
		func() {
			defer func() { panicked = recover() != nil }()
			Opus_opus_extension_iterator_init(nil, &owned, nil, args[0], args[1])
		}()
		if !panicked || owned != before {
			t.Fatal("assert/partial write", args)
		}
	}
}

func TestExtensionRepeatPointers(t *testing.T) {
	owned := func() OpusT_opus_extension_data {
		packet := []byte{7, 11, 5, 22, 33}
		var st OpusT_OpusExtensionIterator
		Opus_opus_extension_iterator_init(nil, &st, &packet[0], 5, 3)
		st.Fcurr_data = &packet[3]
		st.Fcurr_len = 2
		st.Frepeat_len = 2
		st.Fsrc_data = &packet[0]
		st.Fsrc_len = 2
		st.Frepeat_frame = 1
		st.Frepeat_l = 1
		var ext OpusT_opus_extension_data
		entropyInitGrowStack(12)
		runtime.GC()
		if opus_extension_iterator_next_repeat(nil, &st, &ext) != 1 || ext.Fid != 3 || ext.Fframe != 1 || ext.Flen1 != 1 || *ext.Fdata != 22 {
			t.Fatal("first repeat")
		}
		if opus_extension_iterator_next_repeat(nil, &st, nil) != 1 || opus_extension_iterator_next_repeat(nil, &st, nil) != 0 || st.Frepeat_frame != 0 || st.Fcurr_len != 0 {
			t.Fatal("repeat completion")
		}
		return ext
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *owned.Fdata != 22 {
		t.Fatal("payload ownership")
	}
	packet := []byte{7, 11, 5}
	var st OpusT_OpusExtensionIterator
	Opus_opus_extension_iterator_init(nil, &st, &packet[0], 3, 3)
	st.Fcurr_data = (*byte)(unsafe.Add(unsafe.Pointer(&packet[0]), 3))
	st.Fcurr_len = 0
	st.Frepeat_len = 2
	st.Fsrc_data = &packet[0]
	st.Fsrc_len = 2
	st.Frepeat_frame = 1
	st.Frepeat_l = 1
	ext := owned
	if opus_extension_iterator_next_repeat(nil, &st, &ext) != OPUS_INVALID_PACKET || ext != owned || st.Fcurr_len != -1 || st.Fsrc_len != 0 {
		t.Fatal("failure cursor/output")
	}
}

func TestExtensionIncrementPointers(t *testing.T) {
	for _, increment := range []byte{0, 1} {
		packet := [8]byte{3, increment, 7, 11}
		var iter OpusT_OpusExtensionIterator
		Opus_opus_extension_iterator_init(nil, &iter, &packet[0], 4, 2)
		var ext OpusT_opus_extension_data
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_extension_iterator_next(nil, &iter, &ext) != 1 || ext.Fframe != int32(increment) || ext.Fid != 3 || ext.Flen1 != 1 || *ext.Fdata != 11 {
			t.Fatal("frame increment", increment, iter, ext)
		}
	}
	for _, test := range []struct {
		length    int32
		increment byte
	}{{1, 1}, {2, 2}} {
		packet := [8]byte{3, test.increment}
		var iter OpusT_OpusExtensionIterator
		Opus_opus_extension_iterator_init(nil, &iter, &packet[0], test.length, 2)
		ext := OpusT_opus_extension_data{Fid: 77}
		if Opus_opus_extension_iterator_next(nil, &iter, &ext) != OPUS_INVALID_PACKET || ext.Fid != 77 {
			t.Fatal("invalid frame increment", test, ext)
		}
	}
}

func TestExtensionNextPointers(t *testing.T) {
	owned := func() OpusT_opus_extension_data {
		packet := []byte{65, 2, 11, 12, 7, 99, 4, 2, 21, 22, 23, 31, 32, 33}
		guard := struct {
			before uint64
			st     OpusT_OpusExtensionIterator
			after  uint64
		}{before: 77, after: 88}
		Opus_opus_extension_iterator_init(nil, &guard.st, &packet[0], int32(len(packet)), 3)
		var ext, first OpusT_opus_extension_data
		for i := 0; i < 6; i++ {
			entropyInitGrowStack(12)
			runtime.GC()
			if Opus_opus_extension_iterator_next(nil, &guard.st, &ext) != 1 {
				t.Fatal("next", i)
			}
			if ext.Fframe != int32(i/2) || ext.Fid != []int32{32, 3}[i%2] {
				t.Fatal("order", ext)
			}
			if i == 0 {
				first = ext
			}
		}
		if Opus_opus_extension_iterator_next(nil, &guard.st, nil) != 0 || guard.before != 77 || guard.after != 88 {
			t.Fatal("finish/guards")
		}
		return first
	}()
	runtime.GC()
	entropyInitGrowStack(12)
	if owned.Flen1 != 2 || *owned.Fdata != 11 || *(*byte)(unsafe.Add(unsafe.Pointer(owned.Fdata), 1)) != 12 {
		t.Fatal("payload ownership")
	}
	packet := []byte{7, 44, 3, 9}
	var st OpusT_OpusExtensionIterator
	Opus_opus_extension_iterator_init(nil, &st, &packet[0], 4, 3)
	ext := owned
	if Opus_opus_extension_iterator_next(nil, &st, nil) != 1 || Opus_opus_extension_iterator_next(nil, &st, &ext) != OPUS_INVALID_PACKET || ext != owned || st.Fcurr_len != -1 || st.Fcurr_frame != 9 {
		t.Fatal("separator failure")
	}
	if Opus_opus_extension_iterator_next(nil, &st, &ext) != OPUS_INVALID_PACKET || ext != owned {
		t.Fatal("sticky failure")
	}
}

func TestExtensionFindPointers(t *testing.T) {
	owned := func() OpusT_opus_extension_data {
		packet := []byte{65, 2, 11, 12, 7, 99, 4, 2, 21, 22, 23, 31, 32, 33}
		var st OpusT_OpusExtensionIterator
		Opus_opus_extension_iterator_init(nil, &st, &packet[0], int32(len(packet)), 3)
		var ext, first OpusT_opus_extension_data
		for i, want := range []byte{99, 23, 33} {
			entropyInitGrowStack(12)
			runtime.GC()
			if Opus_opus_extension_iterator_find(nil, &st, &ext, 3) != 1 || ext.Fframe != int32(i) || ext.Fid != 3 || ext.Flen1 != 1 || *ext.Fdata != want {
				t.Fatal("find", i, ext)
			}
			if i == 0 {
				first = ext
			}
		}
		before := ext
		if Opus_opus_extension_iterator_find(nil, &st, &ext, 128) != 0 || ext != before {
			t.Fatal("not found modified output")
		}
		Opus_opus_extension_iterator_reset(nil, &st)
		Opus_opus_extension_iterator_set_frame_max(nil, &st, 0)
		if Opus_opus_extension_iterator_find(nil, &st, nil, 3) != 0 {
			t.Fatal("frame limit")
		}
		return first
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *owned.Fdata != 99 {
		t.Fatal("payload ownership")
	}
	packet := []byte{65}
	var st OpusT_OpusExtensionIterator
	Opus_opus_extension_iterator_init(nil, &st, &packet[0], 1, 1)
	ext := owned
	if Opus_opus_extension_iterator_find(nil, &st, &ext, 32) != OPUS_INVALID_PACKET || ext != owned || st.Fcurr_len != -1 {
		t.Fatal("failed find")
	}
	if Opus_opus_extension_iterator_find(nil, &st, nil, 128) != OPUS_INVALID_PACKET {
		t.Fatal("sticky failure")
	}
}

func TestExtensionCountPointers(t *testing.T) {
	for _, test := range []struct {
		data         []byte
		frames, want int32
	}{{nil, 0, 0}, {nil, 48, 0}, {[]byte{7, 44, 65}, 3, 1}, {[]byte{7, 11, 5, 22}, 3, 2}, {[]byte{6, 6, 6, 4}, 48, 144}, {[]byte{7, 11, 2, 7, 22}, 3, 2}} {
		// C-style end cursors must stay inside the Go allocation, not one past it.
		storage := make([]byte, len(test.data)+16)
		copy(storage, test.data)
		for i := len(test.data); i < len(storage); i++ {
			storage[i] = 7
		}
		entropyInitGrowStack(12)
		runtime.GC()
		if got := Opus_opus_packet_extensions_count(nil, unsafe.SliceData(storage), int32(len(test.data)), test.frames); got != test.want {
			t.Fatal(test, got)
		}
		for _, b := range storage[len(test.data):] {
			if b != 7 {
				t.Fatal("guard")
			}
		}
		if len(test.data) == 0 && Opus_opus_packet_extensions_count(nil, nil, 0, test.frames) != 0 {
			t.Fatal("nil empty input")
		}
	}
	for _, args := range [][2]int32{{-1, 1}, {1, 1}, {0, -1}, {0, 49}} {
		panicked := false
		func() {
			defer func() { panicked = recover() != nil }()
			Opus_opus_packet_extensions_count(nil, nil, args[0], args[1])
		}()
		if !panicked {
			t.Fatal("assert", args)
		}
	}
}

func TestExtensionCountExtPointers(t *testing.T) {
	packet := [32]byte{7, 11, 3, 2, 7, 22, 2, 7, 33, 65}
	var counts [50]int32
	for i := range counts {
		counts[i] = 77
	}
	entropyInitGrowStack(12)
	runtime.GC()
	if got := Opus_opus_packet_extensions_count_ext(nil, &packet[0], 10, &counts[1], 4); got != 3 || counts[0] != 77 || counts[5] != 77 || counts[1] != 1 || counts[2] != 0 || counts[3] != 1 || counts[4] != 1 {
		t.Fatal(got, counts)
	}
	before := counts
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("missing assert")
			}
		}()
		Opus_opus_packet_extensions_count_ext(nil, nil, -1, &counts[1], 4)
	}()
	if counts != before {
		t.Fatal("clear before assertion")
	}
	alias := [4]int32{0x2c072c07, 0x2c072c07, 77, 88}
	if got := Opus_opus_packet_extensions_count_ext(nil, (*byte)(unsafe.Pointer(&alias[0])), 8, &alias[0], 1); got != 0 || alias != [4]int32{0, 0x2c072c07, 77, 88} {
		t.Fatal("clear/read alias", got, alias)
	}
	if Opus_opus_packet_extensions_count_ext(nil, nil, 0, nil, 0) != 0 {
		t.Fatal("empty")
	}
}

func TestExtensionParsePointers(t *testing.T) {
	// A larger capacity is not a claim that unused records are backed by storage.
	prefixPacket := [8]byte{7, 11}
	single := OpusT_opus_extension_data{}
	largeCapacity := int32(100)
	if Opus_opus_packet_extensions_parse(nil, &prefixPacket[0], 2, &single, &largeCapacity, 1) != 0 || largeCapacity != 1 || single.Fid != 3 || single.Flen1 != 1 || *single.Fdata != 11 {
		t.Fatal("consumed output prefix")
	}
	owned := func() [2]OpusT_opus_extension_data {
		packet := [32]byte{7, 11, 7, 22, 65}
		guard := struct {
			before uint64
			ext    [2]OpusT_opus_extension_data
			after  uint64
		}{before: 77, after: 88}
		count := int32(2)
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_packet_extensions_parse(nil, &packet[0], 5, &guard.ext[0], &count, 1) != OPUS_INVALID_PACKET || count != 2 || guard.before != 77 || guard.after != 88 {
			t.Fatal("partial/error/guards", count)
		}
		return guard.ext
	}()
	runtime.GC()
	entropyInitGrowStack(12)
	if *owned[0].Fdata != 11 || *owned[1].Fdata != 22 {
		t.Fatal("payload ownership")
	}
	packet := [16]byte{7, 11, 7, 22}
	out := [3]OpusT_opus_extension_data{{Fid: 77}, {Fid: 88}, {Fid: 99}}
	count := int32(1)
	if Opus_opus_packet_extensions_parse(nil, &packet[0], 4, &out[0], &count, 1) != -2 || count != 1 || out[0].Fid != 3 || out[1].Fid != 88 || out[2].Fid != 99 {
		t.Fatal("capacity partial output", out, count)
	}
	out[0].Fframe = 1
	if Opus_opus_packet_extensions_parse(nil, &packet[0], 4, &out[0], &out[0].Fframe, 1) != 0 || out[0].Fframe != 2 || out[1].Fid != 3 {
		t.Fatal("live aliased capacity")
	}
	count = 0
	if Opus_opus_packet_extensions_parse(nil, nil, 0, nil, &count, 0) != 0 || Opus_opus_packet_extensions_parse(nil, &packet[0], 4, nil, &count, 1) != -2 || count != 0 {
		t.Fatal("nil zero capacity")
	}
}

func TestExtensionParseExtPointers(t *testing.T) {
	prefixPacket := [8]byte{7, 11}
	single := OpusT_opus_extension_data{}
	largeCapacity := int32(100)
	oneCount := int32(1)
	if Opus_opus_packet_extensions_parse_ext(nil, &prefixPacket[0], 2, &single, &largeCapacity, &oneCount, 1) != 0 || largeCapacity != 1 || single.Fid != 3 || single.Fframe != 0 || single.Flen1 != 1 || *single.Fdata != 11 || oneCount != 1 {
		t.Fatal("consumed frame output prefix")
	}
	owned := func() [4]OpusT_opus_extension_data {
		packet := [32]byte{7, 11, 5, 22, 33, 9, 44}
		counts := [3]int32{2, 1, 1}
		guard := struct {
			before uint64
			out    [4]OpusT_opus_extension_data
			after  uint64
		}{before: 77, after: 88}
		capacity := int32(4)
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_packet_extensions_parse_ext(nil, &packet[0], 7, &guard.out[0], &capacity, &counts[0], 3) != 0 || capacity != 4 || guard.before != 77 || guard.after != 88 || counts != [3]int32{2, 1, 1} {
			t.Fatal("parse/guards/counts")
		}
		return guard.out
	}()
	runtime.GC()
	entropyInitGrowStack(12)
	for i, want := range []byte{11, 44, 22, 33} {
		if *owned[i].Fdata != want || owned[i].Fframe != []int32{0, 0, 1, 2}[i] {
			t.Fatal("order/ownership", i, owned)
		}
	}
	packet := [32]byte{7, 11, 5, 22, 33, 9, 44}
	counts := [3]int32{2, 1, 1}
	out := [4]OpusT_opus_extension_data{{Fid: 77}, {Fid: 88}, {Fid: 99}}
	capacity := int32(2)
	if Opus_opus_packet_extensions_parse_ext(nil, &packet[0], 7, &out[0], &capacity, &counts[0], 3) != -2 || capacity != 2 || out[0].Fid != 3 || out[1].Fid != 88 || out[2].Fid != 99 {
		t.Fatal("partial gap", out, capacity)
	}
	capacity = 0
	bad := [3]int32{}
	if Opus_opus_packet_extensions_parse_ext(nil, &packet[0], 7, nil, &capacity, &bad[0], 3) != -2 {
		t.Fatal("capacity before bucket assertion")
	}
	if Opus_opus_packet_extensions_parse_ext(nil, nil, 0, nil, &capacity, nil, 0) != 0 {
		t.Fatal("empty")
	}
}

func TestWriteExtensionPointers(t *testing.T) {
	out := [8]byte{77, 77, 77, 77, 77, 77, 77, 88}
	payload := [2]byte{11, 12}
	r := write_extension(nil, &out[0], 7, 1, 32, 2, &payload[0], 0)
	if r != 5 || out != [8]byte{77, 65, 2, 11, 12, 77, 77, 88} {
		t.Fatal(r, out)
	}
	if write_extension(nil, nil, 7, 1, 32, 2, nil, 0) != 5 {
		t.Fatal("size-only")
	}
	// C writes the ID byte before rejecting an invalid or oversized payload.
	if r := write_extension(nil, &out[0], 7, 1, 3, 2, &payload[0], 0); r != -1 || out[1] != 8 {
		t.Fatal("invalid short payload", r, out)
	}
	if r := write_extension(nil, &out[0], 2, 1, 32, 2, &payload[0], 0); r != -2 || out[1] != 65 {
		t.Fatal("capacity", r, out)
	}
}

func TestWritePayloadPointers(t *testing.T) {
	payload := make([]byte, 255)
	for i := range payload {
		payload[i] = byte(i)
	}
	out := make([]byte, 260)
	out[0] = 77
	out[259] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	r := write_extension_payload(nil, &out[0], 259, 1, 32, 255, &payload[0], 0)
	if r != 258 || out[1] != 255 || out[2] != 0 || out[3] != 0 || out[257] != 254 || out[0] != 77 || out[259] != 88 {
		t.Fatal(r, out)
	}
	if write_extension_payload(nil, nil, 259, 1, 32, 255, nil, 0) != 258 {
		t.Fatal("size-only")
	}
	before := append([]byte(nil), out...)
	if write_extension_payload(nil, &out[0], 2, 1, 32, 255, &payload[0], 0) != -2 {
		t.Fatal("capacity")
	}
	for i := range out {
		if out[i] != before[i] {
			t.Fatal("failed write changed data")
		}
	}
}

func TestSkipExtensionPointers(t *testing.T) {
	p := (*byte)(nil)
	h := int32(77)
	if skip_extension(nil, &p, 0, &h) != 0 || h != 0 || p != nil {
		t.Fatal("empty")
	}
	h = 77
	if skip_extension(nil, &p, -1, &h) != -1 || h != 77 {
		t.Fatal("negative")
	}
	data := [6]byte{65, 2, 11, 12, 3, 99}
	p = &data[0]
	h = 77
	if r := skip_extension(nil, &p, 6, &h); r != 2 || h != 2 || p != &data[4] {
		t.Fatal(r, h, p)
	}
	p = &data[0]
	h = 77
	if r := skip_extension(nil, &p, 1, &h); r != -1 || h != 77 || p != &data[0] {
		t.Fatal("error committed", r, h, p)
	}
}

func TestSkipPayloadPointers(t *testing.T) {
	data := [8]byte{3, 1, 2, 3, 4, 5, 6, 7}
	p := &data[0]
	h := int32(77)
	if r := skip_extension_payload(nil, &p, 8, &h, 65, 0); r != 4 || p != &data[4] || h != 1 {
		t.Fatal(r, p, h)
	}
	before := p
	h = 77
	if r := skip_extension_payload(nil, &p, 1, &h, 65, 0); r != -1 || p != before || h != 77 {
		t.Fatal("error committed output")
	}
	var owner *byte
	func() {
		b := make([]byte, 8)
		b[0] = 1
		b[2] = 99
		p := &b[0]
		h := int32(0)
		skip_extension_payload(nil, &p, 8, &h, 65, 0)
		owner = p
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *owner != 99 {
		t.Fatal("interior ownership")
	}
}

func TestRepeatedExtensionIterator(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	payload := []byte{'x'}
	extensions := []OpusT_opus_extension_data{
		{Fid: 3, Fframe: 0, Fdata: &payload[0], Flen1: 1},
		{Fid: 3, Fframe: 1, Fdata: &payload[0], Flen1: 1},
		{Fid: 3, Fframe: 2, Fdata: &payload[0], Flen1: 1},
	}
	packet := make([]byte, 32)
	length := Opus_opus_packet_extensions_generate(tls, &packet[0], int32(len(packet)), &extensions[0], int32(len(extensions)), 3, 1)
	if length <= 0 {
		t.Fatalf("generated extension packet length: got %d", length)
	}

	var iterator OpusT_OpusExtensionIterator
	Opus_opus_extension_iterator_init(tls, &iterator, &packet[0], length, 3)
	for frame := int32(0); frame < 3; frame++ {
		var extension OpusT_opus_extension_data
		if got := Opus_opus_extension_iterator_next(tls, &iterator, &extension); got != 1 {
			t.Fatalf("frame %d iterator result: got %d, want 1; length=%d packet=% x", frame, got, length, packet[:length])
		}
		if extension.Fid != 3 || extension.Fframe != frame || extension.Flen1 != 1 || *(*byte)(unsafe.Pointer(extension.Fdata)) != 'x' {
			t.Fatalf("frame %d extension: %+v", frame, extension)
		}
	}
	if got := Opus_opus_extension_iterator_next(tls, &iterator, nil); got != 0 {
		t.Fatalf("iterator exhaustion: got %d, want 0", got)
	}

	Opus_opus_extension_iterator_init(tls, &iterator, &packet[0], length, 3)
	var found OpusT_opus_extension_data
	if got := Opus_opus_extension_iterator_find(tls, &iterator, &found, 3); got != 1 {
		t.Fatalf("find result: got %d, want 1", got)
	}
	if found.Fframe != 0 || found.Flen1 != 1 || *(*byte)(unsafe.Pointer(found.Fdata)) != 'x' {
		t.Fatalf("found extension: %+v", found)
	}
	if got := Opus_opus_packet_extensions_count(tls, &packet[0], length, 3); got != 3 {
		t.Fatalf("extension count: got %d, want 3", got)
	}
	parsed := make([]OpusT_opus_extension_data, 3)
	count := int32(len(parsed))
	if got := Opus_opus_packet_extensions_parse(tls, &packet[0], length, &parsed[0], &count, 3); got != 0 {
		t.Fatalf("parse result: got %d, want 0", got)
	}
	if count != 3 || parsed[2].Fid != 3 || parsed[2].Fframe != 2 || *(*byte)(unsafe.Pointer(parsed[2].Fdata)) != 'x' {
		t.Fatalf("parsed extensions: count=%d entries=%+v", count, parsed)
	}
	frameCounts := make([]OpusT_opus_int32, 3)
	if got := Opus_opus_packet_extensions_count_ext(tls, &packet[0], length, &frameCounts[0], 3); got != 3 {
		t.Fatalf("per-frame extension count: got %d, want 3", got)
	}
	if want := []OpusT_opus_int32{1, 1, 1}; frameCounts[0] != want[0] || frameCounts[1] != want[1] || frameCounts[2] != want[2] {
		t.Fatalf("per-frame extension counts: got %v, want %v", frameCounts, want)
	}
	ordered := make([]OpusT_opus_extension_data, 3)
	count = int32(len(ordered))
	if got := Opus_opus_packet_extensions_parse_ext(tls, &packet[0], length, &ordered[0], &count, &frameCounts[0], 3); got != 0 {
		t.Fatalf("frame-ordered parse result: got %d, want 0", got)
	}
	if count != 3 || ordered[0].Fframe != 0 || ordered[1].Fframe != 1 || ordered[2].Fframe != 2 {
		t.Fatalf("frame-ordered parsed extensions: count=%d entries=%+v", count, ordered)
	}
}
