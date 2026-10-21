package opuscc

import (
	"bytes"
	"errors"
	"math"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
	"unsafe"
)

func TestQuantAllBandsMergeOwnedLanesPointers(t *testing.T) {
	bands := [2]int16{0, 3}
	raw := [10]float32{77, 1, 2, 3, 88, 4, 5, 6, 99, 111}
	owned := raw
	entropyInitGrowStack(12)
	runtime.GC()
	quantAllBandsMerge(&raw[1], &raw[5], &bands[0], 1, 1, 0)
	quantAllBandsMergeFrom(quantBandFoldOutput{owner: owned[:], offset: 1, owned: true}, quantAllBandsNormRight(owned[:], 5, true), &bands[0], 1, 1, 0)
	want := [10]float32{77, 2.5, 3.5, 4.5, 88, 4, 5, 6, 99, 111}
	if raw != want || owned != want {
		t.Fatal("concrete merge lanes/source/guards", raw, owned)
	}
	alias := [4]float32{1, 2, 3, 4}
	quantAllBandsMergeFrom(quantBandFoldOutput{owner: alias[:], offset: 1, owned: true}, quantAllBandsNormRight(alias[:], 0, true), &bands[0], 1, 1, 0)
	if alias != [4]float32{1, 1.5, 2.25, 3.125} {
		t.Fatal("both merge owners must see sequential aliases", alias)
	}
	live := []float32{math.Float32frombits(3), 2, 77}
	zero := []float32{0, 0, 0}
	quantAllBandsMergeFrom(quantBandFoldOutput{owner: live, owned: true}, quantAllBandsNormRight(zero, 0, true), (*int16)(unsafe.Pointer(&live[0])), 0, 1, 0)
	if math.Float32bits(live[0]) != 2 || live[1] != 1 || live[2] != 77 {
		t.Fatal("both merge owners must retain live bound", live)
	}
}

func TestQuantAllBandsMergeOwnedBoundariesPointers(t *testing.T) {
	bands := [2]int16{0, 3}
	unused := quantBandFoldOutput{offset: math.MaxInt64, owned: true}
	quantAllBandsMergeFrom(unused, unused, &bands[0], 0, 1, 0)
	left := []float32{2, 4, 6}
	short := []float32{0}
	var caught any
	func() {
		defer func() { caught = recover() }()
		quantAllBandsMergeFrom(quantBandFoldOutput{owner: left, owned: true}, quantAllBandsNormRight(short, 0, true), &bands[0], 1, 1, 0)
	}()
	if _, ok := caught.(runtime.Error); !ok || left[0] != 1 || left[1] != 4 || left[2] != 6 {
		t.Fatal("incremental right prefix must follow first completed store", caught, left)
	}
}

// The legacy nil-positive span is fatal under checkptr, not recoverable.
func TestQuantAllBandsMergeOwnedNilRightDeathPointers(t *testing.T) {
	const child = "OPUS_GO_MERGE_NIL_RIGHT_CHILD"
	if os.Getenv(child) == "1" {
		left := []float32{2, 4, 6}
		bands := [2]int16{0, 3}
		os.Stderr.WriteString("merge-nil-right-before-call\n")
		quantAllBandsMergeFrom(quantBandFoldOutput{owner: left, owned: true}, quantAllBandsNormRight(left, 0, false), &bands[0], 1, 1, 0)
		os.Exit(0)
	}
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-test.run=^TestQuantAllBandsMergeOwnedNilRightDeathPointers$", "-test.count=1"}
	run := func(name string, args ...string) ([]byte, error) {
		command := exec.Command(name, args...)
		command.Env = append(os.Environ(), child+"=1")
		return command.CombinedOutput()
	}
	output, err := run(executable, args...)
	if errors.Is(err, syscall.ENOEXEC) && runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		qemu, lookupErr := exec.LookPath("qemu-aarch64")
		if lookupErr != nil {
			t.Fatal(lookupErr)
		}
		output, err = run(qemu, append([]string{executable}, args...)...)
	}
	if err == nil || !bytes.Contains(output, []byte("merge-nil-right-before-call")) || !bytes.Contains(output, []byte("unsafe.Slice: ptr is nil and len is not zero")) {
		t.Fatalf("missing right must retain original raw-span failure, not silently use primary storage: %v\n%s", err, output)
	}
}
