//go:build gcstress

// Command decodestress checks deterministic PCM output under allocation and GC
// pressure. Run separately: unsafe decoder failures may terminate the process.
package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
)

type result struct {
	digest           [sha256.Size]byte
	packets, samples int64
}

func decode(path string, limit int) (result, error) {
	var out result
	f, err := os.Open(path)
	if err != nil {
		return out, err
	}
	defer f.Close()
	r, err := ogg.NewOpusReader(f)
	if err != nil {
		return out, err
	}
	if r.Head.ChannelMappingFamily == 3 {
		return out, fmt.Errorf("projection mapping is not supported by this test")
	}
	d, err := opus.NewDecoderFromHead(r.Head)
	if err != nil {
		return out, err
	}
	defer d.Close()
	const frameSize = 5760
	pcm := make([]int16, frameSize*d.Channels())
	encoded := make([]byte, len(pcm)*2)
	h := sha256.New()
	for limit == 0 || out.packets < int64(limit) {
		pkt, err := r.ReadAudioPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			return out, fmt.Errorf("packet %d: %w", out.packets+1, err)
		}
		if len(pkt.Data) == 0 {
			return out, fmt.Errorf("packet %d: empty audio packet", out.packets+1)
		}
		n, err := d.Decode(pkt.Data, pcm, frameSize, false)
		if err != nil {
			return out, fmt.Errorf("packet %d: %w", out.packets+1, err)
		}
		samples := n * d.Channels()
		for i, v := range pcm[:samples] {
			binary.LittleEndian.PutUint16(encoded[i*2:], uint16(v))
		}
		_, _ = h.Write(encoded[:samples*2])
		out.packets++
		out.samples += int64(samples)
	}
	if out.samples == 0 {
		return out, fmt.Errorf("no PCM decoded")
	}
	copy(out.digest[:], h.Sum(nil))
	return out, nil
}

// churn keeps a bounded live set while continuously replacing heap objects.
// The ring and KeepAlive make the allocations observable to the GC. Touching
// every page also prevents the test from merely reserving virtual memory.
func churn(ctx context.Context, allocated *atomic.Uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	const blockSize = 64 * 1024
	ring := make([][]byte, 32) // 2 MiB retained per worker, not unbounded growth
	var iteration uint64
	for {
		select {
		case <-ctx.Done():
			runtime.KeepAlive(ring)
			return
		default:
		}
		b := make([]byte, blockSize)
		for i := 0; i < len(b); i += 4096 {
			b[i] = byte(iteration)
		}
		ring[iteration%uint64(len(ring))] = b
		iteration++
		allocated.Add(blockSize)
		if iteration%32 == 0 {
			runtime.Gosched()
		}
	}
}

func run(path string, passes, limit, workers, gcPercent int, interval time.Duration) error {
	fmt.Printf("baseline: %s\n", path)
	baseline, err := decode(path, limit)
	if err != nil {
		return fmt.Errorf("baseline: %w", err)
	}
	fmt.Printf("baseline: packets=%d samples=%d SHA256=%x\n", baseline.packets, baseline.samples, baseline.digest)

	oldGC := debug.SetGCPercent(gcPercent)
	defer debug.SetGCPercent(oldGC)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var allocated atomic.Uint64
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go churn(ctx, &allocated, &wg)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				runtime.GC()
			}
		}
	}()
	defer func() { cancel(); wg.Wait() }()

	// Warm up the live set before measuring; warmup GC does not count as a
	// successful stressed run. At least one collection must finish per pass.
	runtime.GC()
	started := time.Now()
	for pass := 1; pass <= passes; pass++ {
		var before, after runtime.MemStats
		runtime.ReadMemStats(&before)
		bytesBefore := allocated.Load()
		got, err := decode(path, limit)
		runtime.ReadMemStats(&after)
		cycles := after.NumGC - before.NumGC
		bytes := allocated.Load() - bytesBefore
		fmt.Printf("stress pass %d/%d: GC-cycles=%d churn-MiB=%.1f total-allocated-MiB=%.1f heap-MiB=%.1f\n", pass, passes, cycles, float64(bytes)/(1<<20), float64(after.TotalAlloc-before.TotalAlloc)/(1<<20), float64(after.HeapAlloc)/(1<<20))
		if err != nil {
			return fmt.Errorf("stress pass %d: %w", pass, err)
		}
		if got != baseline {
			return fmt.Errorf("stress pass %d: PCM changed: packets=%d samples=%d SHA256=%x (baseline SHA256=%x)", pass, got.packets, got.samples, got.digest, baseline.digest)
		}
		if cycles == 0 || bytes == 0 {
			return fmt.Errorf("stress pass %d: insufficient GC/allocation activity; use a longer input or larger -max-packets", pass)
		}
	}
	fmt.Printf("PASS: %d stressed decodes matched baseline in %s\n", passes, time.Since(started).Round(time.Millisecond))
	return nil
}

func main() {
	passes := flag.Int("passes", 3, "number of stressed decodes after baseline")
	limit := flag.Int("max-packets", 0, "packets per decode (0 = entire file)")
	workers := flag.Int("workers", 2, "allocation-pressure goroutines (2 MiB live data each)")
	gcPercent := flag.Int("gc-percent", 10, "GOGC value during stressed decodes")
	interval := flag.Duration("gc-interval", time.Millisecond, "interval between forced GC requests")
	flag.Parse()
	if flag.NArg() != 1 || *passes < 1 || *limit < 0 || *workers < 1 || *gcPercent < 1 || *interval <= 0 {
		fmt.Fprintln(os.Stderr, "usage: decodestress [flags] input.opus (positive passes/workers/gc-percent/gc-interval required)")
		os.Exit(2)
	}
	fmt.Printf("GC stress: GOMAXPROCS=%d workers=%d GOGC=%d forced-GC-interval=%s\n", runtime.GOMAXPROCS(0), *workers, *gcPercent, *interval)
	if err := run(flag.Arg(0), *passes, *limit, *workers, *gcPercent, *interval); err != nil {
		fmt.Fprintln(os.Stderr, "FAIL:", err)
		os.Exit(1)
	}
}
