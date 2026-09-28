# Decode under GC pressure

Standalone, Go-only stress test (no C library or cgo required):

```sh
go run -tags gcstress ./tools/decodestress gogeta.opus
# Shorter smoke test:
go run -tags gcstress ./tools/decodestress -passes 2 -max-packets 100 gogeta.opus
# Longer/repeated stress:
go run -tags gcstress ./tools/decodestress -passes 20 -workers 4 gogeta.opus
```

Run from the repository root. The `gcstress` build tag excludes this program
from normal `go test ./...`. It runs in its own process because invalid pointer
use may cause an unrecoverable runtime crash. Do not use it on a production
service: high CPU and allocation rates are intentional.

The test:

1. Decodes the input once with the Go decoder under ordinary GC settings and
   records packet count, PCM sample count, and SHA-256 of little-endian int16 PCM.
2. Sets GOGC to 10 (configurable with `-gc-percent`).
3. Starts two allocation goroutines (`-workers`). Each continuously allocates
   and touches 64 KiB buffers, retaining a rolling 2 MiB live set so garbage is
   continually reclaimed without retaining all allocations.
4. Starts a goroutine requesting `runtime.GC()` every millisecond
   (`-gc-interval`). Collections also happen naturally from allocation pressure;
   the interval is a request cadence, not a guaranteed collection rate.
5. Reopens and decodes the file three times (`-passes`) with fresh decoder
   state, checking each result against the baseline while workers remain active.

Each pass reports completed GC cycles, background allocation volume, total
allocation volume, and heap size. A pass fails if no collections or background
allocations occurred, even if PCM matches. For very short inputs use a longer
file or increase `-max-packets` (0 means the entire file). `-passes` repeats the
whole decode, not individual packets. A new decoder is created and closed
under pressure for every pass; Go decoding, Ogg parsing, and hashing all run
while the stress workers are active.

PCM includes decoder pre-skip and end padding; no container gain or trimming
is applied. This checks determinism under GC, **not** agreement with native
libopus (use `tools/compareopus` for that). Input/read/decode errors, empty
packets, changed PCM/counts, or insufficient stress produce exit status 1;
invalid arguments produce status 2. Crashes are not caught or suppressed.

Initial full `gogeta.opus` run: all three stressed decodes matched the baseline
(2,907 packets, 16,744,320 samples). Each pass completed 4,209–5,300 GC cycles
with roughly 5.6–13.6 GiB of background allocations. Counts and timing vary by
machine and scheduler. This successful run is evidence, **not proof**, of GC
safety: pointer bugs can depend on input, scheduling, architecture, and Go
version. No codec/runtime safety fixes are made by this test.
