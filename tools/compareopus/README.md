# Native C / Go codec comparison

Run from the repository root (requires cgo and a C compiler):

```sh
go run -tags compareopus ./tools/compareopus test.opus
go run -tags compareopus ./tools/compareopus -max-packets 100 test.opus test2.opus test3.opus
go run -tags compareopus ./tools/compareopus -tolerance 1 -max-packets 100 big.opus
```

`-mode decode` (default) compares decoding; `-mode encode` compares encoding.
With no file arguments, it tries every `*.opus` (decode) or `*.wav` (encode)
in the current directory.
The `compareopus` build tag excludes this program from normal `go test ./...`.

The cgo directives in `main.go` use headers from `../opus/include` and statically
link `../opus/.libs/libopus.a`, relative to the repository root. This is the
native checkout found next to this repository; `../opus-go` is the Go checkout
itself. Build the native library first if necessary. Edit the two `#cgo`
directives if your native checkout is elsewhere. The program prints the linked
libopus version.

## Decode comparison

Each file gets independent, fresh C and Go decoders. The same Ogg audio packets
are decoded through each library's int16 API at 48 kHz, without FEC. Family 0
uses the ordinary decoder; families 1, 2, and 255 use multistream decoding with
the header's mapping. Other families are rejected.

This compares **raw decoder PCM**, including pre-skip and end padding. Neither
side applies OpusHead gain, pre-skip, or granule trimming: container playback
processing cannot mask decoder differences. Sample counts include all channels;
reported packet numbers are one-based, frame/channel indices are zero-based.

By default every sample must match exactly. The report includes the first sample
outside tolerance, total differing samples, samples outside tolerance, maximum
absolute difference, and RMS difference in int16 units. `-tolerance` optionally
allows small differences; exact mismatch counts remain visible. Native build
options, SIMD, and floating-point evaluation can affect bit-exactness; a
mismatch is evidence to investigate, not automatically proof of a Go bug.

Exit status is 0 only when all requested comparisons succeed, 1 for mismatches
or file/decode errors, and 2 for invalid options or no inputs. Empty audio
packets are rejected rather than silently treated as packet-loss concealment.
A read/decode error aborts that file, but subsequent files are still tried.
`-max-packets N` intentionally compares only a prefix; 0 (default) reads to EOF.

Initial decode checks against the local libopus 1.6.1:

- First 100 packets of `test.opus`, `test2.opus`, and `test3.opus`: exact matches.
- First 100 packets of `big.opus`: 26 differing samples, maximum difference 1.
- Full reads of the three test files encounter empty audio packets (packet 5648
  for `test.opus`/`test2.opus`, packet 103 for `test3.opus`), reported as errors.
  `test.opus` and `test2.opus` also first differ by 1 at packet 102.

## Encode comparison

```sh
go run -tags compareopus ./tools/compareopus -mode encode x-gogeta.wav
go run -tags compareopus ./tools/compareopus -mode encode -max-nrmse 0.02 x-gogeta.wav
go run -tags compareopus ./tools/compareopus -mode encode -exact -max-packets 100 x-gogeta.wav
go run -tags compareopus ./tools/compareopus -mode encode -bitrate 96000 -vbr=false -complexity 5 -max-packets 100 big.wav
```

Inputs must be 48 kHz, mono/stereo, signed 16-bit PCM WAV. Unsupported formats
are rejected, not resampled. Both encoders use the audio application, 20 ms
frames, and identical explicit settings: 64000 bps, VBR enabled, complexity 10
by default. `-bitrate`, `-vbr`, and `-complexity` change both encoders together.
Lookahead must match. Final partial frames are zero-padded and encoder delay is
flushed. `-max-packets N` limits source audio to N frames (N * 20 ms), **plus**
the packets needed to flush delay; 0 processes the full WAV.

The tool compares raw Opus payloads, not Ogg files, so serial numbers, tags,
page layout, and checksums cannot cause spurious differences. No output files
are written. It reports unequal packet counts, unequal length counts, total
payload bytes, and relative total size difference (`abs(Go-C)/C`).

Byte equality is useful but not required by Opus: small floating-point changes
can change encoder decisions and entropy-coded bytes. To measure the resulting
audio difference, both packet streams are decoded by **independent native C
decoders**, isolating encoder differences from Go decoder differences. PCM
statistics exclude lookahead and end padding and include every real input
sample. They include maximum absolute error, RMS error in int16 units, and
NRMSE = RMS(Go-encoded output minus C-encoded output) / RMS(C-encoded output).
For a silent C reference, NRMSE is zero only for an identical output, otherwise
infinity. This is a numerical regression metric, not a perceptual quality test
or a comparison of either lossy output with the original WAV.

Encode PASS requires NRMSE <= `-max-nrmse` (default 0.01 = 1%) and relative
payload size difference <= `-max-size-diff` (default 0.05 = 5%). These are
configurable engineering thresholds, not Opus conformance limits. `-exact`
additionally requires every packet payload to be byte-identical. `-tolerance`
is decode-only. Errors or exceeded thresholds exit with status 1.

Full `x-gogeta.wav` with local libopus 1.6.1 and default settings:

- 8,721 packets; 938 unequal payloads, but all lengths identical.
- Both encoders produced 1,373,589 payload bytes.
- 16,743,696 real PCM samples compared; 1,493,686 differed.
- Maximum PCM error 4,714; RMS error 89.1299; NRMSE 0.0140996 (1.41%).
- Fails the default 1% threshold; would pass an explicitly chosen 2% threshold.

Opt-in native tests cover exact silence encodings, mono/stereo, partial frames,
exact frame boundaries, and limited-input delay flushing. They also compare
converted helpers directly against the local C library: packet headers, entropy
lookups, interpolation (both Go codecs), sorting, bandwidth expansion, biquad
filters, downsampling, limiting, vector renormalization, float-to-PCM conversion,
two-band analysis filtering, sum-of-squares energy, variable low-pass cutoff,
VAD initialization, pitch decoding, Laroia NLSF weights, NLSF vector-quantization
errors, high-quality 2× upsampling (including all six state words), CELT LPC
coefficients, SILK gain dequantization, SILK mid/side-to-left/right conversion,
CELT exponential rotation, and fractional entropy-bit accounting (all normalized
16-bit mantissas at four range scales). Entropy decoder tests also compare interval
updates, probability-coded bits, raw tail bits, and unsigned integers, including
normalization, exhausted packets, overlapping front/tail reads, and invalid-value
clamping. These tests copy context fields explicitly across the C boundary;
`opuscc` entropy contexts now carry a typed `Fbuf`; legacy context addresses and
untyped scratch allocations elsewhere still require ownership care. Initialization
compares every numeric field, including the deliberately preserved `ext`, across
short/exhausted packets. Go ownership tests force GC and stack growth with the
context as the sole packet owner. Shared encoder-buffer adaptations compare
shrink/move, final padding/tail writes, and complete numeric state against C.
SILK decoder comparisons also cover NLSF unpacking/reconstruction/stabilization,
LPC coefficient fitting (including input updates), and shell pulse decoding with
all numeric entropy-state fields checked. Additional decoder tests cover 16-bit
ICDFs, both Laplace variants, and all 5,625 SILK stereo predictor-index combinations
plus mid-only flags. Signed Laplace fixtures avoid zero-probability symbols
(e.g. negative values when `p0=32767`). Inverse prediction gain tests cover stable
and unstable coefficients through order 24; NLSF-to-LPC tests cover both decoder
orders, tightly clustered frequencies, and in-place conversion. LPC analysis
filter comparisons include overflow-heavy inputs and overlapping buffers. Pulse
sign tests cover all signal/offset models, sum masking, padded shell blocks,
skipped pulses, exhausted packets, and every numeric entropy-state field.
Decoder reset/init comparisons verify the nonzero defaults and cleared state,
excluding native CPU dispatch (Go uses scalar arch 0). HQ upsampling wrapper
tests check that only IIR state changes. Full pulse decoding covers the ten-LSB
escape limit and shell/sign reconstruction. PLC energy comparisons include the
actual static helper from `PLC.c`; a `compareopus`-only Go bridge exposes its
internal counterpart. These tests check subframe selection, signed narrowing,
saturation, energies, and shifts. CELT coarse/fine/final energy decoder tests
compare exact float32 outputs and all numeric entropy fields across budget
thresholds, prediction modes, optional previous quantization, and nil final
energy output. Band denormalization is compared bit-for-bit across frame scales,
downsampling factors, silence, exp2 underflow, and capped gains; its band table
is passed as a typed pointer rather than read from the legacy mode field.
Allocation cap tests compare mode tables and channel/frame scales. Hybrid folding
uses a scalar C reference for the static helper in `bands.c`, including bitwise
copies of NaN payloads. Pulse-vector decoding compares native `cwrs.c` vectors,
energies, and entropy state across sparse/dense codebooks. Full PVQ reconstruction
compares libopus output bits, collapse masks, and state across spreading modes,
block counts, gains, and exhausted packets. Its pulse scratch is Go-owned.
Outer per-frame entropy, scalar-output, CTL, and silence scratch objects are
pinned while legacy SILK/CELT uintptr interfaces still use their addresses.
This fixes read-chunk and multistream regressions exposed by stack-layout changes;
it does not establish global pointer safety. Pins can be removed as the complete
call chains become typed.
IIR/FIR interpolation is compared against the actual static libopus helper,
including every Q16 fractional phase, saturated inputs, guard samples, and
input/output overlap. The Go helper returns an output count instead of a
one-past-end pointer. The IIR/FIR driver also compares PCM and complete filter
history across empty, short, exact-batch, multi-batch, and consecutive calls,
including the untouched tail of its FIR union. Its scratch buffer is Go-owned.
Down-FIR interpolation compares all six coefficient sets, 18/24/36-tap orders,
every Q16 phase, int32 pair-sum narrowing, saturation, and output guards against
the actual static C helper. The down-FIR driver additionally checks complete
state and PCM across consecutive calls, partial/multiple batches, and C's
unprocessed one-sample remainder. It accepts a typed coefficient table explicitly
and uses Go-owned scratch, leaving `FCoefs` conversion at the legacy caller.
Pitch cross-correlation is compared bit-for-bit with the scalar implementation
compiled from `pitch.c`, covering four-lag groups, scalar tails, minimal input
extents, and zero-length scalar cases. Stereo-angle inputs are compared exactly
with scalar `vq.c`, including equal/opposite channels, tiny-energy thresholds,
large finite magnitudes, and both stereo and independent-energy modes.
Hadamard interleaving and deinterleaving use typed buffers and Go-owned scratch,
with bitwise comparisons against the actual static `bands.c` helpers for every
supported Hadamard stride and plain transpositions, including NaN payloads,
signed zero, guards, and exact round trips.
Projection matrix initialization uses a typed header/data path and is compared
against libopus for metadata, aligned coefficient offsets, and matrix sizes up
to the channel/count limits. Pointer tests also preserve forward-copy aliasing.
Projection float output is compared bit-for-bit across matrix columns, input and
output strides, preloaded accumulators, empty frames, and aliased input/output.
Projection int16 output additionally covers ties-to-even input conversion,
NaN/infinity clamping, Q15 product rounding, and wrapping output accumulation.
Projection int24 output checks finite inputs within C's int32 conversion domain,
including values beyond normalized unity, ties-to-even conversion, 64-bit
products, asymmetric Q15 half rounding, and int32 accumulation narrowing.
Smooth fades compare exactly with a scalar C reference of `opus_decoder.c`,
covering supported sample rates, channel-major stores, and partial buffer overlap.
Packet-duration queries exhaust all TOC/count-byte combinations at supported
rates against libopus, including short headers and the 120 ms limit.
Aggregate SILK init/reset compares the actual `dec_API.c` path, both channel
states and stereo history, while checking that channel-count metadata survives.
Native CPU dispatch is excluded, as in the per-channel reset comparisons.
Soft clipping compares PCM and persistent history bit-for-bit with libopus across
consecutive frames, ramp correction, impulses, clipping runs, zero crossings,
signed zero, NaNs/infinities, multiple channels, and guarded output/state buffers.
Resampler initialization compares the complete state and coefficient selection
against `resampler.c` for all 30 supported encoder/decoder rate pairs; guard tests
also cover the smaller 386 state layout.
The top-level resampler driver additionally compares PCM and the complete state
across consecutive 1/2/10/11/21/30 ms calls for all supported rate pairs, including
copy mode, delay compensation, and empty second batches.
The low-quality 2/3 resampler compares PCM and all six state words, including
full-range initial state, saturation, in-place operation, 480-sample batch edges,
and the un-emitted one/two-sample remainders (C emits two samples per triple).
The exported mapping-matrix data accessor compares aligned coefficient addresses
and contents with C; a Go test retains only the returned pointer across GC.
Float matrix input multiplication compares every selected row, input/output
strides (including zero), empty dot products, guard values and overlapping buffers
bit-for-bit with C.
Int16 matrix input comparisons additionally cover extreme integer products and
255-column accumulations, preserving conversion to float32 after each product.
Int24 matrix input tests include 24-bit limits, float32 integer-rounding boundaries,
and full int32 values; comparisons retain the float32 accumulator and two scaling
steps, with no added 24-bit clipping.
Decoder sample-count forwarding compares actual native decoder queries at all
supported rates, TOC values and representative counts/errors without changing state.
Shared entropy encoder initialization compares all numeric fields with C from
random preloaded states, verifies untouched buffers, and tests sole-context buffer
ownership across GC and stack growth. The separate `opusccenc` copy is unchanged.
Entropy shrinking exhausts valid new sizes, tail counts, and boundary front
counts for buffers through 32 bytes, comparing overlapping moves and all state
fields with C while checking outside guards.
Initial entropy-bit patching compares all 0–8 bit widths and byte values across
finalized/pending-byte, interval, and error branches, including threshold-adjacent
ranges and signed pending state, with exact state and buffer comparisons.
Carry propagation compares the actual static `entenc.c` helper across buffered
bytes, pending carry runs and exhausted buffers; Go tests also cover counter wrap.
Encoder normalization compares the actual static C helper around range thresholds,
multiple normalization passes, carry-producing values and exhausted buffers.
Interval encoding compares complete state and output after each of 400 consecutive
intervals at representative totals, including zero lower bounds, full upper bounds,
normalization and buffer exhaustion.
Binary interval encoding similarly checks 500 consecutive operations with 0/1/2/4/8/15
bit totals, exact numeric state, emitted bytes and guards, including zero-width
identity intervals and output exhaustion.
Probability-coded encoder bits compare 500 consecutive operations at logp
1/2/3/8/15, including arbitrary nonzero signed values and buffer exhaustion.
16-bit ICDF encoder tests compare every symbol in representative 8/15/16-bit
tables, singleton tables and consecutive coding with exhausted buffers, checking
all numeric fields, output and unchanged tables.
Raw encoder tail bits compare every 0–32-bit preloaded window occupancy and
1–25-bit append width, byte flushing, guards and exhausted output. Go tests check
bit-accounting wrap and consecutive appends.
Finalization compares interval termination, pending carries, 0–32 buffered tail
bits, padding and partial-byte front/tail collisions against C, including guards
and exhausted output. A focused Go test round-trips typed range and raw-bit coding.
Unsigned encoder integers compare totals around range/raw-bit split boundaries
through UINT32_MAX, consecutive operations, finalization and exhausted buffers.
Go tests round-trip boundary symbols with sole-context output ownership under GC.
8-bit ICDF encoding compares all symbols in 2/8-bit tables and singleton tables,
consecutive coding and exhausted output. Go tests round-trip symbols and guard
the exact table allocation.
Laplace encoding compares actual C laplace.c state, emitted bytes and in-place
symbol clipping at representative frequencies/decays, signed tail extremes and
exhausted buffers. Go tests decode the clipped symbols exactly.
Laplace-p0 encoding compares signed symbols around each seven-symbol continuation
boundary, minimum decay probabilities, p0 extremes and output exhaustion. Fixtures
exclude zero-probability symbols; Go tests round-trip both signs and continuations.
CWRS indexing compares actual static C icwrs on signed pulse distributions across
all representative PVQ shapes; Go tests invert sampled indices and check guards.
Pulse encoding compares native CWRS/entropy state and bytes across PVQ shapes,
consecutive signed vectors, unchanged inputs and exhausted output. Go tests
round-trip concentrated pulses including the maximum 176-dimensional shape.
SILK shell encoding compares actual C depth-first trees for totals 0–16,
concentrated/random distributions, input preservation and exhausted output.
Go tests round-trip each concentrated position and verify zero trees consume nothing.
SILK sign encoding compares every signal/quantization type, signed int8 extremes,
skipped sums, low-five-bit sum table selection, rounded block counts and 120-sample
padding against C, checking state, bytes, guards and unchanged inputs. Go tests
round-trip signs through the typed decoder.
Fine-energy quantization compares float bit patterns, full encoder state/output,
optional previous quantization, budget skips, clipping and aliased energy/error arrays.
Final-energy quantization compares both priority passes, bit-budget edges, maximum
fine bits, nil/aliased energy outputs, signed zero and NaN sign decisions with C.
Amplitude-to-log conversion compares raw float bits for exponent/mantissa-bin
boundaries, subnormals, signed zero, negative inputs, infinities and NaNs, plus
in-place/partial overlaps and inactive-band fills against C.
Mini-FFT factorization compares actual mini_kfft.c for lengths 1–4096 and large
prime/power-of-two boundaries, including radix order, returned word counts and
untouched factor tails. Go tests also verify products, remainders and guards.
Multistream float channel output compares the actual static C helper for varied
strides/channels, zero-length and nil-source fills, raw float bits and forward overlap.
Multistream int16 output compares clipping, ties-even conversion, NaN/infinity
handling, zero-stride writes, nil fills and guards against the actual static C helper.
Multistream int24 output compares scaling and ties-even conversion without 24-bit
clipping, full representable int32-domain boundaries, strides, nil fills and guards.
Exported 8-bit ICDF decoding compares C decision-threshold neighbors, representative
ranges, singleton/exact-size tables and exhausted input, including full context state
and unchanged inputs. Go tests retain packet ownership solely through the decoder.
SILK gain quantization compares all signed previous-index values, conditional/full
coding, 1/2/4 subframes and aliased previous/index storage against C; Go tests
round-trip valid gain histories through typed dequantization.
PLC parameter updates compare actual static PLC.c for 2/4 subframes, LPC orders
10/16, signal types and LTP gain clamp/narrowing edges, checking all PLC fields,
unchanged control and unused LPC tails.
PLC frame gluing compares concealed-energy capture, unequal energy shifts, integer
sqrt approximation, onset fade/break boundaries and saturation-extreme PCM with C,
checking full Go state, frame guards and zero-length no-op behavior.
Projection demixing-matrix access compares the actual static C accessor's aligned
interior pointer and header fields; Go tests retain the backing allocation solely
through returned matrix/coefficient pointers during GC and stack growth.
Projection multistream-state access compares native alignment for matrix sizes
0–1024 and checks interior-pointer ownership through GC and stack growth.
Projection float output compares actual C callbacks over consecutive input channels,
clear-before-read aliasing, nil sources, frame boundaries and valid matrix/stride extents.
Projection int16 output compares first-channel clearing, later-channel accumulation,
clipping/ties/NaN inputs, nil sources and destination guards against actual C callbacks.
Projection int24 output compares channel clearing/accumulation, rounding boundaries,
representable int32 extremes, nil sources and guards without adding 24-bit clipping.
FFT butterfly tests share `fft_butterfly_helpers_test.go`. Radix-2 compares the
actual static `kiss_fft.c` helper, including guards, subnormals and signed zeros.
Radix-4 also covers its twiddle-free m=1 path, multiple blocks, gaps, zero/repeated
strides, untouched twiddles and float32 product rounding with exact C bit comparisons.
`TestRadix3ScratchRounding` checks explicit multiply-before-add float32 rounding
in both radix-3 implementations; it catches ARM64 FMA fusion without changing
frame goldens. Cross-architecture Go-only validation can run with
`GOARCH=arm64 CGO_ENABLED=0 go test -exec qemu-aarch64 ./...` on Linux with QEMU.

Radix-3 reuses the grouped fixtures and compares epi3 selection, half/scalar
rounding, guards and stack-growth/GC calls with exact native output bits.
Radix-5 uses the same grouped cases to verify ya/yb selection, parenthesized
float32 sums/products, five-way stores and impulse/guard behavior.
Resampler states now retain typed coefficient pointers; initialization/reset uses
barrier-aware struct stores. Native rate-pair comparisons retain coefficient IDs and
check pointer-sized layout, while heap-owner tests drop the original coefficient
slice and compare driver output/history after GC and stack growth.
NLSF codebooks retain all eight typed table pointers. Decoder and encoder states
retain typed codebook references, with direct decode/parameter/index traversal and
sample-rate assignments. Grouped tests check C layouts, both codebook orders,
heap-only owner chains and reset/replacement, alongside the existing native full-state
comparisons. Decoder/encoder pitch-lag and contour table fields are also typed;
heap-only owner tests cover the three lag and four contour tables, repeated entropy
reads, rate/subframe selection and reset. Native full-state indices/parameters and
sample-rate fixtures retain the original table selection and state updates, with
C size/offset checks for both state layouts. Poisoned reset fixtures skip every
converted GC pointer slot and use real sentinels.
CELT band, log-band, window, allocation-vector and pulse-cache index/bit/cap
fields are typed. Native fixtures compare mode/cache sizes, offsets and every table
payload; heap-mode owners retain cloned tables through GC and stack growth.
Band/log/index access preserves signed int16 loads. The quant-partition rate searches
use typed byte caches, retaining the C six-step search, tie order, zero-pulse behavior,
cache reloads and remaining-budget updates. Actual rate.h comparisons cover every
valid LM -1 through 3 cache row, budgets -2 through 400 and every valid pulse count;
a guarded backwards index checks signed offset -1 without exercising C's invalid
sentinel rows. Grouped owner cases also exercise caps and MDCT.
CELT band contexts retain typed mode, band-energy and shared entropy-codec pointers, including
through copied/restored context snapshots. Theta/partition/stereo locals keep these
references typed; the PVQ adapter's codec argument no longer crosses an integer
boundary. Heap-context tests retain mode tables and encoder/decoder packet buffers
through GC/stack growth, exercising one-bin signs, lowband stores and complete codec
state/bytes for zero, tiny and normal encoder capacities. Native bands.c checks
context layout and the one-bin reference fixtures use the context-owned codec.
Theta now takes typed context, split-output, budget and fill pointers, preserving
sequential output stores even when budget/fill alias one another, remaining_bits,
or split fields. Actual bands.c decoder fixtures cover uniform/triangular/intensity
branches, budgets, LM and these output aliases. Mono/stereo band drivers take typed
contexts, and the outer driver uses GC-scanned Go context storage instead of a fixed
80-byte TLS allocation. Decoder native fixtures compare masks, remaining budget,
seed, complete entropy state, spectra, folding output and guards across LM 0–3,
time/frequency changes, budgets and intensity choices.
Theta, recursive partition, mono and stereo spectra/folding/scratch arguments are
now typed throughout the internal band chain. Splits derive interior pointers with
unsafe.Add; direct indexing, copy and clear replace integer-address stores and
memory shims while preserving sequential fold and N=2 stereo store order. The PVQ
integer adapter and obsolete one-bin adapter are removed. The outer band driver
retains explicit mono/stereo legacy spectral adapters and its pseudostack.
Focused checkptr now exercises active spectra, recursive splitting, pulses, zero-fill,
noise/folding, time/frequency changes, optional scratch/output buffers and aliasing,
with stack growth/GC and guards on amd64, 386 and ARM64/QEMU. Actual bands.c encoder
theta fixtures compare split outputs, entropy state/bytes and input/output spectra;
decoder partition fixtures cover recursive budgets and optional folding alongside
the mono/stereo references. Native comparisons run on the host, not ARM64/macOS.
Remaining outer integer boundaries, opaque allocations and pseudostack ownership
mean these tests still do not establish global safety.
Deemphasis now takes typed channel heads, PCM, coefficient and memory pointers.
The generic path uses Go-owned N-sample scratch instead of TLS pseudostack storage;
no unused output cursor is formed when decimation produces zero samples. The common
stereo path remains unchanged, and generic accumulation/non-accumulation retain
separate sum orders and explicit float32 product rounding. Celt decoder callers
still cross an explicit legacy channel-address adapter. Native celt_decoder.c
fixtures compare C=0/1/2 (including the C do-while channel-zero visit), N=0–960,
factors 1/2/3/4/6, accumulation, coefficient choices, output guards and histories.
Focused checkptr exercises active generic/fast paths, GC/stack growth, nil TLS,
short/zero frames and memory/output store order on amd64, 386 and ARM64/QEMU; native
comparisons remain host-only. The initial typed-buffer round checked only the
scratch-free stereo path while the generic scratch was still legacy. Input/output
aliases violate C restrict contracts, so Go store-order tests do not claim C parity
for those aliases. That stage did not migrate CELT synthesis/decode ownership.

CELT synthesis now retains typed mode, spectrum/band-energy and output channel
pointers and uses a Go-owned N-sample interleaved frequency buffer, with no TLS
pseudostack allocation/restoration or integer MDCT addressing internally. Mode
and transform table owners, output-buffer staging, channel-zero do-while behavior,
transient block order, independently rounded downmix halves and the floating-point
SATURATE identity are preserved. Decoder/PLC callers still enter through an
explicit integer-address adapter; its ordinary mono/stereo heads are made typed
before allocation, with uintptr escape annotation for direct legacy arguments.

Actual celt_decoder.c synthesis fixtures use the existing renamed scalar bands.c
and MDCT/FFT implementations rather than the linked library's SIMD transforms.
They compare spectra/energy immutability, full output/TDAC images and guards for
LM 0–3, mono/stereo/upmix/downmix and channel-zero visits, transient/non-transient
blocks, factors 1/2/3/4/6, silence and full/partial/empty band ranges. Normal stereo
shared/partially overlapping output heads are also valid native fixtures because
each MDCT invocation has a separate output restrict scope. Upmix staging/output
alias order is Go-only: it violates the MDCT restrict contract and is not claimed
as C parity. Native comparisons remain host-only.

Grouped synthesis tests own cloned mode/MDCT/FFT tables and channel/spectral
buffers through typed pointers only, with GC and stack growth and exact-sized
output-head arrays. Focused checkptr on amd64/386 and ARM64/QEMU now covers the
complete active synthesis path, nil TLS, an untouched sentinel TLS slot, TDAC
history and output aliases. Early rounds checked geometry/MDCT, denormalization
and staging helpers while full synthesis still depended on legacy scratch. Outer
CELT decoding, concealment and opaque decoder allocation scanning remain legacy;
this is not a global GC-safety proof or direct macOS CI validation.

The CELT concealment prefilter/fold now takes a typed decoder/mode, retains
scanned history interiors, calls the typed comb filter and folds through typed
window/scratch pointers. One Go-owned overlap-length buffer replaces all TLS
pseudostack allocation/restoration and is reused after each channel's fold. Live
per-channel controls, channel-zero do-while visits, odd-overlap truncation and
filter-before-fold order are preserved. Zero overlap consumes no audio cursor,
including exact-sized mono histories with N=0; ordinary outer decode/PLC callers
still use an explicit legacy adapter.

The typed fold initially permitted a new ARM FMA, changing the existing PLC
frame golden (7bc10321 instead of f3e48462). Independent float32 product rounding
restored the original golden. The overlap-three regression is 0xc086cced versus
the fused one-ULP alternative and is checked against actual static C prefilter
output as well as Go. No goldens/tolerances were changed.

Actual celt_decoder.c prefilter fixtures link the renamed scalar celt.c comb path
and compare full normalized decoder/history/guard images over C=0/1/2,
N=0/120/240/480/960, overlap 0/1/2/3/119/120, periods from 0 to 1024, zero/positive/
negative gains and tapset transitions. Grouped Go tests retain heap modes via sole
history interiors, exercise active prefilter/TDAC with nil TLS and forced GC/stack
growth, check exact-sized zero-overlap history and an untouched sentinel TLS slot.
Window/scratch alias-order fixtures are Go-only, not valid enclosing C restrict
inputs. Earlier rounds checked state/history/folding helpers while full prefilter
still used legacy scratch; focused checkptr now covers the complete active path
on amd64/386 and ARM64/QEMU. Surrounding PLC concealment/state scratch, opaque
allocation scanning and extension EOF remain unresolved. Host native comparisons
and QEMU do not establish direct macOS CI or global decoder GC safety.

Four subsequent CELT concealment leaf rounds move band-energy decay, the
25-word autocorrelation noise-floor/lag window, excitation decay energy, and
synthesis explosion/attenuation into typed pointer/fixed-array helpers. Active
legacy concealment calls these helpers at explicit storage boundaries. MAXG and
MIN32 retain ordered comparisons and NaN selection; the energy channel loop still
has C's do/while behavior. Float32 products are explicitly rounded before MACs
and subtraction, preserving ARM behavior. Synthesis computes energy before any
writes, clears explosions (including NaNs) to positive zero, and reloads live
window/output values while applying overlap and tail attenuation.

Grouped Go tests cover exact spans, guards, forced GC/stack growth, nil unused
inputs, zero/odd/full excitation lengths, mono/stereo/zero-channel decay,
threshold branches, NaNs, and live aliases. Grouped native tests use
source-equivalent celt_decoder.c snippets with actual MAXG/MIN32/celt_sqrt macros,
not a whole-C-concealment integration oracle. Scoped checkptr covers only these
migrated leaves. The outer CELT concealment dispatcher, state layout and TLS
scratch remain legacy, as does surrounding CELT decoding. Every round retains
full amd64/386/ARM64-QEMU tests, native/GC-stress comparisons and unchanged
packet/frame/encode/decode goldens and tolerances. Host/QEMU is not macOS CI or
proof of global GC safety or opaque-allocation pointer scanning.

The next four concealment leaf rounds also type excitation history copying,
periodic extrapolation/reference-energy accumulation, reverse 24-word LPC history
gathering, and final decoder-state updates. Excitation copying and LPC gathering
keep individual live loads/stores, not memmove/snapshot behavior. Extrapolation
retains store-before-reference-energy order, period-boundary attenuation and
explicit float32 rounding. Final state updates retain cached loss duration,
live PLC duration, int32 addition/shifts, the 10000 clamp and last-frame-type store.

Grouped tests cover periods 0/1/512/1024, pitch 40/100/511/1024, frame lengths
0/120/240/960, overlap, guards, GC/stack growth and heap mode ownership. Outer
scratch/history aliases are Go-only fixtures; signed overflow and shifts outside
C's domain are also explicitly Go-only. Native source-equivalent leaf snippets
check normal sample spans, sequential energy, LPC history and scalar duration
updates using the actual IMIN macro, without importing pointer-bearing C images.
Scoped checkptr remains leaf-only; the legacy CELT concealment dispatcher, TLS
scratch and integer cursors are not yet fully migrated. Each round runs full
amd64/386/ARM64-QEMU tests, native/GC-stress comparisons and unchanged baselines;
repeated ARM leaf/checkptr and ordinary frame goldens remain separate scopes.

The following four concealment rounds type noise-spectrum generation and
replace all three TLS scratch arrays: the C*N normalized spectrum, the
max_period+24 excitation with its retained negative-index history prefix, and
the exc_length FIR temporary. Spectrum/excitation/FIR pointers and indexing are
now Go-owned slices; FIR copy-back uses copy on disjoint owned arrays. Noise
synthesis receives a scanned two-pointer output table rather than a raw integer
array. Sequential unsigned RNG updates, signed sample conversion, channel/band
order and final RNG store are preserved. Native noise tests use the scalar
vq.c normalization formula, not presumed-SSE native dispatch.

There is no concealment TLS allocation/cursor/save/restore left. Ordinary
full periodic/noise concealment fixtures run with nil TLS and an untouched
sentinel slot, including mono/stereo and LM=0..3. Scoped checkptr covers owned
noise/synthesis and excitation/autocorrelation/LPC/FIR pipelines plus migrated
leaves, not the whole dispatcher: decoder/mode/history integer-addressed views
still remain. The constructor tests keep empty noise/FIR storage and the
prefix-only excitation case separate from consumed interior pointers. Native
FIR spans include length 80/200/1024 and the 24-sample prehistory. All four
rounds retain full amd64/386/ARM64-QEMU tests, native and GC stress, unchanged
baselines/tolerances, and repeated ARM checks. Earlier mentions of concealment
TLS scratch above are historical; opaque allocation scanning and outer CELT
decoding remain unresolved, and host/QEMU coverage is not direct macOS CI.

The next four rounds retain a typed decoder at the concealment entry, a typed
cached mode/eBands owner, scanned history/output views, then typed LPC/window
consumers. Channel geometry uses numeric offsets into the decoder's trailing
float storage; only consumed interiors are materialized (including no unused
mono channel or zero-overlap/N=0 one-past output). Noise and periodic history
shifts use copy with their distinct overlap lengths. Synthesis, pitch search,
comb filtering, autocorrelation, LPC, FIR and IIR now receive typed pointers;
no integer-addressed view or private legacy adapter remains inside concealment.
The surrounding public CELT decode ABI still converts its legacy state once.

Focused checkptr now covers complete active concealment, with nil TLS, an
untouched TLS sentinel, forced GC/stack growth, heap modes/FFT tables, mono/stereo,
LM=0..3, first-loss pitch/LPC, repeated periodic/noise calls and folding into
noise with nonzero postfilters. All prior packet/frame goldens are unchanged.
The actual included celt_decoder.c oracle checks full numeric state/history
images and guards across three calls, periodic reuse, first loss, noise and
prefilter/postfilter cases. Native mode pointers are rebound locally and cleared
before import; Go pointer bytes are never passed to C. Scalar normalization,
FIR/IIR and autocorrelation are selected rather than the linked presumed-SSE
build: an initial noise mismatch exposed this oracle-selection difference.
All four rounds retain full amd64/386/ARM64-QEMU tests, native and GC stress,
unchanged encode/decode baselines and tolerances, and repeated ARM checks.
Earlier leaf-only checkptr limits and legacy concealment views above are now
historical. Outer CELT decoding and opaque allocation pointer scanning still
remain legacy; this is not global GC safety or direct macOS CI coverage.

DRED init pointer migration: Opus_opus_dred_decoder_init now accepts actual
*OpusT_OpusDREDDecoder, not uintptr; removed reconstruction/uintptrescapes directive.
Exported Go parameter type changes, no replacement raw adapter. Existing guarded
12-byte DRED header fixture now calls public entry across GC/stack growth; fields,
magic, guard extents, alias/layout and invalidation oracle unchanged. Existing
private initializer/factory behavior stays (no model/SIMD implementation enabled
or disabled). Full inherited validation, scoped DRED/header/storage amd64/386,
tagged ARM ordinary/checkptr repeats20 and diff checks pass; no failures. Logs
/tmp/opus-dred-init-pointer-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr72->71, unsafe.Pointer76->75. No native GC/global lifetime,
opaque-image GC maps, performance or direct macOS claims.

Decoder typed-destroy cleanup: replaced final Opus_opus_decoder_destroy use
with destroy_typed and removed unused raw-address wrapper. External Go users
must use typed entry. Registry fixture STILL returns only numeric registration
key plus weak references; no strong decoder owner is kept during its original
GC trials. Only AFTER those registry-only checks does RegisteredGoAllocation
recover the actual typed owner, without integer-to-pointer reconstruction, for
cleanup. Added post-destroy registration-absent check BEFORE TLS.Close. No weaker
registry oracle, raw-key dereference, new allocation or fixture extent change.
Existing typed destructor/free/callback semantics and codec references unchanged.
Full inherited validation initially hit recurring native found-pointer-to-free-
object with TestExtensionNextAgainstC active. Complete unchanged retry passed;
failed log retained, no native repair claimed. Scoped create/storage checks
amd64/386 and tagged ARM ordinary/checkptr repeats20 pass. Logs
/tmp/opus-decoder-typed-destroy-{focused,focused-retry,full,full-retry,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr73->72, unsafe.Pointer76 unchanged. Encoder's separate API
unchanged; no global lifetime/GC-map/performance/direct macOS claim.

Float typed-only cleanup: removed unused Opus_opus_decode_float raw wrapper
and uintptrescapes directive entirely. All callers already use
Opus_opus_decode_float_typed, including opus/decoder.go; no caller casts, adapter
or test/oracle changes needed. Typed/private float DSP/error behavior unchanged.
Full inherited validation, scoped float/native-reference amd64/386 and tagged
ARM ordinary/checkptr repeats20 pass; no failures. Logs
/tmp/opus-float-typed-only-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr75->73, unsafe.Pointer79->76. No global/native lifetime,
GC-map, performance or direct macOS claim.

Int24 typed-only entry migration: replaced all remaining opuscc
Opus_opus_decode24 calls with Opus_opus_decode24_typed and actual state/packet/
PCM pointers. Removed unused raw wrapper/uintptrescapes directive, no substitute
integer-address adapter. External Go callers must use typed entry. Nil/error
ordering, mono/stereo real packet/PLC buffers/guards and mixed int16/int24 Go
scratch/TLS cursor checks retain same physical fixtures and expected values;
codec/reference/oracle/tolerance unchanged. Mixed scratch now uses BOTH typed
decode APIs; raw TLS setup remains and is NOT claimed checkptr-safe or repaired.
Full inherited checks and scoped integer pointer/no-pseudostack/native-reference
amd64/386 plus tagged ARM ordinary/checkptr repeats20 pass. Ordinary includes
mixed TLS scratch case. No failures. Logs
/tmp/opus-int24-typed-only-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr77->75, unsafe.Pointer82->79. Encoder separate raw API is
unchanged; earlier remaining-raw-int24 notes below are superseded. No native GC,
global lifetime, opaque-image GC maps, performance or direct macOS claims.

Int16 typed-only entry migration: replaced every remaining opuscc
Opus_opus_decode call with Opus_opus_decode_typed and actual decoder/packet/PCM
pointers. Removed unused integer-address wrapper and its uintptrescapes directive;
no replacement raw adapter or pointer-to-integer conversion. External users must
use typed entry. Production opus/decoder.go already did so and is unchanged.
Nil/zero-frame/FEC validation-order, mono/stereo decode/PLC guards and original
Go-scratch cursor checks retain same physical fixtures and expected results.
Mixed int16/int24 scratch test changes ONLY int16 call; raw int24 ABI and its
existing ordinary coverage remain untouched. No codec, reference, allocation,
rounding, assertion, tolerance or input-size change. Full inherited validation,
scoped int16/integer-pointer/native-C-reference amd64/386 and tagged ARM
ordinary/checkptr repeats20 pass. Ordinary includes mixed legacy scratch case;
no previously unsafe raw int24/TLS-pointer case is relabeled checkptr-safe.
No failures. Logs /tmp/opus-int16-typed-only-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr79->77, unsafe.Pointer85->82. Encoder's separate raw entry
is unchanged; no global/native GC/opaque-image/lifetime/performance claims.

Public native decode pointer migration: Opus_opus_decode_native has NO uintptr
parameters/address reconstruction. Decoder is *OpusT_OpusDecoder, packet *byte,
PCM *float32, offset output *int32; removed obsolete uintptrescapes directive.
DRED remains an ignored unsafe.Pointer argument with unchanged signed offset,
not a fabricated Go DRED layout or new implementation. No relocated raw native
adapter. This CHANGES exported Go parameter types. Existing typed driver and
validation/error order, live inputs/outputs, arithmetic and codec DSP unchanged.

Updated three callers to actual state/buffer pointers, not casts back to integer
addresses. Loss-frame fixture keeps original raw allocation/init/physical extent;
no GC-map repair claimed. Independent native C-reference fixture now exercises
public pointer entry directly for SAME packets, physical buffers, FEC/error/
self-delimited/packet-offset cases, exact hashes/ranges/durations/PCM and original
eight vararg bytes. Existing paired direct/public state/full-PCM/guard checks
remain. New bounded Go case compares valid nonnil opaque DRED byte pointer and
extreme signed offsets against typed core (DRED ignored), retaining original
output canaries; not a native DRED model/restrict/invalid-domain claim.
Full inherited amd64/386/ARM64-QEMU/native/GC stress/reference checks pass;
scoped native/direct/public/C-oracle amd64/386 and tagged ARM ordinary/checkptr
repeats20 pass. Ordinary also includes original raw-allocation loss-frame fixture.
No failures. Logs /tmp/opus-native-typed-public-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr83->79, unsafe.Pointer88->85 (four reconstruction casts
removed, one legitimate opaque pointer parameter). Historical raw-native ABI
comparison descriptions below now refer to typed/public comparisons; no integer
native decode boundary remains in this API. No global lifetime/opaque-image GC
maps/native repair/performance/direct macOS claims.

Typed decoder creation migration: every remaining opuscc caller now uses
Opus_opus_decoder_create_typed and matching destroy_typed. Removed unused
exported integer-returning Opus_opus_decoder_create wrapper; external Go callers
must switch to the typed API. No relocated legacy constructor/opaque alias.
Production opus/decoder.go already used typed creation and is unchanged.

Native decode C-reference fixture now keeps actual decoder/packet/PCM/offset
pointers through existing typed native/int16 drivers, not typed creation followed
by a cast back to integer decoding. Raw native wrapper only forwarded these same
values (its omitted DRED arguments were zero/ignored); all packet bytes, sample
counts, buffers/physical extents, independent C hashes/ranges/offsets/durations,
FEC/error/softclip cases and exact tolerances unchanged. Original eight vararg
bytes stay [2]int32, with equivalent native-word scalar1 encoding, not integer
address reconstruction. Separate raw/native ABI fixtures remain unchanged.
CTL transcript keeps original vararg/output allocations, encodings, transcript
and oracle; only factory/destroy/state/child-origin plumbing becomes typed or
numeric unsafe.Add. Native fixture is now additionally checkptr-tested directly.

Full inherited amd64/386/ARM64-QEMU/native/GC stress/reference validation passes;
explicit native-reference/create/storage scoped amd64/386 plus tagged ARM
ordinary/checkptr repeats20 pass (ordinary also includes original CTL transcript).
No failures. Logs /tmp/opus-decoder-typed-create-{focused,full,amd64-checkptr,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr85->83, unsafe.Pointer89->88. Encoder's separate constructor
is unchanged; no native GC/opaque-image/lifetime/performance/macOS claims.

Custom decoder CTL pointer migration: Opus_opus_custom_decoder_ctl now takes
*OpusT_OpusCustomDecoder and libcshim.VaCursor, removing BOTH uintptr parameters
and BOTH state/base reconstruction casts. Exported Go signature changes; typed
request implementation and lazy helper/argument-read timing stay unchanged. No
replacement integer-address adapter or hidden alias. All migrated public CTL
entry points now use actual state pointers/cursors; DRED keeps no cursor because
it remains unimplemented. Earlier custom-CTL raw-boundary notes are superseded.
New public-entry test uses actual scanned decoder/mode owners and native scalar/
pointer slots across GC/stack growth, setter/getter, range, null, invalid value
before stores, unknown/reset with empty cursor, repeated cursor/live reload,
aliased error clear/store ordering, mode output/write barrier and retained mode
owner. Original typed/native independent oracles/extents/layouts and codec
references/tolerances unchanged. Full inherited validation plus scoped CTL/
cursor amd64/386 and tagged ARM ordinary/checkptr repeats20 pass; no failures.
Logs /tmp/opus-custom-ctl-pointer-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr87->85, unsafe.Pointer91->89. No native lifetime/GC repair,
opaque-image GC maps, performance or direct macOS claims.

Version-string pointer migration: opuscc.Opus_opus_get_version_string now
returns *byte directly, with no integer-address result or pointer-to-integer
conversion. This changes its exported Go return type. Existing _typed entry
remains and returns SAME read-only, NUL-terminated static diagnostic-pool root.
No storage copy, string/pool/layout/NUL change or new compatibility adapter.
Version fixture now holds the public entry's actual pointer across GC/stack
growth and checks original pool identity, exact libopus 1.6.1 text, NUL and
_typed parity without integer conversions. Existing independent oracles/extents,
codec goldens/tolerances and broad validation unchanged. Full inherited checks,
scoped diagnostics amd64/386 and tagged ARM ordinary/checkptr repeats20 pass;
no failures. Logs /tmp/opus-version-pointer-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts: opuscc uintptr89->87, unsafe.Pointer92->91. Encoder package's separate
legacy version API is unchanged; no global ABI/lifetime/performance claims.

Latest DRED CTL decision: removed unused VaCursor again at user request because
DRED remains unimplemented. Signature is now (TLS, *OpusT_OpusDREDDecoder,
request); existing -5 result is unchanged. Updated no-argument test. No DRED
implementation/configuration was disabled or changed. This supersedes the
restoration below. Full inherited validation, scoped amd64/386 and tagged ARM
ordinary/checkptr repeats20 pass; no failures. Logs
/tmp/opus-dred-ctl-remove-cursor-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.

DRED CTL correction: restored VaCursor parameter at user request, preserving
the argument-carrying signature from the common.go CTL migration. DRED was
ALREADY an unimplemented -5 stub before these cursor changes; neither removing
nor restoring an unused parameter enables/disables DRED processing. No DRED
implementation was removed. Current stub still returns -5; implementing actual
DRED is separate work. Cursor holds a pointer and numeric offset, not an integer
address (its shared slot-size expression uses sizeof(uintptr), a width constant).
No model/codec/configuration, ABI layout, lifetime or arithmetic change claimed.
Full inherited validation, scoped amd64/386 and tagged ARM ordinary/checkptr
repeats20 pass unchanged. No failures. Logs
/tmp/opus-dred-ctl-restore-cursor-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts remain89 uintptr/92 unsafe.Pointer. The no-argument signature recorded
immediately below is superseded by this correction.

Disabled DRED CTL follow-up: Opus_opus_dred_decoder_ctl now takes ONLY TLS,
typed DRED decoder and request. Removed unused VaCursor parameter entirely;
no placeholder/opaque alias, varargs parser, pointer-width dependency or new
raw adapter. This is a further exported Go signature change and unused-stub
API cleanup, not an ownership/DSP migration. Every request still ignores state
and returns -5. Other CTL/cursor APIs are unchanged. Renamed/updated no-argument
DRED tests; native fixtures/extents/oracles/codec references/tolerances unchanged.
Full inherited validation initially hit recurring native found-pointer-to-free-
object with TestExtensionNextAgainstC active. Complete UNCHANGED retry passed;
failed log retained, no native fix claimed. Scoped amd64/386 and tagged ARM
ordinary/checkptr repeats20 include DRED and preceding CTL/cursor cases. Logs
/tmp/opus-dred-ctl-no-cursor-{focused,focused-retry,full,full-retry,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
Counts remain89 uintptr/92 unsafe.Pointer; no global safety/performance claims.
Prior DRED cursor-signature notes below are historical.

Focused common.go CTL migration (NOT whole-file integer-address removal):
Opus_opus_decoder_ctl, multistream ctl/ctl_va_list and projection ctl now take
actual concrete decoder pointers and libcshim.VaCursor values. OpusT_va_list now
names that real pointer-plus-offset cursor, not an integer-address alias. No
state/base reconstruction remains in these CTL APIs, no replacement raw adapter
was relocated elsewhere, and no integer slot snapshots/TLS scratch were added.
This intentionally CHANGES exported Go signatures/va_list representation.
Go callers can also use existing *_ctl_typed with OpusDecoderCtlArgs directly.
Disabled DRED ctl receives matching typed state/cursor signature; this is stub
API cleanup only, not a new DRED ownership/DSP claim (still ignores everything,
returns -5). Remaining create/decode/destroy/callback/difference uintptr sites
in common.go are explicitly NOT migrated by this focused CTL step.

Cursor copied per call; opening/reading remains request-lazy. Reset/unknown
consume no slots, invalid stream ID consumes exactly its first native word,
second output slot is deferred until bounds pass. Projection derives MS at
same points as before. Validation/store order, live aliases, child forwarding,
range XOR/clear-before-query, write barriers, reset prefix/gain preservation
and all error values unchanged. Removed unused generated MS forwarding locals
without changing cursor-copy semantics. Original native transcript/hash fixtures
keep SAME allocations/extents/encodings/expected values; only fixture-boundary
state/cursor arguments adapt to new signatures. No production compatibility
casts relocated into new helpers and no oracle/tolerance changes.

New public-entry tests use actual scanned composite decoder parents, pointer
output slots/native scalar words across GC/stack growth, all four ctl routes,
repeated cursor calls/live scalar reloads, null output, unknown/reset with empty
cursor, invalid gain/no stores, actual child output, one-word invalid stream ID,
aliased final range and reset gain preservation. DRED nil/empty cases remain -5.
These are bounded Go cases, not new native invalid/restrict-domain claims; no
opaque image GC maps/legacy vararg pointee lifetime/native GC repair claimed.
Counts: opuscc uintptr99->89; unsafe.Pointer98->92. Full inherited amd64/386/
ARM64-QEMU/native/GC stress/unchanged-codec-reference checks, explicit scoped
public ctl/cursor tests amd64/386, tagged ARM ordinary/scoped repeats20, and diff
checks pass. No failures. Logs
/tmp/opus-common-ctl-cursor-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
No performance/global lifetime/direct macOS claims. Previous raw-boundary notes
below are historical; only custom CELT ctl retains that raw boundary now.

Focused decoder_ctl_typed.go migration: no uintptr types/address reconstruction
remain in that file. msCtlLegacyArgs, opusCtlLegacyArgs and customCtlLegacyArgs
now consume lazy factories for actual pointer-based libcshim.VaCursor values.
Cursor retains a Go-visible base pointer plus numeric byte offset; native-word
Int32 reads preserve low32/sign behavior on 32/64-bit targets, pointer slots
load actual pointers directly, and advancement never forms an unused one-past
pointer. This is not a uintptr alias or relocated integer-word pointer decoder.
Four EXISTING raw ABI entry points keep original exported signatures and defer
ONLY their unavoidable raw va-base conversion until a consumed request opens
the cursor. No new raw adapter, uintptr-storing cursor, TLS scratch or slot copy.

Laziness preserves unknown/reset/unsupported requests with unused varargs;
multistream opens once, reads stream ID first, then checks its original bounds
before accessing a second slot. Request whitelists, delegation, value/null
checks, error order, live output aliases, typed stores/write barriers and DSP
remain unchanged. Upstream multistream va_arg/bounds order inspected. Existing
raw VaList/VaInt32/VaUintptr and public typed CTL arguments are unchanged.
New tests use scanned actual pointer slots/native scalar words across GC/stack
growth, cover all three adapters' setter/getter groups, pointer/null outputs,
unsigned/mode/decoder slots, one-word invalid stream IDs, and unused factories
that fail if opened. Native-word low32/high-bit and offset/base tests cover both
word sizes. No opaque byte image gains GC maps: legacy vararg byte/word buffers
still require external pointee lifetimes. No native GC/lifetime repair claimed.

Counts: opuscc uintptr102->99; unsafe.Pointer100->98 (six integer-to-output
pointer casts replaced by actual pointer loads, four deferred raw-base casts at
existing boundaries). Original fixtures/extents/native independent oracles,
codec hashes/goldens/tolerances and broad validation scope unchanged. Full
inherited amd64/386/ARM64-QEMU/native/GC stress checks pass, explicit cursor/CTL
scoped checkptr amd64/386 and tagged ARM ordinary/scoped repeats20 pass. No test
failures. Logs /tmp/opus-ctl-pointer-cursor-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
No performance/global lifetime/direct macOS claims. Historical counts follow.

Focused Opus_celt_fatal migration: exported str/file parameters are now *byte,
not uintptr. This intentionally CHANGES the exported Go signature: callers pass
typed NUL-terminated C-string pointers; no relocated integer-address adapter was
added. Internal Go-string fatal calls remain unchanged. New libcshim.GoStringPointer
scans from an actual pointer using numeric byte offsets, with SAME nil result,
NUL termination and 1MiB bound as legacy GoString, without integer reconstruction
or a fabricated maximum-length slice. File conversion STILL precedes message;
stderr formatting, signed line and abort behavior remain identical. Legacy
libcshim.GoString and other raw APIs are deliberately unchanged. Earlier historic
notes calling public Opus_celt_fatal a remaining raw-integer adapter are superseded
by this focused change. Counts: unsafe.Pointer100 unchanged, uintptr104->102.

Original two native fatal diagnostic oracles/physical C-string lengths stay;
fixture now uses real Go byte allocations and typed pointers across GC/stack
growth, not malloc addresses or unscanned integer slots. Added nil/embedded-NUL
Go-shim cases, invalid-byte preservation, interior roots and exact scan-limit
coverage. Nil/unterminated bounded Go cases are not new C printf-domain claims.
Pointer reader still requires actual readable storage through NUL or limit;
it does not validate arbitrary pointers or give opaque images GC maps.
Full inherited amd64/386/ARM64-QEMU validation, native comparisons, GC stress,
codec goldens/tolerances and diff checks pass. Explicit fatal/reader scoped
checkptr passes amd64/386 plus tagged ARM ordinary/checkptr repeats20. No failures.
Logs /tmp/opus-celt-fatal-typed-{focused,full,386-checkptr,arm-checkptr-repeat,arm-repeat}.log.
No native GC/extension EOF/raw callback repair, global lifetime safety,
performance or direct macOS claim. Historical batch counts below remain historical.

Four postfilter rounds retain actual output-pointer-list storage, first-block
history destination, tail history destination, then canonical comb window120.
All four connect active decoder/comb consumers. Exact known history and static
window storage are Go-visible owners, not aliases for byte images or displaced
casts. Counts remain100 unsafe.Pointer/104 uintptr; public/raw adapters stay.

Postfilter's actual [2]*float32 output array passes as its real slice, sharing
the existing concrete synthesis-output carrier. Original required first-cell
argument check remains. Full max(1,channels) pointer-list view forms BEFORE any
clamp/channel DSP. Clamp stays per channel; out[c] is reloaded separately for
first and tail, preserving live list aliases, LM guard and channel-zero do/while
behavior. History arrays/origins, state/period/gain/tapset reads and arithmetic
remain unchanged. Raw and typed output-list routes share identical dispatch.

First and tail use actual history as BOTH comb source and destination only
AFTER existing exact live-pointer/root/origin match. Mismatched/nil/unbound paths
retain original raw fallback. Tail STILL forms original full N prefix and
required shortMdctSize cell BEFORE ownership matching and BEFORE zero-gain
identity guard; needed scalar tail identity/fallback boundary is deliberately
retained. No skipped required index, eager destination views or guessed root.
Source-before-output view order, no-work/identity behavior, signed unused
compatibility, per-channel first-before-tail, float32 rounding/products and
sequential stores stay unchanged. No spectrum/image/whole-decoder migration claim.

Comb window binding occurs at original window-view point AFTER history, output
and initial delayed-sample loads, never at entry or before zero-gain fast path.
Only exact &window120[0] plus overlap0..120 binds actual window120[:]. Copy,
interior, oversized, negative and nil cases retain legacy pointer/length;
unsupported classification does NOT manufacture fake larger spans. Owner flag
never silently falls back. Overlap suppression and original deferred read timing
stay. Window coefficients remain live through destination aliases; no snapshot.
No C/SIMD/QEXT enabling, table/layout/ABI or arithmetic change. Upstream
celt_decoder.c first/tail loop and celt.c comb guards/window arithmetic inspected.

Existing postfilter history fixture keeps same buffers/guards and full bit/state
oracle, now exercises real pointer-list owner under GC/stack growth. Existing
first/tail fixtures retain independent raw comb oracle and physical extents,
add equal-sized owned-destination full-bit/state comparisons across gains/taps/
N. Bounded new tail check requires original interior cell even with zero gains,
before stores/state changes. Canonical window tests prove actual root and exact
identity/extent classification; raw copied table is independent same-120-word
reference for whole-buffer live-window alias comparisons. Canonical table is
restored after EVERY pair and on exit. Go-only aliases/invalid bounded cases do
not claim new C restrict/const/invalid domains. No native fixture/hash/bytes/
bit-count/seed, codec reference/tolerance/scope change or pointer-slot poisoning.

Round1 initially failed compilation because actual out_syn is [2]*float32;
corrected ONLY array-to-slice argument (no copied pointer list), then complete
validation hit native found-pointer-to-free-object with TestExtensionNextAgainstC
active. COMPLETE UNCHANGED next retry passed. All logs retained:
/tmp/opus-round-postfilter-go-pointer-list{,-retry,-retry2}.log.
No oracle weakening or production native GC fix. Other round logs:
/tmp/opus-round-postfilter-first-go-destination.log,
/tmp/opus-round-postfilter-tail-go-destination.log,
/tmp/opus-round-comb-canonical-go-window.log and
/tmp/opus-round-postfilter-owners-arm-{checkptr,golden}-repeat.log.
Each round passes full amd64/386 and ARM64/QEMU, broad/targeted scoped checkptr,
tagged typed PartA/PartB/storage/owner/cell checks, native comparisons, GC stress,
codec references and diff checks. Final tagged ARM ordinary/scoped repeats20,
explicitly including postfilter/tail/window and preceding owner fixtures.
No other batch failures. No native GC/extension EOF/raw callback/opaque-image
repair, global lifetime safety, performance or direct macOS claims.

Four norm-input rounds stage concrete copy sources, mono input lanes, stereo
input lanes, then connect BOTH actual normalization lanes in quant-all-bands.
Rounds1/2/3 stage owned private APIs with direct equivalence fixtures; round4
connects actual _norm storage for normal decode, mono/stereo/dual paths and both
opuscc encode-only RDO trials (not opusccenc). Counts remain100 unsafe.Pointer/
104 uintptr: real scanned slice/offset flow, not opaque aliases or token removal.

The existing concrete scratch/lowband carrier now accepts a concrete copy source;
legacy preparation delegates with explicit legacy discriminator. Required copy
still forms destination prefix BEFORE source prefix BEFORE stores and switches
prepared source only AFTER copy. The original owned zero-length unused-address
boundary stays at same point and ignores unconsumed source geometry. Go overlap
semantics and exact required snapshot/copy extent remain; no extra allocation or
sample snapshot. Source presence never silently falls back from owned nil.

Mono carries incoming lane directly as prepared lowband until SAME scratch-copy
predicate replaces it. Existing Haar/deinterleave/recursive partition source
owners now survive even when scratch is disabled and NO copy occurs. Context
loads/division BEFORE N1 shortcut, N1 unused-input behavior, transform/copy/read
order, fill/gain/seed/entropy/mask, partition validation and sequential folding
remain unchanged. Stereo forwards SAME source through N1/N2/intensity and either
mid/side order; originally nil side source/scratch/output remain nil. Scalar raw
APIs/wrappers retain original branches and consume identical shared kernels.

Actual input binding keeps effective==-1 OR N==0 return BEFORE lane/presence/
root access; missing work-required lane uses original raw nil-span boundary,
not noise fallback or guessed primary ownership. Present lane uses a real Go
suffix of _norm at cached numeric lane origin, with relative effective offset.
Original int32 offset+N endpoint is computed BEFORE full-prefix clipping. FULL
N input prefix is validated at SAME argument point BEFORE output/table binding
and BEFORE DSP/N1 shortcuts, not deferred to only samples eventually consumed.
Right origin is cached at original norm2 setup point, never reconstructed from
addresses or reloaded from mutable bands. Parent/suffix slices retain actual
allocation ownership; no fake prefix. No pulse/codec/byte snapshot or RDO rollback/
tie change. Upstream bands.c source selection/copy/transform/dual dispatch checked.
Remaining X/Y, raw/image/callback and unused-coordinate boundaries are NOT
claimed migrated or lifetime-repaired; scalar norm heads still serve other uses.

Existing mono/stereo full-buffer/guard/context/COMPLETE entropy comparisons now
exercise actual incoming owners as well as destination/scratch owners, retaining
same physical extents and original raw/native expected results. Context rebases
ONLY separate entropy-instance pointer. Additional owner-copy overlap/manual
oracle, source-prefix-before-stores failure, zero-work ignored source, mono/
stereo N1 unused origins, true suffix/root/offset identity, unused guards and
input-before-output full-prefix error checks use real bounded storage. A no-
scratch shared-input/output fixture compares whole states and float bits to
original raw sequential folding under GC/stack growth. Existing native physical
fixtures/hashes/bytes/bit counts/seeds, codec references/tolerances and validation
scope remain unchanged. No pointer-slot poisoning, fixture repair/enlargement or
oracle weakening; Go-only invalid/alias checks are not new C restrict/domain claims.

Each round passes full amd64/386 and ARM64/QEMU, broad/targeted scoped checkptr,
tagged typed PartA/PartB/storage/owner/cell checks, native comparisons, GC stress,
codec references and diff checks. Final tagged ARM ordinary/scoped repeats20,
explicitly including all incoming-lane/copy/no-scratch and prior owner fixtures.
Logs /tmp/opus-round-norm-input-owned-copy-source.log,
/tmp/opus-round-mono-go-norm-input.log,
/tmp/opus-round-stereo-go-norm-input{,-retry}.log,
/tmp/opus-round-all-bands-go-norm-input.log and
/tmp/opus-round-norm-input-owners-arm-{checkptr,golden}-repeat.log.
Round3 initially hit native found-pointer-to-free-object with
TestExtensionNextAgainstC active; COMPLETE UNCHANGED retry passed. Failed log
retained. Other rounds/final checks have no validation failures. No native GC/
extension EOF/raw callback/opaque-image repair, global lifetime safety,
performance or direct macOS claims.

Four norm-mutation rounds retain primary hybrid storage, secondary hybrid
storage, primary merge storage, then secondary merge storage. Each connects an
actual quant-all-bands _norm Go owner (not an opaque alias), using the previously
cached numeric right origin. Both hybrid entry points, including opuscc RDO
restore, and dual-to-intensity resynthesis merge keep roots. Public/raw adapters
remain; opusccenc is untouched. Counts remain100 unsafe.Pointer/104 uintptr:
active dataflow migration, not token removal, relocated casts or layout recovery.

The existing concrete fold-output carrier supplies mutable lane spans directly.
Hybrid retains original eager bands view and int32 width calculations, equal-
width guard BEFORE all lane views, primary full-n2 view then copy, and secondary
view/copy ONLY after primary copy and ONLY when dual stereo is nonzero. No
whole-buffer/table snapshots, new samples or allocations. Required first-copy
completion before second-view failure is unchanged. Right selection preserves
whether setup actually produced a right lane; absent right remains legacy nil,
never silently aliases primary. Present right uses actual root/cached origin,
not an address-derived offset or a new setup-boundary reload.

Merge retains boundary lookup in EVERY loop condition (including final check),
original int32 M*boundary-offset and j+1 geometry, primary prefix BEFORE secondary
prefix on EACH iteration, addition BEFORE separate float32 half product, and
sequential stores. No eager full-bound spans, bound caching, data snapshots or
changed expression/read order. Aliased right reads observe prior left writes;
Go-only output/table alias retains live-bound updates. Resynth/dual/intensity
conditions, call points, hybrid/RDO copy/restore order, entropy, masks, contexts,
seed, tie choices and all surrounding arithmetic stay unchanged. Upstream
bands.c hybrid and intensity transition inspected. Norm input, X/Y and remaining
raw/image/callback adapters are NOT claimed migrated or lifetime-repaired.

Original manual hybrid n+guard allocations/extents stay; equal-size primary-only
and both-owned runs compare full float bits and original expected arrays under
GC/stack growth. Equal-width/nil/unused-origin guard, missing-right selection,
copy-before-bounded-secondary-error and Go-only cross-lane sequential alias
oracles are added on real allocations. Existing merge numeric/guard/alias/live-
bound fixtures retain physical extents, add primary-owner full-bit comparisons;
new both-lane fixtures independently check complete storage, unchanged source/
guards, sequential averaging and live-bound behavior. Incremental short-right
failure must occur AFTER the first completed store. No fake prefixes, pointer-
slot poisoning, original native fixture/hash/bytes/bit-count/seed changes or
codec golden/tolerance/scope weakening. Go-only invalid bounded/alias cases are
not new C-valid/restrict-domain claims.

Missing-right positive work retains original raw nil-span failure, not recovery
or ownership fallback. Its NEW Go assertion initially incorrectly expected
recover() under checkptr; the runtime terminated the test process. Failed log
/tmp/opus-round-merge-secondary-go-norm-owner.log retained. Corrected ONLY that
new assertion to a subprocess death test checking exact nil-positive-span text
and a pre-call marker (nonzero exit required). Same check runs ordinarily and
under scoped checkptr; cross-ARM exec-format failure explicitly launches QEMU,
without skipping or suppressing validation. No production correction or weakened
oracle. COMPLETE unchanged validation rerun passed in
/tmp/opus-round-merge-secondary-go-norm-owner-retry.log. No other batch failures.
Precision text-edit uniqueness/overlap mismatches resolved BEFORE validation.

Each round passes full amd64/386 and ARM64/QEMU, broad/targeted scoped checkptr,
tagged typed PartA/PartB/storage/owner/cell checks, native comparisons, GC stress,
codec references and diff checks. Final tagged ARM ordinary/scoped repeats20,
explicitly including hybrid, merge, death tests and preceding owner fixtures.
Other logs /tmp/opus-round-hybrid-primary-go-norm-owner.log,
/tmp/opus-round-hybrid-secondary-go-norm-owner.log,
/tmp/opus-round-merge-primary-go-norm-owner.log and
/tmp/opus-round-norm-mutation-owners-arm-{checkptr,golden}-repeat.log.
No native GC/extension EOF/raw callback/opaque-image repair, global lifetime
safety, performance or direct macOS claims.

Four fold-output rounds stage concrete mono/singleton destinations, propagate
them through stereo, connect primary norm output, then connect dual-stereo norm
output with cached numeric lane origin. Rounds1/2 stage owned private APIs and
direct equivalence fixtures; rounds3/4 connect actual quant-all-bands _norm Go
storage in normal decode, stereo/dual paths and both opuscc encode-only RDO
trials. Not opusccenc. Counts remain100 unsafe.Pointer/104 uintptr: genuine
owner/offset migration, not opaque aliases, cast hiding or token cleanup.

quantBandFoldOutput holds actual []float32, int64 numeric origin, explicit owned
flag and separate legacy pointer. Nil-owned never falls back. Mono preserves
context/division BEFORE N1 shortcut, original scratch/partition/transform order,
and resynth/presence guards. Scaling retains sqrt(N0) at original point, full
destination prefix BEFORE EACH source read, float32 product then sequential
store, without new sample snapshots/views before guards. N1 keeps encode caching,
channel/sign/entropy/resynth ordering and final live X copy AFTER both channels;
legacy first-cell store stays direct, owned store uses actual root. No changed
float32 products, wrapping, cache reloads, pulse behavior, masks, seed or context.

Stereo forwards SAME destination through N1, selected N2 mid, either mid/side
recursion order and intensity handling. Originally absent side fold output
remains absent; fold occurs before original reconstruction/inversion. Raw public/
private adapters remain, with no added slice-view validation on legacy N1.
Primary norm now supplies root/origin directly to mono/stereo/both RDO trials;
dual left and right execute in original order. Norm input/hybrid folding/merge,
spectrum X/Y, entropy and external-image adapters still retain legacy boundaries;
this is NOT a whole-norm/spectrum/opaque-image or native GC repair claim.

Owned output binding keeps last/N==0 guards BEFORE table/lane/root access and
forms original FULL N output prefix at original argument point, AFTER fold input
view and BEFORE DSP entry, even when resynth will skip writing. Numeric band
index retains original int32 M*boundary-offset; independent lane/index/length
coordinates combine in int64 for bounds. Relative negative-index checks remain.
Each actual store forms its on-demand span without copying data. Right lane is
cached at the EXACT original setup read that bound norm2, not recovered from
addresses or reloaded later from mutable bands. Length-zero return, allocation,
left binding, conditional right-table read and required right-first-cell check
remain in order. Existing scalar norm heads remain for other consumers. RDO
snapshots/restores/folding and dist0>=dist1 tie choice are unchanged. Upstream
bands.c singleton/mono/stereo fold and normalization allocation inspected.

Existing mono/stereo fixtures retain physical extents and compare COMPLETE
entropy/context (rebasing ONLY separate entropy-instance pointer), masks and
whole float-bit buffers/guards; stereo includes N1/N2/intensity and block/TF
cases under GC/stack growth. New bounded N1 final-copy alias oracle checks
independent buffers/entropy and exact negative-zero/NaN payload copy bits.
New Go-only four-bin shared storage compares full states/bits plus an independent
sequential-doubling oracle, ensuring no input snapshot/hoisted reads. Existing
band-view fixtures add actual root/origin/full-prefix identity, full argument
bounds and nil/unused guards without enlarging allocation. Existing pointer-only
norm GC oracle runs BEFORE an independent equal-sized cached-lane fixture;
mutable setup boundary is restored after checking already-bound right origin.
No native fixture/hash/bytes/bit-count/seed, codec golden, scope or tolerance
change; no fake prefixes, pointer-slot poisoning or original fixture repair.
Go-only invalid bounded/alias cases are not new C-valid/restrict-domain claims.

Each round passes full amd64/386 and ARM64/QEMU, broad/targeted scoped checkptr,
tagged typed PartA/PartB/storage/owner/cell checks, native comparisons, GC stress,
codec references and diff checks. Final tagged ARM ordinary/scoped repeats20,
explicitly including all fold/root/alias fixtures. Logs
/tmp/opus-round-mono-go-fold-output.log,
/tmp/opus-round-stereo-go-fold-output{,-retry}.log,
/tmp/opus-round-primary-norm-go-fold-output.log,
/tmp/opus-round-dual-norm-go-fold-output{,-final}.log and
/tmp/opus-round-fold-output-owners-arm-{checkptr,golden}-repeat.log.
Round2's initial native run hit found-pointer-to-free-object while
TestExtensionNextAgainstC was active; COMPLETE UNCHANGED retry passed, failed
log retained. Other rounds/final passes have no validation failures. A unique-
text edit mismatch was resolved BEFORE round3 validation, with no test weakening.
No native GC/extension EOF/raw callback/opaque-image repair, global lifetime
safety, performance or direct macOS claims.

Four PVQ pulse-codec rounds retain actual encoder search output, actual CWRS
decoder output, stage owned CWRS ranking, then connect owned pulse encoding.
Round3 stages ranking with direct owned fixtures; round4 connects it to the
actual opuscc alg_quant pulse allocation. Search/mask/encode/normalization now
keep the SAME encode slice; decode/normalization/mask keep the SAME decode
slice. Neither alg_quant nor alg_unquant manufactures a pulse head pointer.
Public/raw APIs remain; opusccenc is untouched. Counts remain100 unsafe.Pointer/
104 uintptr: actual owner dataflow, not token cleanup or opaque pointer aliases.

PVQ search shares unchanged span arithmetic. Legacy views remain X BEFORE pulse
output; owned search clips actual pulses[:N] at SAME point BEFORE temporary y/
signs allocation. Signs/absolute stores, sum/fallback guard, projection/floor,
N+3 wrapping, two separate emergency energy additions, strict tie comparison,
ordered energy updates and final signed pulse stores remain verbatim. Encoder
still allocates N+3 physical pulse words, decoder N; no padding/root enlargement,
new algorithm snapshots or changed feature/arch behavior. Input spectrum remains
a legacy scalar boundary; this does NOT claim spectrum ownership migration.

celtPulseVector carries actual []int32 plus separate legacy pointer and explicit
discriminator. Owned spans NEVER fall back to legacy, including nil owners.
CWRS decoder forms output view at SAME first-body point BEFORE existing K/N
assertions; rows/position, integer wrapping, int16 captures, unsigned subtraction,
zero stores, sequential writes and separate float32 energy updates are unchanged.
Pulse decoding computes SAME table totals and consumes entropy BEFORE forming
output prefix. No eager output checks, including on the public raw adapter.

CWRS rank keeps n>=2 assertion BEFORE source view, last-cell sign/absolute reads,
reverse iteration, live table calls and repeated y[j] sign reads after k updates.
No pulse/table snapshot. Encoding preserves K assertion BEFORE totals, totals
BEFORE ranking, ranking BEFORE entropy writes. Root clips stay deferred to
original consumers. Scalar raw wrappers use identical shared kernels without
binding external byte images or guessed owners. Upstream cwrs.c/vq.c search,
rank/reconstruction, entropy calls and macros inspected; no arithmetic changes.

Original hand-computed two-bin search/silence fixtures now add equal-size owned
results; added Go variants retain N+2 guard shape and compare whole pulse arrays,
full spectrum float bits and energy with GC/stack growth across budgets/sizes.
Existing exhaustive CWRS reconstruction rows retain n+2 extents and independent
L1/energy/guard oracle while comparing complete owned vectors/energy. Existing
rank fixtures retain their index oracle and extents, add actual-root reads and
input immutability checks. Existing encode streams retain 512-byte buffers and
vectors; equal-size independent owned runs compare ALL bytes and COMPLETE
entropy after EACH vector and finalization, rebasing ONLY the comparison copy's
independent buffer pointer. Original decode/round-trip oracle remains unchanged.

Bounded Go-only checks preserve owned-nil discrimination, output-view BEFORE
CWRS assertions, rank assertion BEFORE source view, decoder entropy BEFORE a
short output error, encoder K diagnostic BEFORE input, and input failure BEFORE
entropy writes. No fake prefixes, out-of-allocation addresses, pointer-slot
poisoning, original fixture repair/enlargement or weaker oracle/tolerance. These
new invalid bounded Go cases are not claimed as new C-valid/restrict domains.
Existing native vectors/hashes/bytes/bit counts/seeds, codec references and
normalization tolerances are unchanged. No new lifetime/opaque-image claim.

Each round passes full amd64/386 and ARM64/QEMU, broad/targeted scoped checkptr,
tagged typed PartA/PartB/storage/owner/cell checks, native comparisons, GC stress,
codec references and diff checks. Final tagged ARM ordinary/scoped repeats20,
explicitly including search, CWRS, encoding and all preceding owner fixtures.
Logs /tmp/opus-round-pvq-search-go-pulse-output.log,
/tmp/opus-round-cwrs-decode-go-pulse-output.log,
/tmp/opus-round-cwrs-rank-go-pulse-input.log,
/tmp/opus-round-cwrs-encode-go-pulse-input.log and
/tmp/opus-round-pvq-codec-pulse-owners-arm-{checkptr,golden}-repeat.log.
No validation failures in this batch. No native GC/extension EOF/raw callback/
opaque-image repair, global lifetime safety, performance or direct macOS claims.

Four partition/PVQ rounds stage concrete lowband recursion, connect prepared
scratch to it, retain actual pulse inputs through residual normalization, then
retain actual pulse inputs through collapse-mask extraction. Round1 stages the
owned partition API with direct equivalence fixtures; round2 connects actual
allocated/borrowed prepared scratch from mono/stereo/RDO to recursive folding.
Rounds3/4 connect BOTH opuscc alg_quant and alg_unquant pulse allocations, not
opusccenc. Counts remain100 unsafe.Pointer/104 uintptr: real owner/offset flow,
not opaque aliases, removed ABI, moved casts or token cleanup.

quantPartitionLowband carries actual []float32 root, int64 numeric origin,
separate legacy pointer and explicit discriminator. Splits preserve the original
required second-half first sample check AFTER theta/budget updates and BEFORE
recursion, with independent child coordinates combined in int64. No eager full
parent/child views. Both recursion orders, rebalance, gain rounding, fill shifts,
cache reloads, local split contexts, entropy, mask and seed remain unchanged.
Fold leaves form the SAME full consumed source prefix at each original sample
read AFTER seed/noise work, preserving destination-view-before-source and live
sequential output aliases. No added samples/snapshots or hidden scalar heads.
Raw adapters and absent-source/noise behavior remain. Known readable prepared
roots project directly; unreadable signed/unused origins explicitly retain the
original scalar boundary. Actual transform-to-partition lowband data no longer
round-trips through a head pointer. Spectrum X/theta/recursive output, initial
uncopied fold sources, pulse codecs/search and external/raw ABI are NOT claimed
migrated. No unbound/opaque image gains a Go map or registered ownership here.

Residual normalization forms actual pulse[:N] BEFORE legacy output view, then
uses the SAME reciprocal-square-root float32 gain, conversion, products and
sequential stores. Raw helper shares unchanged span arithmetic. Fixed-point-
only shift stays ignored; no new zero-work shortcut. Collapse keeps B<=1 return
BEFORE touching pulse/N geometry, unsigned division BEFORE full-N prefix, exact
block slices and signed-to-uint32 OR/bit shifts. Trailing remainder remains
ignored; no source snapshot. Mask/encode/resynth ordering and normalization/
inverse-rotation/mask ordering stay unchanged. Encode retains N+3 physical pulse
padding and decode N physical words at SAME allocation points. Only consumers'
head conversions disappear; decode/encode/search bridges remain deliberately.
Upstream bands.c recursive split/fold and vq.c helpers/macros inspected.

Existing partition fixtures compare raw/owned mask, COMPLETE entropy/context
(rebasing ONLY the separate entropy-instance pointer), full spectrum and source
bits/guards on SAME extents, including GC/stack growth and original Go-only
sequential overlap. Direct child tests verify actual root/origin, live reads,
no owned-to-legacy fallback and required interior-cell failure. Prepared root
projection and bounded unused identity checks preserve physical allocations.
Existing residual manual numeric oracle adds equal-sized full-bit/guard/input-
immutability comparisons; existing mask manual per-block oracle adds owned
results, source immutability, sign-bit/32-block/remainder and unused-guard cases.
No changed native physical fixtures/hashes/bytes/bit counts/seeds, codec goldens,
tolerances or scope. Go-only bounded/alias cases are not new C-valid/restrict-
domain claims. No fake prefixes, pointer-slot poisoning or production snapshots.

Each round passes full amd64/386, ARM64/QEMU, broad/targeted scoped checkptr,
tagged typed PartA/PartB/storage/owner/cell checks, native comparisons, GC stress,
codec references and diff checks. Final tagged ARM ordinary/scoped repeats20,
explicitly including partition/source/pulse helper fixtures. Logs:
/tmp/opus-round-partition-owned-lowband-recursion{,-retry}.log,
/tmp/opus-round-partition-prepared-lowband-owner.log,
/tmp/opus-round-residual-normalise-go-pulse-owner.log,
/tmp/opus-round-collapse-mask-go-pulse-owner.log and
/tmp/opus-round-partition-pulse-owners-arm-{checkptr,golden}-repeat.log.
First round initially hit native found-pointer-to-free-object with
TestExtensionNextAgainstC active; COMPLETE UNCHANGED full retry passed. Failed
log retained. Round2's new focused compatibility test initially put root-1
outside its allocation and failed checkptr; retained reproduction log is
/tmp/opus-round-partition-prepared-unused-fixture-focused-failure.log. Corrected
ONLY that new Go test to an interior slice of SAME five-element allocation so
both fallback coordinates lie inside real storage, then ran COMPLETE unchanged
validation. No original fixture/oracle/tolerance weakening, suppressions or
production fix. Rounds3/4 have no validation failures. No native GC/EOF/callback/
opaque-image repair, global lifetime safety, performance or direct macOS claims.

Four lowband-transform rounds extract unchanged Haar/deinterleave span kernels,
then retain prepared lowband roots through Haar and deinterleave respectively.
Rounds1/2 stage kernels; rounds3/4 connect actual allocated/borrowed scratch from
the preceding owner work, including mono/stereo/RDO paths. A required copy now
returns real owner/origin metadata instead of immediately reducing it to a
scalar pointer. The existing scalar bridge remains at recursive quant_partition;
normalization/folding inputs and band output X are NOT claimed migrated here.
Counts remain100 unsafe.Pointer/104 uintptr; no alias/cast hiding or token cleanup.

The existing concrete scratch carrier now supplies consumed spans and explicit
legacy/unused pointer boundaries. copyFrom compatibility remains, but active
preparation retains metadata and performs destination-prefix BEFORE source-view
then copy, without extra sample snapshots or allocation. Its owned zero-work
case preserves the original unused-address boundary at the original point.
Presence and disabling still distinguish owned nil from legacy storage. Scalar
conversion is deferred to the remaining DSP boundary; known roots are never
recovered from integer addresses or serialized/opaque byte images.

Haar keeps pairs=int(N0>>1), nonpositive pair/stride guard BEFORE views, native-int
2*pairs*step extent arithmetic (including386 narrowing), column-then-pair order,
separately rounded .70710678f products and sum-before-difference stores. Only
complete pairs are consumed; odd trailing rows remain untouched and are NOT
invented as readable storage. Deinterleave retains stride-positive assertion
BEFORE N0==0 guard, int32 N0*stride geometry BEFORE independent root-origin
combination, values view BEFORE temporary allocation/table access, identical
row/table/column iteration, and full temporary fill BEFORE copy back. New span
kernels preserve arithmetic and order verbatim. Legacy raw adapters remain.

Mono's prepared source starts as the original lowband; only the SAME scratch
predicate replaces it after a required copy. Haar recombination/time-resolution
and deinterleave receive its real owner until partition. Context/division/N1
shortcut, fill updates, X transforms, recursive partition, inverse transforms,
fold scaling, gain, entropy, seed, masks, and RDO trial/rollback ordering stay
UNCHANGED. Source views are still formed only at original consumer points; no
eager full-band root checks, changed validation/diagnostics/error ordering or
new snapshots. Upstream bands.c Haar/deinterleave/copy/transform sequence and
floating macros inspected. Interleave and recursive input ABI remain intact.

Existing Haar manual fixtures now compare raw/shared-span/owned full-buffer bits,
odd rows and guards under GC/stack growth on SAME extents; additional TWO-sample
owned storage proves a third odd row is not needed. Nil owned roots with extreme
unused origins skip views under original no-work conditions. Existing manual
permutation fixtures add shared/owned full-bit comparisons, preserving NaN/SNaN
payloads, table orders and guards on SAME allocations/extents. Nil zero-work and
stride-before-zero exact diagnostic checks cover the owned deinterleave route.
Existing mono/stereo context/entropy/mask/full-buffer comparisons, native spectra/
packets/hashes/bit counts/seeds, codec goldens and tolerances remain unchanged.
Prepared-root tests check actual owner/origin rather than a pointer alias. No
fake spans, fixture repair/enlargement, pointer-slot poisoning or weakened oracle;
Go-only bounded/alias checks are not new C restrict/invalid-domain claims.

Each round passes full amd64/386 and ARM64/QEMU, broad and targeted scoped
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks. Final tagged ARM ordinary/scoped
repeats20, explicitly including Haar/deinterleave/interleave and band fixtures.
Logs /tmp/opus-round-haar-live-span-kernel.log,
/tmp/opus-round-deinterleave-live-span-kernel.log,
/tmp/opus-round-lowband-haar-go-prepared-root.log,
/tmp/opus-round-lowband-deinterleave-go-prepared-root.log and
/tmp/opus-round-lowband-transform-owners-arm-{checkptr,golden}-repeat.log.
No validation failures in this batch. No native GC/extension EOF/raw callback/
opaque-image repair, global lifetime safety, performance or direct macOS claims.

Four lowband-scratch rounds introduce concrete mono scratch, propagate that
owner through stereo, connect actual allocated quant-all-bands scratch, then
connect actual borrowed spectrum-tail scratch. Rounds1/2 stage the owned private
API with direct fixtures while preserving raw adapters; round3 connects the
allocated encode/resynth path in opuscc (not opusccenc), round4 connects normal
decoding/other borrowed-scratch paths when spectrum ownership is known. Legacy
spectra retain legacy scratch. Counts remain100 unsafe.Pointer/104 uintptr:
active owner/offset representation, not token removal or opaque alias types.

quantBandScratch holds actual []float32, numeric origin, separate legacy pointer
and explicit owned discriminator. Presence uses actual owner nilness for owned
storage, never silently falls back to legacy. Required lowband copy forms the
destination prefix BEFORE source; positive owned prefixes combine independent
coordinates in int64, with unchanged int32 band geometry. Copy completes BEFORE
returning the matching scalar lowband bridge to existing Haar/reordering/recursive
consumers. That legacy DSP boundary is deliberately NOT claimed as migrated.
Copies retain Go overlap semantics; no added samples, allocations or snapshots.
Zero-length identity uses existing explicit unused-address compatibility without
validating an unconsumed full prefix; positive work remains bounded by owner.

Mono still loads context and divides block width BEFORE its N==1 shortcut;
scratch is considered only after that shortcut and the SAME recombination/
time-resolution/block predicate. Original copy-then-switch-lowband order stays.
Stereo forwards the SAME carrier through original x2/mid/intensity/side ordering;
its originally nil side scratch remains nil. No changed theta, bit allocation,
float32 products, transforms, gain, seed, recursion, mask, context, assertions,
error/validation order or snapshots. Upstream bands.c mono/stereo, allocation,
scratch-disable conditions and OPUS_COPY inspected. No new C restrict/invalid-
domain or memcpy-overlap claims from Go-only alias tests.

Quant-all-bands allocates SAME storage and extent at SAME point; no scalar head
is manufactured for its owned scratch. Borrowed scratch retains actual X root
and numeric M*boundary origin. Its setup STILL requires the original first
sample immediately, BEFORE later buffers/band iteration, but does not eagerly
form a future N-sample view. Raw source setup retains original pointer helper.
Last/non-RDO and effective-band disables clear the whole carrier at original
points; dual mono, stereo, both RDO trials and mono route retain owners. No
stale fallback pointer survives disabling. RDO snapshots/restores/tie selection,
folding, spectrum reloads and native numerical/packet oracles remain unchanged.

Existing mono fixture extents/guards stay; additional equal-sized owner copies
compare mask, complete entropy and context (rebasing ONLY the separate entropy-
instance pointer), and full X/lowband/output/scratch float bits. Added stereo
scratch fixtures cover LM2/3, block/time-frequency variants, budgets0/24/400 and
intensity0/21 with same full comparisons, GC and stack growth. Original native
fixtures/hashes/bytes/bit counts/seeds and codec references/tolerances unchanged.
Owned-scratch identity/presence, unused origin and Go-only overlap are checked.
Original pointer-only storage GC oracle is retained BEFORE independently testing
an equal-sized allocated carrier. Existing borrowed-view fixtures add exact
root/origin/legacy identity and full last-band copied bits with original guards;
a bounded nil-owner fixture preserves setup's required first-sample failure.
No physical fixture enlargement, fake prefix, pointer-slot poisoning, fixture
repair, weakened scope/oracle or new allocation/image lifetime claim.

Each round passes full amd64/386 and ARM64/QEMU, broad plus targeted scoped
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks. Final tagged ARM ordinary/scoped
repeats20, explicitly including mono/stereo/scratch fixtures. Logs:
/tmp/opus-round-quant-band-mono-owned-scratch.log,
/tmp/opus-round-quant-band-stereo-owned-scratch.log,
/tmp/opus-round-quant-all-allocated-lowband-owner.log,
/tmp/opus-round-quant-all-borrowed-spectrum-scratch-root.log and
/tmp/opus-round-lowband-scratch-owners-arm-{checkptr,golden}-repeat.log.
No validation failures in this batch. No native GC/extension EOF/raw callback/
opaque-image repair, global lifetime safety, performance or direct macOS claims.

Four deemphasis rounds extract the unchanged stereo live-span kernel, retain
the actual decoder input pointer list, retain actual preemphasis filter memory,
then carry actual decode_mem history roots plus numeric origin to sample
consumers. Both lost and normal decoder exits use the owned route. The original
raw driver and scalar stereo adapters remain; matching input pointers are NOT
reconstructed from integer addresses. Counts remain100 unsafe.Pointer/104
uintptr: active owner representation, not token removal, ABI deletion, new
storage, fixture repair or cast centralization.

Pointer-list storage reuses the concrete owned list carrier; entries remain
live and are reloaded at original per-channel points, never copied to a new
list. Concrete memory carrier has explicit owned/legacy distinction; its views
form only when originally consumed. Decoder passes its actual two-float field,
not a scalar head. History channel carrier holds real Go slice plus origin;
only exact live-pointer identity, channel and consumed extent match permit
owned reads. Unknown/mismatched pointers, signed origins, unsupported extents
and unused samples retain legacy pointer behavior; no address registry, opaque
alias or new ownership claim for caller byte images is introduced. Numeric
root-coordinate combinations widen to int64, not DSP/stride arithmetic.

Stereo dispatch still forms the two-entry list BEFORE coefficient read, skips
memory/data/PCM views for N<=0, otherwise forms memory before left/right/PCM
consumption. Kernel captures both memories once, computes both channel samples
before stores, preserves x+VERY_SMALL+m and separately rounded products/scaling,
left-before-right PCM stores and final memory commits. Generic driver retains
scratch allocation BEFORE coefficient read/division, memory read BEFORE live
input pointer read, C==0 do-while channel zero, per-sample view timing and the
accumulating x+m+VERY_SMALL versus nonaccumulating x+VERY_SMALL+m distinction.
Per-channel memory commit remains BEFORE downsampled PCM stores, including
memory/output aliases; scratch remains one N-float buffer reused per channel.
No added sample snapshots, coefficient reloads, changed FMA/grouping, buffer
validation, diagnostics, error order or public ABI. Positive history binding
checks metadata only; samples stay live across overlapping PCM writes. Existing
signed/unused PCM cursor compatibility remains unchanged. Upstream
celt_decoder.c stereo/general deemphasis inspected against these contracts.

Existing guarded stereo/manual/chunk fixtures additionally exercise the shared
span kernel on SAME extents. Existing generic N/C/downsample/accum manual matrix
now compares complete owned-list/memory and history-bound PCM bits plus state
against raw/manual values under GC/stack growth. Existing nil/zero-work cases
also cover owned memory, preserving fast-stereo nil-memory behavior versus the
generic channel-history requirement. Existing memory/PCM alias checks include
the actual owned memory route. New history tests verify exact root/origin,
geometry/identity/signed/unused fallbacks and live pointer-list mutation; a
Go-only overlapping input/PCM case checks sequential live reads for accumulating
and nonaccumulating paths. These aliases are NOT new C restrict-domain claims.
No physical fixture enlargement, fake spans, pointer-slot poisoning, expected
value/hash/packet/sample/range changes or weakened native/codec tolerances.

Each round passes full amd64/386 and ARM64/QEMU, broad and scoped deemphasis
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks. Final tagged ARM ordinary/scoped
repeats20. Logs /tmp/opus-round-deemphasis-stereo-live-span-kernel.log,
/tmp/opus-round-deemphasis-go-input-pointer-list.log,
/tmp/opus-round-deemphasis-go-filter-memory-owner.log,
/tmp/opus-round-deemphasis-go-channel-history-roots.log and
/tmp/opus-round-deemphasis-owners-arm-{checkptr,golden}-repeat.log. No validation
failures in this batch. No native GC/extension EOF/raw callback/opaque-image
repair, global lifetime safety, performance or direct macOS claims.

Four count/renormalization rounds reduce tagged non-test token counts from
101 unsafe.Pointer/105 uintptr to100/104. This modest reduction is HONEST cleanup:
round1 removes disabled-DRED's unreachable initialization-failure free/cast and
a redundant uintptr(uint32(0)) stub conversion. It does NOT represent live pointer
migration or new lifetime safety. Most remaining tokens are deliberate public
ABI, callback/varargs, opaque-layout or signed/unused-address compatibility; no
new aliases, reflect conversions or relocated casts are used to hide them.

Disabled DRED initialization keeps loaded0, arch0, magic0xD8EDDEC0 store order,
then OPUS_OK; there is no Go model/SIMD initializer that can return failure.
Creator keeps nil-TLS/registration allocation checks and errors; it directly
calls this guaranteed-success initializer instead of an unreachable free branch.
The disabled allocation stub still returns zero/-5. Actual destruction/free,
magic invalidation, registration cleanup, raw/typed public APIs and12-byte layout
remain. Upstream opus_decoder.c's ENABLE_DRED/init/create gates inspected; this
cleanup does NOT apply to an enabled native DRED model. Existing guarded layout,
factory, registered/weak lifetime checks stay; the stub gains exact zero/error
coverage. No caller data, ABI or allocation-failure oracle is weakened.

Round2 extracts the unchanged renormalization span kernel and stages a concrete
owner+numeric-origin entry, round3 connects actual PLC-noise spectra, round4
connects actual anti-collapse spectra. Consumed spans are formed on demand with
independent root/origin combinations in int64; original band/count arithmetic
remains int32. Existing public raw vector adapter stays. N<=0 returns BEFORE
owned root/origin validation. Kernel preserves scalar float32 square/product/
accumulation order, E=float32(1e-15)+energy, float64 sqrt conversion, float32
reciprocal then gain multiplication, and sequential g*x stores. No extra sample
copy, snapshots, allocation, norm approximation or arithmetic reassociation.
Upstream vq.c, PLC and anti-collapse renormalization consumers/macros inspected;
existing numerical/hash oracles and native floating tolerances stay unchanged.

PLC still fills each positive-width band with sequential LCG values BEFORE
renormalization, skipping zero/nonpositive widths without any spectrum cursor;
channel/band order and final seed commit remain. Anti-collapse keeps repair/mask/
history/seed behavior and calls renormalization only after repair. It explicitly
retains its original REQUIRED first-sample argument check before normalization,
even if a nonpositive count would otherwise skip all work. No eager band views,
changed live reloads, partial-store behavior or assertion ordering. Raw signed
history/unused ends and opaque storage elsewhere remain untouched and unclaimed.

Grouped normalization manual-value/guard fixtures now additionally compare ALL
owned-vs-raw output bits on SAME extents with GC/stack growth, including zero,
small/mixed-sign vectors and gains0/.5/1/2. Nil zero/negative-count owned roots
with INT_MIN/0/INT_MAX unused origins do not form views. Existing full PLC-noise
and anti-collapse/native fixtures retain seeds, spectra, masks and expected
values; PLC adds a nil-spectrum unused zero-width band, anti-collapse adds a
bounded negative-width required-sample check. Those two are Go-only guard tests,
NOT new C-valid-domain claims. No enlarged fake history, pointer-slot poisoning,
fixture ownership repair or changed golden values/codec tolerances. opusccenc is
untouched; this is active decoder span migration plus separately identified
count cleanup, not four rounds of cast/token removal.

Each round passes full amd64/386 and ARM64/QEMU, broad plus targeted scoped
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks. Final tagged ARM ordinary/scoped
repeats20, explicitly including DRED and normalization pointer fixtures. Logs:
/tmp/opus-round-disabled-dred-unreachable-cast-cleanup.log,
/tmp/opus-round-renormalise-live-span-kernel.log,
/tmp/opus-round-plc-noise-go-renormalise-root.log,
/tmp/opus-round-anti-collapse-go-renormalise-root.log and
/tmp/opus-round-renorm-count-arm-{checkptr,golden}-repeat.log. No validation
failures in this batch. No native GC/extension EOF/raw callback/opaque-image
repair, global lifetime safety, performance or direct macOS claims.

Four quant-all-bands RDO rounds retain actual X_save/Y_save and X_save2/Y_save2
snapshot DESTINATIONS, actual X_save/Y_save DOT sources, actual initial/winning
spectrum RESTORE sources, then actual norm_save2 plus its existing _norm parent
for normalization save/restore. These are active theta-RDO ENCODE branches in
opuscc, validated by existing scalar-C encode fixtures; normal decoding disables
RDO and opusccenc is untouched. This is NOT a decoder-hot-path or production
encoder migration claim. Fourteen scratch-head round trips disappear there;
folding, lowband, recursive spectrum and raw external boundaries remain.

Concrete []float32 parameters are actual pre-existing Go storage, not pointer
aliases or new allocations. Save copies still form destination BEFORE source;
restore copies likewise form destination BEFORE source. Zero work returns before
any root/coordinate checks. Dot forms saved prefix BEFORE spectrum view, retains
ordered sum+float32(product) with positive-zero empty result; its shared kernel
copies no samples. Normalization forms SAME offset:offset+N window BEFORE saved
prefix and retains original int32 geometry, left-lane root and conditional last-
band skipping; _norm is actual backing storage, not reconstructed from norm.
No eager RDO buffer views or snapshots beyond upstream trial requirements.

Entropy/band-context snapshots, X-before-Y copies, downward trial before scoring,
first-result snapshots, norm save, byte save, initial entropy/context/X/Y restore,
conditional hybrid refold, upward trial/scoring and dist0>=dist1 tie choice are
UNCHANGED. Winning mask/context/entropy/X/Y/norm/byte restore sequence also stays;
no float32 product/grouping, narrowing/wrapping, live alias/reload, recursion,
assertion, diagnostic or validation-order changes. Upstream bands.c theta-RDO and
OPUS_COPY inspected; Go overlapping-copy tests remain Go-only, not new memcpy/
restrict-domain claims. Raw helper adapters and all public APIs stay intact.

Existing grouped initial/trial/norm storage fixtures now use the owned helpers
on SAME allocations and extents. Existing manual ordered-dot/cancellation/input-
alias checks also exercise owned sources. Existing signed-zero/NaN-payload copy
fixture additionally compares WHOLE owned snapshot/restore buffers and guards;
overlapping Go copy is compared for both directions of ownership transfer.
Normalization fixtures add whole-buffer save/restore/overlap equivalence and
nil zero-work roots with unused99/INT_MIN/INT_MAX origins; original values,
physical extents, guards and native C spectrum hashes/masks/packets/bit counts/
seeds remain unchanged. No fake spans, pointer-slot poisoning, fixture repairs,
weakened oracles or changed encode/decode references/tolerances. Counts remain101
unsafe.Pointer/105 uintptr: active scratch representation work, not token or
dead-adapter cleanup; no broad ownership claim for recursive/opaque consumers.

Each round passes full amd64/386 and ARM64/QEMU, broad and targeted scoped
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks. Final tagged ARM ordinary/scoped
repeats20. Logs /tmp/opus-round-quant-rdo-go-snapshot-destinations{,-retry}.log,
/tmp/opus-round-quant-rdo-go-dot-sources.log,
/tmp/opus-round-quant-rdo-go-restore-sources.log,
/tmp/opus-round-quant-rdo-go-normalization-owners.log and
/tmp/opus-round-quant-rdo-owners-arm-{checkptr,golden}-repeat.log. Initial snapshot
run retains recurring native "found pointer to free object" with
TestExtensionNextAgainstC active; COMPLETE unchanged retry passes. No reduced
scope or claim to repair native GC/extension EOF/raw callbacks/opaque images,
global lifetime safety, performance or direct macOS support.

Four autocorrelation rounds retain its actual signal as xcorr INPUT, then as
xcorr HISTORY, retain actual correlation output arrays in PLC/pitch whitening,
then bind its canonical window to the actual window120 Go array ONLY on exact
pointer identity and supported0..120 overlap. Other/custom/equal-valued-copy,
interior/oversized windows preserve the original legacy pointer path. This is
bounded canonical-owner reuse, not generalized mode/window ownership recovery.

pitchXCorrSamples holds concrete float slice owner, separate legacy pointer and
explicit discriminator; xcorr adapters share assertion and input/history/output
view formation, not address casts or opaque alias types. The existing all-owned
pitch-search path stays all-owned. During rounds1/2, active autocorrelation loses
its respective scalar signal-head round trips; after round2 neither input nor
history reduces signal to a pointer. Round3 carries actual ac arrays through
xcorr and the lag tail, without reconstructing output views from &ac[0]. Both
PLC and pitch-downsample whitening use it. Public/raw adapters remain intact.

Preserved autocorrelation assertions n>0 BEFORE overlap>=0 BEFORE input view,
windowed signal allocation/copy BEFORE window view, two separately rounded end
products using ORIGINAL input even when end regions overlap, fastN arithmetic,
xcorr then per-lag correction read/add/store order, all live input/window/output
aliases and return0. No extra scratch allocation, sample/table snapshots or
arithmetic reassociation. Canonical root selection forms metadata only; its
actual overlap prefix is deferred to original window consumption, and overlap0
still never reads window data. Xcorr keeps positive-lag assertion BEFORE input,
exact length+maxPitch-1 history extent (zero history for nonpositive scalar
length), output view BEFORE unrolled len>=3 assertion, four-lag temporary/load/
product/store order and scalar tails. Integer narrowing/wrapping unchanged.
Upstream celt_lpc.c and pitch.c/macros inspected; no new C invalid/restrict-domain
claims, altered assertions, diagnostics, validation order or public ABI claims.

Existing grouped xcorr whole-buffer bit/guard/alias matrix now also compares
mixed input-owner and input/history-owner routes against raw and fully owned
routes on SAME physical extents; original numerical/manual oracles stay.
Existing autocorrelation raw/input-span matrix now additionally exercises owned
output for n8/32, lag0/1/3, overlap0/2/n/2/n, and separate/in-place/shifted output
with window/input aliases; exact diagnostic checks also cover nil owned roots.
Canonical-window tests verify actual array identity and all fallback gates,
then compare whole120-float table bits for live window/input/output aliases,
lag0/1/3/24 and overlap0/2/60/120 under GC/stack growth. Table is restored after
EVERY pair and on exit; no enlarged fake spans, pointers poisoned or changed
fixtures/oracles/tolerances. These aliases are Go-only, not new C restrict-domain
claims. Existing native comparisons/packets/samples/ranges/hashes and full
encode/decode references unchanged. Counts remain101 unsafe.Pointer/105 uintptr:
active owner work, not token/dead-adapter cleanup or fixture ownership repair.

Each round passes full amd64/386 and ARM64/QEMU, broad plus targeted scoped
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks. Final tagged ARM ordinary/scoped
repeats20. Logs /tmp/opus-round-autocorr-xcorr-go-input-owner.log,
/tmp/opus-round-autocorr-xcorr-go-history-owner.log,
/tmp/opus-round-autocorr-go-correlation-output.log,
/tmp/opus-round-autocorr-go-canonical-window-owner.log and
/tmp/opus-round-autocorr-span-owners-arm-{checkptr,golden}-repeat.log.
No validation failures in this batch. No native GC/extension EOF/raw callback/
opaque-image repair, global lifetime safety, performance or direct macOS claims.

Four inverse-MDCT/FFT rounds extract the unchanged backward live-span kernel,
retain its FFT output as actual float-lane owner plus numeric origin, introduce
exact strided owned MDCT input (private staged API in round3), then CONNECT the
actual synthesis frequency scratch to owned MDCT/copy consumers in round4.
Synthesis no longer manufactures freqHead, including upmix's staging copy source;
no extra scratch allocation or sample copies introduced. Upmix's LEFT transform
still reads borrowed output staging through the original legacy numeric-origin
path; output pointer-list loads remain live. Public/raw APIs, signed history,
unused-address compatibility and opaque caller storage remain unclaimed.

MDCT lookup/trig then input/output/window/reversal views retain formation order.
Kernel retains explicit float32 products, real/imag bitrev store order, FFT timing,
both-end capture BEFORE post-rotation stores (including odd-N4 middle pair), and
sequential TDAC window reads/stores. FFT-state lookup precedes its original
required first-cell output check; that check now validates an index without
forming a pointer. FFT carrier holds the actual output slice and scalar origin;
each butterfly still forms ONLY its consumed prefix after unchanged guard/assert
order. No factor/twiddle/window/sample snapshots, complex reinterpretation of
float cells, changed assertions, rounding, DSP arithmetic or sequential stores.

Owned MDCT source holds a concrete slice, separate legacy pointer and explicit
owned discriminator; nil owners never silently become legacy pointers. On-demand
strided span uses SAME int32 last/first/end/width arithmetic and base correction,
only combining independent root coordinates in int64. Negative/zero strides and
singleton extreme strides retain exact minimal consumed extents: not a universal
N-element contiguous prefix. No new eager spectrum/frequency views. Synthesis
keeps mono/stereo/upmix/downmix block order, independent half-products, silence
handling, output reloads and the SAME one N-float scratch; staging copy still
forms destination BEFORE source and uses overlapping Go copy. No fallback or
new ownership claim for signed borrowed staging or external opaque images.
Upstream mdct.c backward and kiss_fft.c stage dispatch inspected; this is Go
representation work, not a new C restrict/invalid-domain contract.

Extended grouped fixtures retain physical extents, guards, expected values and
independent existing native/hash oracles. All four canonical MDCT shifts and
strides -2/0/1/2 compare whole output/guard bits across public raw, numeric input,
positive/negative numeric output, shared span kernel and owned input. Existing
whole-buffer complex-vs-float FFT checks now also compare actual owned lanes,
including guards. Owned strided identity tests cover forward/backward/zero and
one-cell INT_MIN/INT_MAX strides under GC/stack growth; no broader fake history.
Added staging-copy whole-buffer checks cover lengths0/1/3/120 and identical,
left/right overlapping origins under GC/stack growth. Existing alias tests,
packets, samples, ranges, hashes and encode/decode goldens/tolerances unchanged.
No fixture repair, pointer-slot poisoning, cast centralization or token cleanup;
counts remain101 unsafe.Pointer/105 uintptr, including tagged non-test sources.

Each round passes complete amd64/386 and ARM64/QEMU, broad and targeted scoped
checkptr, tagged typed PartA/PartB/storage/owner/cell checks, native comparisons,
GC stress, codec references and diff checks; final tagged ARM ordinary/scoped
repeats20. Logs /tmp/opus-round-mdct-backward-live-span-kernel.log,
/tmp/opus-round-mdct-fft-go-output-lanes.log,
/tmp/opus-round-mdct-backward-go-strided-input.log,
/tmp/opus-round-synthesis-mdct-go-frequency-root.log and
/tmp/opus-round-mdct-span-owners-arm-{checkptr,golden}-repeat.log. No failures in
this batch. No native GC/extension EOF/callback/opaque-image repair, global
lifetime safety, performance or direct macOS claims.

Four comb/prefilter rounds extract the unchanged constant-comb slice kernel and
use it directly in the active tail, retain actual prefilter scratch output in
comb filtering, retain the same scratch slice as TDAC input, then reuse the rooted
history span with numeric TDAC output origin. Prefilter no longer reduces its
scratch to a scalar pointer or manufactures an active positive output cursor.
Concrete destination fields hold actual Go slice owners; legacy fields preserve
original pointer/identity/signed/unused-address boundaries. Counts remain101
unsafe.Pointer/105 uintptr: active representation work, not cast/adapter cleanup,
new scratch/allocation, fixture ownership repair or global lifetime safety proof.

Constant-comb keeps four history temporaries, sample load/product/rolling update
order and sequential stores; its existing Go zero-work guard remains. Main comb
keeps zero-gain identity gating and overlapping Go copy, explicit legacy-x behavior
with foreign history, period clamp, tapset/product order, overlap selection, source
history BEFORE output view, initial history loads, live window reads, transition
stores and g1-zero tail copy. Owned output formation is deferred; zero-length
spans do not validate unused output origins. No new sample/table snapshots or
arithmetic reassociation. The active constant tail receives its already-proven
source/destination spans instead of reconstructing them from interior addresses.

TDAC keeps zero-fold guard BEFORE mode/window/output/scratch access; otherwise
window, output, scratch views form in original order. Products round separately,
live window/sample aliases and sequential stores remain. Its overlap-three
rounding golden stays EXACT0xc086cced. Prefilter still selects required per-channel
history starts before the overlap==0 return, preserves the original minimal
positive first-sample extent check BEFORE complete history formation (now without
creating a pointer), reuses one overlap-sized scratch, and finishes each comb
before fold. Controls remain live per channel; CC==0 visits channel zero. Signed
origins retain explicit legacy cursor fallback; only independent owned output
coordinate combinations widen to int64. Public/raw entries stay. Upstream celt.c
and celt_decoder.c inspected; no C-domain/ABI/rounding/error-order changes claimed.

Grouped existing constant-comb/manual and prefilter/TDAC fixtures retain extents,
expected values, guards and live alias semantics. Added owned-output WHOLE-buffer
full-bit comparisons cover lengths0/1/3/8/17/64, gain pairs, transition/constant
paths, zero/partial overlap, separate/in-place/shifted/history-overlapping output
and nil zero work under GC/stack growth. Existing TDAC matrix and alias fixtures
now use the fully owned path on SAME allocations; its one-float rounding result
is a one-element Go array (same4 bytes), not an enlarged fixture. Extreme unused
owned TDAC origins leave nil roots untouched at overlap0/1. No pointer-bearing
slots are poisoned. Existing native/packet/sample/range/hash and encode/decode
goldens/tolerances remain; Go-only aliases are not new C restrict-domain claims.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped comb/prefilter checkptr, typed PartA/PartB and tagged owner/cell checks
on3 architectures, native comparisons, GC stress, codec references and diff checks.
Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-comb-constant-live-span-kernel.log,
/tmp/opus-round-prefilter-comb-go-output-owner.log,
/tmp/opus-round-prefilter-tdac-go-scratch-input.log,
/tmp/opus-round-prefilter-tdac-go-history-output{,-retry}.log and
/tmp/opus-round-prefilter-span-owners-arm-{checkptr,golden}-repeat.log.
First history-output full run retains recurring native "found pointer to free
object" with TestExtensionNextAgainstC active; COMPLETE unchanged retry passes.
No scope/fixtures/tolerances reduced, and no claim to repair native GC/extension
EOF/raw callbacks/opaque images or provide performance/direct macOS guarantees.

Four quant-all-bands rounds retain actual collapse-mask, pulse-budget, TF-flag
and spectrum owners. Active normal decode now passes all five work buffers
(left/right spectra, masks, pulses, flags) as concrete Go slices. Explicit private
sources carry real typed owners plus owned flags; raw pointer fields serve only
unchanged compatibility entries. Numeric band offsets select only the REQUIRED
first sample at existing spectrum/lowband consumers, not a new universal span.
No active owner is reconstructed from an address. Counts remain101 unsafe.Pointer/
105 uintptr: active carrier changes, not cast/adapter deletion, new scratch,
fixture ownership repair or proof of global lifetime/performance/macOS safety.

All mask/pulse/flag accesses remain on demand, not eagerly bounded/copied at entry.
Every original repeated pulse read in budget clamps AND the post-mask balance
update stays live. TF is loaded at its original point. Folding masks are read at
the original left/right indices; byte narrowing and TWO sequential stores remain
(left then right, including mono right-overwrites-left). A bounded private test
checks that the first store survives failure at the second index. No all-buffer
validation is moved ahead of these stores or empty-band work.

Spectrum sources retain raw nil-right identity or owned nil-right absence for
channel/RDO decisions. Decoder's new slice channel helper preserves N==0 returning
nil views and requires the original interior right first sample otherwise. Band/
lowband offsets retain int32 arithmetic and exact one-sample geometry. Seed/context,
entropy tells, RDO snapshots/restores/byte ranges/tie choice, normalization owners,
recursion, scratch allocation, resynthesis/last-band/folding and store/reload order
remain unchanged. No arithmetic reassociation, snapshots beyond upstream scratch,
new assertions or global ABI/callback/opaque-image lifetime claims. Upstream bands.c
quant_all_bands/RDO/folding/store contracts inspected.

Existing grouped active LM0..3 mono/stereo fixtures now use owned APIs on identical
allocations/extents/guards. Existing whole encode/RDO/dual/mono/decode fixtures
retain EXACT scalar-C spectrum hashes, mask bytes, entropy bit counts, packet bytes
and seeds; no golden/tolerance changes. Pulse and TF tests check live mutation and
original signed widths; masks check mono/stereo narrowing/store order. Spectrum
fixtures compare original pointer identity to owned band views and actual decoder
slice roots under GC/stack growth, without new allocations or pointer-slot
poisoning. Empty-band tests leave all nil work owners unused while preserving seed.
The legacy spectrum helper stays for compatibility; no dead-adapter cleanup claim.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped quant-all-bands/spectrum checkptr, typed PartA/PartB and tagged owner/cell
checks on3 architectures, native comparisons, GC stress, codec references and diff
checks. Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-quant-bands-go-mask-owner{,-retry}.log,
/tmp/opus-round-quant-bands-go-pulse-owner.log,
/tmp/opus-round-quant-bands-go-tf-owner.log,
/tmp/opus-round-quant-bands-go-spectrum-owners.log and
/tmp/opus-round-quant-band-work-owners-arm-{checkptr,golden}-repeat.log.
The first mask run retains a BUILD failure, "undefined: slices", from the new
fixture's missing import. Only the import was added; the COMPLETE unchanged retry
passes. No production/fixture/golden/scope weakening, and no claim to repair prior
native GC/extension EOF/raw callback/opaque image issues.

Four allocation rounds extract the unchanged interpolation slice kernel, retain
actual bits1/bits2/threshold/cap rows, retain active decoder offsets/caps, then
retain pulse/fine/priority outputs. Normal decode now passes ALL five allocation
arrays as concrete Go slices instead of scalar addresses. Internal curve scratch
already allocated in Go also remains slice-valued through interpolation. Explicit
allocationOutputs holds real slice owners with an owned flag; legacy scalar
pointers serve only the original raw compatibility path. Counts remain101
unsafe.Pointer/105 uintptr: active carrier work, not cast/adapter deletion,
fixture ownership repair, new scratch/storage or global lifetime/performance proof.

Driver still forms offset then cap views BEFORE total clamp/mode/reservations,
retains bit/intensity/dual-stereo reservation order, exact curve scratch lengths,
[start,end) initialization, binary searches, vector/table reloads and int32 shift/
product/wrapping arithmetic. Outputs are NOT clipped early at driver entry:
interpolation forms pulse, fine, priority views BEFORE row1,row2,threshold,cap
views, then runs the same live loops. Skip signaling, entropy calls, scalar
intensity/dual aliases, sequential pulse/fine/priority stores, remaining-balance
store BEFORE skipped-band updates and assertions remain ordered. No table/audio/
context snapshots; cap/offset/output overlaps remain live. The original pointer
interfaces and no-TLS-scratch behavior remain. Upstream rate.c inspected.

Existing grouped driver/interpolation fixtures now exercise actual owner slices
against unchanged legacy/static-mode references with SAME seven23-int arrays,
guards, buffers and expectations. Coverage retains channels1/2, LM0..3, start0/18,
negative/zero/high budgets, trims0/10, encode/decode, complete entropy context and
buffer comparisons, GC/stack growth, exact one-cell/nil-entropy/TLS-sentinel cases.
All seven alias routes compare owned vs raw: pulse/offset, intensity/dual,
balance/final-fine, priority/fine, intensity/pulse, pulse/cap and fine/cap. A private
bounded setup test checks that a short pulse view fails before nil rows/mode and
without stores. No pointer slots are poisoned; Go-only invalid/alias cases are
not new native valid-domain claims. Existing codec/native packet/sample/range/
hash, assertions, fixtures, goldens and encode/decode tolerances are unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped allocation checkptr, typed PartA/PartB and tagged owner/cell checks on3
architectures, native comparisons, GC stress, codec references and diff checks.
Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-allocation-interp-live-span-kernel.log,
/tmp/opus-round-allocation-go-interpolation-rows{,-retry}.log,
/tmp/opus-round-allocation-go-offset-cap-owners.log,
/tmp/opus-round-allocation-go-output-owners.log and
/tmp/opus-round-allocation-span-owners-arm-{checkptr,golden}-repeat.log.
First row-owner full attempt retains recurring native comparison "found pointer
to free object", with TestExtensionNextAgainstC active. The COMPLETE unchanged
retry passes; no scope/fixture/tolerance reduction. Extension EOF/global native
GC issues, opaque images, raw callbacks and direct macOS/global safety/performance
guarantees remain outside these typed numeric buffer interfaces.

Four synthesis/denormalization rounds extract the unchanged denormalization
kernel, retain actual frequency scratch, retain active spectrum owners, then
retain actual output-pointer lists. Both normal and PLC synthesis now receive X
as []float32 and out_syn[:] as []*float32. Explicit private input/output carriers
contain real concrete slice owners and an explicit owned flag; legacy fields are
only for original pointer/negative-history/unused-address compatibility. They do
not promote caller bytes or reconstruct active owners from addresses. Counts
remain101 unsafe.Pointer/105 uintptr: active representation work, not dead/cast
cleanup, new scratch/allocation, fixture ownership repair or global GC-map safety.

Denormalizer still reads table/geometry/downsample bound, forms output, clears
and returns on silence BEFORE input/energy views, then forms input and energy
views and runs the same live loop. Prefix clear, per-band gain clamp/Horner/
integer exponent arithmetic, sequential coefficient stores, late start<=end
assertion and final tail clear remain ordered; no snapshots or float32 product
changes. Numeric origins combine in int64 only at owned span formation, not in
DSP arithmetic. Zero-length owned spans ignore unused origins; raw signed-history
and unused-end behavior remains. Spectrum input bounds are formed only when
needed, so silent nil/unused roots stay unused.

Synthesis still captures geometry and allocates exactly ONE N-float scratch,
then forms its output list at the original point. Upmix stages before transforms,
finishes ALL left blocks before right blocks; downmix uses existing output scratch
and rounds the two .5 products independently; ordinary channels stay sequential.
MDCT strided input/output consumers and numeric origins are unchanged, not
replaced with universal contiguous spans. Output pointers remain live list loads,
not a copied pointer array. Legacy pointer APIs, PCM/callback/layout boundaries,
seed/error/assertion behavior and the floating saturation identity remain.
Upstream bands.c and celt_decoder.c contracts inspected.

Existing grouped denormalization fixtures additionally compare live shared input/
output/energy owner slices on EXACT same18-float parents, start/downsample variants
and guards. Existing synthesis matrix now exercises the owned path against its
unchanged legacy/static-mode reference: LM0..3, all mono/stereo/upmix/downmix and
zero-channel layouts, transient/silence, downsample1/2/3/4/6 and band limits. The
fixture factory exposes its ORIGINAL actual X allocation as a slice (same size/
guards), not a pointer-to-slice repair or enlarged allocation. Existing alias and
manual upmix ordering fixtures use owned inputs/output lists; Go-only overlap
restrictions remain documented, with no new C alias-validity claim. Codec goldens,
packet/sample/range/hash, extents and encode/decode tolerances are unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped denormalize/synthesis checkptr, typed PartA/PartB and tagged owner/cell
checks on3 architectures, native comparisons, GC stress, codec references and diff
checks. Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-denormalise-live-span-kernel{,-retry}.log,
/tmp/opus-round-synthesis-go-frequency-owner.log,
/tmp/opus-round-synthesis-go-spectrum-owner{,-retry}.log,
/tmp/opus-round-synthesis-go-output-list-owner.log and
/tmp/opus-round-synthesis-span-owners-arm-{checkptr,golden}-repeat.log.
First kernel and spectrum-owner full attempts retain recurring native
"sweep increased allocation count" failures, with TestExtensionNextAgainstC
active. Both COMPLETE unchanged retries pass; no fixtures/scope/tolerances reduced.
No claim to repair those native GC/EOF issues, opaque images, raw callbacks,
performance, or direct macOS/global lifetime safety.

Four coarse pitch-xcorr rounds extract the unchanged kernel, then retain output,
input and history owners. Active pitch search now passes x4, y4 and corr directly
as concrete Go slices, without scalar-pointer round trips for coarse correlation.
The public raw xcorr entry remains. Counts remain101 unsafe.Pointer/105 uintptr:
active carrier work, not adapter/cast deletion, new allocation/scratch, fixture
ownership repair or proof of global lifetime/performance/macOS safety.

Max-pitch assertion stays BEFORE all views; setup remains input, history, output;
the unrolled len>=3 assertion stays AFTER those views and before output stores.
Input extent is length. History extent is exactly length+maxPitch-1 when length>0,
otherwise zero, preserving original int32 arithmetic. This is ONE LESS than the
subsequent candidate-energy extent, so y4 remains its full owner for candidate
selection; no fixtures or reads are extended. Scalar zero-length lags1..3 still
write zero correlations without touching either input. Unsupported short unrolled
lengths still assert. Four accumulators, rotating history temporaries, original
sample/remainder order, sequential four-lag stores and scalar tail stores stay;
no audio/context snapshots, algebraic dot-product replacement or reassociation.
Original ignored arch behavior remains. Upstream pitch.c and pitch.h inspected;
existing Go public validation and assertion behavior remain unchanged.

Existing grouped xcorr fixtures keep the independent scalar expected values for
lengths0..19 and lags1..13 on exact history extents. Added whole-buffer/full-bit
comparisons cover remainders, output overlapping input/history (including shifted
starts), guards, GC and stack growth. These numeric aliases are Go compatibility
cases, not new independent C alias-domain claims. Directed tests retain original
max-pitch and unrolled diagnostics, bounds failure before the unrolled assertion,
no stores on failure, zero-length unrolled failure with nil inputs, and an exact
6-sample history yielding14/20/26/32 (not a7-sample enlargement). Existing codec
packet/sample/range/hash, encode/decode goldens and tolerances are unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped xcorr/search/PLC checkptr, typed PartA/PartB and tagged owner/cell checks
on3 architectures, native comparisons, GC stress, codec references and diff checks.
Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-pitch-xcorr-live-span-kernel.log,
/tmp/opus-round-pitch-xcorr-go-output-owner{,-retry}.log,
/tmp/opus-round-pitch-xcorr-go-input-owner.log,
/tmp/opus-round-pitch-xcorr-go-history-owner.log and
/tmp/opus-round-pitch-xcorr-owners-arm-{checkptr,golden}-repeat.log.
The first output-owner full run retains the recurring native comparison
"found pointer to free object" failure (TestExtensionNextAgainstC was active).
The COMPLETE unchanged retry passes; no fixture/scope/tolerance reduction and no
claim to repair extension EOF/global native GC, raw callbacks or opaque images.

Four pitch-search rounds extract live best-candidate and search kernels, retain
PLC's actual lowpass owner with numeric input/history origins, then pass actual
correlation/downsampled-history buffers to both candidate searches. Active PLC
no longer reduces lowpass to two interior scalar pointers; candidate selection
no longer reconstructs its slice arguments from scratch-array addresses. Raw
public/search/best adapters remain. Counts remain101 unsafe.Pointer/105 uintptr:
active representation changes, not dead-adapter/cast deletion, new scratch,
fixture ownership repair or global lifetime/performance/macOS guarantees.

Length assertion precedes max-pitch assertion, BOTH before views, and original
diagnostics remain. Pointer adapter forms input then history views; owned entry
uses the same half-rate extents with numeric origins, widening coordinate
combinations to int64 but retaining original int32 length+maxPitch arithmetic.
A zero-consumption input does not validate an unused origin. The same quarter-
rate scratch arrays, decimation loads, coarse correlation, refined candidate
window, -1 floor, ordered float32 sums, strict ratio comparisons/ties and .7
pseudo-interpolation remain. Final pitch store stays LAST. Candidate views are
correlations then samples, including len+max samples: the energy update after
THE LAST candidate still reads its extra terminal sample. Initial best stores,
per-iteration correlation loads, live sample/output alias effects and energy
floor ordering remain. Upstream pitch.c contracts inspected; no arithmetic or
new valid-domain/unsupported-length behavior changes are claimed.

Existing grouped best/LPC/pitch fixtures keep manual candidate/tie/NaN/zero-row,
silence and local native expected results. Added full-bit/guard comparisons cover
lengths16/17/32, max pitches8/9/16/18, shared/shifted/disjoint lowpass origins,
result words overlapping input, unused extreme input origins and assertions
before invalid owners. Best-output aliases exercise both correlation and history
words; a bounded short history fails before best initialization, preserving the
required terminal extent. Numeric alias fixtures touch no pointer-bearing slots
and are Go-only compatibility cases, not new C restrict/alias-validity claims.
Existing packet/sample/range/hash, codec goldens/tolerances and fixture extents
are unchanged. Coarse xcorr raw adapter, opaque images, callbacks, EOF and global
native GC issues remain outside these typed private boundaries.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped pitch-search/best/PLC checkptr, typed PartA/PartB and tagged owner/cell
checks on3 architectures, native comparisons, GC stress, codec references and diff
checks. Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-best-pitch-live-span-kernel.log,
/tmp/opus-round-pitch-search-live-span-kernel.log,
/tmp/opus-round-plc-pitch-search-go-lowpass-origins.log,
/tmp/opus-round-pitch-search-go-candidate-owners.log and
/tmp/opus-round-pitch-search-owners-arm-{checkptr,golden}-repeat.log.

Four PLC pitch-downsample rounds separate left sampling, stereo addition and
whitening into shared live slice kernels, then retain actual lowpass output,
left history and right history owners. Active PLC receives decode_mem as a concrete
[2][]float32 and passes both history slices and lowpass[:] directly. The final
private downsample entry has only Go slice audio arguments; original raw pitch/
downsample adapters stay for compatibility. Counts remain101 unsafe.Pointer/
105 uintptr: active carrier work, not cast deletion, new allocation/scratch,
fixture ownership repair or a global lifetime/performance/macOS safety claim.

Setup remains left view, output view, ALL left stores (indices1..length-1 THEN0),
then, ONLY for C==2, right view and its same-order additions. No eager right
geometry validation or audio snapshot; all other channel counts ignore right.
Whitening then retains autocorrelation,1.0001 noise floor, ordered float32 lag
window products, LPC, sequential .9 chirp, .8 zero and FIR5 coefficient snapshots
and MAC/history/store order. The original unsupported zero-length behavior is
not replaced by a shortcut. PLC lag-search input/index extents, local lowpass
array, seed/state non-use and final lag subtraction are unchanged. Upstream
pitch.c and celt_decoder.c inspected; public validation/error order and native
rounding/tolerances remain. Existing raw adapters use the original pointer entry,
not eager reconstruction of an optional right history.

Grouped existing LPC/pitch fixtures retain manual silence/periodic/whitening and
native expected values. Added full-bit WHOLE-buffer comparisons cover lengths8/16,
factors2/4, channel counts0/1/2, separate buffers, output overlapping either input,
shifted overlap and all buffers shared, with unchanged pairwise extents/guards.
Owned PLC silence/shared/periodic tests retain actual history roots under GC and
stack growth, and do not mutate inputs. Private bounded missing-right tests
confirm all left stores precede right-view failure AND whitening; nonstereo
counts0/1/-1/3 ignore absent right. These invalid/private and alias cases are
Go compatibility tests, not new upstream C valid-domain claims. Existing native
packet/sample/range/hash, encode/decode goldens and tolerances remain unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped downsample/PLC pitch checkptr, typed PartA/PartB and tagged owner/cell
checks on3 architectures, native comparisons, GC stress, codec references and diff
checks. Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-pitch-downsample-live-kernels.log,
/tmp/opus-round-plc-downsample-go-lowpass-owner.log,
/tmp/opus-round-plc-downsample-go-left-history-owner.log,
/tmp/opus-round-plc-downsample-go-right-history-owner.log and
/tmp/opus-round-pitch-downsample-owners-arm-{checkptr,golden}-repeat.log.
Opaque caller images, EOF, raw callbacks, global native GC failures, and direct
platform/performance guarantees remain outside these private slice interfaces.

Four anti-collapse rounds extract the unchanged slice kernel, then retain active
spectrum, collapse-mask and pulse-budget owners. Normal decode now passes X,
collapse_masks and pulses directly as concrete Go slices; no scalar-pointer
round trip for those three buffers. Existing public/raw entries and band/energy/
history roots stay. Counts remain101 unsafe.Pointer/105 uintptr: active private
representation changes, not adapter deletion, fixture ownership repair, new
scratch/storage or cast-count proof of global lifetime safety/performance.

The early start>=end return precedes every view; disabled decode avoids state/
mode/owner access. Otherwise setup remains bands, spectrum, masks, pulses, energy,
then both histories, with the original max(C,1), max(end*C,1), end and mono-decode
TWO-channel history extents. Kernel depth addition/division/shifts, float32 exp2/
sqrt/product/LM3 adjustment, ordered min/max comparisons, live history reads and
numeric channel/band positions are unchanged. Masks are read for each block;
pulses for each band. Noise stores, per-sample RNG progression, per-channel
renormalization and band/channel/block order stay. Seed is local, not committed
to decoder state. Upstream bands.c inspected; existing Go behavior and public
assertion/validation behavior remain, with no new valid-domain claims for aliases.

Grouped existing anti-collapse fixtures preserve manual spectral goldens and exact
extents. Added full-bit/guard comparisons cover LM0..3, channels1/2, encode/decode,
zero/full masks, two seeds and spectrum/history overlap. Last round additionally
covers mixed0x55/0xaa masks. Go-only numeric alias fixtures exercise mask bytes and
pulse words changed by earlier spectrum stores; no pointer slots are reinterpreted
or poisoned. Pulse aliases can leave the upstream valid pulse domain and are NOT
new C-oracle claims. Disabled/empty nil owners and unchanged decoder seed/state
are checked. Native comparisons, packet/sample/range/hash and codec goldens/
tolerances are unchanged. No direct macOS, opaque-image, raw callback, EOF or
global native-GC repair claims.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped anti-collapse checkptr, typed PartA/PartB and tagged owner/cell checks
on3 architectures, native comparisons, GC stress, codec references and diff checks.
Last round's expanded mixed-mask fixture set also passes the COMPLETE sequence.
Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-anti-collapse-live-span-kernel.log,
/tmp/opus-round-anti-collapse-go-spectrum-owner.log,
/tmp/opus-round-anti-collapse-go-mask-owner.log,
/tmp/opus-round-anti-collapse-go-pulse-owner{,-mixed-masks}.log and
/tmp/opus-round-anti-collapse-owners-arm-{checkptr,golden}-repeat.log.

Four CELT energy rounds separate fine-energy arithmetic into a live slice
kernel, migrate active fine quantizers to actual Go storage, separate final-energy
refinement, then migrate active final quantizer/priority storage. Normal decode
passes fine_quant/fine_priority slices directly, rather than reducing them to
scalar pointers and reconstructing views. Existing pointer/source APIs and energy
history roots stay. Counts remain101 unsafe.Pointer/105 uintptr: active private
carrier work, not cast deletion, public ABI removal, new scratch/storage, pointer
map repair or a global lifetime/performance/macOS safety claim.

Fine setup still forms energy, extra and optional previous views in that order.
The shared loop captures int16 extra BEFORE its full-int eligibility/budget reads,
then snapshots int16 previous at its original point. It retains per-band entropy
tell, full-width budget arithmetic, per-channel bit consumption, float32 shift/
product boundaries and live energy/mode accesses. Final setup still forms quant,
priority and optional energy views in order; priority passes, MAX_FINE_BITS checks,
remaining-budget checks and bit decrements stay ordered. Quant values are reloaded
AFTER each entropy bit for offset calculation, not snapshotted. Nil final output
still consumes exactly the same bits. Decode's tell capture and int32 length*8
wrapping precede view setup as before. No early empty-band shortcuts were added.
Upstream quant_bands.c contracts inspected; existing Go loop behavior preserved.

Grouped existing fine/final fixtures keep their manual expected samples and
contexts. Added full-context/bit/guard comparisons cover channels1/2, previous
quantizer present/absent, storage and refinement budgets, skips, MAX_FINE_BITS,
nil final output and int32 budget wrapping. Go-only aliases use ONE ACTUAL int32
entropy field, never pointer slots: fine captures extra before the aliased entropy
counter changes; final reloads quant after each bit (expected stereo offsets
.125/.0625, not a snapshotted .25). Existing frame/native packet/sample/range/hash
and encode/decode goldens/tolerances, exact fixture extents and assertions stay.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped fine/final energy checkptr, typed PartA/PartB and tagged owner/cell
checks on3 architectures, native comparisons, GC stress, codec references and diff
checks. Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-fine-energy-live-span-kernel.log,
/tmp/opus-round-fine-energy-go-quantizer-owners.log,
/tmp/opus-round-final-energy-live-span-kernel.log,
/tmp/opus-round-final-energy-go-quantizer-owners.log and
/tmp/opus-round-energy-quantizer-owners-arm-{checkptr,golden}-repeat.log.
Opaque caller images, EOF, raw callbacks and recurring native GC failures remain
unclaimed and independent of these typed numeric buffer interfaces.

Four LPC/FIR owner rounds extract the unchanged LPC slice kernel, migrate
PLC/downsample LPC outputs and autocorrelations to actual Go owners, then retain
PLC FIR and downsample FIR5 coefficient owners rather than scalar pointers.
Final active calls pass lpc/ac/lpc2 slices and numeric channel origins; existing
public/raw LPC/FIR/FIR5 interfaces and signed-history/unused-source identity
boundaries remain. Counts remain101 unsafe.Pointer/105 uintptr: active DSP
carrier changes, not removal of necessary public adapters, cast-count safety,
new allocation/GC-map repair or performance/macOS claims.

LPC reads error from ac[0] BEFORE clearing output, then rereads LIVE ac[0] for
threshold and bailout, preserving alias effects. All float32 reflection products,
paired symmetric updates, early30dB bailout and remaining-zero taps stay ordered.
Owned output view precedes correlation view, with no ac snapshot; p==0 can ignore
nil output/large unused origin but still requires ac[0]. PLC retains the original
lpc channel index check before the call. FIR keeps x!=y identity assertion BEFORE
coefficient views, reverses all coefficients before input/output work and retains
original scalar/four-lane store order; ord==0 ignores nil/unused coefficient root.
FIR5 snapshots exactly5 coefficients BEFORE length<=0 guard or sample view, then
uses existing MAC/history/store order. No added scratch or copies beyond original
required reversal/snapshots. Upstream celt_lpc.c and pitch.c contracts inspected.

Grouped existing fixtures preserve their manual/native expectations and add span
output/guard comparisons; LPC orders0/1/2/4/8/24, thresholds/NaNs and full-memory
live output/correlation aliases; FIR zero-order unused coefficient and exact
identity assertion ordering, plus coefficients overlapping source/output on the
same unchanged extents; FIR5 aliased coefficient snapshots, empty/negative sample
lengths and5-cell validation before zero-work guard. Existing goldens, float
narrowing/tolerances, byte layouts, packet/sample/range/hash and GC-stress settings
remain. No opaque-image/EOF/raw callback/global native GC safety claims.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped LPC/FIR/FIR5/PLC checkptr, typed PartA/PartB and tagged owner/cell checks
on3 architectures, native comparisons, GC stress, codec references and diff checks.
Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-lpc-live-span-kernel.log,
/tmp/opus-round-plc-downsample-lpc-go-owners.log,
/tmp/opus-round-plc-fir-go-coefficient-owner.log,
/tmp/opus-round-downsample-fir5-go-coefficient-owner{,-retry,-retry2}.log and
/tmp/opus-round-lpc-fir-coefficient-owners-arm-{checkptr,golden}-repeat.log.
The first TWO full FIR5 validation attempts retain recurring native comparison
"found pointer to free object" failures; the first retry reports it during
TestExtensionNextAgainstC. The SECOND COMPLETE retry passes, with unchanged scope
and fixtures. These extension/global GC issues are not claimed repaired.

Four active PLC IIR rounds separate unchanged arithmetic into celtIIRSpans,
then migrate the private PLC input, coefficient and filter-memory carriers to
actual Go slices. Final celtIIROwned takes real input/coefficient/output owners,
numeric origins and the actual lpc_mem slice; no scalar-pointer reduction or
unsafe.Slice reconstruction for those private buffers. Existing public IIR,
explicit-history and numeric-pointer entries retain signed caller-owned history
and their source interfaces. Counts remain101 unsafe.Pointer/105 uintptr because
the necessary external compatibility fallbacks stay. This is active DSP carrier
work, not another fixture-ownership repair, cast deletion, new GC map or global
lifetime proof; no performance or direct macOS claim.

The order assertion precedes all views. N==0 ignores even a nil input owner with
large unused origin; ord==0 ignores nil coefficients with large unused origin.
Output/memory views retain original setup order and geometry. Coefficients reverse
before memory-history initialization; the original zeroed N+ord history and
reversed-coefficient scratch remain, with no extra snapshots/storage. Each lane
keeps explicit float32 products, live coefficient patch reloads and sequential
history/output stores, including the original positive scalar-tail history vs
negative four-lane history distinction. Memory tail view is formed AFTER all
output stores, then memory updates in forward order from live reverse output
indices, even N==0/N<ord and overlapping memory/history. No deferred/snapshot
memory writeback. PLC's original coefficient index check stays before the call.
Origins combine as before (int64 span coordinates, original int32 N-ord and N+ord
arithmetic). Upstream celt_lpc.c IIR and PLC call contracts inspected.

Existing grouped fixtures retain public/manual expectations and add bitwise full
output/memory/guard comparisons for raw span and owned paths, orders0/4/8,
lengths0/1/3/4/5/8/17, separate/in-place input, shifted coefficient roots,
coefficient/input/output/memory overlaps, nil unused owners and exact assertion
ordering/diagnostics. New short-output/history/memory alias cases explicitly retain
sequential stores for N<ord instead of reverse-copy snapshots. Physical extents,
all native packet/sample/range/hash and codec goldens/tolerances remain unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped IIR/LPC/PLC checkptr, tagged owner/cell checks on3 architectures, typed
PartA/PartB, native comparisons, GC stress, codec references and diff checks.
Final tagged ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-iir-live-span-kernel.log,
/tmp/opus-round-plc-iir-go-input-owner.log,
/tmp/opus-round-plc-iir-go-coefficient-owner.log,
/tmp/opus-round-plc-iir-go-memory-owner.log and
/tmp/opus-round-plc-iir-full-owners-arm-{checkptr,golden}-repeat.log.
No opaque caller-image, raw callback, EOF or recurring native GC safety claims.

Four CTL owner/cell rounds remove FIVE more unsafe.Pointer sites (106→101;
uintptr105 unchanged, source totals include compareopus bridges). First, MS and
projection comparison images use canonical scanned Go children, with decoded
geometry validated before temporary mode installation; no caller-byte→state
casts. Then single/custom scalar aliases and MS/projection aliases resolve real
Go int32/uint32/mode fields, not unsafe.Add addresses into serialized bytes.
All five such casts are gone from decoder_validation_compare.go. Necessary raw
public integer adapters, callbacks, signed history and production layout fallbacks
remain. This combines fixture ownership repair and a real private typed-output
capability, NOT five production pointer-map fixes or removal of uintptr ABIs.

Private opusCtlScalar holds a genuine *int32 OR *uint32 cell. Explicit numeric
conversions preserve word bits without pointer-type punning. Getter stores occur
at their old points; the error getter stores BEFORE clearing state; MS final-range
resets output BEFORE reading live stream count and reloads the actual destination
AFTER every child call before XOR/store. No scratch-output/deferred commit or
snapshots, changes to clamps/wrapping/validation/error ordering, entropy or goldens.
The existing I32/U32 output slots keep precedence and wrong-slot nil behavior.
Internal opusCtlArgs carries the optional cell; final public OpusDecoderCtlArgs
retains its original FIVE fields/positional source shape, enforced by a compile
fixture. The third commit initially put the extension in the public record; the
fourth isolates it in private dispatch to preserve final source/layout compatibility.
No external codec API or raw vararg format is changed.

Canonical comparison storage includes actual projection/MS/child Decoder/Silk/Celt
fields; tests retain weak nested modes under GC/stack growth and check child
identities, XOR and exact serialized guards. Scalar tests cover sign-bit/all-ones
bits, error self-alias clearing, opposite-signed real cells, nil validation and
existing slot precedence. MS range aliases to first/second child words prove live
reset/read/XOR behavior (expected2/0, NOT copied input3). Aliasing the live stream
count to the zeroed range output must terminate immediately, also through the
projection-relative base. Native existing live-layout alias matrices are unchanged.
Typed fixture aliases require actual supported field cells (no pointer-slot
poisoning, padding or arbitrary byte-address promotion); ordinary public raw ABI
fallbacks are unchanged. Source/layout/runtime GC limitations outside these paths
remain unclaimed. Extra dispatch/reflection cost is unbenchmarked.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and typed cell/tagged-owner checkptr on3 architectures, typed PartA/PartB, native
comparisons, GC stress, unchanged codec references and diff checks. Final tagged
ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-ms-ctl-comparison-scanned-owner.log,
/tmp/opus-round-projection-ctl-comparison-scanned-owner.log,
/tmp/opus-round-ctl-concrete-scalar-cells.log,
/tmp/opus-round-ms-ctl-typed-alias-cells{-focused,-retry}.log and
/tmp/opus-round-ctl-owner-cells-arm-{checkptr,golden}-repeat.log.
The focused fourth-round log retains an unused local st build error after a
fixture initializer extraction; only that unused binding was removed. COMPLETE
validation retry passes. No fixture/expectation/golden weakened, no performance,
direct macOS or global opaque-image/EOF/raw callback/native GC safety claims.

Four decoder comparison-owner rounds REMOVE four byte-image→decoder
unsafe.Pointer sites: ComparePrefilterFold, CompareCustomCtl,
CompareCeltDecodePrefilterImage and CompareOpusCtl. Decoder source counts (including
compareopus-tagged bridges) fall110→106 unsafe.Pointer; uintptr remains105.
Ordinary production decoder representation/API code is unchanged. Necessary public
integer ABI adapters, raw callbacks, signed history, mutated-layout production
fallbacks and external layouts remain. This is specifically fixture-ownership
repair, not four active DSP migrations or a claim to have removed uintptr ABIs.

Comparison byte images are now serialized scalar transport, NOT Go decoder
allocations. A real reflected Go allocation contains a concrete custom header or
canonical Decoder/Silk/Celt fields, scalar tail and original guard bytes. Header
pointers come from actual Go fields through Interface/type assertions, never casts
from caller bytes. Pointer slots must be null before import and export. Modes or
codebooks are attached only to the scanned owners; after processing they are
cleared before exact byte serialization. Positive scalar CTL aliases target the
new actual state at the original numeric offsets. Opus comparison images require
canonical channels/offsets; production noncanonical/opaque fallback behavior is
untouched. There is STILL one shared unsafe reflection→byte view boundary for
serialization of the genuine typed owner: not a zero-unsafe or mere cast-count
safety claim, nor a promotion of arbitrary bytes into a pointer-bearing object.

The transport view's length AND capacity remain exactly the supplied fixture
extent. No input, guard or golden is enlarged. Pointer-aligned Go structs may have
unexposed trailing alignment padding for four-byte-aligned variable C tails;
that padding is never added to the transport view or DSP consumed lengths. All
numeric bits/guards are copied exactly; only already-null pointer slots cross the
byte transport. A grouped compareopus-tagged owner test file was necessary because
ordinary opuscc tests cannot refer to tagged-only bridges. It tests weak heap
mode/window and nested SILK codebook retention under GC/stack growth, pointer-free
round trips, non-null/short/misaligned transport rejection, exact views including
Go alignment padding, scalar alias/store/range bits and prefilter dispatch images.
Existing native fixture matrices, byte comparisons, layouts, assertions, entropy
states and sample/range/hash/encode/decode goldens/tolerances remain unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
checkptr, tagged owner checkptr on3 architectures, typed PartA/PartB, native
comparisons, GC stress, codec references and diff checks. Final tagged ARM ordinary
and scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-prefilter-comparison-scanned-owner{,-retry}.log,
/tmp/opus-round-custom-ctl-comparison-scanned-owner.log,
/tmp/opus-round-prefilter-dispatch-comparison-scanned-owner.log,
/tmp/opus-round-opus-ctl-comparison-scanned-owner{,-focused}.log and
/tmp/opus-round-decoder-comparison-owners-arm-{checkptr,golden}-repeat.log.
The first full prefilter log retains a new constructor's overstrict assertion that
Go struct size equal logical C fixture extent. Corrected to allow mandatory,
unexposed Go trailing alignment padding, with original view extents unchanged;
COMPLETE retry passes. No fixture or expectation weakened, and no global native
GC/opaque-image/EOF/raw callback safety, performance or macOS claims.

Four CELT DSP span rounds retain PLC noise output, PLC synthesis attenuation,
downsample FIR5 samples and autocorrelation inputs as the actual existing Go
buffers. Active PLC/downsample callers stop converting them to scalar pointers
and rebuilding their views. Compatibility adapters retain their original source
interfaces and extents. Counts remain110 unsafe.Pointer/105 uintptr: actual
span-carrier changes, not cast deletion, new storage, pointer-map repair or proof
of global lifetime safety. No performance/macOS claims or new scratch/copies.

Noise retains the seed snapshot/writeback even for no bands/channels, per-channel
band iteration, uint32 LCG wrap, signed seed shift and live arch reload for each
normalization. Attenuation retains length<=0 before any view, ordered float32
energy accumulation, NaN-clearing comparison and live window reads after prior
output stores. Its PLC caller explicitly retains the old scalar argument's index
check before passing the Go suffix. FIR5 snapshots all5 coefficients BEFORE its
length guard/sample view; shared arithmetic retains separate MAC/store/history
order and original input sample history. Autocorrelation asserts n>0 THEN
nonnegative overlap before either input view, retains original-input window
multiplications when ends overlap, existing required window scratch, unrolled
xcorr followed by scalar tails, output-view timing and live output/input aliases.
Original contracts inspected in celt/pitch.c, celt/celt_lpc.c and celt_decoder.c.

Grouped fixtures keep their manual/native references and add full noise state/
bit/guard comparisons for all shifts/channels/seeds; attenuation finite/NaN,
nil zero-work and live window/output aliases; FIR5 lengths and aliased coefficient
snapshots; autocorrelation zero/full/overlapping windows, lag choices, full-memory
input/window/output alias comparisons and exact assertion diagnostics/order.
Pointer slots are never poisoned. Public codec APIs, external layouts and all
native packet/sample/range/hash, encode/decode goldens and tolerances unchanged.
Opaque images, raw callbacks, EOF ownership and recurring native GC issues remain
unclaimed, independent of these typed numeric-buffer interfaces.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped PLC/FIR5/autocorr/downsample checkptr, typed PartA/PartB on3 architectures,
native comparisons, GC stress, codec references and diff checks. Final ARM ordinary
and scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-plc-noise-go-output-span.log,
/tmp/opus-round-plc-attenuate-go-output-span.log,
/tmp/opus-round-downsample-fir5-go-sample-span.log,
/tmp/opus-round-autocorr-go-input-span{-focused,-retry}.log and
/tmp/opus-round-dsp-work-spans-arm-{checkptr,golden}-repeat.log.
The focused autocorrelation log retains a fixture failure: expected diagnostic
mistakenly prepended "assertion failed:" to static text ALREADY containing it.
Only the new expected-string construction was corrected; production diagnostics
and prior expectations were unchanged. The COMPLETE validation retry passes.

Four normal CELT decode work-buffer rounds retain history movement, boost
caps/offsets, TF flags and initialized caps as the existing actual Go slices.
Active callers stop reducing these arrays to scalar pointers and reconstructing
views. Scalar compatibility adapters remain; shared arithmetic loops consume the
same extents and live samples. No new scratch/allocations or input snapshots,
layout/ABI changes or pointer-map/lifetime/performance claims. Counts remain110
unsafe.Pointer/105 uintptr: this is a real span-carrier/interface change, not
cast deletion or a fix for opaque images, callbacks or EOF ownership.

History movement remains overlapping memmove, with length==0 before extent math.
Boost total-shift wrapping and entropy tell precede the empty-band guard/views;
remaining/tell are passed to a common loop exactly once. Width/quanta wrapping,
strict budget comparison, cap reloads, probability resets, sequential offset
stores and final entropy context remain. TF budget/storage/tell/logp/reservation
capture precedes output view construction; the shared loop retains live entropy
reloads and the separate selected-table mapping pass. Empty TF bands MUST NOT
skip the reservation/table-selection behavior. Caps retain band/cache/output
view order, empty guard, signed width narrowing/shifts and live byte-table reads
before each int32 output store. Original C contracts inspected in celt.c and
celt_decoder.c; public init_caps and existing tf_decode source entries unchanged.

Grouped tests retain prior manual/native expectations and add bitwise history
span/guard comparisons, full boost context/scalars/output comparisons and live
cap/offset overlap. TF tests cover all4 shifts, transient choices,0/1/32-byte
budgets, empty/nonempty bands, exact guards, nil zero extent and a single numeric
entropy-field alias (no pointer slots). Caps compare direct span output/guards,
empty/nonpositive bands and sequential byte-cache/output alias reads. Existing
native packet/range/sample/hash and encode/decode goldens/tolerances unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad relevant
and scoped history/boost/TF/caps checkptr, typed PartA/PartB on3 architectures,
native comparisons, GC stress, codec references and diff checks. Final ARM ordinary
and scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-decode-history-move-go-span.log,
/tmp/opus-round-decode-boost-work-go-spans.log,
/tmp/opus-round-decode-tf-flags-go-span.log,
/tmp/opus-round-decode-caps-go-output-span{,-focused}.log and
/tmp/opus-round-decode-work-spans-arm-{checkptr,golden}-repeat.log.

Four PLC-buffer-span rounds move active excitation-history copying, excitation
decay, periodic extrapolation and LPC-memory history preparation to the actual
existing Go slices. Pointer-based private fixture adapters remain and form exactly
the old consumed extents, then delegate to the same arithmetic. Active PLC no
longer converts these Go buffers to scalar pointers and reconstructs their views.
This is an interface/span-carrier change, NOT new storage, a repaired byte-image
pointer map or a cast deletion; counts remain110 unsafe.Pointer/105 uintptr.

Excitation history still performs forward loads/stores, not memmove; LPC memory
still reads descending source indices while writing ascending destinations. Decay's
half<=0 guard and extrapolation's int32 N+overlap<=0 guard precede all source views,
including nil with large unused geometry. The per-sample wrap/subtract, pitch
attenuation updates, live excitation/history aliases and energy load AFTER each
history store retain order. Decay E1/E2 accumulation, MIN32-style NaN selection and
float32 products/Sqrt narrowing remain; no added snapshots, scratch or copies.
The public codec APIs, legacy pointer boundary extents and field layouts stay.

Grouped fixtures retain their old scalar/manual references and add all-bit span
comparisons, exact guarded storage and source immutability. Forward copy overlap,
reverse LPC-memory overlap and history/excitation extrapolation overlap explicitly
check sequential live reads/writes, not copied inputs. Nil zero-work geometry,
recent/older NaNs, pitch lengths, decay bounds and GC/stack growth remain. Existing
PLC/native packet/sample/range/hash goldens and tolerances are untouched. No weakened
fixtures or opaque-layout/EOF/raw callback/global GC/macOS safety claims.

Each separately committed round passes full amd64/386, ARM64/QEMU, scoped PLC and
broad relevant checkptr, typed PartA/PartB on3 architectures, native comparisons,
GC stress, unchanged codec references and diff checks. Final ARM ordinary/scoped
fixtures repeat twenty times. Logs:
/tmp/opus-round-plc-excitation-history-go-spans.log,
/tmp/opus-round-plc-excitation-decay-go-span.log,
/tmp/opus-round-plc-extrapolate-go-spans.log,
/tmp/opus-round-plc-lpc-history-go-span{,-focused,-retry}.log and
/tmp/opus-round-plc-buffer-spans-arm-{checkptr,golden}-repeat.log.
The first full LPC-history log retains the recurring native comparison fatal
"found pointer to free object"; the COMPLETE validation retry passes. This GC
failure is not claimed repaired by these scalar-span interface changes.

Four comb-source rounds add combFilterSource with real *float32 legacy input,
actual []float32 owner, numeric int32 origin and an explicit owned-path flag; then
migrate prefilter, normal postfilter and PLC postfilter source consumption. Owned
paths read copy/history samples through Go slices, not an input scalar pointer.
Source/destination address comparison is used only for the original copy/skip
identity rules. The compatibility path keeps explicit x for both zero-gain copy
and identity EVEN if a separate history owner was supplied, preserving legacy
mixed-argument behavior. Owned/source history is never inferred by pointer math.

Normal/PLC postfilters choose owned source only for an exact readable owner/index
match with the existing output address. Nil, displaced-end and nonmatching history
keep the old private path; this preserves zero-gain in-place skip and opaque or
foreign caller layouts. Target pointers remain because they are writable outputs;
prefilter's TDAC target is unchanged. Source-history slices retain old int32 extent
arithmetic; direct owned copy root additions widen. Required signed/backward and
unused identity boundaries remain explicit. Counts stay110 unsafe.Pointer/
105 uintptr: ACTIVE source carriers, not cast deletion, pointer-map creation,
lifetime proof or performance improvement (extra matching is unbenchmarked).

Upstream celt.c zero gains skip period/tapset/window/history work; period clamps,
tap products, equal-filter overlap suppression, live delay reloads, x4→x1 shifts,
g1==0 suffix-copy/skip and constant-filter order remain. Go in-place/overlapping
source/output/window aliases and state reload/store order are unchanged. New
grouped tests extend gain/overlap/length cases with the concrete source path, plus
overlapping zero-gain copy, nil zero identity, explicit-x legacy copy despite a
foreign owner, exact owner matching and normal-postfilter foreign-history skip.
All prior prefilter/normal/PLC native frame/sample/range/hash goldens and guards
remain; no snapshots/copies beyond the existing required copy, scratch additions,
changed tolerances or caller-image/EOF/raw callback/global GC/macOS claims.

Each separately committed round passes full amd64/386, ARM64/QEMU, scoped comb/
prefilter/postfilter and broad relevant checkptr, typed PartA/PartB on3 architectures,
native comparisons, GC stress, unchanged codec references and diff checks. Final
ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-comb-concrete-source-view.log,
/tmp/opus-round-prefilter-owned-comb-source.log,
/tmp/opus-round-postfilter-owned-comb-source.log,
/tmp/opus-round-plc-postfilter-owned-comb-source{,-focused}.log and
/tmp/opus-round-comb-owned-sources-arm-{checkptr,golden}-repeat.log.

Four matching CELT FIR/PLC rounds separate unchanged FIR arithmetic into a live
span kernel, pass PLC excitation as its real Go owner + numeric source origin,
carry LPC coefficient origin numerically, then pass FIR scratch as an actual Go
output slice rather than a scalar pointer. Public celt_fir and existing explicit
history signatures keep their assertions, reverse coefficient preparation and
signed backward-history fallback. The extraction round alone is arithmetic
organization; the following rounds change actual private PLC owner/span carriers.
Counts remain110 unsafe.Pointer/105 uintptr, not cast deletion or new pointer maps.

The owned entry uses source identity only for the original x!=y assertion (and
retains an explicit unused-source/end compatibility boundary). Filtering reads the
Go history span directly. Original source, coefficient and history argument bounds
checks remain at the PLC call, before the callee; coefficient reversal still
precedes input/output views. Complete four-lane RHS/correlation sums precede output
stores; scalar tails, float32 products and sequential overlap reads/writes are
unchanged. Source/root-relative span additions widen where needed; PLC's bounded
source index/channel*order and all DSP loop arithmetic retain their original types.
No new signal snapshot/copy, history padding or public ABI/tolerance changes.

Grouped FIR fixtures retain old public/explicit-history cases and add span/owned
comparisons for orders0/4/8, N0/1/4/12, output guards, numeric coefficient roots and
GC/stack growth. Existing unsupported ord0 unrolled assertion remains. Added partial
source/output overlap tests compare all bits and surrounding storage to the same
public implementation, preserving its live Go semantics rather than claiming new
upstream restrict guarantees. PLC excitation/FIR packet/range/float/hash and native
references remain unchanged; no fixture weakening, pointer-slot poisoning,
opaque-image/EOF/raw callback/global GC/macOS claim. Unused end and external
negative-history boundaries intentionally stay explicit.

Each separately committed round passes full amd64/386, ARM64/QEMU, scoped FIR/LPC/
IIR and broad relevant checkptr, typed PartA/PartB on3 architectures, native
comparisons, GC stress, unchanged codec references and diff checks. Final ARM
ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-fir-live-span-kernel.log,
/tmp/opus-round-plc-fir-owned-excitation-origin.log,
/tmp/opus-round-plc-fir-numeric-coefficient-origin.log,
/tmp/opus-round-plc-fir-go-output-owner{,-focused}.log and
/tmp/opus-round-plc-fir-owned-spans-arm-{checkptr,golden}-repeat.log.

Four CELT IIR owned-history/origin rounds add an optional concrete output/history
slice, migrate PLC's IIR output to its actual decode-buffer owner + numeric origin,
then defer PLC input and LPC coefficient origins numerically until readable Go
views are needed. The public celt_iir entry delegates with zero origins/no output
owner, retaining the original external signed-history fallback. Owned N<ord tails
read the real preceding samples through the Go owner, not by subtracting from an
output scalar pointer. The active PLC call no longer forms input/output/coefficient
sub-pointers. The original coefficient index bounds check remains before the call.
This is private concrete-owner/span handling, not a new allocation or pointer map
for unbound caller byte images. Counts remain110 unsafe.Pointer/105 uintptr.

ord&3 assertion still precedes views; input, coefficient, output and memory view
order stays. Span root-relative additions widen, while N+ord, N-ord, channel*order,
loops and all DSP int32 operations retain their old semantics. Upstream celt_lpc.c
reverse-coefficient snapshot, initial negative memory, explicit four-lane products,
live coefficient patch reloads, sequential stores, positive scalar-tail history
and reversed output→memory commit persist. No full input/output/coefficient snapshot,
new scratch/copy, ABI field extent or encode/decode tolerance change.

Grouped IIR fixtures cover orders0/4/8 and N0/1/3/4/5/8/17, separate/in-place inputs,
exact guarded output/history/memory owners, numeric input/coefficient origins and
bitwise/full-memory equivalence to the same public entry across GC/stack growth.
The pre-existing unsupported ord0 four-lane assertion remains. Added coefficient/
input/output alias cases also share memory with the same scalar owner to verify
live patch and final-store semantics. Earlier signed-history/nil/order/native
fixtures and PLC/frame goldens remain unchanged. No fixture weakening, pointer-slot
poisoning, caller-image/EOF/raw callback/global GC/macOS safety claim.

Every separately committed round passes full amd64/386, ARM64/QEMU, scoped IIR/PLC
and broad relevant checkptr, typed PartA/PartB on3 architectures, native comparisons,
GC stress, unchanged codec references and diff checks. Final ARM ordinary/scoped
fixtures repeat twenty times. Logs:
/tmp/opus-round-iir-owned-output-history.log,
/tmp/opus-round-plc-iir-owned-output.log,
/tmp/opus-round-plc-iir-numeric-input-origin.log,
/tmp/opus-round-plc-iir-numeric-coefficient-origin{,-focused}.log and
/tmp/opus-round-plc-iir-owned-origins-arm-{checkptr,golden}-repeat.log.

Four active deemphasis-output-origin rounds migrate direct accumulate, direct
replace, downsample accumulate and downsample replace stores to the actual PCM
root + numeric channel/sample coordinates. The last round removes the now-unused
per-channel y pointer and its early cursor formation; zero N or Nd naturally
never invokes the output accessor. No zero/end pointer is fabricated. Channel and
sample coordinates stay separate until the consumed scalar is needed: j*C retains
int32 multiply/wrap, then channel addition is widened to match the original two
pointer additions rather than introducing another int32 wrap. Zero offset keeps
identity; positive stores use a minimal Go prefix; signed history retains the
explicit compatibility address boundary. Counts remain110 unsafe.Pointer/
105 uintptr: ACTIVE consumer-origin changes, not cast removal or new GC maps.

Upstream celt_decoder.c direct accumulate's x+m+VERY_SMALL order differs from
replace/downsample's x+VERY_SMALL+m, and those orders stay. Per-sample preemphasis
state update, per-channel mem commit BEFORE downsample stores, scratch reuse,
sequential live accumulate loads/stores and do-while C0/zero-stride behavior are
unchanged. The common stereo early return, nil zero-work behavior, coefficient
load/negative scratch/division errors, all input/memory/output aliases and sample
rounding remain; no additional copies/snapshots or public API/tolerance changes.

Grouped deemphasis tests add exact rooted output identity, last-readable/signed/
zero-channel cancellation, write-through guards and GC/stack growth. Existing N0/
Nd0/nil, all output branches/downsample factors/channel counts, scalar reference
bits, alias, chunking and zero-work coefficient fixtures persist unchanged. No
caller-byte layout/EOF/raw callback/global GC/macOS safety claim.

Every separately committed round passes full amd64/386, ARM64/QEMU, relevant
scoped checkptr, typed PartA/PartB on3 architectures, native comparisons, GC stress,
unchanged codec references and diff checks. Final ARM ordinary/scoped fixtures
repeat twenty times. Logs:
/tmp/opus-round-deemphasis-direct-accum-numeric-output.log,
/tmp/opus-round-deemphasis-direct-replace-numeric-output.log,
/tmp/opus-round-deemphasis-downsample-accum-numeric-output.log,
/tmp/opus-round-deemphasis-downsample-replace-numeric-output.log and
/tmp/opus-round-deemphasis-numeric-origins-arm-{checkptr,golden}-repeat.log.

Four active CELT pulse-cache row rounds add modePulseRow{bits *byte,offset int16},
then migrate quant_partition split-threshold reads, initial bits→pulses/cost helpers
and EACH remaining-budget retry to that concrete root + numeric coordinate.
modePulseCacheRow captures Fbits before the signed index load, matching the old
pointer constructor, but positive rows do not form a displaced cache pointer.
Only the current byte's minimal root prefix is formed. Negative row offsets keep
the explicit signed-history/unused-boundary compatibility constructor. Legacy
pointer/cache helpers remain for fixtures; factoring their boundary is not cast
removal. Counts stay110 unsafe.Pointer/105 uintptr: ACTIVE positive row consumers,
not hidden pointer integers or new pointer maps/lifetime guarantees.

The scalar rate.h bit search retains all int32 operations, bits--, low=-1, midpoint
loop count and <= tie order; zero pulses still costs0 without byte access. bands.c
split short-circuit order and lookup index arithmetic stay, including LM−1 cache
selection before the condition. The original Go per-helper reload points persist:
threshold, bits search, initial price and EVERY retry after refund/decrement.
Mode pointer/band cached at entry remain as before; no whole-mode/index/table
snapshot or source copy is introduced. Existing live source/entropy/budget aliases,
remaining-bit wrap/store order, pulse conversion, FFT/MDCT/golden bits remain.

Grouped mode-pulse tests cover every valid static cache row/pulse, bits−2..260,
nil zero pulse, signed backwards rows, exact consumed bytes, GC/stack growth and
live data mutation. New reload fixtures independently change signed row index,
byte-root field and earlier row contents to check metadata capture versus live
reads. Native rate/partition/quant-all comparisons and existing scalar C packet/
mask/seed/bit/FNV goldens remain. No buffer enlargement, pointer-slot poisoning,
new audio scratch or tolerance changes; no opaque-image/EOF/raw callback/global
GC/macOS claim. Negative/unused cache boundaries intentionally remain unresolved.

Each separately committed round passes full amd64/386, ARM64/QEMU, relevant scoped
checkptr (mode pulse, partition/quant-all, SILK rows), typed PartA/PartB on3
architectures, native comparisons, GC stress, unchanged codec references and diff
checks. Final ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-celt-pulse-rooted-row-view.log,
/tmp/opus-round-partition-rooted-split-threshold.log,
/tmp/opus-round-partition-rooted-initial-pulse-cost.log,
/tmp/opus-round-partition-rooted-retry-pulse-cost.log and
/tmp/opus-round-celt-rooted-pulse-rows-arm-{checkptr,golden}-repeat.log.

Four active SILK pulse/shell numeric-row rounds select rate-level ICDFs from the
original2x9 byte table root + numeric row, pulse-count ICDFs from the10x18 table
root + numeric row, escape rows from that same root + numeric start, and shell
split rows from the actual shell slice root + numeric table offset. The decoder
no longer forms those selected row pointers before entropy decoding. Row/start
and shell offset/slice bounds checks remain BEFORE entropy or output mutation;
only live symbol entries are consumed by ecDecICDFAt. No table snapshots/copies,
extra PCM scratch or public codec APIs change; counts remain110 unsafe.Pointer/
105 uintptr. This is ACTIVE row-origin handling, not cast deletion or a new map
for opaque images; signed/narrowing/range normalization semantics persist.

Upstream decode_pulses.c and shell_coder.c order remains: rate decode precedes
frame/block validation, sums precede shell decoding, shift10 suppresses another
escape via start1, LSBs/signs follow, and split child1 is committed before reloading
it for child2 (even when outputs alias). The helpers retain original fixed-array
row and slice bounds validation; zero/negative split totals skip table/entropy
and write both zeros. No changed codec packets/samples/ranges/tolerances.

Grouped pulse fixtures compare all2 rate rows and10 count rows (48 encoded symbols
and EVERY context snapshot per row) against the same old pointer-based reference.
Fixture symbol enumeration uses ONLY positive-probability ICDF entries, preserving
the earlier rate-row sequence. Both escape starts are covered; a crafted10-shift
pulse stream exercises the forbidden-eleventh-escape boundary,160 zero LSBs, exact
zero output/guards and full entropy equality against the original row decoder.
Shell fixtures cover all4 tables/totals1..16, complete entropy equality, sequential
and aliased child stores, guards and invalid table/offset rejection before mutation.
Existing120-sample padded shell, pulse scratch, real SILK/frame/native goldens and
GC/stack-growth coverage persist. No pointer poisoning or weakened fixtures; no
EOF/caller-image/raw callback/global GC/macOS claim.

Each separately committed round passes full amd64/386, ARM64/QEMU, relevant scoped
checkptr (including pulse/shell/ICDF/index cases), typed PartA/PartB on3 architectures,
native comparisons, GC stress, unchanged codec references and diff checks. Final
ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-pulse-rate-numeric-row.log,
/tmp/opus-round-pulse-count-numeric-row.log,
/tmp/opus-round-pulse-escape-numeric-row{,-focused}.log,
/tmp/opus-round-shell-split-numeric-row{,-focused}.log and
/tmp/opus-round-pulse-shell-numeric-rows-arm-{checkptr,golden}-repeat.log.

Four active byte-ICDF numeric-row rounds add private ecDecICDFAt, move SILK NLSF
first-stage and residual selection from displaced byte pointers to original table
roots + numeric offsets, then move SILK gain selection to the global24-byte table
root + numeric8-byte row origin. Public ICDF signatures still delegate at offset0.
Negative caller-owned row origins use the existing signed table boundary before
entropy snapshots; zero/positive rows never create a guessed full-row cursor.
The kernel builds only the prefix through the currently consumed symbol entry.
Offsets are widened when combined with symbol position; source row arithmetic,
SILK signed narrowing and gain<<3 wrapping remain unchanged.

Upstream entdec.c snapshots rng/val before scanning, uses IMUL32 per live table
entry, commits val/rng then normalizes. The same sequence remains, including
source tables aliasing entropy values. Upstream decode_indices.c CB1 row product,
residual ec_ix order/EXT branches, fields/codebook reloads and gain decode order
remain. Gain rows explicitly retain the original Go array bound check BEFORE
entropy access; invalid rows panic without touching the context. No full table
snapshot, padded suffix, extra audio scratch/copy or public codec API change.
Counts stay110 unsafe.Pointer/105 uintptr: these are ACTIVE numeric row consumers,
not cast relocation or repair of opaque pointer maps/EOF/raw callbacks/global GC.

Grouped ICDF tests compare200 encoded symbols and EVERY full context snapshot for
positive numeric and signed-history rows against the same public decoder; exact
last-readable table entries, source guards and GC/stack growth persist. Gain
fixtures cover all3 rows and24 symbols each, complete context equality, plus
negative/out-of-range row rejection before entropy mutation. Existing codebook
weak roots, table/entropy aliases, source guards, real SILK/frame references and
native/golden bits remain; no fixture/golden/tolerance changes or macOS claim.

Each separately committed round passes full amd64/386, ARM64/QEMU, relevant scoped
checkptr including numeric ICDF/gain/index cases, three-architecture typed frame
PartA/PartB, native comparison, GC stress, unchanged codec references and diff checks.
Final ARM ordinary/scoped fixtures repeat twenty times. Logs:
/tmp/opus-round-byte-icdf-numeric-origin.log,
/tmp/opus-round-nlsf-first-icdf-numeric-row.log,
/tmp/opus-round-nlsf-residual-icdf-numeric-row.log,
/tmp/opus-round-gain-icdf-numeric-row{,-focused}.log and
/tmp/opus-round-icdf-numeric-rows-arm-{checkptr,golden}-repeat.log.

Four remaining synthesis-scratch-origin rounds add private denormalization output
origins, pass downmix's outputs[0]+overlap/2 as a concrete root+numeric origin,
copy mono→stereo scratch through a rooted output-offset view, then keep its left
MDCT source as outputs[1] + NUMERIC overlap/2/block coordinates. No synthesis
scratch pointer is formed by these active callers. Existing public and private
reference signatures delegate at zero offset; compatibility cursor helpers stay
for boundary fixtures. Denormalized/copy views use the actual N samples; MDCT
views wait for strided kernel geometry. No extra audio storage/snapshot is added.

The combined scratch/block INPUT origin is int64: the original two pointer
additions must not become a new int32(overlap/2+b) wrap. Individual overlap/2,
b, NB*b, MDCT strided products/width and output extents retain int32 semantics.
Signed history/unused/invalid fallback boundaries persist; displaced zero spans
form no readable pointer. Silent input/energy skips, narrowing/rounding, live
spectrum/energy/output/window aliases and copy-left-before-right/store order stay.
All relevant sample/frame/range/bits/payload/hash expectations remain unchanged.
Counts110 unsafe.Pointer/105 uintptr are unchanged: ACTIVE root-relative scratch
access, not cast centralization, new caller-image pointer maps or lifetime proof.

Grouped denormalization tests now compare numeric OUTPUT origins in live aliased
spectrum/energy storage against the same old public results, plus nil all-displaced
empty silence. Existing cursor tests add rooted overlapping-copy equivalence,
outer guards and nil empty copy with large displacement. MDCT's all-shift/stride
fixtures use the widened private input coordinate without changing buffers or
expected results; existing synthesis transient/mono/stereo/downmix/CC0/silent and
exact terminal aliases persist. No EOF/opaque-image/raw callback/global GC/macOS
claim. Public Opus APIs and native tolerances are unchanged.

Each separately committed round passes full amd64/386, ARM64/QEMU, relevant scoped
checkptr, native comparisons, GC stress, unchanged codec references and diff checks.
Final typed PartA/PartB scoped tests pass on3 architectures; ARM ordinary/scoped
checks repeat twenty times. Logs:
/tmp/opus-round-denormalise-numeric-output-origin.log,
/tmp/opus-round-downmix-rooted-denormalise-output.log,
/tmp/opus-round-synthesis-copy-numeric-output.log,
/tmp/opus-round-mono-stereo-rooted-scratch-origin.log,
/tmp/opus-round-synthesis-scratch-origins-typed-frames.log and
/tmp/opus-round-synthesis-scratch-origins-arm-{checkptr,golden}-repeat.log.

Four matching active numeric MDCT-OUTPUT-origin rounds add a private
opusMDCTBackwardOffsets entry, then move mono→stereo, downmix and general-channel
synthesis outputs from displaced block pointers to the original output root plus
numeric NB*b origin. Public MDCT and existing private input-offset/zero-offset
entries delegate unchanged at outputOffset0; fixture helper signatures persist.
Input/output origins remain independent, so live aliased storage stays live rather
than being copied/snapshotted. Within the kernel the output extent is exactly
N2+overlap/2, not the full synthesis frame N: prefix/subslice formation uses that
known readable span. Zero offset preserves the original direct view; nonpositive
extents explicitly keep the unused/invalid address boundary, and signed history
keeps its existing compatibility arithmetic. NB*b and extent remain int32 before
root-relative widened addition. Pre-rotation store order, FFT real/imag lanes,
post-rotation paired reloads/stores and TDAC overlaps are untouched.

Existing all4-shift,−2/0/1/2 stride bit fixtures now also compare numeric positive
output origins and signed negative origins to the same public-call results,
including outer guards and GC/stack growth. All existing source/output/window
alias, exact synthesis terminal extent, CC0, transient and silent fixtures remain.
No pointer slot poisoning, enlarged buffers, new audio scratch/copies or changed
expected counts/bits/ranges/native tolerances. Counts stay110 unsafe.Pointer/
105 uintptr: ACTIVE output span formation replaces early pointer subviews without
moving or hiding casts; no opaque caller-image pointer-map/lifetime/macOS claim.

Every separately committed round passes full amd64/386, ARM64/QEMU, relevant scoped
checkptr, native comparisons, GC stress, unchanged codec references and diff
checks. Final typed PartA/PartB scoped tests on3 architectures and ARM ordinary/
scoped twenty-repeat checks pass. Logs:
/tmp/opus-round-mdct-numeric-output-origin.log,
/tmp/opus-round-mono-stereo-mdct-numeric-outputs.log,
/tmp/opus-round-downmix-mdct-numeric-outputs.log,
/tmp/opus-round-general-mdct-numeric-outputs.log,
/tmp/opus-round-mdct-numeric-outputs-typed-frames.log and
/tmp/opus-round-mdct-numeric-outputs-arm-{checkptr,golden}-repeat.log.

Four active numeric MDCT-input-origin rounds add mdctStridedSliceAt and a private
opusMDCTBackwardAt entry, then move mono→stereo copy transforms, downmix transforms
and general per-channel synthesis transforms to real float roots + numeric block
origins. Public MDCT signatures delegate at offset0; the existing private zero-
offset helper remains for reference fixtures. No frequency block pointer is formed
by synthesis on these paths. Only the MDCT kernel forms the actual strided input
view, after lookup geometry establishes the consumed endpoints. Synthesis output
cursors/history, copy-left-before-right transform order, silence, CC0 do-while,
mode/bitrev/twiddle/window state and denormalization offsets remain unchanged.

The strided endpoint product and width retain original int32 wrap/narrowing; only
root-relative addition is widened. Positive readable origins use a minimal Go
prefix and subslice, including a negative stride whose history is within the Go
root. Offset0 identity, zero stride, signed backward history and unused/invalid
extent boundaries remain. No audio snapshot/copy or FFT/MDCT rounding/store-order
changes occur; live input/output/window aliases continue through the same kernel.
Counts remain110 unsafe.Pointer/105 uintptr: actual positive input origin handling
changes, not lexical cast deletion or new maps for opaque caller images.

Grouped MDCT fixtures compare ALL4 shifts at strides−2/0/1/2 bitwise against the
unchanged public offset0 call using exact final-readable owners, seeded output,
outer guards and GC/stack growth. Existing native MDCT and synthesis source/output
alias, transient channel geometry, mode/header/twiddle aliases and golden bits
remain. No arbitrary external byte-image/EOF/raw callback/global GC/macOS claim.

Every round passes full amd64/386, ARM64/QEMU, relevant scoped checkptr, native
comparisons, GC stress, unchanged codec references and diff checks; final typed
PartA/PartB scoped tests on3 architectures and ARM ordinary/scoped twenty-repeat
checks pass. Retained /tmp/opus-round-mono-stereo-mdct-numeric-blocks.log records a
compile failure from changing a private helper signature used by existing alias
fixtures. Preserve that helper as an offset0 delegate and add the At helper;
full ...-retry.log passes without weakening fixtures. Other logs:
/tmp/opus-round-mdct-numeric-input-origin.log,
/tmp/opus-round-downmix-mdct-numeric-blocks.log,
/tmp/opus-round-general-mdct-numeric-blocks.log,
/tmp/opus-round-mdct-numeric-inputs-typed-frames.log and
/tmp/opus-round-mdct-numeric-inputs-arm-{checkptr,golden}-repeat.log.

Four active CELT readable-span rounds form positive displaced denormalization
views directly as Go prefixes of offset+length (widened numeric addition), retain
contiguous synthesis copy/downmix scratch prefixes, then use live Go TDAC window
views and exact destination(overlap/2)/filtered(overlap) views. No copy, snapshot
of audio/window contents or new ownership map is introduced. Silent and displaced
empty denormalization returns remain before views; zero-offset/negative-history
and negative-length compatibility remain. TDAC skips all views for overlap/2<=0,
including nil mode/data at overlap0/1; per-iteration live aliases and original
multiply/add order, float32 rounding and sequential writes remain. Upstream
celt_decoder.c TDAC loop consumes exactly those spans. Public signatures, factory
layouts, raw adapters, source aliasing and codec goldens/tolerances are unchanged.
Counts remain110 unsafe.Pointer/105 uintptr: these are ACTIVE positive span/data
access changes, not deletions, cast hiding or new caller-byte GC pointer maps.

Existing grouped denormalization zero/silent/signed/live-alias, synthesis transient
mono/stereo/downmix and exact guarded scratch, and TDAC odd/even/nil/rounding/live
window/filtered aliases persist. Contiguous cursor tests add exact final-readable
spans, signed/zero/unused-end fallbacks, GC/stack growth and write-through identity.

Retained /tmp/opus-round-synthesis-readable-consumed-spans.log is a FAILED scoped
checkptr attempt: treating every celtSynthesisCursor's consumed argument as a full
contiguous length overstates strided IMDCT frequency/output tails. No fixture or
buffer extent was enlarged. The original minimal-prefix generic cursor remains;
only proven contiguous copy/downmix scratch uses the new full-prefix helper.
Full ...-retry.log passes. This distinction is a consumer contract, not lexical
pointer safety. No opaque-image/EOF/raw callback/global GC/macOS claim.

Each separately committed round passes full amd64/386, ARM64/QEMU, relevant scoped
checkptr, native comparisons, GC stress, unchanged encode/decode references and
diff checks. Final three-architecture typed PartA/PartB scoped tests and ARM
ordinary/scoped twenty-repeat checks also pass. Other logs:
/tmp/opus-round-denormalise-direct-rooted-span.log,
/tmp/opus-round-tdac-readable-window-view.log,
/tmp/opus-round-tdac-readable-sample-views.log,
/tmp/opus-round-celt-readable-spans-typed-frames.log and
/tmp/opus-round-celt-readable-spans-arm-{checkptr,golden}-repeat.log.

Four SILK/quantization typed-entry/reference-owner rounds convert Opus_silk_Decode
to *silk_decoder, *silk_DecControlStruct, *ec_dec, *float32 PCM and *int32 count,
then give all3 SILK mono/stereo packet/PLC reference fixtures actual scanned Go
decoder owners. The reference helper checks reported ABI size against the Go
state extent, initializes the same zeroed image and retains the existing scanned
pseudostack producer. The packet runner keeps control/entropy/PCM/count roots
concrete through GC/stack growth before EACH frame. Same packets, sample counts,
PCM bits/hashes, pitch, range, PLC state and all historical goldens remain.

Opus_quant_all_bands now takes concrete mode, float spectra/energy, byte masks,
int32 pulse/TF resolution, ec_ctx and uint32 seed pointers. The 2 scalar C-reference
encode/decode calls use real pointers; optional mono Y remains typed nil. The
following round replaces ALL xmallocArray byte-image fixtures for the fake mode,
its tables, entropy contexts and arrays with actual make([]T) Go storage. The
obsolete integer-laundering helper/comment and redundant byte clears disappear;
make supplies the same zero contents and array extents with genuine pointer maps.
No scratch/different buffer copy changes RDO/live source aliases or entropy state.
Same no-intrinsics upstream C cases,4 encode payloads, masks, seeds, bit budgets,
X/Y FNVs and decode resynthesis remain. Existing pseudostack setup/reset/zero
scratch assertions persist; Go-owned modes/entropy buffers get GC/stack growth
before recursive quantization. The original encode/decode loops are unchanged.

Counts124→110 unsafe.Pointer/116→105 uintptr:5 SILK and9 quantizer crossing casts
and11 uintptr parameter spellings are removed, not relocated. These generated
C-internal leaf signatures change Go source API; callers of their old integer
forms must pass real pointers. Public Opus codec APIs, necessary adapters, raw
custom callbacks and independent opusccenc entries remain untouched. Typed entries
were already the active private DSP; owner repairs here are specifically the
LEGACY REFERENCE fixture producers, not new production lifetime fixes or global
opaque-allocation/EOF/GC safety. No encode/decode tolerance changes or macOS claim.

Each separately committed round passes full amd64/386, ARM64/QEMU, broad/scoped
checkptr, three-architecture typed PartA/PartB, native comparison, GC stress,
unchanged codec references and diff checks. SILK references gain scoped checkptr
after their scanned-owner round; quant-all scalar references gain it after their
Go-owner round. Final ARM ordinary/scoped fixtures repeat twenty times. Retained
/tmp/opus-round-silk-concrete-decode-entry.log is a compile failure exposing an
old int16 lost-frame output passed to the float API; change320 int16 to160 float32,
PRESERVING the exact640-byte extent and all expectations, then the full ...-retry.log
passes. No golden/fixture was weakened. Other logs:
/tmp/opus-round-silk-scanned-reference-decoders.log,
/tmp/opus-round-quant-all-concrete-entry.log,
/tmp/opus-round-quant-scanned-reference-owners{,-focused}.log and
/tmp/opus-round-concrete-silk-quant-entries-arm-{checkptr,golden}-repeat.log.

Four generated leaf-entry rounds delete the unreferenced Opus_celt_decode_with_ec
integer forwarding alias, then its unreferenced Opus_celt_decode_with_ec_dred
integer adapter, and convert VAD plus delayed-NLSF quantizer leaf entry signatures
to real Go pointers. Standard CELT frame/audio/native comparison paths already
call the concrete private decoder, so the two deletions are DEAD adapter cleanup,
not lifetime repairs. Repository-wide searches find no decoder/tools references;
same-named opusccenc functions belong to the independent encoder package and stay.

Opus_silk_VAD_GetSA_Q8_c now receives *silk_encoder_state and *int16 input;
Opus_silk_NLSF_del_dec_quant receives *int8 indices, int16 sample/weight/index
pointers and byte predictor/rate pointers. Both delegate unchanged to the existing
concrete kernels; removed uintptrescapes annotations/integer→pointer conversions
are not relocated. VAD Go/reference-native fixture calls and the delayed quantizer
fixture now pass actual pointers. Existing scratch/state aliasing, rounding/
wrapping, signed narrowing, search/tie order, error/validation conditions and
pseudostack non-consumption semantics remain. VAD native whole-state/fields and
quantizer RD15284434/indices are unchanged; delayed fixture adds GC/stack growth
before the same golden call. No encoder-state embedded legacy-pointer bitmap or
unrelated table backward-boundary safety is claimed repaired.

These generated C-INTERNAL Go leaf exports are not the public Opus codec API:
external Go users of the removed CELT leaf wrappers or old uintptr VAD/NLSF
signatures must adapt; this is a source-level leaf API cleanup. Public Opus
create/init/decode/CTL/destroy, required used adapters and custom callback ABI
remain untouched. Counts136→124 unsafe.Pointer/125→116 uintptr:4 cast deletions
are unused CELT cleanup and8 are the concrete leaf parameter crossings;9 uintptr
parameter spellings disappear. No opaque aliases, raw casts hidden in helpers,
full-buffer copies or changed codec goldens/tolerances.

Every separately committed round passes full amd64/386, ARM64/QEMU, broad/scoped
checkptr, typed PartA/PartB on all three architectures, native comparisons, GC
stress, unchanged codec references and diff checks. VAD/NLSF typed entries also
pass their specific scoped checkptr runs; final ARM ordinary/scoped fixtures
repeat twenty times. No global opaque-image/extension EOF/GC or macOS claim.
Logs: /tmp/opus-round-remove-unused-celt-decode-alias.log,
/tmp/opus-round-remove-unused-celt-dred-adapter.log,
/tmp/opus-round-vad-concrete-leaf-entry.log,
/tmp/opus-round-nlsf-concrete-leaf-entry{,-focused}.log and
/tmp/opus-round-concrete-celt-vad-nlsf-leaves-arm-{checkptr,golden}-repeat.log.

Four ACTIVE canonical accessor rounds select MS decode children, projection's
embedded MS header, demixing matrix header, then coefficient cells (and MS typed
CTL child consumers) from actual registered Go views. opusMSStorageViews.Offsets
contains numeric reflection-field/array stride displacements, not pointer
subtraction. MS decode/CTL selection requires the exact parent root, original
live stream/coupled geometry and an exact consumed offset match; otherwise it
uses the unchanged explicit byte-layout accessor. CTL offset increments and live
header reloads, alias outputs, child order and validation/error order remain.

Projection storage records the actual Go Multistream field offset. Its accessor
compares the original live signed/narrowed matrix-size alignment computation to
that numeric offset: mutated layouts still use the original boundary. Demixing
matrix header and coefficient location are FIXED fields independent of mutable
rows/cols/matrix-size values; registered views return the actual matrix header
and int16 cells. Matrix registered aliases join the projection base allocation
group, so typed/raw/raw-allocator free or TLS.Close clears matrix, MS and codec
bindings together, including prior interior MS cleanup. External matrix/header
images, nil/foreign TLS and unbound storage keep the original ABI accessor casts.
No audio/control/matrix samples, modes or entropy are snapshotted or copied.

Grouped fixtures compare bound/nil-TLS child and matrix/coefficient identities,
MS GET_DECODER_STATE returned children, live coupled geometry and projection
matrix-size mutations forcing offset fallback, matrix fixed-field identity after
header mutations, GC/stack growth and matrix alias cleanup on every destructor
path. Native MS/projection matrix/source aliases, caller buffer geometry, child
strides/padding, reinitialization/errors and codec goldens/tolerances persist.

Each round passes full amd64/386, ARM64/QEMU, scoped storage/pointer checkptr,
three-architecture typed PartA/PartB references, native comparison, GC stress,
unchanged encode/decode references and diff checks; final ARM ordinary/scoped
fixtures repeat twenty times. Counts remain136 unsafe.Pointer/125 uintptr:
these are ACTIVE opaque layout-recovery bypasses for Go factories, not cast
relocation, new pointer maps or global byte-image/EOF/raw callback safety. Registry
metadata lookup/linear child matching costs are not benchmarked; no performance
or direct macOS success is claimed. Logs:
/tmp/opus-round-ms-concrete-child-accessors.log,
/tmp/opus-round-projection-concrete-ms-accessor.log,
/tmp/opus-round-projection-concrete-matrix-header.log,
/tmp/opus-round-projection-concrete-coefficients-focused.log,
/tmp/opus-round-projection-concrete-coefficients.log and
/tmp/opus-round-canonical-decoder-accessors-arm-{checkptr,golden}-repeat.log.

Four registered codec-view rounds add exact-root typed metadata APIs in libcshim,
bind single-decoder factory SILK/CELT views for frame/control use, propagate MS
child views, then propagate projection MS/children with alias-group cleanup.
Normal factory-created decoders with their TLS now obtain SILK/CELT state from
actual canonical reflection fields, not decode-time unsafe.Add layout recovery.
Private opusCodecSilkState/opusCodecCeltState require the exact registered root,
correct metadata type and unchanged corresponding live offset; mutated offsets,
nil/foreign TLS, caller buffers and other unbound images retain the original
layout accessors. Matching views do not snapshot mode, codebook, entropy or audio
state. Public decode/control signatures, initialization/validation/error ordering,
caller-buffer extents, codec arithmetic and ALL goldens/tolerances remain.

Metadata attaches ONLY to a registered actual *T owner of the same type/identity:
byte images and mismatched root/value types cannot be promoted. TLS registry
integer keys (including alias group keys) are allocation identity/ABI transport;
new typed lookup accepts *T and never reconstructs a pointer from those keys.
Metadata contains actual Go views/closures, not opaque pointers. MS/projection
interior aliases are associated with the actual base registry group, never
independent allocations. Typed and raw codec destruction removes bindings;
underlying raw Xfree or TLS.Close also removes all aliases. Removing an interior
alias leaves the group base/siblings intact. Projection cleanup tolerates prior
interior MS destruction; malformed projection allocation/error fallback stays.

Grouped fixtures cover metadata-only weak-owner retention, type mismatch, opaque
byte rejection, nil/idempotent cleanup, alias-group raw free, single codec views
through GC/stack growth, live-offset mutations and nil-TLS fallback, MS/projection
child views and typed/raw/raw-allocator/TLS cleanup. Existing reinitialization,
foreign mode/codebook owners, child stride/padding, matrix aliases, native/error
snapshots and codec goldens persist. The canonical scanned owner allocations
predate these rounds: this is ACTIVE decode/control representation bypass plus
registry lifetime bookkeeping, NOT new GC pointer maps, global opaque-image
safety or cast relocation. Production counts stay136 unsafe.Pointer/125 uintptr;
necessary external-layout/ABI casts stay visible. Registry locking/metadata cost
is not benchmarked; no performance or direct macOS claim is made.

Each round passes full amd64/386, ARM64/QEMU, scoped pointer/storage/registry
checkptr, native comparisons, GC stress, unchanged codec references and diff
checks. The final three codec rounds also check typed PartA/PartB on all three
architectures; final ARM scoped/ordinary fixtures repeat twenty times. Initial
/tmp/opus-round-concrete-registry-view-metadata.log retains recurring native
runtime failure (fatal error: s.allocCount != s.nelems && freeIndex == s.nelems);
complete ...-retry.log passes, not a repair claim. Other logs:
/tmp/opus-round-single-bound-decode-codec-views.log,
/tmp/opus-round-ms-bound-decode-child-codecs.log,
/tmp/opus-round-projection-bound-codecs-focused.log,
/tmp/opus-round-projection-bound-decode-child-codecs.log and
/tmp/opus-round-registered-codec-views-arm-{checkptr,golden}-repeat.log.

Four ACTIVE denormalization/history rounds add celtDenormaliseBandsAt and its
concrete float span selector, migrate stereo→mono channel spectrum/energy offsets,
then general synthesis channel offsets, and finally migrate prefilter's per-channel
history header. Public Opus_denormalise_bands retains its signature and zero-offset
semantics. Synthesis passes actual *float32 bases and numeric int32 offsets instead
of computing X+c*N/oldBandE+c*nbEBands addresses before calling a silent/no-work
consumer. Original multiplication/wrapping, table/size/bound reads and output-span
selection remain. Only AFTER the original silence return are input and energy
spans selected, in their original order, before initial output clearing and live
sample/energy loads. No values are copied/cached: output/spectrum/energy aliases
retain sequential behavior. Nonzero-offset EMPTY spans form no unused cursor;
zero-offset spans retain the original Go slice shape. Signed history and invalid
negative-span compatibility remain explicit boundaries, not guessed readable views.

Prefilter history headers select actual configured channel-first samples (each
channel owns DEC_PITCH_BUF_SIZE samples even when overlap is zero). Zero offset
still preserves pointer identity; the right-channel start is interior, not EOF.
Existing mode/window/history-owner GC/weak retention, zero overlap/CC0 and filter
control fixtures remain. Grouped offset fixtures compare full float bits/guards
against zero-offset calls across start/downsample cases and output overlapping BOTH
spectrum and energy; they check signed history, empty displaced nil spans, silent
nil bases with large numeric offsets and no extra stores. Grouped silent synthesis
fixtures cover upmix/downmix/general/CC0 and transient/nontransient nil spectrum/
energy bases. Native/reference cases and ALL codec goldens/tolerances remain.

All4 rounds pass full amd64/386, ARM64/QEMU, scoped denormalization/prefilter/
synthesis and broad pointer checkptr, native comparisons, GC stress, unchanged
codec references and diff checks. Final validation also checks typed PartA/PartB
on all three architectures, with final ARM scoped/ordinary fixtures repeated20
(times, unchanged scopes/counts). Counts remain136 unsafe.Pointer/125 uintptr:
these are active deferred-offset/readable-view changes; the arithmetic cast
remains ONLY for genuine compatibility branches, not hidden by opaque aliases.
No global exact-allocation EOF, raw-image/GC or direct macOS safety claim follows.
Logs: /tmp/opus-round-denormalise-offset-kernel.log,
/tmp/opus-round-downmix-denormalise-offsets.log,
/tmp/opus-round-general-denormalise-channel-offsets.log,
/tmp/opus-round-prefilter-history-consumed-header.log and
/tmp/opus-round-denormalise-history-offset-arm-{checkptr,golden}-repeat.log.

Four ACTIVE prefilter/synthesis cursor rounds migrate TDAC memory/window/filtered
sample accesses, prefilter overlap input, mono→stereo synthesis cursors, then
stereo→mono and general synthesis block cursors. These use minimal readable Go
prefixes rather than positive unsafe.Add-derived addresses. Every selected TDAC
lane is consumed inside the overlap/2 loop, with its original live window/sample
read order, two separately rounded products and sequential store. Prefilter input
is selected only after the original overlap==0 return; signed preceding-history
cases still use the explicit compatibility boundary and filter history owner.

Synthesis copied/downmixed temporary starts and IMDCT interleaved spectrum/output
block starts use celtSynthesisCursor: positive consumer length selects readable
prefixes; zero/nonpositive consumer length retains the original unused-end
identity. Offset zero remains identity (including nil). Negative offsets retain
caller-owned history semantics. No guessed full suffix, additional scratch copy,
source snapshot, cached controls or changes to channel-zero do-while behavior,
left-before-right transforms, output-as-input aliases, half-product rounding,
window/overlap order, SATURATE identity or entropy/arithmetic are introduced.
Remaining celtNormAdd users (e.g. potentially unconsumed denormalization channel
inputs/energy and history headers) remain honest explicit boundaries. Production
counts stay136 unsafe.Pointer/125 uintptr: these are ACTIVE arithmetic consumer
migrations, not dead adapter removal, cast relocation or global lifetime proofs.

Grouped synthesis cursor fixtures cover nil/zero identity, last readable sample,
negative history, live alias writes and a terminal ABI-image cursor INTERIOR to
an actual guarded Go owner through GC/stack growth. They do not claim exact-Go-
allocation EOF safety. Existing TDAC rounding/window-memory aliases, prefilter
N0/zero overlap/controls/history, synthesis branches/transients/source-output
aliases/mode reloads and native/reference fixtures/goldens/tolerances persist.

Each separately committed round passes full amd64/386, ARM64/QEMU, scoped
prefilter/synthesis and broad pointer checkptr, native comparisons, GC stress,
unchanged encode/decode references and diff checks. Final validation also checks
typed PartA/PartB references, and final ARM scoped/ordinary fixtures repeat twenty
times. No performance/global opaque-image/extension EOF/GC or macOS claim.
Logs: /tmp/opus-round-prefilter-tdac-consumed-samples.log,
/tmp/opus-round-prefilter-consumed-overlap-input.log,
/tmp/opus-round-synthesis-mono-stereo-consumed-cursors.log,
/tmp/opus-round-synthesis-general-downmix-consumed-cursors.log and
/tmp/opus-round-synthesis-consumed-cursors-arm-{checkptr,golden}-repeat.log.
The combined360-second final shell timed out DURING the ordinary ARM repetition
(after full validation and the scoped20-repeat ARM run passed in176.472s).
The interrupted ordinary log is retained; rerun the SAME ordinary scope/count
separately to completion in
/tmp/opus-round-synthesis-consumed-cursors-arm-golden-repeat-retry.log. No scope,
fixtures, golden or tolerances are changed for that time-budget retry.

Four partition/deemphasis/unused-adapter rounds change active positive pointer
advances and remove unreferenced generated C-internal leaf adapters. Round1
indexes the split partition's second spectrum through a minimal readable prefix
and deletes the unused Opus_silk_decode_core integer wrapper; round2 does the
same for the split lowband and unused Opus_silk_PLC wrapper. Round3 migrates
consumed general deemphasis channel starts and removes the unused
Opus_silk_decode_frame wrapper. Round4 migrates all four general deemphasis
output stores (direct/downsample, replace/accumulate) and deletes the unused
Opus_clt_compute_allocation wrapper and its obsolete uintptr-local comment.

Both split pointers are formed only in the existing N>2 branch AFTER halving N;
lowband's nil guard and pointer-formation position after theta/budget updates
remain. No guessed full suffix, unused EOF pointer or scratch copy is introduced.
Deemphasis preserves channel zero/nil identity, no-work/downsample guards,
coefficient/memory reload ordering, float32 recurrence/rounding and sequential
accumulate reads/stores. celtNormConsumedAt indexes readable positive prefixes;
zero and negative caller-owned history explicitly retain celtNormAdd compatibility.
The other synthesis/history/boundary consumers of celtNormAdd are NOT claimed
migrated. New grouped consumed-norm fixtures verify nil/zero identity, last
readable interior cell, negative history, GC/stack growth and live write aliases.
Existing scalar/native partition, spectrum/lowband alias, deemphasis scratch/zero
work/C0/downsample/accumulate fixtures and codec goldens remain unchanged.

All four removed wrappers have no call sites in this repository's opuscc decoder,
its fixtures or tools: active calls already use the concrete private SILK/core/PLC/
frame and allocation kernels. Same-spelled functions in opusccenc are distinct
and untouched. This removes generated C-internal Go exports, NOT public Opus
codec create/init/decode/CTL/destroy APIs or used leaf compatibility adapters.
Counts157→136 unsafe.Pointer and134→125 uintptr. The21 cast deletions are DEAD
adapter removal, not lifetime fixes or relocated casts;8 uintptr declarations
and1 obsolete comment disappear. Active cursor changes reduce address arithmetic
but not lexical unsafe.Pointer counts because necessary backward/EOF boundaries
remain. No performance, global checkptr/opaque-image/extension EOF or macOS claim.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
GC stress and unchanged encode/decode references. Final validation additionally
checks partition/deemphasis paths and typed frame references on amd64/386/ARM;
final ARM ordinary/scoped fixtures repeat twenty times. The initial
/tmp/opus-round-partition-remove-obsolete-address-helper.log fails compilation:
the attempted celtNormAdd removal inspected only partition-local references and
missed active celt_decoder.go synthesis users. Restore the genuine remaining
boundary helper (not just a test helper), migrate only proven consumed sites,
then the full /tmp/opus-round-consumed-deemphasis-channel-prefix-retry.log passes.
No fixture/golden was weakened. Other logs:
/tmp/opus-round-partition-consumed-{spectrum,lowband}-prefix.log,
/tmp/opus-round-deemphasis-consumed-output-prefix.log retains a native comparison
runtime failure (fatal error: sweep increased allocation count); the complete
/tmp/opus-round-deemphasis-consumed-output-prefix-retry.log passes. This recurring
native/GC failure is NOT repaired by a retry or these cursor/adapter changes.
/tmp/opus-round-consumed-norm-deemphasis-arm-{checkptr,golden}-repeat.log retain the
successful final repeats.

Four decoder diagnostic rounds finish migrating PRIVATE static-pool consumers
from integer C-string addresses to GC-visible Go strings and numeric pool offsets:
28 decoder/frame checks;25 MS/matrix/CELT quantization checks;20 extension checks;
then10 remaining CELT/SILK checks and5 mini-FFT assertion sites. The mini-FFT sites
use libcshim.AssertFail(string,string,line); legacy X__assert_fail still converts
assertion before file and emits the identical panic. Opus_celt_fatal remains a
raw public adapter with file-before-message conversion. No static call sites
cross either integer-string adapter anymore.

The inferred pointer-valued uintptr __ccgo_ts and its reflect.StringHeader cast
are DELETED, not renamed/relocated. One real residual consumer was public version
query: it now uses Opus_opus_get_version_string_typed's actual static *byte, with
an explicit integer conversion ONLY in the public adapter. That preserves static
address identity, text and the trailing NUL, without allocation or guessed data.
Because the old address variable's type was inferred, lexical totals are
157 unsafe.Pointer/134 uintptr (previous157/133): the removed StringHeader cast
is offset by the explicit version adapter cast, and uintptr's extra spelling is
that public conversion. This is active integer-address removal, not a claim of
reduced lexical counts or repaired DSP heap ownership. No opaque pointer aliases
are introduced; the original immutable pooled string remains Go text.

All conditions, evaluation/error ordering, diagnostic text/file/line, multiline
messages, abort-vs-assert panic shape and valid-path arithmetic remain. Grouped
fixtures assert exact stderr/abort for decoder/quantization/extension checks,
capacity-before-ID assertions, exact mini-FFT panic strings, Go-text assertion
and ordinary raw ABI compatibility, full pool splitting and static version
identity/text/NUL after GC/stack growth. Native checks/goldens/tolerances persist.
Each separately committed round passes full amd64/386, ARM64/QEMU, scoped checkptr,
native comparison, GC stress and codec references. Final ARM scoped/ordinary
fixtures including typed frame references repeat twenty times. No global raw
string adapter/opaque-image/extension EOF/GC safety or direct macOS claim follows.

Retained failures: /tmp/opus-round-extension-go-diagnostics.log has a NEW fixture
compile error (undefined OPUS_BUFFER_TOO_SMALL); use the unchanged port code-2,
then the complete ...-retry.log passes. The initial
/tmp/opus-round-mini-assert-go-text-focused.log exposed the separate legacy
GoString raw ABI checkptr limitation AND the missed public version pool consumer.
The new raw ABI assertion fixture is an ordinary compatibility test, separate
from the Go-text scoped test; this does NOT fix/suppress legacy GoString or weaken
an existing test/golden. Version now uses the typed static pointer; the complete
focused retry passes. Logs: /tmp/opus-round-decoder-frame-go-diagnostics.log,
/tmp/opus-round-codec-matrix-go-diagnostics.log,
/tmp/opus-round-extension-go-diagnostics-retry.log,
/tmp/opus-round-mini-assert-go-text-focused{,-retry}.log,
/tmp/opus-round-mini-assert-static-version-go-text.log and
/tmp/opus-round-diagnostic-go-text-arm-{checkptr,golden}-repeat.log.

Four active CELT FFT/MDCT output rounds replace direct complex-only butterfly
access with concrete complex/float-lane output views: radix2, radix4, radix3/5,
then the common dispatcher and backward MDCT consumer. Backward MDCT now passes
its actual *float32 output to FFT; the float→complex unsafe.Pointer reinterpret
cast is REMOVED, not moved to a helper or hidden in an opaque alias. The output
descriptor contains concrete *kiss_fft_cpx and *float32 alternatives; consumed
buffer views contain corresponding Go slices. There is no whole-buffer scratch
copy/snapshot that would break live aliases. Counts fall158→157 unsafe.Pointer;
uintptr remains133. The prior rounds were required kernel representation work
before deleting this cast on the normal decoder's active synthesis path.

Upstream celt/kiss_fft.c, mdct.c and complex macros guide lane/store semantics.
No-work/assertion guards precede prefix construction. Each stage uses its original
consumed extent, factor/shift arithmetic and live state-factor reads; twiddle
loads remain live, while epi3/ya/yb retain their original caches. Whole-pair stores
capture both RHS lanes before writing real then imaginary; lane updates keep
sequential reload/store ordering. Radix3/5 explicit float32 product rounding,
ARM fusion safeguards, middle-pair MDCT dependencies and window/overlap order
remain unchanged. Native FFT/MDCT fixtures and decoder goldens/tolerances persist.

Grouped fixtures compare complete complex vs scalar float bits/guards across
all radices, m1/general paths, multiple groups and output/twiddle aliases through
GC/stack growth. They also check radix2's assertion-before-no-work and null
operands on no-work paths, and full FFT scalar/complex bit identity at all four
codec shifts. Test-only twiddle alias casts describe genuine ABI/layout aliases;
no production output cast was displaced into them. Reflection/FFT accessor
costs are not benchmarked or claimed to improve performance.

Every separately committed round passes full amd64/386, ARM64/QEMU, scoped
checkptr, native comparisons, GC decode stress, unchanged encode/decode references
and diff checks. Final validation additionally checks typed PartA/PartB frame
references and FFT/MDCT fixtures on amd64/386/ARM; final ARM ordinary/scoped tests
repeat twenty times. No global opaque-image, extension EOF/GC or raw-callback
safety claim follows. Logs: /tmp/opus-round-fft-concrete-float-radix{2,4,35}.log,
/tmp/opus-round-mdct-float-fft-focused.log,
/tmp/opus-round-mdct-concrete-float-fft-output.log and
/tmp/opus-round-mdct-float-fft-arm-{checkptr,golden}-repeat.log.

Four DRED/frame-reference rounds add concrete aligned Go storage for the
port's disabled12-byte DRED scalar header, register that owner in the factory,
then migrate PartA and PartB frame-reference fixtures to typed create/destroy,
float decode, and (PartA) gain-control entries. DRED scalar fields, magic,
16-byte allocator alignment, nil-TLS allocation error and cleanup remain;
required ABI size stays12 despite possible alignment-array storage. No enabled
DRED model/weights behavior is claimed. The prior opaque malloc→DRED cast is
actually removed from the factory, not relocated. Lexical production counts
remain158 unsafe.Pointer/133 uintptr because that implicit allocator cast did
not spell unsafe.Pointer; most other changes are fixture crossings.

Both frame fixtures retain EVERY packet, step, sample count, float bits, hash,
range, mode, redundancy and CELT PLC golden/tolerance. They now keep decoder,
packet and PCM roots concrete instead of reconstructing a decoder from a raw
handle and passing data/PCM through uintptr. PartA gains use the same int32
values256/0 without legacy varargs staging. New DRED fixtures verify zero/init
scalar images, alignment, GC/stack growth, registry-only retention via weak
references, typed lookup/free and nil-TLS errors. Existing disabled/native DRED
validation/ABI fixtures and public integer adapters persist.

Every separately committed round passes full amd64/386, ARM64/QEMU, scoped
checkptr, native comparisons, GC decode stress, unchanged codec references and
diff checks. Rounds3/4 additionally run the unchanged PartA/PartB goldens under
checkptr on amd64/386/ARM, and final ARM scoped/ordinary fixtures (including
BOTH reference parts) repeat twenty times. The previous raw PartB fixture
checkptr blocker at opus_decode_frame_ref_test.go:168 is now removed; this is a
fixture representation repair, NOT a claim that all legacy tests or opaque
allocations are globally checkptr-safe. Extension EOF/GC/raw-callback issues
remain outside this scope.
Logs: /tmp/opus-round-dred-canonical-scalar-storage.log,
/tmp/opus-round-dred-registered-concrete-factory.log,
/tmp/opus-round-frame-reference-{a,b}-typed-entries.log,
/tmp/opus-round-frame-reference-{a,b}-focused-checkptr.log and
/tmp/opus-round-dred-typed-frame-arm-{checkptr,golden}-repeat.log.

Four decoder pseudostack-compatibility rounds resolve registered headers as
actual typed roots via libcshim.RegisteredGoAllocation, replace both scratch/cursor
uintptr fields with *byte, construct resampler reference headers as registered
scanned Go objects, and do the same for quant-all-bands reference headers/cursor
resets. Upstream celt/stack_alloc.h declares scratch_ptr/global_stack as char*.
The two pointer fields retain their two-word ABI on amd64/386/ARM; they are not
opaque aliases. Cleanup uses typed nil stores/write barriers, preserving its
scratch-nonnull conditional, free/key ordering, nil TLS and idempotence. Integer
TLS keys remain real compatibility handles. Unknown external headers still need
an explicit ABI cast; registering/casting byte images does not supply a GC map.
Decoder kernels still do not consume this legacy pseudostack for DSP scratch.

Counts change156→158 unsafe.Pointer and138→133 uintptr. Two actual pointer-valued
field types are removed; three further uintptr tokens were stale comment/numeric
offset cleanup. The two added unsafe casts are honest external-header/free
boundaries replacing previously hidden integer loads/stores, not opacity
concealed in a helper. The generic registry lookup returns the real typed owner
only (byte allocations/type mismatches return nil); it reconstructs no pointer.
Sentinel fixtures now use real Go byte addresses instead of invalid integer
123/456 pointers, while testing the same untouched/reset cursor behavior. Codec
PCM/state/range goldens and native VAD fixtures retain their expectations.

Grouped fixtures verify field size/offset ABI, scratch/cursor live alias, byte
owners through weak references/GC/stack growth, registered header cleanup,
conditional empty-scratch cursor preservation, deleted registry/key and repeated
free. Legacy producer migrations are test-storage ownership repairs, not new
codec pseudostack work. Retained failures:
/tmp/opus-round-pseudostack-concrete-byte-fields.log caught one native VAD fixture
still assigning uintptr to the new fields; concrete byte-pointer assignments fix
it and the complete -retry.log passes. /tmp/opus-round-resampler-pseudostack-partb-checkptr.log
shows the old setupResamplerPseudostack issue is bypassed, but PartB still fails
checkptr at opus_decode_frame_ref_test.go:168 on its separate raw decoder fixture.
It remains an ordinary golden test, not a claimed global checkptr success.
/tmp/opus-round-quant-scanned-pseudostack-producer.log caught its old Xmemset
integer argument after scratch became *byte; clearing the exact known byte
allocation with a typed slice preserves its zeroed reference scratch, and the
complete -retry.log passes.

Every separately committed round passes full amd64/386, ARM64/QEMU, scoped
checkptr, native comparisons, GC decode stress, unchanged encode/decode references
and diff checks; final ARM ordinary/scoped fixtures repeat twenty times. No
extension EOF/GC, opaque-image or raw-callback global safety claim is made.
Logs: /tmp/opus-round-pseudostack-typed-registry-cleanup.log,
/tmp/opus-round-pseudostack-concrete-byte-fields-retry.log,
/tmp/opus-round-resampler-scanned-pseudostack-producer.log,
/tmp/opus-round-quant-scanned-pseudostack-producer-retry.log and
/tmp/opus-round-pseudostack-byte-roots-arm-{checkptr,golden}-repeat.log.

Four multistream-offset/pthread-declaration rounds replace positive consumed
MS packet-offset arithmetic with byte indexing, express the fixed MS header's
alignment as a numeric int offset, remove unused generated mutex/condition
layouts, and remove unused rwlock/barrier layouts plus pthread_t. Counts fall
157→156 unsafe.Pointer and144→138 uintptr. ONE active cursor representation
change, ONE numeric-size bookkeeping change, and TWO dead-declaration rounds
are distinct: zero-length pointer-array alignment metadata and an unused thread
handle declaration are not repaired pointer lifetimes. Removed generated Go
type names have no repository consumers; real TLS/pthread emulation, public
integer adapters and allocator synchronization are unchanged.

Upstream src/opus_multistream_decoder.c verifies the current packet cursor is
consumed only with positive remaining length outside PLC. opusMSPacketAt returns
nil before any offset arithmetic for no work and preserves exact offset0
identity (including nil); positive offsets index only through the current byte,
not a guessed/unused suffix. Numeric MS header alignment still gives the port's
fixed272-byte offset on amd64/386/ARM. Existing owner/packet fixtures now also
check zero-offset identity, a minimal consumed prefix and a source guard, plus
numeric header offset vs actual child layout. No EOF pointer is synthesized.
Malformed unrepresentable offset/extent geometries are not claimed supported.

Each separately committed round passes full amd64/386, ARM64/QEMU, scoped
checkptr, native comparisons, GC decode stress, unchanged encode/decode references
and diff checks; final ARM scoped/ordinary fixtures repeat twenty times. No
extension EOF/GC, opaque-allocation or raw-callback safety claim is made.
Logs: /tmp/opus-round-ms-consumed-packet-offset.log,
/tmp/opus-round-ms-numeric-header-alignment.log,
/tmp/opus-round-remove-unused-mutex-cond-address-layouts.log,
/tmp/opus-round-remove-unused-rwlock-barrier-thread-address-layouts.log and
/tmp/opus-round-ms-offset-pthread-arm-{checkptr,golden}-repeat.log.

Four packet/dead-declaration rounds index readable packet-parse cursors directly
through the already consumed byte view, remove unused generated stdio cookie
callback-address fields, delete duplicate unused compiler/isoc varargs aliases,
and delete unused generated channel-in/channel-out/downmix address aliases.
Counts fall158→157 unsafe.Pointer and153→144 uintptr. This is ONE active typed
packet-cursor change plus THREE dead-metadata deletions, not nine repaired
pointer lifetimes. The packet cursor's exact EOF/backward identity still uses
the existing explicit extensionByteAt compatibility boundary; its cast was not
hidden in a new wrapper. Public integer parse/decode/control/callback adapters,
the used OpusT_va_list and real custom callback ABI remain intact. Only unused
generated Go declaration names are removed; no replacement opaque typedefs.

Upstream src/opus.c confirms packet-size/frame/padding cursor ordering. The
existing typed packet view already covers the input; new indexing declares no
guessed span and changes no parsing reads, output stores, padding initialization,
validation/errors, narrowing, private suppression of unused EOF output, or live
aliases. Grouped packet fixtures check readable and empty-frame cursor identity,
payload/packet offsets, semantic EOF padding identity with a real allocation
guard, and GC/stack growth. Existing parse/native comparisons and codec goldens
persist. The declaration deletions have no runtime behavior claim.

Every separately committed round passes full amd64/386, ARM64/QEMU, scoped
checkptr, native comparisons, GC decode stress, unchanged codec references and
diff checks; final ARM ordinary/scoped fixtures repeat twenty times. Guarded
semantic EOF fixtures do not establish global exact-allocation EOF/extension GC
safety, and other external opaque/callback boundaries remain.
Logs: /tmp/opus-round-packet-consumed-cursors.log,
/tmp/opus-round-remove-unused-stdio-cookie-addresses.log,
/tmp/opus-round-remove-unused-varargs-address-aliases.log,
/tmp/opus-round-remove-unused-codec-callback-address-aliases.log and
/tmp/opus-round-packet-declarations-arm-{checkptr,golden}-repeat.log.

Four decoder callback-type removal rounds bind legacy multistream float,
int16 and int24 destinations as concrete captured *float32/*int16/*int32 roots,
then bind custom legacy callbacks and remove the obsolete opusMSChannelCopy
opaque-PCM callback type, opusMSDecodeNative wrapper and intermediate binder.
The public integer native adapter now calls opusMSDecodeBound directly. This
also removes the format-specific erased destination adapters; standard PCM
captures never cross the raw uintptr callback trampoline. Only the boundary
that identifies an externally supplied callback/destination and genuinely
custom callbacks retain an erased destination. The latter keeps that pointer
as a GC-visible captured root and retains the real function value, user data,
src/dst integer ABI and uintptrescapes invocation. Raw function-address ownership
is not globally repaired. Counts fall162→158 unsafe.Pointer; uintptr remains153.
This combines concrete active legacy-format routing with deletion of its now
unneeded private erased bridge; no necessary boundary casts were moved to tests.

Upstream src/opus_multistream_decoder.c confirms per-call live channel strides,
left/right/mono sequential copies, muting and the float/short/int24 kernels.
Their arithmetic, signed narrowing, rounding, clipping, validation order,
source live aliases, decoded state, PLC/FEC, final ranges and guards are unchanged.
Grouped fixtures test concrete bindings across forced GC/stack growth and
muting/guards, standard callbacks ignoring user data, and the custom callback's
arguments/function owner through GC. Existing format/projection comparison
fixtures now invoke the same concrete kernels with bound destinations instead
of using the deleted erased wrapper; all expected PCM/state/ranges persist.
Native channel/codec comparisons and reference tolerances are unchanged.

Retained staging failure /tmp/opus-bind-focused.log: prematurely dropping the
old float-binder arm while old fixtures still used it sent a standard float
through the uintptr fallback and failed checkptr. Its typed compatibility arm
was restored until final consumer migration; the complete round1 validation
passes. /tmp/opus-ms-bound-final-focused.log caught two now-unused test unsafe
imports after their opaque projection sinks were removed; imports were removed,
and /tmp/opus-ms-bound-final-focused-retry.log passes. No goldens or fixtures
were weakened. Every separately committed round passes full amd64/386,
ARM64/QEMU, scoped checkptr, native comparisons, GC decode stress, unchanged
codec references and diff checks; final ARM ordinary/scoped fixtures repeat
twenty times. These scopes do not establish global raw-callback, external-byte
allocation or extension EOF/GC safety.
Logs: /tmp/opus-round-ms-legacy-{float,short,int24}-bound-pcm.log,
/tmp/opus-round-ms-remove-opaque-pcm-callback-type.log and
/tmp/opus-round-ms-bound-pcm-arm-{checkptr,golden}-repeat.log.

Four typed decoder-clearing rounds add field-wise SetZero descriptors for actual
canonical Go images, use them in single factory initialization, extend them to
standalone MS children, and enable them for projection children. A descriptor
captures its real typed reflection image, not an opaque address. Each declared
field (including numeric padding arrays and the float tail) is cleared by its
actual type; pointer-bearing SILK/CELT fields therefore use Go's typed clearing
barriers. The initializer keeps its pre-existing CELT mode nil and full typed
SILK zero ordering. Public caller-buffer initializers still clear the ABI byte
image after those pointer stores; all valid Go factory paths now bypass that
byte-wise clear. Counts remain162 unsafe.Pointer/153 uintptr: their necessary
external-buffer fallback is not removed or disguised.

The entire enclosing reflection struct is deliberately NOT zeroed. A single
image may have final Go padding outside the reported decoder size; for MS/projection
that padding can live in the child stride. Clearing declared fields preserves
those bytes, the MS header/mapping, projection matrix, and sibling children.
Grouped fixtures install live foreign mode/codebook/resampler pointers, verify
exact zero ABI images through GC/stack growth, poison only scalar outside-image
padding/neighbor/stride bytes, and confirm neighboring bytes and sibling foreign
owners survive. Single init/re-init checks count descriptor invocation and verify
invalid rates do not clear anything. Existing complete initializer image/status,
projection source-alias, native init/reset/validation and registry fixtures also
exercise this path without changed expectations. Reflection/closure metadata
costs are not benchmarked or claimed to improve performance.

Every separately committed round passes full amd64/386, ARM64/QEMU, scoped
checkptr, native comparisons, GC decode stress, unchanged codec references and
diff checks; final ARM ordinary and scoped fixtures repeat twenty times. No
global opaque-allocation, extension EOF/GC or raw-callback safety claim is made.
Logs: /tmp/opus-round-decoder-typed-field-clear.log,
/tmp/opus-round-single-factory-typed-clear.log,
/tmp/opus-round-ms-factory-typed-child-clear.log,
/tmp/opus-round-projection-factory-typed-child-clear.log and
/tmp/opus-round-typed-child-clear-arm-{checkptr,golden}-repeat.log.

Four concrete decoder-codec initialization rounds expose Decoder/SILK/CELT
views from canonical Go images, bind single-decoder factory initialization to
those fields, bind standalone multistream children to complete codec views, and
propagate the same bindings into projection factory children. Views are obtained
from actual typed reflection fields, not integer address reconstruction or opaque
pointer wrappers. Valid Go factory paths no longer recover either codec pointer by layout
arithmetic. Public caller-buffer initializers, decode-time accessors and necessary
ABI byte clearing retain explicit boundaries. Existing header-only child helpers
remain compatible. Counts stay162 unsafe.Pointer/153 uintptr: these changes
bypass active factory casts, not delete their required caller-buffer fallbacks.

Upstream src/opus_decoder.c, opus_multistream_decoder.c and
opus_projection_decoder.c supply the size/offset and initialization semantics.
This retains the port's Fs/channel checks, initialization/error ordering and
GC-required CELT mode nil then complete typed SILK zero before the ABI clear. Single initialization still computes/stores the exact ABI
offsets; the storage views choose only its actual codec fields. Child channel
selection and numeric strides, forward mapping aliases, projection source
snapshot/header/matrix-before-rate validation, malformed factory byte allocation
and registry cleanup stay unchanged. Views are temporary initializer metadata;
registered interior headers still root the genuinely scanned compound allocation.
The additional codec-view vectors have no performance/memory improvement claim.

Grouped fixtures verify concrete slot identities, foreign mode/codebook/resampler
owners through GC and stack growth, single reinitialization after installing live
pointer-bearing state, unchanged invalid-rate images, and complete single/MS/
projection ABI-image/status comparisons across mono/coupled geometries, valid
rates, invalid rates and matrix-size errors. Existing projection source/header
alias tests also exercise the new codec binding. Native initialization/reset/
validation comparisons and all codec goldens persist. Each separately committed
round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons, GC
decode stress, unchanged codec references and diff checks; final ARM ordinary and
scoped fixtures repeat twenty times. This does not repair external opaque-buffer
GC maps, raw callbacks, extension EOF/GC or decode-time layout boundaries.
Logs: /tmp/opus-round-decoder-concrete-codec-views.log,
/tmp/opus-round-decoder-bound-codec-init.log,
/tmp/opus-round-ms-bound-child-codec-init.log,
/tmp/opus-round-projection-bound-child-codec-init.log and
/tmp/opus-round-codec-bound-init-arm-{checkptr,golden}-repeat.log.

Four decoder-package mini-FFT rounds replace the real-FFT byte suffix with
canonical complex cells, add an actual Go complex guard outside the reported ABI
image, bind Go tmp/super views directly to those cells, and read packed real input
through float32 lanes instead of a float→complex pointer reinterpretation. The
shared recursive FFT work retains its complex-input path and gains a scalar-lane
path with the same factor reads, leaf load-before-store order, live partial input
aliases, recursion, twiddle reads, butterflies and explicit float32 arithmetic.
The in-place assertion uses typed first-real-field identity, including nil/nil,
and retains its original diagnostic/line. Counts fall 163→162 unsafe.Pointer;
uintptr remains153. Typed tail/views are representation improvements even without
further lexical deletions; caller-supplied byte headers retain their ABI casts.

Upstream celt/mini_kfft.c confirms substate→tmp→super layout and real even/odd
packing. The nfft2 allocation has a scalar substate of size1 with zero consumed
super twiddles, but its native super pointer is at the reported allocation end.
The Go guard preserves that exact pointer identity inside declared Go storage,
so a scanned pointer field does not point past its actual compound allocation.
Required sizes/lenmem queries and caller-byte layouts do not include the guard.
No support for the upstream-unsupported scalar size1 transform is invented.

Grouped fixtures verify complex tail types, header/substate/tmp/super geometry,
guard placement/zeroing/identity, GC and stack growth, supported transforms,
scalar/complex path bitwise equivalence across radices and forward/inverse modes,
live partial aliases/whole-buffer guards and in-place assertions. Existing
upstream native FFT/real-FFT comparisons, encode/decode goldens and tolerances
persist. Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native
comparisons, GC decode stress, codec references and diff checks; final ARM scoped
and ordinary fixtures repeat twenty times. This standalone mini-FFT work does
not repair external opaque images, extension EOF/GC or raw callback lifetimes.

Retained failures: /tmp/opus-round-mini-fftr-scanned-terminal-guard.log failed a
new fixture that incorrectly attempted the unsupported scalar size1 transform;
the allocation/guard test stays, and transform checking uses supported size4.
Its complete retry /tmp/opus-round-mini-fftr-scanned-terminal-guard-retry.log passes.
/tmp/opus-round-mini-fftr-bound-complex-buffers.log caught a new test's uint vs
uintptr size-expression compile mismatch; explicitly numeric uint sizes fix it,
and the full -retry.log passes. Neither production behavior nor goldens were
relaxed. Remaining round logs are /tmp/opus-round-mini-fftr-complex-go-tail.log
and /tmp/opus-round-mini-fftr-typed-real-input.log.

Four decoder extension-cursor rounds index short/terminal payload advances,
lacing/payload advances, ID headers, and normal/repeated payload starts through
consumed byte prefixes plus numeric offsets. extensionByteAt returns the exact
input at offset0; readable positive offsets use an actually consumed prefix.
Exact EOF, backward and otherwise unproven cursors retain one explicit unsafe
boundary. Counts fall 169→163 unsafe.Pointer, with uintptr153 unchanged. The six
lexical deletions centralize seven necessary cursor sites into one fallback;
readable consumers genuinely use typed indexing, but this is not a global EOF or
GC lifetime repair and is not counted as deleting public ABI adapters.

Upstream src/extensions.c confirms failure/output ordering and lacing semantics.
The lacing loop tracks its physical suffix separately: logical remaining length
has already deducted payload bytes not yet advanced. This lets its header and
payload cursors use exact physical budgets, without guessed extents. Payload-start
spans are cached before output stores only to choose the consumed view; payload
length still computes the live pointer difference after those stores, and both
iterator pointer/remaining-length assertions remain unchanged. Repetition, frame
limits, sticky errors and header/payload IDs are unchanged.

Grouped fixtures cover zero/nil identity, readable offsets, guarded semantic EOF
and backward boundaries, multi-byte lacing with cached 255-byte payload and
malformed lengths, untouched failure outputs, ID-header work, normal/repeated
payload starts, GC/stack growth and whole-packet guards. A Go-only output alias
into scalar iterator counters (not its GC pointer slots) verifies final live
length reloads; it does not claim C cross-struct-alias validity. Existing native
extension/init/next/find/parse/generate comparisons and codec goldens remain.
Each round passes full amd64/386, ARM64/QEMU, existing decoder plus new cursor
scoped checkptr, native comparisons, GC decode stress, unchanged encode/decode
references and diff checks; final ARM ordinary/scoped fixtures repeat twenty
times. These scopes do not establish global opaque-buffer/EOF safety. Native
extension GC failures, raw callbacks and malformed external allocations remain
separate issues.

Four projection/initialization ownership rounds add canonical projection Go
storage, migrate valid projection factories to registered scanned owners, bind
standalone multistream factory initialization to concrete child headers, then
bind projection factory matrix/coefficient/MS/child initialization to concrete
views. The projection runtime type contains the actual scalar projection header,
canonical matrix header plus int16 coefficient/padding cells, and the canonical
multistream compound type. All successful projection factory states now scan
their embedded decoder pointers, superseding the earlier factory limitation.

Factory creation still queries size and attempts allocation before initialization
validation. Invalid geometries keep the old byte-allocation/error path rather
than inventing a reflective layout or changing error ordering. Nil-TLS behavior,
coefficient little-endian/signed conversion, complete input snapshot before
header writes, forward matrix stores, mapping validation before child rate checks,
coupled-before-mono order, cleanup and 16-byte alignment are preserved. Bound
factory initialization no longer recovers matrix/MS/child pointers with layout
arithmetic; public initializers and malformed-layout paths retain explicit
compatibility fallbacks. Decode-time layout accessors remain necessary boundaries.
Production counts remain 169 unsafe.Pointer/153 uintptr; these are genuine Go
allocation/ownership and concrete-view improvements, not lexical cast deletion.

Upstream src/opus_projection_decoder.c and mapping_matrix.c establish the header,
matrix, child layout and snapshot/error order. Grouped canonical/registered
projection fixtures cover exact sizes/offsets/alignment and weak-reference survival
of both coupled/mono children with only integer handles plus TLS. Bound init
fixtures compare complete ABI images and return codes against compatibility
initializers under GC/stack growth, including aliased projection/matrix source
headers, size mismatch and invalid rate. Existing native layouts/initialization,
format comparisons and codec goldens are unchanged. Each round passes full
amd64/386, ARM64/QEMU, scoped checkptr, native comparisons, GC decode stress,
encode/decode references and diff checks, with twenty final ARM ordinary/scoped
repetitions. Scanned-alignment fallback overhead and reflective type creation
remain as documented below. Caller-supplied/invalid byte images, extension EOF/GC,
raw callbacks and other opaque allocations are not globally repaired.

Four decoder ownership rounds clear every embedded SILK pointer via typed
stores before composite initialization's ABI-image clear, construct canonical
Go single-decoder storage, register/migrate single factories, then construct and
register canonical multistream factories. SILK has ICDF, codebook and resampler
pointer slots in both channels, not just the CELT mode slot; its complete typed
zero assignment preserves write barriers before the byte clear. Caller-supplied
byte images still require external referent lifetime management.

opuscc/decoder_storage.go uses actual runtime-created structs containing the
canonical decoder/SILK/CELT types and float32 tail, with fixed-eight ABI padding.
Multistream images contain canonical coupled/mono child arrays at the original
aligned strides. Unlike casts over byte/word allocations, these allocations have
real GC pointer maps. An interior typed header roots the complete actual Go
allocation; libcshim.RegisterGoAllocation records that Go root for legacy integer
handles, while free/Close remove the registry entry. XmallocPointer itself remains
opaque byte allocation. Single and standalone multistream factories migrate;
projection factories and caller-owned ABI images remain separate unresolved
scanning boundaries. Production opuscc counts stay 169 unsafe.Pointer/153 uintptr:
this batch repairs representation/lifetime, not necessary layout/ABI casts.

Sixteen-byte allocation alignment is preserved without pretending a byte image
is scanned. Normally an aligned reflective image is used directly; otherwise a
scanned array of canonical images with controlled stride covers every alignment
residue. This fallback can cost two images on amd64/ARM64 or four on 386, and the
first unaligned allocation becomes GC-reclaimable. Single storage may have final
Go padding outside the reported ABI image; multistream child stride/report sizes
remain exact. Reflection/type creation is factory work, not per-packet decode;
no allocation-performance or huge/malformed-layout success claim is made.

Upstream src/opus_decoder.c and opus_multistream_decoder.c confirm validation,
initialization, cleanup and coupled-before-mono layout. Nil-TLS allocation failure,
mapping-before-child-rate validation, error cleanup, public APIs and numerical
codec behavior remain intact. Grouped geometry/alignment/weak-reference fixtures
prove that dynamically allocated CELT modes, SILK codebooks and resampler arrays
survive GC/stack growth through the actual compound owner, including both channels
and coupled/mono children. Registry fixtures retain only integer handles plus TLS;
the generic shim test roots a pointer-bearing parent through its scalar interior
header. Existing native layout/init/reset/control/format comparisons persist.
All four rounds pass full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
GC decode stress, unchanged encode/decode references and diff checks; final ARM
scoped/ordinary fixtures repeat twenty times. Opaque projection allocation,
extension EOF/GC and raw callback ownership are not globally repaired.

Four encoder sort/ID rounds migrate int32 increasing top-K sort, int16 full
sort, float32 decreasing top-K sort, and gains-ID into concrete private kernels.
Public integer adapters retain six explicit casts and uintptrescapes. Production
opusccenc counts fall 11276→11240 unsafe.Pointer and 6649→6607 uintptr; decoder
opuscc remains 169/153. Active loads/stores use numeric indices and typed views.
Original assertion conditions/diagnostics/order, stable strict ties, float NaN
comparisons/signed-zero copies, initial index writes, untouched top-K tails and
separate value→index stores persist. Singleton sorts consume only the index
output (or no pointers for int16); nonpositive ID counts consume no bytes. Gains
use signed int8 extension and uint32 shift/int32 wrapping addition.

Upstream silk/sort.c, silk/float/sort_FLP.c and silk/gain_quant.c/macros are the
oracles. Native fixtures now test both integer ports, float sorting (including
NaNs, signed zeros and whole-buffer/index guards), and signed/wrapping gain IDs.
Grouped private fixtures cover GC, aliases, singleton nil values and empty IDs.
All four rounds pass full amd64/386, ARM64/QEMU, scoped decoder/private-encoder
checkptr, native comparisons, GC decode stress, unchanged codec references and
diff checks; final ARM ordinary/scoped fixtures repeat twenty times.

The first sort32 run failed its new alias comparison at L16/K5/shift1:
/tmp/opus-round-encoder-sort32.log. A direct native probe confirmed the encoder
rewrite matched C (/tmp/opus-round-encoder-sort32-alias-native-probe.log). The
existing decoder sort's parallel Go assignment incorrectly cached the index
read before C's preceding value store. Changing its four assignment sites to
sequential stores repairs that live alias; preserving the fixture and extending
the C alias oracle to both ports gives a passing full
/tmp/opus-round-encoder-sort32-retry.log. Disjoint codec paths/goldens remain
unchanged; this is an upstream-order correction, not weakened expectations.

The float-sort full run hit the recurring native extension GC failure, fatal
error: found pointer to free object, in TestExtensionNextAgainstC (extensions
iterator stack), retained in /tmp/opus-round-encoder-sort-float.log. Its complete
retry /tmp/opus-round-encoder-sort-float-retry.log passes; no extension/allocator
repair is claimed. Encoder opaque-state scanning/pseudostack, extension EOF/GC
and raw callback lifetime remain outside these scalar-leaf migrations.

Four encoder FLP-vector rounds migrate scale-copy, in-place scale, inner product
and energy into concrete float32-pointer private kernels, with numeric indices
and consumed-prefix slices. Public integer adapters retain uintptrescapes and
six explicit casts. Production opusccenc counts fall 11305→11276 unsafe.Pointer
and 6684→6649 uintptr; decoder opuscc remains 169/153. These are actual load/store
and address-representation changes, not deletion of dormant implementations.

Upstream silk/float/{scale_copy_vector_FLP,scale_vector_FLP,inner_product_FLP,
energy_FLP}.c confirms forward stores/live overlapping aliases and the scale
helpers' literal 0xFFFC mask, not ~3. In particular size=-1 still consumes 65532
samples; size=-65536 consumes none. Views use max(size, masked bulk), preserving
that behavior and >65535 tails. Reduction bounds retain int32 size-3 arithmetic;
the numeric extent oracle tests the signed wrap without constructing enormous
addresses. Native signed-overflow/unrepresentable buffer geometries are not
claimed supported. Four-product groups are added to the running float64 total,
not four independent updates; products widen both operands first and retain
explicit rounding. Energy's translated disabled assertion remains a no-op.

Grouped private fixtures exercise aliases/whole-buffer guards, GC, empty/nil
unused inputs, signed zeros, Inf/NaN, subnormal double products and a cancellation
case distinguishing group accumulation. New upstream C bridges compare finite
outputs/results bitwise for unrolled/tail/large/negative-mask work and random
wide-exponent reductions; native fixtures/goldens/tolerances stay unchanged.
Each round passes full amd64/386, ARM64/QEMU, decoder and private encoder scoped
checkptr, native comparisons, GC decode stress, codec references and diff checks,
plus final repeated ARM scoped/ordinary runs. No opaque encoder-state scanning,
pseudostack, extension EOF/GC or raw callback lifetime repair is claimed.

The first scale-copy run failed its new fixture: the zero-work forward-overlap
case formed &buffer[2] in a length-2 buffer before calling the kernel. Giving the
fixture the same extra terminal guard as its shifted cases fixed the fixture;
production code/goldens were not changed. The failure is retained in
/tmp/opus-round-encoder-flp-scale-copy.log and full retry passed in
/tmp/opus-round-encoder-flp-scale-copy-retry.log. Later round logs are
/tmp/opus-round-encoder-flp-{scale,inner-product,energy}.log.

Four encoder LPC-analysis rounds migrate the independent encoder leaf's output
stores/late prefix clearing, fixed six prediction taps, remaining taps/prediction
subtraction, then its private parameters/cursor to concrete int16 pointers and
numeric indices. The public integer API remains an explicit three-cast adapter
with uintptrescapes; active math no longer constructs addresses. Production
opusccenc counts fall 11320→11305 unsafe.Pointer and 6696→6684 uintptr. Decoder
opuscc counts remain 169/153. The last unused integer cursor assignment/declaration
is separately dead-code cleanup; the load/store changes are actual representation
improvements, not merely moving casts to another helper.

Upstream silk/LPC_analysis_filter.c and its wrap macros confirm six initial taps,
paired remaining taps, uint32 wrapping adds/subtraction, signed rounded shift,
saturation, and zeroing the first d samples only after every prediction. The
original three assertions/diagnostics/order persist. Input/coefficient views are
formed only for len>d; no-work calls consume only output. Coefficients and input
remain live views, including output overlaps. Native C fixtures now exercise both
ports across orders6/8/10/16/24, empty/short/long work, extrema/random values,
forward/backward/in-place input/output overlap and coefficient/output overlap;
whole buffers/guards are compared. Grouped typed-encoder fixtures cover GC,
overlap and nil unused input/coefficient pointers under scoped checkptr.

Each round passes full amd64/386, ARM64/QEMU, decoder scoped checkptr, native
comparisons, GC decode stress, unchanged encode/decode references and diff checks.
The completed private encoder kernel additionally passes amd64/386/ARM scoped
checkptr and final repeated ARM ordinary/checkptr fixtures. Goldens/tolerances
are unchanged. This does not repair the encoder's legacy pseudostack or opaque
state scanning, nor the known extension EOF/GC or raw callback lifetime issues.

Four postfilter-owner rounds bind normal CELT first/tail postfiltering and
noise-PLC first/tail postfiltering to their existing per-channel history slices.
Normal dispatch carries the concrete history array into first/tail helpers;
compatibility helpers still accept standalone output pointers through nil-owner
fallbacks. Noise PLC passes its already-bound history directly. All active comb
calls in celt_decoder.go now bypass backward pointer recovery. Production
lexical counts remain 169 unsafe.Pointer/153 uintptr: required public/standalone
compatibility boundaries are retained rather than artificially concealed.

The numeric current offset is DEC_PITCH_BUF_SIZE-N (plus the short-MDCT length
for tail work). Channel-zero do-while behavior, clamp→first→tail order, LM==0
tail suppression, source/destination live aliases, control/window reads, cached
mode and float32 arithmetic remain unchanged. Grouped owned-dispatch fixtures
compare complete state and bitwise histories/guards against standalone kernels
for LM0..3 and channels0/1/2 under GC/stack growth. Existing noise PLC, normal
decode, native postfilter and entropy/PCM goldens persist without tolerance or
assertion changes. Each round passes full amd64/386, ARM64/QEMU, scoped checkptr,
native comparisons, codec references, GC stress and diff checks, with final
repeated ARM scoped/ordinary fixtures. This is active scalar-owner binding, not
opaque-allocation scanning, extension EOF/GC or raw callback lifetime repair.

Four owned-history rounds add history-aware CELT FIR and comb kernels, then
bind periodic PLC FIR to its complete Go excitation prefix and prefilter/fold
comb filtering to its known decoder history owner plus numeric current offset.
These active consumers no longer recover preceding samples by subtracting from
an interior pointer. Public pointer APIs still use the explicit backward-layout
fallback. Production lexical counts remain 169 unsafe.Pointer/153 uintptr:
this batch bypasses required compatibility boundaries rather than hiding or
deleting them. FIR's zero-order fallback directly uses its typed input pointer.

The excitation prefix begins ord samples before the consumed source; the comb
owner is the existing DEC_PITCH_BUF_SIZE+overlap per-channel history. Live channel
state reads, initial coefficient reversal, assertions, gain/tapset lookup and
zero-gain early returns, overlap/window reads, in-place dependencies, terminal
copy suppression and float32 products remain in their original order. Existing
upstream FIR/comb and PLC/prefilter native/golden fixtures persist. Grouped owner
fixtures compare exact output/guards with compatibility kernels across empty,
short/unrolled work, valid orders, gain combinations, overlap, live in-place
aliases and GC/stack growth. No goldens/tolerances/assertions changed.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references, GC stress and diff checks, plus final repeated ARM scoped/
ordinary fixtures. Owned scalar prefixes do not repair opaque allocation
scanning, extension EOF/GC or raw callback lifetime; other comb/history callers
still retain their necessary boundary until their owners are separately bound.

Four further mini-storage rounds index externally supplied real-FFT submemory
through its validated byte prefix, derive temporary storage from the canonical
scalar FFT twiddle/temporary complex prefix, give Go-created scalar flexible
tails an actual complex-cell array type, and allocate Go-created real FFTs with
canonical typed outer and scalar substate headers. The latter initializes its
substate through miniFFTInit without a byte-buffer reinterpretation; supplied
C-layout memory still uses its explicit scalar-header cast. The shared initializer
preserves all stores, factor/twiddle operations and ordering. The outer header
continues to have a real GC bitmap; this does not scan arbitrary caller bytes.

Production unsafe.Pointer counts decrease 171→169; uintptr stays 153. Two offset
casts are removed; the typed-tail/subheader rounds improve Go allocation and
consumer representation without claiming additional lexical deletion. Grouped
geometry fixtures prove flexible twiddles end at the scalar header boundary,
check temporary offsets for real sizes 2/4/6/16, typed scalar tail element kinds,
canonical substate displacement, empty/short tails and GC/stack ownership.
Existing foreign weak-owner, caller-buffer capacity/guard, scalar/native mini
FFT and codec fixtures remain. Numerical phases, reported sizes, layout,
assertion/validation ordering and supplied-buffer aliases remain unchanged.
Dynamic Go allocations may have extra unconsumed final padding; oversized/
invalid allocation configurations retain the previously documented limits.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references, GC stress and diff checks, with final repeated ARM scoped/
ordinary fixtures. No goldens/tolerances/assertions changed. These are standalone
mini-FFT paths, not active decoder FFT hot loops or global opaque-state scanning,
extension EOF/GC or raw callback lifetime repair.

Four mini-FFT storage rounds replace Go-created scalar/real FFT word-slice
images with dynamically sized, contiguous Go objects containing canonical typed
headers, bind real-FFT substorage through the owning Go byte tail, and index
consumed super-twiddle prefixes through concrete complex pointers. Dynamic
reflect.StructOf allocation gives the real FFT header an actual runtime GC
pointer bitmap; this is not a typed cast over unscanned bytes. Its scalar flexible
suffix remains pointer-free. Caller-supplied opaque byte layouts still use
explicit casts, and empty/backward/unused-end super-twiddle boundaries remain.
Production unsafe.Pointer counts decrease 173→171; uintptr stays 153. The latter
two rounds improve consumed paths without claiming lexical boundary deletion.

Upstream mini_kfft.c's header→substate→temporary→super-twiddle layout, reported
size, factor/twiddle arithmetic, assertion order and caller-buffer capacity
checks remain unchanged. Go allocation may include additional end padding for
zero-length tail fields, outside the reported/consumed ABI image. Nonrepresentable
Go allocation sizes fail before initialization; huge/invalid FFT configurations
are not an allocation-success compatibility claim. Dynamic type creation is on
standalone mini-FFT allocation paths, not active decoder FFT hot loops.
Grouped layout/tail/GC/stack tests cover typed ownership and exact consumed
prefixes; a noinline weak-owner fixture proves a foreign Go temporary stored in
a Go-created real-FFT header survives GC. This scan guarantee does not extend
to externally supplied byte buffers or the codec's other opaque allocations.

The first final-round run failed the new nil/zero-offset identity fixture:
`/tmp/opus-round-mini-fftr-consumed-super-prefix.log`. The arithmetic boundary
produced a nil-looking pointer yet failed the inlined nil comparison. The helper
now returns the concrete input directly at offset zero, without unsafe.Add;
fixtures/assertions were retained and a complete retry passed. This is not a
codec golden/tolerance adjustment.

Each round ultimately passes full amd64/386, ARM64/QEMU, scoped checkptr, native
comparisons, codec references, GC stress and diff checks, with final repeated
ARM scoped/ordinary fixtures. Existing native mini-FFT/allocation references and codec
goldens/tolerances remain unchanged. Extension EOF/GC and raw callback lifetime
remain independently unresolved.

Four numeric/private-entry rounds share projection's matrix→MS numeric
displacement with the native-reference CTL bridge, track returned MS decoder
children by concrete pointer identity plus recorded numeric displacement, then
migrate private opus_decode_frame's state/packet and PCM arguments to concrete
pointers in separate steps. No pointer-address subtraction remains in these
reference bridges, and the private frame forwarder has no integer handles,
reinterpretation casts or uintptrescapes directive. Its eight grouped fixtures
now pass concrete owners directly; registered byte-backed test state retains
an explicit allocation-boundary cast. Public integer entry APIs remain intact.

Production lexical counts decrease 180→173 unsafe.Pointer and 159→153 uintptr:
four casts/four integer tokens from reference address subtraction, three casts/
two tokens from private entry conversion. The numeric projection helper retains
uint32 addition/alignment then int32 narrowing and machine-word conversion.
Grouped numeric fixtures cover ordinary sizes, high-bit/negative/wrapped values
without forming invalid addresses. Existing native CTL child/matrix/alias
matrices and frame PCM/error/GC/stack/golden fixtures remain unchanged. Child
identity tracking is reference-only; opaque state allocations are not globally
made GC-scanned. Decoder/packet/PCM validation and argument order persist.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references, GC stress and diff checks, with final twenty-repeat ARM scoped/
ordinary fixtures. Goldens/tolerances/assertions unchanged. Layout boundaries,
extension EOF/GC and raw callback lifetime remain independently scoped.

Four pointer-removal rounds delete the unused private integer packet-parser
adapter (no callers in opuscc), then migrate multistream, single and custom
CTL native-reference bridges to concrete int32/uint32 output cells instead of
reinterpreting one uint32 cell as *int32. Numeric signed→unsigned conversion
preserves result bits; only final-range requests consume unsigned outputs.
Real state/output alias fixtures retain their explicit raw-layout pointers.
Mode/decoder outputs, nil-output requests, initial 77 sentinels and error order
remain unchanged. Single/MS child lookup reuses typed layout helpers; this
centralizes three casts rather than eliminating their underlying boundaries.
The MS header alignment is numeric int arithmetic, not a uintptr word mask.

Production lexical counts decrease 196→180 unsafe.Pointer and 167→159 uintptr.
Ten pointer casts/seven integer tokens belonged to the deleted dead adapter;
three scalar casts are genuinely removed, three layout casts centralized and
one size mask is numeric. Public parser/escape APIs remain intact. These CTL
changes are tagged native-reference bridge code, not decoder hot-path/allocator
scanning improvements. Native matrices already cover signed gain and high-bit
range output; added grouped custom CTL tests compare signed/unsigned 0, max,
minimum and all-one result bits plus complete state images against C.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references, GC stress and diff checks, with final twenty-repeat scoped/
ordinary ARM fixtures. No goldens/tolerances/assertions changed. Opaque byte
allocation scanning, extension EOF/GC and raw callback ownership remain open.

Four state/recovery diagnostic rounds migrate the remaining seventeen CELT
state-validation assertions, six SILK frame assertions, five SILK core/CNG/PLC
assertions and five NLSF/gain helper assertions. All 33 sites now pass Go strings:
66 raw integer string-address arguments replaced, with 65 integer-addition
expressions removed (one argument was the pool base at numeric offset zero).
Production lexical counts remain 196 unsafe.Pointer/167 uintptr. This remains
static diagnostic representation work, not allocator/GC scanning repair.
The existing all-pool/GC string fixture covers these numeric offsets. Conditions,
source lines, file strings, format/abort behavior, assertion timing, arithmetic,
error order and public legacy adapters remain unchanged. Every round passes
full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons, codec references,
GC stress and diff checks; final ARM scoped/ordinary fixtures repeat twenty
times. Goldens/tolerances unchanged.

Four further diagnostic rounds migrate eight CELT control, six LPC filter,
six MDCT and six SILK API assertion sites to the existing Go-string fatal
helper. This removes another 52 integer-address constructions at 26 sites;
production lexical counts remain 196 unsafe.Pointer/167 uintptr. This is
representation cleanup of static diagnostic arguments, not a lifetime or
allocator-scanning repair. The grouped all-pool string/GC fixture remains.
All conditions, source lines, filenames, fatal format, abort calls, scalar
arithmetic and assertion positions remain unchanged. Each round passes full
amd64/386, ARM64/QEMU, scoped checkptr, native comparisons, codec references,
GC stress and diff checks, with final repeated scoped/ordinary ARM fixtures.
No goldens/tolerances/assertions changed; public legacy adapters remain.

Four diagnostic-string rounds introduce a Go-string fatal helper and static
pool numeric-offset string lookup, then migrate three CELT setup assertions,
three FFT assertions, five entropy assertions and twelve SILK resampler
assertions. These 23 sites eliminate 46 integer-address constructions and raw
string-address arguments; other sites and public Opus_celt_fatal remain legacy.
Production lexical counts stay 196 unsafe.Pointer/167 uintptr. This is diagnostic
representation cleanup, not allocator/scanning/callback lifetime repair.

Upstream arch.h supplies the unchanged format, assertion conditions, filenames,
line numbers and abort behavior. The public adapter retains file-before-message
conversion order; private string lookup returns the same NUL-terminated pool
prefix. A grouped fixture checks every pool string boundary under GC/checkptr.
Assertions remain at their existing points, with unchanged narrowing/branches,
codec arithmetic and error ordering. Each round passes full amd64/386,
ARM64/QEMU, scoped checkptr, native comparisons, codec references and GC stress,
with final repeated scoped/ordinary ARM fixtures. Goldens/tolerances unchanged.

Four Go-staging/API rounds move Decoder's int16 PCM, float PCM and packet
scratch into concrete Go slices, then expose typed single/MS int24 and all three
projection PCM entry points. Decoder no longer uses cBuf, cPointer/cBufferSlice,
shim scratch allocations or packet/PCM reinterpretation. It still stages rather
than writing caller PCM directly, preserving error-return isolation. Core opuscc
lexical counts remain 196 unsafe.Pointer/167 uintptr because legacy public entry
adapters remain; this batch improves actual outer allocation/consumer paths.

Existing int24/projection scanned-owner/native/golden/guard fixtures now also
exercise the newly exported typed entries. Grouped wrapper tests add int16/float
scratch reuse, invalid-packet caller-PCM preservation, packet reuse and nil PLC
with retained backing, plus release of all Go scratch at close. Existing deep
stack, full player seek/chunk, native encode/decode and GC fixtures persist.
Validation/frame-size/buffer/error order, typed format/FEC dispatch, staging-
before-decode, returned-prefix copying and repeated-close behavior remain
unchanged. PCM scratch is now separate by format; its stale contents are not an
input to valid non-accumulating decode. Encoder stable raw-interface buffers and
integer compatibility helpers remain unchanged. Go scratch uses normal GC
release rather than shim free; no whole output view is formed before validation.

Each round passes full amd64/386, ARM64/QEMU, scoped opuscc and typed-wrapper
checkptr, native comparisons, codec references and GC stress, with final repeated
scoped/ordinary ARM fixtures and wrapper repetitions across all three arches.
No goldens/tolerances/assertions changed. Decoder state remains byte-backed;
this is typed scalar scratch, not global embedded-pointer scanning, extension
EOF/GC or raw callback ownership repair.

Four public-wrapper rounds expose typed single and multistream int16/float
entry points, retain Decoder's single/MS state in concrete pointer fields with
typed creation/destruction, and retain scratch in cBuf's *byte owner with typed
allocation/copy/view operations. Decoder packet/PCM calls now bypass all integer
escape entry points; mapping is synchronously consumed directly from its Go
slice. Legacy encoder cBuf ensure/copyIn/cSlice integer adapters remain explicit.
Production opuscc counts remain 196 unsafe.Pointer/167 uintptr: this batch moves
the outer Go wrapper rather than deleting its retained public C-style adapters.
Decoder.go now has neither uintptr nor unsafe.Pointer references. Lexical counts
of cbuf.go include newly explicit encoder adapters, not an active decoder path.

Existing grouped scanned-owner int16/float single/MS fixtures now exercise the
new exported entries without changing goldens or guard/native comparisons.
Wrapper tests verify concrete state ownership, GC, both PCM-format PLC paths,
close/repeated close, scratch reuse/growth/signed values and empty-input owner
retention. Decoder validation/closed/error order, format/FEC dispatch, allocation
before decode, positive-return copy gating and actual returned-prefix views
remain unchanged. Full scratch views are not formed before decode, preserving
existing invalid-size/error ordering. Encoder calls and stable shim allocation
behavior remain unchanged; no caller stack slice crosses a raw integer encoder
interface. Mapping no longer needs its temporary shim copy.

Every round passes full amd64/386, ARM64/QEMU, scoped opuscc checkptr, native
comparisons, codec references and GC stress. Final typed wrapper/scratch tests
also pass scoped checkptr on amd64/386/ARM64, with repeated scoped/ordinary ARM
fixtures. No goldens/tolerances/assertions changed. Typed cBuf holds byte-backed
PCM/packet storage, not a general scanned object allocator. Decoder state still
uses existing opaque byte allocations: embedded pointer scanning, extension
EOF/GC and raw custom callback ownership are not claimed globally repaired.

Four CELT reset/init rounds bind old logs through existing concrete energy
views, reset the named scalar suffix before clearing trailing storage, clear the
allocation remainder through typed float32 words, and initialize the pointer/
scalar header through named fields plus that same memory clear. Production
unsafe.Pointer references decrease 199→196; uintptr remains 167. No byte-pointer
clear or integer-address energy/log traversal remains in these private paths.

Upstream celt_decoder.c init/reset remains the reference. Size-call/validation
order, GC-visible mode nil/rebind stores, all zeroing before initialization/
child dispatch, oldLog2→oldLog -28 writes and skip_plc/frame-type stores remain
unchanged. The 40-byte scalar prefix and 64-byte reset suffix have explicit
layout fixtures; the complete trailing region is 4-byte float words, including
header tail padding when present. Raw-byte-clear comparisons cover prefix/
flexible-element preservation, padding, terminal guards and empty memory clears.
Named header stores deliberately avoid whole-struct assignment, which could
write a flexible element/padding excluded by a minimal zero-channel allocation.
Existing whole-state native/init/reset/scanned-owner/golden fixtures persist.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with final repeated scoped/ordinary ARM fixtures.
Physical scalar stores replace bulk byte clearing, with no reads or externally
observable callbacks between those stores; live dependencies and final byte
images remain unchanged. No goldens/tolerances/assertions changed. Typed clears
do not turn opaque byte allocations into scanned state or repair extension
EOF/GC or legacy raw callback lifetime.

Four typed-state rounds read the IIR resampler union history through numeric
native-endian int32/int16 conversion, write its final history the same way
(removing the array-pointer reinterpretation), clear the standard decoder's
scalar reset suffix through its named fields, and preserve celtNormAdd's concrete
pointer identity at zero displacement. Production unsafe.Pointer references
decrease 201→199; uintptr remains 167. Nonzero norm offsets still retain genuine
signed history/unused-end-pointer support at the existing boundary.

Upstream resampler_private_IIR_FIR.c, opus_decoder.c reset and bands.c norm use
remain the reference. NativeEndian preserves both signed halfwords rather than
assuming little endian. Initial history is copied before input processing, and
final history comes from separate Go scratch: grouping its eight int16 stores
into four int32 words has no intervening reads/aliasing source or changed final
bits. Only the first four union words change; the tail is preserved. Tests use
native-memory union load/store oracles, signed extrema and roundtrip/tail checks.
Big-endian hardware is not validated (amd64/386/ARM64 only).

The reset suffix is ten contiguous 32-bit words, verified by a new layout and
raw-byte-clear comparison fixture. Prefix fields remain untouched and all suffix
stores precede child reset; short-channel/range/softclip behavior is unchanged.
Zero/nil/singleton and signed interior norm tests retain the old nonzero address
semantics without creating a forward slice for potentially unused EOF pointers.
No assertions, output/load/store dependencies, codec goldens or tolerances change.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with final repeated scoped/ordinary ARM fixtures.
This removes a real union reinterpretation and reset byte span, not a global
opaque-state scanning, extension EOF/GC or raw callback ownership repair.

Four cleanup-consolidation rounds route single-decoder, multistream and
projection factory init-error release through their existing concrete typed
destructors, then remove 23 obsolete generated `var _ unsafe.Pointer` markers
(and the resulting unused unsafe import in silk.go). Production unsafe.Pointer
references decrease 227→201: three duplicate factory allocator casts centralized,
23 metadata-only markers removed. uintptr remains 167. Marker removal does not
change a pointer representation or improve lifetime safety; the actual allocator
casts remain in the typed destruction boundaries.

Upstream create/destroy paths in the single/multistream/projection C sources
remain the reference. Prevalidation, size/allocation-before-init ordering,
init→release→error construction, nil/success behavior, one-base-only release and
public legacy signatures remain unchanged. Typed destructors perform exactly
the same allocator release without inspecting or resetting partially initialized
child states. Existing grouped create/destroy/allocation/validation/native and
scanned-owner fixtures remain unchanged. Single-decoder init error after valid
prevalidation remains only defensive/unforced coverage; no failure injection or
new opaque-allocation ownership claim is made.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with final repeated scoped/ordinary ARM fixtures.
No code arithmetic, assertions, goldens, tolerances or aliases changed. This
centralizes real boundary conversions and removes misleading lexical noise;
opaque allocation scanning, extension EOF/GC and raw callback lifetime remain
unresolved separately.

Four disabled-DRED lifecycle rounds canonicalize OpusDREDDecoder's state alias,
derive its existing 12-byte size from the concrete type, expose a typed factory
that calls the typed initializer without an integer allocation handle, and
expose typed destruction with magic invalidation before allocator release.
Public integer factory/destruction adapters remain. Production uintptr decreases
171→167; unsafe.Pointer increases 224→227 because explicit allocator/error-cleanup
and public-adapter boundaries are now visible. This is an internal integer-handle
removal, not a claim that every lexical pointer counter must fall.

Upstream disabled DRED layout/get_size/init/create/destroy in opus_decoder.c
remain the reference. Allocation→nil/error→init→cleanup/return ordering, size
uint32/uint64 conversion, loaded=0/arch=0/init magic, destruction magic-before-free,
nil free, return/error behavior and exported legacy signatures remain unchanged.
Grouped tests verify canonical alias/12-byte geometry across architectures,
allocation failure with nil TLS, registered typed creation/init with GC/stack
growth and typed release, plus enclosing guards and Go-only invalidation under
nil TLS (free is a shim no-op, not reading freed registered storage). The init-
failure cleanup is retained but unreachable in the current disabled init.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, followed by repeated scoped/ordinary ARM fixtures.
No goldens/tolerances changed. These are disabled DRED header/lifecycle fixtures,
not enabled native DRED model/SIMD parity or model lifetime coverage. The malloc
backing is still bytes (this state has only scalars), not a general scanned
allocator repair; extension EOF/GC and raw callback ownership remain unresolved.

Four bound-destination rounds separate opusMSDecodeBound from the explicit
legacy/format-erased opusMSDecodeNative adapter, capture standard multistream
int16/int24/float32 PCM as concrete pointers, capture projection integer PCM
with its typed matrix, and complete projection float capture. The shared active
core now accepts only a bound copy callback, with no opaque PCM destination
argument. Production unsafe.Pointer references decrease 233→224 (nine removed);
uintptr remains 171. Actual legacy custom callbacks/standard token binding still
use the explicit opaque adapter; raw closure-address ownership is not repaired.

Upstream multistream/projection decoder callback and copy functions remain the
reference. Validation/packet/child error order, positive-return gating, live
mapping/channel reloads, left→right→mono→unmapped copy order, nil-source clearing,
short-only clipping, matrix column-zero clear and accumulation order remain
unchanged. No staging PCM or numeric conversion changes; captured pointers keep
the same live aliases. Existing grouped format/guard/native/scanned-owner/GC
fixtures remain unchanged except the projection float callback-owner fixture
now binds its concrete destination instead of passing an unsafe destination at
each call; dropped matrix owner, stack growth, GC and nil-source clear persist.
The old opaque private adapter remains for compatibility fixtures and public
custom callback crossings, rather than reinterpreting those handles as a
fictional concrete destination type.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with final repeated scoped/ordinary ARM fixtures.
No goldens/tolerances changed; opaque allocation scanning and extension EOF/GC
remain outside this callback-binding improvement.

Four typed-binding rounds give multistream CTLs numeric uint child sizes/cursors
and reuse opusMSDecoderAt, bind single-decoder CTL children through the existing
typed frame helpers, select SILK ICDF entries through consumed byte prefixes,
and isolate disabled DRED initialization behind a concrete decoder pointer.
Production unsafe.Pointer references decrease 242→233 (nine removed); uintptr
171 replaces the previous 172. CTL changes centralize existing necessary layout
casts rather than eliminating allocation boundaries. SILK negative row offsets
retain a signed backward boundary. DRED's public integer adapter remains (with
an explicit uintptrescapes annotation); creation/destruction remain legacy.

Upstream multistream/single-decoder CTLs, decode_indices.c and DRED init in
opus_decoder.c were inspected. Size-call order, fixed8/uint32/word-width wrapping,
live coupled/stream reloads, child selection→next→dispatch, output/error order,
unconditional single-decoder interior derivation before dispatch, ICDF first/
residual row selection and assertion timing remain unchanged. Grouped tests add
zero/nil/singleton/final ICDF entry and signed interior selection with GC/stack
growth, plus disabled DRED loaded/arch/magic initialization and enclosing guards.
DRED has no model/SIMD initialization in this port; this Go-only header test is
not parity or model ownership coverage for an enabled native DRED build.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with final repeated scoped/ordinary ARM fixtures.
Goldens/tolerances unchanged. Opaque allocation scanning, extension EOF/GC and
raw callback ownership remain independent unresolved boundaries.

Four owner/numeric-consumer rounds keep the multistream validation packet owner
with a numeric word-width cursor, express the projection multistream displacement
as uint rather than pointer-valued uintptr, read forward band energies through
consumed float32 prefixes, and expose forward delayed-NLSF rate rows through
consumed byte prefixes. Production unsafe.Pointer references decrease 243→242;
uintptr decreases 173→172. Negative energy/rate offsets retain explicit signed
backward boundaries; their forward-path improvements do not reduce lexical
unsafe.Pointer counts.

Upstream opus_multistream_decoder.c/opus_projection_decoder.c, bands.c and
NLSF_del_dec_quant.c remain the reference. Validation/error/sample comparison
order and terminal-pointer suppression, fixed 8-byte geometry/uint32 wrapping/
int32 narrowing/word conversion, live energy reads and NLSF row selection/store
order remain unchanged. New grouped tests cover singleton energy backing,
signed interior energy, exact single and final NLSF row extents/capacity/identity,
and GC/stack growth. Negative helper geometry fixtures are Go-only, not new
coverage of upstream negative band or generated NLSF table indices.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with repeated scoped/ordinary ARM fixtures after
completion. No goldens, tolerances, assertions or aliases are weakened. Opaque
allocation scanning, extension EOF/GC and raw callback boundaries are not
claimed repaired.

Four CELT typed-consumer rounds select comb-filter tail pointers from existing
history/output slices (suppressing the unused zero tail), read coarse-energy
encoder probability pairs through consumed byte prefixes, preserve MDCT's typed
base for nonnegative/zero strides, and select the IIR forward history tail from
its existing output slice. Production unsafe.Pointer references decrease
247→243 (four removed); uintptr remains 173. Negative MDCT strides and N<order
IIR tails still use the necessary signed backward pointer boundary; lexical
counts do not capture these forward-path improvements.

Upstream celt.c, quant_bands.c, mdct.c and celt_lpc.c remain the reference.
Comb transition load/store order and rounded products, coarse-energy budget/
clamp/probability order, wrapped stride geometry/zero-stride store order, and
IIR's four-sample/tail sign convention and reversed live memory stores remain
unchanged. New grouped MDCT tests cover singleton prefixes at extreme signed
strides, forward/backward/zero owner identity with GC/stack growth. IIR tests
retain N=0 preceding-output behavior and add N=1<order backward history and
empty order-zero cases. Existing overlap-only comb and native energy/MDCT/IIR
fixtures and codec goldens remain unchanged.

All four rounds pass full amd64/386, ARM64/QEMU, scoped checkptr, native
comparisons, codec references and GC stress, followed by repeated scoped/ordinary
ARM fixtures. Genuine backward and opaque/public ABI boundaries remain; this
is not global opaque-state scanning, extension EOF/GC or raw callback repair.

Four numeric/typed-owner rounds replace the frame gain union pointer casts with
math.Float32bits/Float32frombits, obtain multistream right-channel PCM directly
from its Go audio owner, and advance down-FIR/IIR-FIR outputs through a shared
concrete int16 consumed-prefix helper. Production unsafe.Pointer references
decrease 252→247 (five removed); uintptr remains 173.

Upstream mathops.h/opus_decoder.c, opus_multistream_decoder.c and the two
resampler drivers remain the reference. Gain's existing float32 polynomial,
floor/underflow test, signed narrowing/wrap, exponent shift and sign mask remain
unchanged. Multistream only selects audio[1] after a positive child result and
retains callback ordering, live mappings and channel counts. Resampler output
views occur only when another block remains; zero count preserves identity,
including nil, and terminal/lone-remainder suppression is unchanged. Tests add
exact-sized 80-sample down-FIR and 322-sample IIR-FIR output backing and helper
zero/positive identity checks, retaining existing state/guard/chunk comparisons.
The actual i32/i16 resampler state union boundary is intentionally unchanged.

Each round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, followed by repeated scoped/ordinary ARM runs.
No goldens, tolerances, assertions or live aliases are changed. Round four's
initial native suite failed with `found pointer to free object`, with
TestExtensionNextAgainstC on the stack; retained at
/tmp/opus-round-iir-fir-typed-output-cursor.log. The complete retry and ARM repeats
passed. Typed storage and numeric bits here do not repair opaque allocation
scanning, extension EOF/GC or legacy raw callback ownership.

Four typed-view rounds replace the remaining ICDF8/ICDF16 pointer walks with
consumed-prefix indexing, the stereo deemphasis memory reinterpretation with
Go's typed slice-to-array-pointer conversion, and the down-FIR coefficient
byte offset with a typed int16 prefix. Production unsafe.Pointer references
decrease 256→252; uintptr remains 173. Public escape adapters and actual backward
pointer/history and opaque allocation boundaries remain unchanged.

Upstream entdec.c, celt_decoder.c and resampler_private_down_FIR.c remain the
reference. ICDF entry/load order and cached range/value state remain unchanged;
no table-wide guessed extent is formed. Stereo N<=0 still loads the coefficient
but consumes no memory/PCM, covered by nil-memory and coefficient-panic cases.
Down-FIR only exposes the row for a consumed valid-order interpolation; zero
work may provide only the two AR coefficients (new exact-sized fixture), and
invalid order still reaches the existing interpolation assertion. MAC/rounding,
state copies, batch/remainder behavior and output stores remain unchanged.

All rounds ultimately pass full amd64/386, ARM64/QEMU, scoped checkptr, native
comparisons, codec references and GC stress, plus final repeated ARM scoped and
ordinary fixtures. Goldens/tolerances unchanged. Round three's first two native
runs failed in the extension-test GC area: `found pointer to free object`, then
`s.allocCount != s.nelems && freeIndex == s.nelems` with
TestExtensionNextAgainstC on the stack. Both logs are retained under
/tmp/opus-round-deemphasis-typed-memory-view{,-retry}.log; the complete retry2
passed. Round four's initial native run also hit `found pointer to free object`
with TestExtensionNextAgainstC on the stack; retained at
/tmp/opus-round-down-fir-typed-coefficient-view.log. Its complete retry passed.
These failures are not repaired by the typed-view changes, and no global
extension EOF/GC, raw callback or opaque-state lifetime claim is made.

Four SILK API direct-pointer rounds remove redundant casts from control/rate
reads, VAD/LBRR/mid-side entropy calls, stereo/frame dispatch and channel-count
selection, and pitch/history finalization. decControl/psDec/channel_state/
psRangeDec already have concrete types; their fields and typed calls now use
those pointers directly. Production unsafe.Pointer references decrease 288→256
(32 removed); uintptr remains 173. The public integer escape adapter is unchanged.

Upstream dec_API.c and existing grouped whole/real-packet, resampler/control,
LBRR entropy, channel-storage/views, output-count/PCM and packet-state fixtures
remain the reference. Stereo-to-mono decision, sample-rate validation, LBRR
consumption, conditional frame coding, live channel/count/sample-rate reloads,
resampling/history, pitch-lag multiplication and final recovery flags keep their
original ordering and narrowing/wrapping. No storage/layout, assertion, aliases,
fixtures, codec golden or tolerance changes.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, with repeated scoped/ordinary ARM runs after
completion. This is direct-cast cleanup in the already typed active SILK API,
not a new opaque allocation scanning, extension EOF/GC or raw callback repair.

Four CELT direct-pointer rounds remove redundant casts from typed decoder
entry/lost dispatch, entropy/energy/allocation calls, synthesis/finalization,
and concealment fields/autocorrelation/finish. Concrete decoder/mode/entropy/
float32/byte pointers now pass directly; celt_glog is already a float32 alias.
Production unsafe.Pointer references decrease 392→288 (104 removed); uintptr
remains 173. Public escape adapters, actual allocation-boundary clear and stereo
memory reinterpretation remain unchanged.

Upstream celt_decoder.c order and existing scanned-state decode/concealment,
mode/history/energy views, nil/validation/error, prefilter/postfilter/deemphasis
and native/packet fixtures remain the reference. Live downsample/channel/arch/
period/pitch/frame-type loads, cached tell timing, coarse→fine→bands→anti-collapse
sequence, output/energy clear/store order and final range/error handling remain
unchanged. Rounded float32 products and ARM fusion fixes are untouched. No
storage/layout/algorithm, assertions, alias fixture, golden or tolerance changes.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, followed by repeated scoped/ordinary ARM runs.
This removes redundant casts from already typed active decoder paths, not a
new opaque allocator scanning, extension EOF/GC or raw callback fix.

Four SILK direct-field rounds remove redundant same-type unsafe.Pointer casts
from decoder core state/control reads, normal frame dispatch/history, lost-frame/
finalization state and PCM/control calls, and PLC concealment state reads/pitch
updates. Concrete *OpusT_silk_decoder_state, *OpusT_silk_decoder_control,
*OpusT_silk_PLC_struct, *OpusT_ec_dec and *int16 pointers are used directly.
Production unsafe.Pointer references decrease 463→392 (71 removed); uintptr
remains 173. Public escape adapters and real table/address boundaries remain.

Upstream decode_core.c/decode_frame.c/PLC.c field/operation order and existing
core scratch/output/control/history/whitening aliases, frame loss/finish/history,
PLC native reference and complete scanned-owner Opus frame fixtures remain the
reference. Live frame/LPC/subframe/recovery loads, assertion conditions/sites,
LTP scale narrowing, fixed-point multiply/wrap/shift order, random seed/pitch
drift/clamp and history→PLC/CNG/glue→lag→count ordering are unchanged. These
owners were already typed; no storage, aliases, fixtures or goldens change.

Every round passes full amd64/386, ARM64/QEMU, applicable scoped checkptr, native
comparisons, codec references and GC stress; repeated scoped/ordinary ARM runs
finish the batch. Goldens/tolerances remain unchanged. Active SILK/private frame
paths benefit from simpler typed code, but opaque allocator scanning, extension
EOF/GC and raw callback ownership are not newly fixed.

Four quant-all-bands direct-pointer rounds remove redundant casts from mode/
lowband/hybrid folding setup, dual-band dispatch, RDO candidate saves/dots and
restore/final mono/stereo dispatch. Private m/X_/X/Y/norm/norm2 pointers already
have the required *OpusT_OpusCustomMode or *float32 types; celt_norm aliases float32.
Call sites now pass those owners directly instead of round-tripping through
unsafe.Pointer. Production references decrease 502→463 (39 removed); uintptr
remains 173. Integer-valued public escape adapters are retained unchanged.

Upstream bands.c folding/save/restore ordering and grouped existing whole-quantizer
fixtures remain the reference. Spectrum/norm copy order, theta -1/+1/0 transitions,
dot/weight float32 evaluation, complete entropy/band snapshots, byte rollback
window, dist0>=dist1 tie rule, last-band suppression and nil/mono/dual/stereo
branch timing are unchanged. Existing heap-owner/alias/GC/stack-growth and actual
native quantization comparisons remain unchanged; no fixtures/goldens are weakened.
RDO is an encode-mode branch of this helper, not new active decoder RDO coverage;
opusccenc's independent implementation is unchanged.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, followed by repeated scoped/ordinary ARM runs.
Goldens/tolerances remain unchanged. This is redundant-cast cleanup, not a new
opaque-allocation scanning, extension EOF/GC or raw callback repair.

Four direct-pointer rounds remove redundant typed→unsafe.Pointer→same-typed
round trips from VAD setup/input, VAD analysis/output fields, allocation mode
reads, and private allocation/band-quantization entropy calls and RDO snapshots.
Pointers are passed/read directly as *OpusT_silk_encoder_state, *int16,
*OpusT_OpusCustomMode and *OpusT_ec_ctx. Go entropy enc/dec/ctx aliases have
identical structural types, as upstream entcode.h aliases the same ec_ctx.
Production unsafe.Pointer references decrease 556→502 (54 removed); uintptr
remains 173. Public ABI escape conversions are retained.

No pointer owners/layouts, assertion sites, field reloads, arithmetic order,
float32 rounding, cached tells or entropy snapshot/restore fields change. Grouped
existing VAD, allocation alias/scratch/mode, quantization and native parity tests
remain unchanged; a grouped entropy fixture additionally proves compile-time
ctx/encoder/decoder pointer identity, typed snapshot/restore and heap-buffer
retention after dropping original references, stack growth and GC.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons,
codec references and GC stress, followed by repeated scoped/ordinary ARM runs.
Goldens/tolerances remain unchanged. VAD remains standalone in opuscc, while
allocation/band entropy calls are active decoder paths. These owners were already
typed, so the batch simplifies casts rather than fixing opaque allocator scanning,
extension EOF/GC or raw callback ownership.

Four concrete-view rounds replace pulse-cache byte address casts with typed
byte prefixes, form forward cache rows as *byte prefix elements, replace the
resampler's next-output cast with a typed int16 prefix element, and read the
extension frame increment through a two-byte prefix. Production unsafe.Pointer
references decrease 559→556 (3 removed); uintptr tokens remain 173. These paths
already had concrete owners; this removes casts rather than inventing opaque
*byte stand-ins for unrelated libc handle/backend objects.

Pulse cache zero offsets return the original pointer (including nil); signed
negative interior offsets retain the necessary unsafe.Add boundary and the
existing backwards-offset fixture unchanged. Only forward consumed rows use
prefix indexing. Resampler views are formed only for a nonempty second phase;
zero output displacement remains an identity. Valid initialized rate states and
room for the produced second phase remain the upstream contract. Grouped driver
tests add five rate pairs at 1ms+1 sample, exact ceil-sized output plus guards,
stack growth/GC. The initial fixture incorrectly initialized 48k→8k in decoder
mode and hit the existing initialization assertion; using the required encoder
mode fixes that fixture without weakening the assertion (failed log retained).

Grouped extension tests add zero/nonzero frame increments, truncated increments
and frame overflow with output unchanged. Validation still precedes the two-byte
view; iterator EOF representation is unchanged and fixtures use padded backing.
All four rounds ultimately pass full amd64/386, ARM64/QEMU, scoped checkptr,
native cache/rate/resampler/extension comparisons, codec references and GC stress,
followed by repeated scoped/ordinary ARM runs. Goldens/tolerances remain unchanged.
Opaque allocator scanning, extension EOF/GC and raw callback ownership remain
separate; truly opaque pointer boundaries are not blindly changed to scalar types.

Four dormant compatibility-layout rounds change locale_t and timer_t aliases
to unsafe.Pointer, tm's timezone abbreviation to *byte, and musl __ptcb's callback/
argument/link members to unsafe.Pointer/unsafe.Pointer/*__ptcb. Actual Linux
libc typedefs, tm header and vendored modernc musl pthread.h were inspected.
Native locale/timer widths and host tm size/offsets are checked; __ptcb is checked
against an explicit C mirror of the musl record, not glibc's different cleanup
buffer. Exported alias types intentionally become pointer types. Production
opuscc uintptr tokens decrease 179→173 (6 removed).

Grouped layout tests drop original heap references, grow stacks/force GC, check
opaque handle/string/cleanup-chain ownership and clearing, and compare existing
Go record size/offsets to the old integer-word shape. The cleanup function-value
slot is explicitly a Go-only opaque payload fixture, not a C function address
or a repair of the raw legacy callback ABI. These layouts are unused by the
codec and no locale/timer/pthread libc operations are newly implemented.
The existing tm int64 gmtoff representation is unchanged on 386; host-native
C tm parity is measured on amd64, not claimed for that port's 386 long layout.

Final initial attempts hit a nested C block-comment preamble syntax error, then
a recursive anonymous type-alias build error. A line comment and a named __ptcb
record (same layout) resolve those errors; complete validation is rerun, with
the original failure logs retained.
Full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons, codec golden/
tolerance checks and GC stress ultimately pass each round, followed by repeated scoped/
ordinary ARM runs. No active decoder ownership coverage, opaque byte-allocation
scanning, extension EOF/GC or raw callback fix is claimed.

Four private decode offset rounds introduce numeric opusFramePCMAtBytes,
migrate frame zeroing/SILK/recursive PCM cursors, fade/redundancy/transition/gain
accesses, then native packet/FEC/PLC PCM dispatch and payload traversal. The
old opusFrameSilkPCM integer-word adapter is deleted once every private caller
and grouped fixture uses the numeric helper. All byte displacements retain the
old uint-width wrapping, multiply order and /4 truncation; nil/zero handling and
consumed-prefix view timing remain unchanged. Frame/payload owners were already
typed, so this is numeric clarity cleanup, not new lifetime coverage of a formerly
opaque owner. Public decoder escape adapters remain unchanged.
Production opuscc uintptr tokens decrease 209→179 (30 removed).

Grouped frame tests add nil/zero and heap-owner/alias checks for numeric PCM
views; existing transition/fade/chunk fixtures are updated without changing
assertions, data or goldens. Existing native tests extend unused negative-length
payload-offset suppression. Real SILK/hybrid/CELT, PLC/FEC, fade/softclip and
validation/state/range references remain unchanged. The third initial validation
attempt failed to compile because six existing fixtures still used the retired
helper; updating those fixture calls to the numeric equivalent resolved it,
followed by a complete successful validation rerun (original log retained).

Every round ultimately passes full amd64/386, ARM64/QEMU, scoped checkptr,
native comparisons and GC stress; repeated scoped/ordinary ARM runs finish the
batch. Codec goldens/tolerances are unchanged. No global opaque scanning,
extension EOF/GC or raw callback fix is claimed.

Four extension-output/multistream rounds replace bitstream- and frame-ordered
extension record address multiplication with typed consumed-prefix stores,
reuse opusAlignSize8 for the private multistream header/child strides, then
express multistream packet/child offsets as numeric uint instead of uintptr.
The uint change preserves native word-width arithmetic/wrapping on amd64/386/
ARM64; no addresses are stored in these numeric cursors. Child advance before
error gating and packet advance before child-return gating remain unchanged.
Production opuscc uintptr tokens decrease 221→209 (12 removed).

Grouped existing extension tests preserve live capacity reloads, prefix-sum
snapshotting, bucket/assertion/error ordering, partial writes, GC-visible payload
owners and guards; singleton output tests ensure only consumed records are
viewed, not an entire nominal larger capacity. Upstream extensions.c record
store/error order is unchanged. Frame counts must satisfy the upstream count_ext
contract; malformed negative cumulative output indices are not a supported array
access contract. Multistream tests retain exact-stride scanned child geometry,
add packet-owner retention through numeric offsets and zero/negative-length
terminal-offset suppression. Existing numeric alignment boundary tests remain.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native extension/
multistream comparisons, codec references and GC stress, followed by repeated
scoped/ordinary ARM runs. Goldens/tolerances are unchanged. Extension iterator/
zero-length EOF pointer representation is deliberately unchanged: scoped parser
fixtures use padded packet backing. This batch does not fix that separate EOF/GC
issue, opaque allocator scanning or raw legacy function-pointer capture.

Four mode/resampler rounds replace the mode log, pulse-cache index and band
boundary helpers' integer byte-address calculations with typed int16 consumed
prefix views, then change the resampler output byte displacement's cast from
uintptr to numeric uint. The latter preserves machine-word width/wrapping and
existing typed unsafe.Add behavior exactly; it is numeric cleanup, not an owner
or EOF fix. Production opuscc uintptr tokens decrease 225→221 (4 removed).

Existing grouped mode tests retain heap-owner GC/stack-growth coverage, all table
rows and signed values; singleton views now explicitly prove that nominal larger
band counts do not cause fabricated full-table extents. Existing resampler
four-rate-pair driver guards/history tests at 1/2/10/21ms now also grow the stack
and force GC before every call. The terminal remaining==0 guard, output offsets,
driver selection and load/store order are unchanged. Negative/out-of-range table
indices are outside the upstream table-access contract. Upstream rate/bands
index expressions and resampler.c were inspected before changing these helpers.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native mode/rate/
resampler comparisons, codec references and GC stress, followed by repeated
scoped/ordinary ARM runs. Goldens/tolerances remain unchanged. These owners were
already typed; no new opaque-allocation scanning, raw callback or extension EOF
safety is claimed.

Four entropy consumed-view rounds replace remaining integer address-offset
casts in front/back byte readers, front/back byte writers, 8/16-bit encoder ICDF
row loads, and final partial-byte OR with typed unsafe.Slice prefix indexing.
Only the prefix through the consumed element is formed; a claimed full storage
extent is not needed for individual byte operations. Exhaustion/collision checks
still precede views, counter mutations still precede the buffer load/store, rate
loads keep their existing current/previous caching, and final window/collision/
sticky-error/clear behavior is unchanged. Production opuscc uintptr tokens
 decrease 234→225 (9 removed). These are active entropy operations, though the
 buffer/table owners were already typed before this batch.

Grouped existing entropy tests add guarded front/tail sequences, empty/full
buffers with unchanged counters/bytes, single-byte consumed prefixes despite
larger claimed storage, counter-byte live aliases, and partial-tail empty/shared
byte behavior. The cached rate-row state alias is explicitly Go-only: upstream
C rereads the row after its intervening val store, whereas this port already
cached these entries before this refactor. It is not new C alias parity.
Existing native entropy snapshots/round trips, allocation and complete codec
references remain unchanged; every round passes full amd64/386, ARM64/QEMU,
scoped checkptr, native comparisons and GC stress, with repeated scoped/ordinary
ARM validation after completion. Codec goldens/tolerances are unchanged.
Negative/out-of-range symbols are outside the upstream table-access contract;
no global opaque scanning, raw callback or extension EOF fix is claimed.

Four numeric-size rounds simplify single decoder, multistream decoder, mapping
matrix and projection decoder size composition. The generated uintptr(0)+8
alignment expressions become a shared numeric opusAlignSize8 helper, preserving
this port's fixed 8-byte ABI geometry, uint32 add/divide/multiply wrapping and
final int32 narrowing. Header constants, child-size call order, invalid-input
checks, matrix multiply/threshold behavior, projection matrix-before-decoder
validation and int32 final sum/product order are unchanged. This removes numeric
uintptr noise, not pointer ownership defects: tokens decrease 242→234 (8 removed).

Grouped common tests cover alignment boundaries including negative/MinInt32/
MaxInt32, valid/invalid compositions and Go-only wrapping cases. Grouped native
validation tests compare all four size APIs with actual upstream C, including
bounded negative matrix dimensions accepted by the original size routine and
matrix capacity limits. Cases that would overflow signed C intermediate arithmetic
remain Go-only and are not represented as C parity. Upstream align uses platform
union alignment; this port intentionally retains its existing 8-byte alignment
on 386 instead of substituting host sizeof/alignment.

Every round passes full amd64/386, ARM64/QEMU, scoped checkptr, native comparison,
codec golden/tolerance and GC-stress validation, with repeated scoped/ordinary
ARM runs after completion. No assertion/validation was strengthened, no goldens
changed and no opaque scanning/extension EOF/raw callback repair is claimed.

Four pointer-layout rounds type repacketizer frames and paddings as [48]*byte,
unify the duplicate OpusRepacketizer alias with the canonical typed layout, then
type FFT architecture linkage as *OpusT_arch_fft_state and its opaque private
payload as unsafe.Pointer. Actual upstream opus_private.h/kiss_fft.h define these
fields as pointers. Sizes/offsets stay unchanged (native repacketizer all-field
and architecture-state layout comparisons added; existing FFT state offset tests
remain). Exported Go field types intentionally change to typed/scanned ownership.
Production opuscc uintptr tokens decrease 248→242 (6 lexical tokens removed,
including array declarations representing many pointer slots).

Grouped layout tests store heap-backed frame/padding arrays in all 48 slots,
drop original slice references before stack growth/GC, verify the canonical alias
retains both owners, and verify FFT state→architecture→opaque Go payload retention
and clearing. Full amd64/386, ARM64/QEMU, applicable scoped checkptr, native
layout/FFT/codec comparisons and GC stress pass every round with unchanged
codec goldens/tolerances, plus repeated scoped/ordinary ARM runs. Repacketizer
operations are not implemented/used by the opuscc decoder and architecture FFT
backend use is disabled: these are dormant-layout ownership improvements, not
additional active decode coverage. No opaque allocator scanning or extension
EOF/raw callback fix is claimed.

Four small CELT metadata rounds change the dormant mini complex/real FFT cfg
aliases from uintptr to the corresponding state pointer types, remove unused
function-name pointer conversions/storage at four FFT assertion sites, and delete
an unused uintptr temporary plus unused generated numeric locals/blank-use tuple
from quant_partition. Upstream mini_kfft.c defines both cfg types as state
pointers; active allocation/transform functions were already typed. These alias
changes intentionally update the exported Go type aliases, not allocator layouts.
libcshim.X__assert_fail ignores its final function-name argument, so supplying
zero leaves existing assertion condition, expression/file/line and panic text
unchanged. No assertion was removed or weakened.

Grouped FFT tests exercise both typed cfg holders after GC/stack growth using
nil TLS, and explicitly recover/check diagnostics for all four failure sites.
Existing allocation/stride/layout/transform native comparisons and complete
quantization/packet/concealment references remain unchanged. Every round passes
full amd64/386, ARM64/QEMU, scoped checkptr, native comparisons and GC stress with
unchanged codec goldens/tolerances; repeated scoped/ordinary ARM runs complete
validation. Production opuscc uintptr tokens decrease 255→248 (7 removed).
This is a small duplicate-metadata/API-type cleanup, not additional active decoder
ownership coverage or a repair of opaque scanning/raw callback/extension EOF.

Four VAD analysis rounds introduce typed encoder/input ownership behind the
public escape adapter, replace integer sample-buffer addressing with typed slices
and numeric band offsets, allocate the reusable decimation/energy samples in Go,
then remove obsolete TLS cursor setup/snapshot/restore and dead integer locals.
The private silkVADAnalysis path has no pointer-valued uintptr or pseudostack
operation. Production opuscc uintptr tokens decrease 273→255 (18 removed).
Assertions retain their original order/sites; live frame-length reads, overlapping
filter-bank passes, reverse HP differentiation, per-store int16 narrowing,
energy accumulation/saturation, noise update and output-field ordering are
unchanged. This standalone opuscc VAD analysis is not an active decoder call;
opusccenc retains its separate implementation.

Grouped existing VAD analysis tests retain original field goldens, add typed
heap-state/input lifetime checks, and complete nil-TLS/checkptr paths at lengths
80/120/160/240/320 across four calls with stack growth/GC/input guards and
assertion-before-mutation coverage. A native bridge invokes actual upstream
silk_VAD_GetSA_Q8_c with a numeric VAD state image and compares the entire VAD
state plus speech activity, tilt and all quality bands for the same lengths and
four consecutive calls. Full amd64/386, ARM64/QEMU, applicable scoped checkptr,
native comparisons, codec references and GC stress pass each round with unchanged
goldens/tolerances, plus repeated scoped/ordinary ARM runs after completion.
The final initial validation attempt again hit runtime 'sweep increased allocation
count' in the ordinary native suite, with TestExtensionNextAgainstC on the stack;
this is retained in the log and not attributed to VAD or claimed globally fixed.
A complete rerun and VAD-specific native repetitions pass. Opaque allocation
scanning/raw callback/extension EOF issues remain separate.

Four delayed-decision NLSF quantizer rounds replace integer-address input
loads with typed live x/weight/predictor/rate-index views, use typed nine-byte
rate-row views, type index output/final adjustment stores, then introduce the
fully typed private silkNLSFDelayedQuant entry behind the public escape adapter.
Fixed Go scratch arrays, candidate order/sorting, signed narrowing, wrapping
MACs, row snapshots, disabled original assertion expressions and late output
store order remain unchanged. There is no TLS scratch/pin/private escape
annotation. Production opuscc uintptr tokens decrease 292→273 (19 removed).
This standalone quantizer is not an active decoder call; opusccenc has its own
unchanged implementation.

Grouped parameter tests cover nil TLS/whole-private-entry checkptr, order10/16,
Go inputs/rate tables/guarded index output, stack growth/GC, deterministic state
and a Go-only output/input alias proving late stores. Native comparison tests
invoke the actual upstream silk_NLSF_del_dec_quant for 200 seeded bounded
fixtures, comparing exact RD return/index bytes/guards without altering goldens.
The exported comparison bridge is typed; public legacy ABI remains available.
Full amd64/386, ARM64/QEMU, applicable scoped checkptr, native comparisons,
codec references and GC stress pass every round with unchanged tolerances, plus
repeated scoped/ordinary ARM tests after completion. Existing opaque allocation,
raw callback and extension EOF boundaries are not claimed repaired.

Four CELT cleanup rounds delete the unused CELT_PVQ_U_ROW integer-address
array (active PVQ already uses numeric offsets), unused comb/prefilter adapters,
unused deemphasis/synthesis integer channel-array adapters, and unused inverse
MDCT/anti-collapse adapters. Repository-wide Go reference searches found only
the deleted definitions; active typed helpers/numeric row data and consumers are
unchanged. Production opuscc uintptr tokens decrease 333→292 (41 removed).
The obsolete exported Go row variable is removed; public decoder legacy escape
adapters remain. This deletes duplicate representations rather than introducing
new arithmetic, allocations, validation or decode coverage.

Existing native PVQ/filter/MDCT/synthesis/anti-collapse comparisons, complete
scanned-state pointer tests and packet/concealment references cover the remaining
implementations. Full amd64/386, ARM64/QEMU, scoped checkptr, GC stress and
unchanged codec goldens/tolerances pass each round; repeated scoped/ordinary ARM
runs and repeated native comparisons complete the batch. The synthesis cleanup's
first validation attempt hit a runtime 'sweep increased allocation count' in the
ordinary compareopus suite; a complete rerun passed. This is recorded rather than
claiming every first attempt passed or that deleting unused code globally repairs
legacy unsafe fixtures/opaque allocation scanning.

Four SILK LTP ownership rounds replace all four [3]uintptr pointer tables with
GC-scanned typed arrays: bit-cost/ICDF/vector-gain [3]*byte, and coefficient
[3]*[LTP_ORDER]int8 row pointers. The active index decoder forwards the typed
ICDF directly, and the parameter decoder consumes the typed coefficient owner
using fixed 8/16/32-row geometry. Signed int8 row snapshots, per-subframe index
reloads, tap order and int32 shift→int16 Q14 narrowing remain unchanged. Bit-cost
and vector-gain tables have no active opuscc decoder consumer; their migration
removes dormant integer-address owners rather than claiming new decode coverage.
The separate opusccenc copies are unchanged.

Grouped existing parameter-test files rebind each table slot to a heap clone,
drop the original clone slice, then force stack growth/GC before consumption and
restore table slots after the test. ICDF entropy symbols/all context fields match
an independent original byte copy; every 8/16/32 vector row decodes all four
subframes with signed Q14 coefficients checked. Native tests copy all 56 bit
costs, 56 ICDF entries, 56 vector gains and 280 signed coefficient bytes from the
actual upstream libopus tables and compare exactly. Existing renamed upstream
indices/parameters source comparisons, whole packet/concealment references,
full amd64/386 and ARM64/QEMU tests, scoped checkptr and GC stress pass each
round with unchanged codec goldens/tolerances, plus repeated scoped/ordinary ARM
runs. Production opuscc uintptr tokens decrease 349→333 (16 removed); the four
exported Go table variables intentionally acquire typed-pointer element types.
Public legacy decoder escape adapters are retained; opaque allocations/raw custom
callback addresses and extension EOF remain separate unresolved boundaries.

Four uintptr-reduction rounds type projection int16 and int24 forwarding with
captured matrix owners, delete all three unused private integer projection
callbacks and get_multistream_decoder_legacy (plus the redundant integer callback
bridge), and remove integer user-data from opusMSChannelCopy/opusMSDecodeNative
and every standard/projection private callback/call site. All standard and
projection format wrappers now call the typed multistream core; only public
legacy signatures and the custom callback escape binder retain integer data
addresses. The binder captures legacy user-data once at the public boundary,
without piping it through the decode loop. Standard callbacks still ignore it;
custom callbacks still receive its exact original value. Numeric byte offsets
remain numeric rather than being misclassified as pointer-valued uintptr.

Compared to this batch's initial tree, common.go uintptr tokens fall 207→180;
all opuscc production .go files (excluding *_test.go) fall 376→349, measured
consistently with rg -o '\\buintptr\\b'. This is a lexical reduction, not a proof
that all remaining integer words are unsafe or that all ownership is solved.
Projection int16 retains OPTIONAL_CLIP (tested with amplified CELT output),
projection int24 retains zero clipping, and both retain pointer-derivation,
callback clear/accumulation, rounding and validation ordering. Existing grouped
full scanned-state/nil-TLS/checkptr tests cover normal/PLC/FEC, independent
existing-entry PCM/state/count parity, errors, guards and GC/stack growth. Custom
callback retention/argument tests verify the captured user-data value after GC;
standard binding tests verify nonzero user-data is still ignored. No C golden,
tolerance, assertion, public ABI or custom callback contract changed.

Every round passes full amd64/386, ARM64/QEMU, applicable scoped checkptr, existing
native mapping/callback/codec comparisons, multistream C references and GC stress,
with repeated scoped/ordinary ARM runs after completion. Opaque byte-backed
allocation scanning and the raw custom function-address boundary remain separate
legacy issues; no global GC-safety/direct macOS CI claim is implied.

Four format-wrapper rounds introduce typed private multistream float/int16/
int24 entries, each forwarding directly to opusMSDecodeNative with a typed
standard callback (no integer function-address binding), followed by a typed
private projection float entry. Public format signatures/escape adapters remain.
Clipping is unchanged: OPTIONAL_CLIP for int16, zero for float/int24/projection
float; no additional validation, gain, quantization or sample-count operation is
introduced. Projection still derives multistream then matrix views at the original
forwarding point, but a Go closure captures the typed matrix owner instead of
passing its address through integer callback user-data. Matrix/source/destination
loads and clear-before-accumulation order remain in the existing typed leaf.

Grouped whole wrapper checkptr/nil-TLS tests use scanned exact-stride two-stream
state, reordered/duplicated/muted channels, real SILK/synthetic active CELT packets,
normal/PLC/FEC, guarded outputs, stack growth/GC and independent existing-native
entry PCM/count/state parity. The int16 fixture also raises gain to exercise the
retained soft-clipping flag. Projection float tests use a scanned header/matrix/
multistream/children composite with asserted offsets, nontrivial 4x4 coefficients,
full active normal/PLC/FEC parity and invalid-packet/minimum-frame ordering.
A separate matrix callback drops the original composite owner before GC/stack
growth, then verifies captured-coefficient output and nil-source clearing.
Existing native mapping/callback comparisons and enclosing MS C references,
full amd64/386 and ARM64/QEMU tests, scoped checkptr and GC stress pass each round
with unchanged codec goldens/tolerances, plus repeated scoped/ordinary ARM runs.
Projection integer wrappers/custom callback user-data and opaque byte-backed
allocator scanning remain separate legacy boundaries; no global GC-safety or
direct macOS CI claim is implied.

Four enclosing multistream rounds replace the decode sample-rate varargs CTL
with its existing typed equivalent, use Go Fs/packet-offset scalar storage,
remove all multistream decode pseudostack initialization/snapshot/restores, then
introduce opusMSDecodeNative with typed state/payload/output/callback parameters
behind the public escape adapter. Child and packet traversal remain numeric;
payload views are formed only when consumed, and the validator no longer forms
an unused final EOF view. Validation still precedes frame-size rejection; sample
rate/query assertion, frame cap, allocation, length/packet validation, child
advance-before-error, packet-offset advancement-before-return, live mapping/
channel bounds and muted-channel timing are retained.

Full active standard-callback private-entry checkptr/nil-TLS fixtures use a
scanned exact-stride two-coupled-stream composite (computed decoder/SILK/CELT
padding/tails, with size assertions on all architectures), self-delimited real
SILK packets, reordered/duplicated/muted channels, float/int16/int24 destinations,
normal decode, 120ms PLC/FEC, guards, stack growth/GC and unchanged C-derived
sample/range expectations. Argument failures leave PCM/child state untouched;
ordinary public adapter tests cover capped PLC and nil TLS. Existing enclosing
multistream C references and native leaf comparisons, full amd64/386,
ARM64/QEMU, scoped checkptr and GC stress pass each round with unchanged codec
goldens/tolerances, plus repeated scoped/ordinary ARM runs. No multistream decode
pseudostack or opaque scalar scratch remains. Legacy callback user-data remains
an explicit uintptr contract only for custom/projection fallbacks; public format/
projection wrappers and opaque byte-backed state scanning are separate from
these tested scanned-state standard-callback paths.

Four multistream-consumer rounds type child native dispatch/packet-offset
forwarding, replace integer child addresses with numeric byte traversal and typed
consumed child views, move the reusable 2*frame_size float PCM scratch into Go
storage, and bind typed standard output callbacks/source views. The header/child
alignment and coupled/mono stride calculations, advance-before-error order,
self-delimited flag/live stream-count read, packet-offset reset/advancement,
positive return gating, mapping iteration and muted-channel timing are retained.
The frame cap/Go allocation remains after the original sample-rate query and
before packet validation. Only used child addresses are formed; terminal walk
positions stay numeric. Source left/right/mono pointers and muted nil source now
stay typed through standard float/int16/int24 callbacks, with unchanged rounding,
strides and live copy loops. A bound Go function retains the custom callback
owner; arbitrary/custom and projection callback fallbacks preserve the public
legacy user-data/uintptr contract through explicit escape bridges.

Grouped scoped checkptr/nil-TLS fixtures cover native child normal/PLC and exact
self-delimited packet offset, scanned MS-interior child lifetime after GC/stack
growth, Go stereo scratch/typed consumers, all standard callback bindings,
guards/strides/muting and custom fallback argument/owner retention. Ordinary MS
C-reference/adapter fixtures exercise the enclosing loop; full amd64/386,
ARM64/QEMU, native comparisons, GC stress and unchanged codec goldens/tolerances
pass each round, plus repeated scoped/ordinary ARM tests. The initial scratch
fixture attempted the raw legacy float callback under checkptr and correctly
failed at its uintptr-to-pointer representation. That boundary is kept as an
ordinary legacy fixture; typed scratch consumers have scoped coverage, and the
following callback round bypasses that boundary for all three standard formats.
A forced-GC custom-closure probe also exposed that converting a movable Go
closure through __ccgo_fp does not preserve its original owner/lifetime. Added
opusMSBindLegacyCopy to accept and capture a Go function directly; its scoped
owner-retention fixture uses that typed path, while public integer callback
addresses retain the legacy requirement for stable caller-owned function storage.
The raw __ccgo_fp closure boundary is not claimed safe or repaired globally.
No checkptr suppression, pin or golden change was used. This is not a complete
multistream/projection ownership or checkptr claim: enclosing payload/output/
state entry, CTL varargs scratch, pseudostack setup/restores and projection matrix
user-data remain legacy, as does opaque byte-backed state scanning.

Four integer-entry rounds remove int16 pseudostack setup/snapshot/restores,
remove the corresponding int24 operations, then introduce fully typed private
opusDecodeInt16 and opusDecodeInt24 entries behind the original public escape
adapters. Both private paths accept nil TLS and use Go float scratch with typed
native forwarding and conversions; no integer owner, pin, private escape
annotation or cursor operation remains. Frame-size rejection still precedes state
access; packet sample-count validation precedes channel assertions; duration
trimming, assertion sites, clipping flags, architecture read, positive-return
conversion gating and post-decode live channel count loads are unchanged.

Grouped full active integer-path checkptr tests use scanned composite decoders,
Go packets/output arrays, nil TLS, forced GC/stack growth, mono/stereo normal/
PLC/FEC/multiframe/empty payload cases and SILK/hybrid packets. API8/12/16/24/48kHz
and guarded int16/int24 output counts/range goldens pass; malformed packet
validation still precedes the channel assert, and failed decode leaves caller
PCM unchanged. Ordinary public-adapter fixtures retain untouched TLS sentinels.
Existing native RES2INT24 and float2int16 comparisons, native codec/frame
references, full amd64/386 and ARM64/QEMU tests, GC stress, codec goldens/tolerances
and repeated scoped/ordinary ARM runs pass every round. All three private format
entries and their frame/native decoder consumers are now typed. This is coverage
of tested scanned-state paths, not opaque byte-backed allocation scanning, direct
macOS CI, public legacy adapter removal or multistream/projection outer ownership.

Four decoder-format rounds type the private float entry, int24 PCM conversion,
int16 temporary float storage and int24 temporary float storage. opusDecodeFloat
forwards directly to the complete typed native decoder, with frame_size<=0 checked
before state access; the public float escape adapter remains. The int24 loop uses
typed live views, the original float32 scale/evaluation and lrintf→int32 narrowing.
Native tests invoke the upstream RES2INT24 macro with arch.h/float_cast.h; ties,
signs, guards and valid representable C values match, with effective-type-invalid
float/int live aliases explicitly Go-only.

Both integer APIs retain their original packet-duration trimming, channel
assertion, return gating, clipping flag (int16 only) and conversion count reload.
Their float temporaries are independent Go storage, with direct typed native
forwarding/typed conversions and public uintptrescapes tracking. The legacy
integer wrapper scratch alignment/capacity allocation blocks disappear, but
setup/snapshot/early/final cursor restores remain. Grouped nil-TLS scoped checkptr
scratch-consumer fixtures and whole float SILK/hybrid reference goldens run with
GC/stack growth; ordinary integer wrapper fixtures retain untouched TLS cursor
sentinels and guards. Full amd64/386, ARM64/QEMU, native comparisons, references,
GC stress, unchanged codec goldens/tolerances and repeated scoped/ordinary ARM
runs pass each round. This is not whole integer-wrapper checkptr/nil-TLS proof;
those entries/cursors, opaque allocation scanning and other outer APIs remain.

Four decode-native entry rounds retain a typed decoder, payload/numeric byte
cursor, packet-offset output and PCM owner. opusDecodeNative is fully typed with
no private escape annotation; the public Opus_opus_decode_native signature and
uintptrescapes adapter remain. Recursive PLC/FEC calls forward typed owners.
Unused DRED arguments remain accepted by the public adapter but are not forwarded
in this build, matching the upstream disabled-deep-PLC branch. Payload cursor
advancement preserves native unsigned wrapping, and only length>1 consumed views
are formed: zero/tiny frames still select loss, with no unused EOF pointer.
Parser output timing, metadata commit/error order, duration rollback, frame flags,
soft clipping and live channel reads are unchanged.

The first 386 full checkptr run caught an exact-sized empty-packet EOF pointer
in the existing public parser's padding output. Decode-native now asks the parser
for numeric consumed/payload/size descriptors without requesting padding, derives
padding length numerically, and forms a padding view only for nonzero lengths.
Caller packet-offset writes still occur in the parser at their original point;
when the caller omits that output a private numeric slot captures consumption.
Actual C descriptor comparisons verify sizes/offsets/nonempty padding. The public
parser retains its C EOF-pointer contract, so this fixes the active consumer
without weakening empty-packet tests or claiming extension EOF is globally solved.

Complete active private decode-native checkptr now runs with nil TLS and scanned
composite state/Go payload, PCM and offset storage. Grouped fixtures cover
mono/stereo SILK/hybrid/CELT, single/multiple/empty/padded packets, normal/PLC/FEC
recursion, validation gates, self-delimited packet-offset guards and too-small
output errors without premature metadata commit. Real SILK/hybrid packets reuse
unchanged reference PCM FNV/range goldens. A 10ms self-delimited test initially
used a 40ms TOC, correctly returning buffer-too-small; its TOC was corrected, not
the validation. Full amd64/386 and ARM64/QEMU, native comparisons, whole reference
fixtures, GC stress and codec goldens/tolerances pass each round, with separate
repeated typed-path checkptr and ordinary ARM golden runs. This covers the tested
typed native decoder paths, not opaque byte-backed allocation scanning, direct
macOS CI, exact-sized extension EOF consumption or remaining float/int decoder,
multistream/projection outer ownership.

Four decode-native consumer rounds use the typed packet parser and all three
typed frame dispatches: PLC, FEC suffix and ordinary packet sequence. Parser toc,
48 sizes, payload/packet offsets and padding outputs now use Go slots directly;
padding retains a typed byte owner and extension-ignore still clears it before
the parser-error check. Frame dispatch helpers forward typed decoder/data/PCM
owners and derive only consumed numeric offset views, keeping int32 product
narrowing/uintptr stride wrapping, PLC/FEC remaining sizes and literal flags.
Duration rollback, metadata commit timing, frame-size assertions, sequential
payload advancement and soft-clip ordering are unchanged. The public native
boundary receives uintptrescapes because it still accepts legacy addresses.

Grouped scoped checkptr/GC fixtures cover padded packet descriptors/extension
ownership, mono/stereo PLC consumption windows, untouched FEC prefixes, and an
actual SILK packet dispatched at a nonzero PCM offset using unchanged scalar-C
FNV/range goldens. Full amd64/386 and ARM64/QEMU tests, actual packet/CTL/helper
native comparisons, existing whole native/frame references, codec baselines and
GC stress pass each round, plus separate repeated ARM typed-consumer checkptr and
ordinary golden runs. There are no legacy parser/frame calls left in decode-native,
but its enclosing input/output/state/payload cursors and recursive calls remain
integer representations; this is not enclosing decode-native checkptr coverage.
Opaque allocation scanning and exact-sized extension EOF remain separate risks.

Four private frame entry rounds retain the decoder as a typed pointer, the
payload as a typed byte pointer, PCM operations as numeric typed views, then PCM
as a typed pointer at the entry. opus_decode_frame remains the uintptrescapes
legacy adapter; opusDecodeFrame is entirely typed, has no escape annotation,
pins or TLS cursor operations, and recursively forwards typed owners. The long
PLC loop accumulates only a numeric byte offset, constructing views before each
consumed call instead of an unused terminal EOF pointer. Live channel bounds,
store/load order, unsigned stride wrapping, gain rounding, frame validation,
length<=1→nil payload gates, entropy tells and final range XOR are unchanged.

Complete active private frame checkptr now runs on scanned composite decoder
storage with Go payload/output arrays and nil TLS. Grouped fixtures cover
mono/stereo CELT LM0..3 packet/loss sequences, 40ms recursive PLC, short/10/20ms
SILK PLC, actual SILK normal/FEC payloads, validation gates and untouched TLS
cursors. Real SILK NB and hybrid FB stereo packets reuse the unchanged frame
reference FNV/final-range goldens (SILK values scalar-C-derived, hybrid PCM Go
golden as documented in the original reference). These run with forced GC/stack
growth inside scoped checkptr, separately from the old malloc/pseudostack golden
fixtures. Full amd64/386 and ARM64/QEMU, native comparisons, existing transition/
redundancy/gain references, GC stress, codec goldens/tolerances and ARM repeated
typed-path/ordinary golden runs pass every round. This covers tested typed frame
paths, not direct macOS CI, opaque byte-backed embedded-pointer scanning or the
still-integer Opus decode-native entry, packet descriptors and enclosing owners.

Four frame pseudostack cleanup rounds remove scratch initialization, all six
early-return cursor restores, normal-return restore, then snapshot/metadata
allocation and dead temporaries. No frame pseudostack/TLS cursor access remains;
codec work/validation/recursive PLC/CTL/final-range/error ordering is unchanged.
Scratch initialization and cursor metadata side effects are intentionally gone.
The legacy uintptrescapes entry still forwards integer input/output arguments.

Whole ordinary frame fixtures now exercise nil TLS on scanned composite storage:
mono/stereo CELT LM0..3 packet→PLC→PLC→packet, recursive 40ms PLC, short/10/20ms
SILK PLC, validation gates and untouched TLS sentinels. The test composite includes
explicit computed 8-byte decoder/SILK padding: the initial 386 geometry assertion
caught native alignment vs Go's 4-byte struct alignment, which was fixed rather
than skipped. Packet pointers stay typed across forced stack growth/GC until the
annotated call boundary. Existing whole native comparisons, frame references,
codec baselines/goldens/tolerances, full amd64/386 and ARM64/QEMU tests and GC
stress pass every round. ARM repeated scoped typed-consumer checkptr and ordinary
whole-frame/golden repetitions are run separately: whole legacy frame uintptr
entry/PCM operations are not claimed to pass checkptr. Opaque byte-backed scanning
and outer decode-native ownership remain unresolved; next work can migrate the
frame entry and its payload/output operations to typed views.

Four frame scratch rounds replace SILK short-frame PCM, CELT transition PCM,
SILK transition PCM and redundant audio pseudostack allocation blocks with
independent Go float32 storage at the original allocation points. Zero sizes
produce nil unused storage. Short SILK frames retain the 10ms temporary then copy
only the requested samples. Transition size selection and redundancy suppression
are unchanged; recursive concealment precedes the same validation/CTL points,
and prefix-copy/fade and final-range ordering remain live. Typed transition and
redundancy views use numeric offsets and retained Go owners. The four legacy
cursor alignment/capacity blocks disappear; frame setup/save/restores remain.

The legacy opus_decode_frame boundary now has go:uintptrescapes: recursive calls
pass the Go-owned transition output through that still-integer entry, so escape
tracking is required before removing its surrounding legacy adapters. This is a
lifetime fix, not a checkptr workaround or a full typed-entry claim. Grouped nil
TLS/checkptr/GC scratch-consumer fixtures cover mono/stereo CELT concealment and
redundancy, short SILK truncation, independent transition owners, prefix copies
and fade guards. Existing actual C/transition/redundancy frame references, codec
baselines/goldens/tolerances and full amd64/386, ARM64/QEMU, GC stress pass every
round, with repeated typed-consumer checkptr and separate ordinary ARM goldens.
Opaque byte-backed scanning and enclosing frame input/PCM/cursor ownership are
still separate. Go zero initialization replaces C's uninitialized scratch only
where consumers already initialize the active data before it is read.

Four further Opus-frame owner rounds retain the SILK interior as a typed
pointer, represent SILK PCM with a typed base/current view plus numeric byte
offset, call silk_Decode directly with typed control/entropy/PCM/sample-count
arguments, then retain the CELT interior as a typed pointer throughout dispatch
and CTLs. Interiors use unsafe.Add on the retained decoder and the original live
signed offsets at the original derivation point. The PCM offset preserves native
uintptr stride wrapping; only consumed loop views are formed, not an unused
terminal EOF pointer. PLC error clears still reload the live channel bound and
store in the original order. Entropy tell gates read the typed context fields at
the original three points, with unchanged cached values/rounding/precedence.

All frame-local Pinner slots are now removed: dec and silk_frame_size no longer
cross integer interfaces, and the runtime import is gone. Public legacy adapters
remain for other callers. Grouped scanned composite owner fixtures exercise both
interiors, numeric PCM chunk views/guards, SILK PLC dispatch with typed count and
untouched entropy, and active CELT decode/range/mode retrieval. Scoped checkptr,
GC/stack growth, unchanged whole native SILK/CELT comparisons, enclosing frame
references, full amd64/386 and ARM64/QEMU tests and codec baselines pass each
round, with repeated ARM typed-consumer and separate ordinary golden runs.
Removing pins is not an enclosing frame GC-safety/checkptr proof: the entry,
payload/PCM/transition/redundancy owners, pseudostack storage/cursors and fades
still include legacy representations. Byte-backed decoder allocation scanning
also remains separate; these fixtures deliberately use scanned composite storage.

Four Opus-frame CELT CTL rounds replace all twelve legacy varargs calls with
the existing typed CTL interface: six band/channel setters, two resets, three
final-range outputs and mode retrieval. Setter values, reset/start-band order,
assertion sites and the deliberately ignored main range-query return are
unchanged. Redundant range outputs retain typed local pointers; main range output
uses the actual outer decoder field rather than byte offset96. The retrieved
mode and fade window are typed owners. Unused va_list scratch disappears, along
with redundant_rng, celt_mode and va Pinner slots; only dec and silk_frame_size
remain pinned for their legacy SILK crossings.

Grouped scoped checkptr/GC tests cover setter validation/no partial writes,
reset barriers/energy-log initialization, unsigned range extremes/live field
aliases, mode retrieval, retained heap window ownership and fade guards. Existing
actual-C custom CTL comparison fixtures and transition/frame goldens remain
unchanged; full amd64/386, ARM64/QEMU, native comparisons, GC stress and codec
baselines pass each round. Repeated ARM typed-consumer checkptr remains separate
from ordinary enclosing Opus-frame goldens. This does not yet make frame decoder
interiors, PCM/redundancy/fade offsets or remaining SILK uintptr arguments typed,
and opaque allocation scanning remains unresolved.

Four Opus-frame CELT dispatch rounds forward typed decoder, payload, PCM and
entropy owners to the internal decoder for the main/FEC call, CELT→SILK redundant
frame, hybrid→SILK silence frame and SILK→CELT redundant frame. Redundant packet
suffixes use numeric byte slices; length<=1 forms no unused interior and still
selects loss through the original length predicate. Main FEC keeps nonzero→nil
payload selection with the original length and provided entropy. Silence uses
the live two-byte Go array, retains accumulation forwarding and no longer needs
its frame Pinner slot. The other five legacy local pins remain. Existing CTL
start-band/reset/range queries and redundancy fade ordering are unchanged.

Grouped full typed-boundary checkptr/GC tests compare dispatch with direct CELT
decoding for normal/FEC/loss, packet suffixes, silence accumulation, PCM guards,
reset/redundancy state and final range; native suffix tests use scalar C window
copies. Existing transition/redundancy frame references and native codec
comparisons preserve all goldens/tolerances, with full amd64/386 and ARM64/QEMU
validation each round and separate repeated typed-path/ordinary golden runs.
These typed consumers do not prove enclosing Opus-frame checkptr: decoder offsets,
PCM/redundancy allocation, varargs CTL/fades and other integer owners remain.
Opaque byte-backed embedded-pointer scanning is still a separate blocker.

Four outer CELT pseudostack cleanup rounds remove scratch initialization,
lost-frame cursor restore, normal-frame cursor restore, then the remaining cursor
snapshot/metadata allocation and dead temporaries. Decoder work remains ordered:
validation/frame-size selection, packet/PCM checks, concealment or entropy decode,
deemphasis, packet finalization and terminal error handling. The public uintptr
escape adapters remain, but the private celt_decode_with_ec_dred entry and all its
active consumers now require no pseudostack storage and accept nil TLS on valid
paths. Scratch/cursor initialization side effects are intentionally eliminated.

Full active decoder checkptr now runs on scanned decoder-tail storage with Go
packet/PCM/entropy buffers, mono/stereo LM0..3 packet→PLC→PLC→packet sequences,
forced GC/stack growth, PCM guards, validation gates, and provided-vs-local entropy
state/history/PCM parity. A typed TLS cursor sentinel is untouched through both
normal and lost decoding. The build's original standard-mode pointer-identity
assertion remains: a cloned heap mode is invalid at this entry and was rejected,
not accommodated by weakening validation. Native scalar helper/whole concealment
comparisons plus unchanged enclosing frame C-reference/Go goldens and codec
baselines still pass. Full per-round amd64/386 and ARM64/QEMU, repeated active-path
checkptr/ordinary goldens, GC stress and unchanged tolerances pass. This proves
coverage of these typed decoder paths, not global Opus-frame uintptr ownership,
macOS CI or scanning of embedded pointers in opaque byte-backed allocations.

Four final quant-all-bands owner rounds type the left spectrum base, right
spectrum base and per-band/fallback X/Y pointers, then remove the unused TLS
pseudostack setup/save/final restore and its dead translated temporaries. The
public uintptr escape ABI remains, but the private quantizer has only typed
pointer arguments and no integer-addressed owners or views. The typed CELT
caller forwards both spectrum channels directly; the private escape annotation
is no longer necessary. All band offsets use numeric indexing; fallback X/Y
remain live aliases of the typed norm lane. No private call allocates or touches
pseudostack storage, and valid paths accept nil TLS.

Full active typed-quantizer checkptr now runs with scanned heap mode/cache/table
owners, Go arrays/entropy/buffers and nil TLS. The new grouped whole fixture
reuses unchanged scalar C goldens for four encoder scenarios (mono, stereo,
dual stereo and theta-RDO) plus decoder resynthesis, including encoded bytes,
bit counts, spectra, masks, guards and seeds. An additional active mono/stereo
LM0..3 fixture exercises pulse/noise bands, TF changes, heap modes, GC/stack growth
and guards; one case supplies an untouched legacy TLS cursor sentinel. Other
legacy C-reference fixtures remain ordinary tests; this expanded checkptr claim
is for the typed quantizer, not the enclosing outer CELT decoder's still-legacy
pseudostack or public decoder storage. Full per-round amd64/386, ARM64/QEMU,
native comparisons and GC stress preserve encode/decode goldens/tolerances.
Opaque byte-backed allocations still do not scan embedded pointers globally.

Four main-norm rounds type the dual-stereo-to-intensity merge with live eBands
bound reloads and ordered float stores, then switch dual-stereo and normal
mono/stereo dispatch to typed quant_band/quant_band_stereo calls with numeric
optional folding views. The final round allocates the main norm buffer in Go,
retains typed left/right pointers, skips unused mono/empty second-lane interiors,
and switches RDO dispatch to the same typed band views. Obsolete private mono
and stereo uintptr adapters are removed. Quant-all-bands no longer allocates any
TLS scratch; initial pseudostack setup/save and final cursor restore remain for
a later boundary cleanup. Spectral bases/X/Y fallback still have explicit integer
views; all norm offsets/merges/copies/folding consumers are now numeric/typed.

Grouped scoped checkptr tests cover optional -1 inputs, last-band nil outputs,
zero/nonzero norm extents, channels1/2, LM0..3, nonzero offsets, live merge aliases
and retained lane owners across GC/stack growth. An int16/float merge-bound alias
is explicitly Go-only, because C effective-type rules exclude it. Native fixtures
compare scalar merge operations and source-equivalent folding/allocation geometry,
not a whole-driver pointer oracle. Whole-band scalar C references still verify
mono/stereo/dual/intensity/theta-RDO entropy/spectra/seed. Each round passes full
amd64/386, ARM64/QEMU, native comparisons and GC stress with unchanged baselines;
repeated ARM checkptr stays scoped to typed helpers/consumers, separate from
ordinary whole-band/frame goldens. Spectrum integer views and byte-backed
embedded-pointer scanning still prevent a global GC-safety claim.

Four quant-all-bands float-storage rounds replace initial X/Y snapshots,
trial X/Y snapshots, the norm snapshot and encoder lowband scratch with separate
Go-owned float slices of the original resynth_alloc length. Zero length yields
nil unused storage. All six float scratch TLS allocation/alignment/capacity
blocks are removed; snapshots and dot/norm consumers now receive SliceData
pointers directly, never a reconverted snapshot integer address. Lowband scratch
is a typed pointer forwarded through the remaining mono/stereo legacy adapters.
On decode/non-resynthesis it preserves the live last-band spectrum alias using
numeric indexing; an empty band loop does not form an unused interior view.

Grouped GC/stack-growth/checkptr tests check zero/1/4/64/960 lengths, independent
initial/trial lanes, exact snapshot/restore consumers, norm offsets/guards and
LM0..3 last-band scratch aliases. Native storage pipeline checks use actual owned
Go storage with scalar C copy/norm operations and source-equivalent band offset
geometry; they are not a whole-driver pointer oracle. Existing whole-band scalar
C-reference cases still verify theta-RDO bytes, spectra, entropy and seed. Each
round passes full amd64/386, ARM64/QEMU, scoped checkptr, native codec comparison
and GC stress without changing goldens/tolerances. Only the main norm buffer
still allocates quantizer TLS scratch; spectrum/folding/norm views and initial
pseudostack setup/final cursor restore remain legacy. Full driver checkptr and
opaque byte-backed embedded-pointer scanning are still not claimed.

Four quant-all-bands RDO rounds replace spectrum snapshot/restore memcpy calls
with typed float copy consumers; replace four integer-addressed distortion loops
with numeric, ordered scalar dot products and explicitly rounded float32
products; replace norm snapshot/restore memcpy with typed numeric-offset views;
and replace entropy byte scratch with Go-owned 1275-byte storage when theta-RDO
is nonzero. The original entropy offs/storage snapshot selects a retained byte
window; later context rollback does not rederive its owner. Zero-byte windows
remain nil and never form unused one-past pointers. Only the byte scratch TLS
allocation/alignment/capacity block is removed; float scratch still uses TLS.

Grouped tests cover exact float bits, NaN/signed-zero copies, ordered products,
guards, zero-length nil inputs, byte offsets/full extent/EOF and retained payload
owners across GC/stack growth. Overlapping float/norm copy cases are explicitly
Go-only (C OPUS_COPY's memcpy overlap is invalid). Native copy/window fixtures use
the source-equivalent OPUS_COPY operations; dot fixtures use scalar ordered
products, not libopus SIMD reduction. Whole-band scalar C-reference scenarios
still check theta-RDO rollback/bytes/spectra/entropy/seed. Each round passes full
amd64/386, ARM64/QEMU, scoped helper checkptr, native comparisons and GC stress
with unchanged encode/decode goldens/tolerances. Full driver checkptr is not
claimed: float scratch owners and spectral/lowband views remain integer-addressed.

Four quant-all-bands view rounds type TF flags, pulse budgets, collapse masks
and the cached eBands owner. All accesses to these views in the driver now use
numeric unsafe.Slice indexing, including allocation geometry, folding searches,
RDO norm-copy offsets and the final balance reload. Private TF/pulse/mask arguments
are forwarded directly by the typed CELT decoder; the public escape adapter is
retained. Pulse loads stay live at every original comparison/reload rather than
being cached across entropy work. Mask stores narrow to byte, first then last
lane, so mono retains the original last-store-wins behavior. Band reads sign
extend int16, with the same M multiplication, int32 narrowing and uintptr offsets
for the still-legacy spectrum/norm consumers; arithmetic was not canonicalized.

Grouped GC/stack-growth/checkptr tests cover live mutations, signed extremes,
heap-mode endpoint ownership, mono/stereo narrowing/store order and guards.
Native source-equivalent scalar array accesses/stores plus existing whole-band
C-reference cases and native codec comparisons pass, as do full amd64/386 and
ARM64/QEMU, GC stress and unchanged encode/decode goldens/tolerances. Repeated ARM
checkptr is scoped to typed helpers/consumers; separate ordinary runs retain the
whole-band and frame goldens. Spectrum/RDO/norm/lowband scratch, memcpy operations
and pseudostack cursors remain legacy, so full quant-all-bands checkptr is not
claimed. Opaque byte-backed pointer scanning remains a separate blocker.

Four quant-all-bands owner rounds introduce a private quant_all_bands entry with
typed mode, entropy, seed and band-energy pointers. The public uintptr escape ABI
remains; the typed CELT decoder now forwards those owners directly rather than
reconverting them through the public adapter. Integer spectrum/mask/pulse/TF
arguments retain an explicit escape annotation on the private entry. Context
mode/entropy/energy assignments occur at their original points, preserving the
scanned band context's write barriers. Seed load and final store stay in place;
channel-weight lane loads now use numeric indexing, left before right and before
any output writes, including valid energy/weight aliases.

Each round runs the existing whole-band scalar C-reference scenarios (including
encoder theta-RDO entropy snapshots/rollback and final seed), full amd64/386,
ARM64/QEMU, native comparisons and GC stress without changing encode/decode
baselines. Grouped typed-owner/seed/weight tests force GC and stack growth; native
weight comparisons use source-equivalent bands.c MIN32/ADD32 float branches and
check aliases/guards. Scoped checkptr covers these helpers, not the whole-band
legacy fixtures. Norm/lowband/RDO scratch, integer-addressed spectra and arrays,
pseudostack setup/cursor restore remain unmigrated. Typed mode roots do not make
opaque byte-backed allocations scan their embedded child pointers globally.

Four internal-entry rounds introduce celt_decode_with_ec_dred with typed decoder,
payload, entropy and float PCM arguments, behind the retained public uintptr
adapters. Both public forwarding ABIs explicitly escape pointer arguments. The
legacy quant-all-bands seed argument is now the direct address of the typed rng
field, not a decoder-base integer offset. Packet continuity, loss-packet predicate
and PCM/length validation helpers preserve their original ordering. Frame-size
matching uses a typed mode and the same live maxLM/short-size loop, before packet
argument validation and without moving N computation ahead of it.

Per-round full amd64/386 and ARM64/QEMU suites, native comparisons, GC stress and
unchanged encode/decode baselines exercise the private entry via the public ABI.
Grouped scoped checkptr/native tests cover continuity, nil payload/PCM,
length bounds, zero/one-byte loss predicates and frame-size matching. This is a
typed entry/owner migration, not yet a full active-path checkptr proof: initial
pseudostack setup/cursor restore and quant-all-bands still use legacy integer
addresses. Existing redundant typed-pointer casts do not turn those owners back
into integer addresses. Opaque byte-backed decoder allocations still do not scan
embedded pointer fields; public ABI retention alone does not fix that storage.

Four outer-owner rounds retain typed mode and entropy, replace energy-history
integer addresses with typed numeric views, and replace decode-history integer
addresses with scanned channel slices/output pointers. The mode getter snapshots
bands/overlap/eBands at the same point after decoder validation. Entropy retains
an existing context untouched or initializes a scanned local context with its
typed payload pointer; only quant-all-bands reconverts mode/context at its escape
ABI. Energy/log/previous/background views use stride*channels then 2*bands lane
increments. History slices use (DEC_PITCH_BUF_SIZE+overlap) stride and outputs at
DEC_PITCH_BUF_SIZE-N, preserving mono unused lanes and zero-frame nil outputs.

Grouped tests force GC/stack growth with heap modes, scanned decoder-tail
fixtures and typed entropy contexts, verify stores/owners/unused views, and
compare mode metadata, entropy consumers and C-source numeric view geometry.
Native view fixtures compare formulas, not a whole outer-frame pointer oracle;
float-backed numeric images never acquire embedded Go owners. Full amd64/386,
ARM64/QEMU, scoped typed-helper/consumer checkptr, native and GC stress preserve
existing encode/decode goldens/tolerances. Repeated ARM checkptr remains scoped:
public decoder/payload/PCM uintptr entry and legacy quant-all-bands/pseudostack
remain, and opaque byte-backed embedded pointers are still not scanned globally.

Four allocation/finalization rounds type trim selection, fractional bit-budget
and anti-collapse reservation, final-energy forwarding, and anti-collapse bit/
dispatch boundaries. Trim keeps the cached fractional tell plus six-bit guard and
default five without consuming unused entropy. Reservation keeps int32 length*8
shift, tell-frac subtraction, minus one, nonzero-transient/LM>=2 threshold and
one-bit refund before allocation. Final energy samples integer tell at its original
position and forwards typed mode/energy/fine/priority/entropy owners. The reserved
bit still decodes before final energy; conditional anti-collapse still runs after
it with live seed/arch and typed spectra/masks/history/pulse owners.

Grouped tests cover budget gates, negative/zero flags, nil unused entropy/mode/
state inputs, mono/stereo final-energy guards and complete entropy state. Existing
owned-mask pipelines now exercise the dispatch helper with full LM0..3 native
scalar anti-collapse parity. C trim/reservation/raw-bit fixtures use actual
celt.h tables/entropy kernels; final energy reuses actual quant_bands.c. Each
round retains full amd64/386, ARM64/QEMU, scoped typed-helper/consumer checkptr,
native, GC stress and unchanged encode/decode goldens/tolerances, followed by
separate repeated ARM pointer and ordinary golden tests. These boundaries do not
make the remaining outer CELT/quant-all-bands integer views globally GC-safe.

Four header rounds type silence recognition/budget exhaustion, postfilter
parameter decoding, transient/intra flags, and spreading selection. All retain
typed entropy/payload owners and numeric int32 tell arithmetic. Silence keeps
the initial tell snapshot unless exhaustion advances total nbits. Postfilter
retains the start/budget gate, octave uint/raw-bit/gain/tapset decoding order and
unsigned pitch arithmetic, with a tell refresh only inside that gate. Global
flags refresh tell after transient decoding but deliberately not after intra;
spreading likewise returns its pre-ICDF tell snapshot. These cached values and
exact inequalities remain unchanged at the outer call sites.

Grouped GC/stack-growth/checkptr and native tests cover budgets immediately
below/at/above gates, omitted postfilter for nonzero start, LM0..3, all eleven
entropy fields and every scalar output. Source-equivalent C header fixtures use
actual entropy kernels, celt.h tapset/spread tables and SPREAD_NORMAL. They remain
leaf oracles rather than full enclosing CELT decoder checkptr coverage. Full
amd64/386 and ARM64/QEMU, native comparisons, GC stress, unchanged encode/decode
goldens/tolerances and repeated scoped ARM pointer/ordinary golden tests are
retained per round/batch. Quant-all-bands and outer integer-addressed owners
still require subsequent migration.

Four subsequent boundary rounds type initial mono energy merging, conditional
prefilter dispatch, decoder deemphasis forwarding and final entropy/error
completion. Mono merging keeps MAXG's second-operand tie/NaN choice and only
writes lane zero. Prefilter dispatch retains the live nonzero fold predicate and
calls the fully typed prefilter directly. The decode output array is now scanned
[2]*float32 storage shared by synthesis, postfilter and normal/lost deemphasis;
coefficient and preemphasis memory pointers use actual mode/state fields rather
than offsets. Terminal error handling retains int32 tell/length arithmetic,
returns -3 before sticky-error writes, leaves entropy unchanged and only sets
state error to one for a nonzero entropy error after passing the budget check.
Its position after packet reset and legacy cursor restore is unchanged.

Grouped tests cover mono NaNs/ties/infinities/guards, prefilter flags 0/1/-1/7,
heap mode/history/state owners through GC/stack growth, mono/stereo deemphasis
factors 1/2/3/6 and accumulation, and exact entropy budget/error ordering. Native
prefilter comparisons reuse whole scalar folding for active flags, with unchanged
images for zero flags; deemphasis reuses actual scalar decoder C. Packet-error
fixtures use actual ec_tell for valid nonzero ranges; signed-overflow and zero-
range edge fixtures are explicitly Go-only. Per-round full amd64/386,
ARM64/QEMU, scoped typed-helper/consumer checkptr, native and GC stress retain
all original encode/decode goldens/tolerances. Final ARM repeats are still not
whole legacy CELT/quant-all-bands checkptr or a global opaque-scanning proof.

Four postfilter rounds type period clamps, first-short-frame comb inputs,
tail comb inputs and the channel driver. Clamping retains current-then-previous
stores each channel. First and tail calls read live decoder gains/periods/tapsets,
mode short size/window and arch separately; the tail uses numeric indexing into
the retained N-sample output span, and is only invoked for nonzero LM. The driver
uses the scanned synthesis output-pointer array, reloads each channel pointer for
each call, and preserves the original do-while channel semantics. The outer
caller no longer uses comb_filter_legacy in this path; postfilter-state finalization
remains after the complete channel loop.

Grouped tests retain heap mode/state/history/window owners through GC and stack
growth; compare full guarded histories, untouched state fields, gain/tap choices,
mono/stereo LM0..3, and zero-channel leaf do-while behavior (not a claim that a
zero-channel decoder passes validation). Native first/tail tests use the actual
scalar celt.c comb implementation. The source-equivalent C channel driver uses
that same kernel and compares resulting periods and complete sample buffers.
Each commit passes full amd64/386, ARM64/QEMU, native and GC stress with original
goldens/tolerances. Final ARM repeats keep scoped typed postfilter/helper/PLC
checkptr separate from ordinary golden runs. Legacy enclosing decoder views,
quant-all-bands scratch and opaque pointer scanning remain outside that proof.

The following four recovery/finalization rounds type per-channel missing/safety
controls, individual loss-recovery energy prediction, the full two-channel
recovery loop, and normal packet-state reset. Safety keeps signed loss_duration
shift and the ten-frame cap. Prediction retains ordered MAXG/MING/MAX32 branches,
NaN/tie selection, explicit float32 difference/half/product/subtraction rounding,
the -20 clamp and the final post-store safety subtraction. Recovery uses typed
numeric-indexed energy/log/history spans, the original intra/loss short circuit,
exactly two channels (including mono), and a fresh loss-duration read per channel.
The normal reset writes loss, PLC duration, frame type, then fold flag after
deemphasis and before the unchanged cursor restore and entropy/error checks.

Grouped pointer/native tests cover signed loss extremes, LM0..3, partial/empty
bands, no-op nil inputs, finite/clamped branches, NaNs, infinities and signed
zero, guards, GC/stack growth and unchanged unrelated decoder fields. The
float/count alias that changes channel-two safety is Go-only, not a C effective-
type claim. Native fixtures remain source-equivalent celt_decoder.c snippets
using its actual comparison macros. Each commit retains full amd64/386,
ARM64/QEMU, scoped typed helper/recovery/PLC checkptr, native, GC stress and
unchanged encode/decode baselines/tolerances. Batch repeats separate typed-path
ARM checkptr from ordinary frame goldens; enclosing CELT/quant-all-bands still
contains legacy integer-addressed views and is not covered by that claim.

Four further decode-operation rounds type history memmove, silence energy
filling, dynamic boosts and postfilter finalization. History shifting uses one
exact trailing-storage slice and overlapping copy; zero length consumes no view.
Silence writes -28 in cached C*nbEBands order (Go count overflow cases are not C
parity claims). Dynamic boosts retain a typed cached eBands owner, typed live
cap/offset/entropy pointers, six-bit quanta bounds, logp evolution, strict budget
and cap checks, and tell reload after every flag. The helper returns both budget
and tell to the unchanged trim decision. Postfilter finalization preserves the
three previous-state stores, three new-state stores and nonzero-LM previous-state
reload/stores, with no float arithmetic or NaN/signed-zero selection changes.

Grouped tests force GC/stack growth, verify guards/exact history spans and nil
unused inputs, exercise boost budgets/partial and empty bands/caps/LM0..3/mono
and stereo, and cover postfilter negative/nonzero LM and exceptional gain bits.
Source-equivalent celt_decoder.c C snippets use actual IMIN/IMAX and entropy
kernels; boost comparisons include all eleven entropy fields, offsets and both
returned scalars. These are leaf-operation oracles, not a new whole-decode
comparison or enclosing checkptr proof. Per-round full amd64/386, ARM64/QEMU,
native, GC stress and unchanged encode/decode goldens/tolerances are retained;
repeated ARM typed-operation checkptr and ordinary golden runs remain separate.

The next four scratch rounds migrate pulses, fine priorities, contiguous C*N
spectra and C*nbEBands collapse masks into Go-owned slices. Allocation writes
typed pulse/priority outputs, and final-energy decoding consumes typed priorities.
Spectral channel views retain the base and stereo N offset with numeric indexing;
mono and empty views never materialize an unused interior. The decode caller now
uses typed anti-collapse and synthesis directly, with a scanned two-pointer
synthesis output array; only quant-all-bands converts these scratch owners back
to uintptr through its explicit escape ABI. Every outer decode temporary array
is now Go-owned. Its initial pseudostack setup/save/restore is deliberately still
present: the legacy quant-all-bands body expects an initialized scratch stack.
This is not yet a complete typed outer CELT decode or quant-all-bands migration.

Grouped tests cover empty/positive storage geometry, retained channel views,
GC/stack growth, allocation pulse/priority writes, final-energy priority reads,
mono/stereo LM0..3 synthesis and anti-collapse with guarded outputs. Native
fixtures compare actual rate.c output/scalar/entropy state, quant_bands.c final
energy, scalar synthesis and anti-collapse. All four rounds run full amd64/386,
ARM64/QEMU, scoped helper/consumer checkptr, native comparisons and GC stress
without changing goldens/tolerances. Final repeated ARM checkptr remains scoped
to typed helper/consumer/PLC paths, separately from ordinary frame goldens.

Four later outer CELT decode scratch rounds replace pseudostack allocations for
TF flags, caps, boost offsets and fine-energy bits with Go-owned int32 slices.
Each removes its TLS allocation/alignment/cursor block. TF and caps initialize
through the existing typed kernels; caps reads and boost-offset stores use
numeric indexing. Fine-energy allocation now calls the typed allocation driver
with typed local outputs, followed by typed fine/final-energy consumers. Remaining
pulses, priorities, spectra and collapse-mask storage are still pseudostack-backed.
The quant-all-bands uintptr ABI explicitly escapes pointer arguments so new TF
storage survives its recursive legacy consumers; this is a boundary adapter,
not migration of the quant-all-bands body. Empty caps skip unused table views,
matching the C zero-iteration loop and accepting nil unused pointers.

Grouped owned-storage tests force GC/stack growth, cover empty and positive
storage sizes, TF transient/LM/partial-band combinations, mono/stereo caps and
fine/final-energy consumers. Native tests reuse actual TF/caps kernels, rate.c
allocation and quant_bands.c energy decoding, including complete entropy state.
Scoped checkptr covers these typed helper/consumer pipelines and migrated PLC,
not the enclosing legacy CELT decode/quant-all-bands path. Full amd64/386,
ARM64/QEMU, native, GC stress and unchanged encode/decode goldens/tolerances run
per commit; repeated ARM pointer and ordinary golden tests follow the batch.

Four subsequent outer CELT decode-history rounds type mono energy duplication,
current/previous log-energy updates, background energy tracking and out-of-band
clearing. Mono duplication copies the contiguous second lane; nontransient logs
copy previous<-current before current<-energy. Transient updates retain MING's
ordered comparison/tie/NaN selection, not Go's builtin floating minimum. The
background increment keeps int32 loss_duration+M, the 160 cap, exact float32
.001 scaling, and sequential live loads/stores. Clearing retains two channel
passes, prefix then suffix ranges, energy positive-zero followed by previous-log
then current-log -28 stores; overlapping ranges still execute in source order.

Grouped Go/native tests cover bands 0/1/3/21/25, zero/nonzero/negative transient
flags, loss/LM increments, full/empty/overlapping active bands, guards, GC/stack
growth, quiet NaNs, signed zero and infinity. Copy/history aliases and signed
overflow are explicitly Go-only, not memcpy-overlap or C-overflow parity claims.
Native fixtures use source-equivalent celt_decoder.c snippets and its actual
MING/IMIN macros; they are leaf oracles, not a new complete decode-frame oracle.
Scoped checkptr covers these typed history helpers alongside full migrated
concealment, not the still-legacy outer CELT decode dispatcher or quant-all-bands.
Each round keeps full amd64/386/ARM64-QEMU tests, native and GC stress, unchanged
packet/frame/encode/decode goldens/tolerances and repeated ARM checks. This does
not resolve opaque pointer scanning, global GC safety or direct macOS CI.

The outer SILK API now retains typed decoder/channel, control, entropy, float PCM
and output-count pointers behind Opus_silk_Decode's explicit uintptr escape ABI.
Packet-start frame counters use typed state and a live channel-count pointer.
Stereo-start state clears use clear; resampler cloning uses typed struct assignment
rather than the hard-coded 400-byte amd64 image, respecting 386 size and barriering
the coefficient pointer. Clear/clear/copy order and channel-transition predicates
are preserved. This fixes the adjacent-state overwrite risk on 386 without changing
C's sizeof-based behavior.

LBRR flag reconstruction clears the live three-word field, preserves the no-entropy
one-frame case, and uses typed pointers for both ICDF table rows. Normal entropy
forwarding retains typed contexts. Tests compare all flags and eleven entropy
fields for 1/2/3 frames, zero/nonzero/negative LBRR flags and 32 input patterns.

Count and PCM helpers use typed outputs and numeric indexing. Count calculation
keeps int32 multiplication/division and signed int16 rate narrowing; float PCM
conversion preserves the exact float32 1/32768 scaling. Mono/stereo interleave,
collapsed-stereo right-channel output and sequential mono duplication retain live
count reloads, channel order and guards. Native fixtures cover counts at API rates
8/12/16/24/48 kHz, lengths 0/1/17/80/240/960 and all output channels. A Go-only
float/count alias verifies stopping after a count-changing store; it is not a
C effective-type parity claim.

The subsequent four API scratch rounds type channel slices, replace channel
storage with one Go-owned int16 array, type delayed-input/resampling views, then
replace resampling storage with a Go slice. Unused mono channel pointers are not
materialized. Sequential history copies, side clears, channel passes, collapsed
stereo resampling and live output/count reloads retain source order. All API TLS
allocation/cursor/save/restore operations are gone; the public escape ABI remains.

Focused checkptr now covers the complete typed API with nil TLS, GC/stack growth,
heap codebooks, normal/FEC/loss calls, transitions, exact guards, validation order
and an untouched TLS sentinel. Original real-packet and PLC goldens are unchanged;
a full-API float/count alias is explicitly Go-only. Earlier helper-only checkptr
limits and references below to legacy outer SILK API scratch are historical.

The renamed upstream dec_API.c oracle compares full numeric decoder/control
images, PCM/counts and all eleven entropy fields at internal rates 8/12/16 kHz,
API rates 8/24/48 kHz, all channel combinations and normal/FEC/loss flags. It also
checks channel transitions and six consecutive real 60-ms packet frame calls.
Native initialization selects scalar arch=0 to match Go, not host SIMD arch=4.
Numeric images strip pointers; native tables are rebound locally, cleared before
import, and Go table owners are restored by typed assignment. Delayed resampling
views also match C at all five supported decoder API rates and 10/20-ms lengths.
All four rounds retain full amd64/386/ARM64-QEMU tests, original API/frame goldens,
end-to-end native/GC-stress comparisons and unchanged baselines/tolerances. This
is neither direct macOS CI nor a global GC-safety proof for opaque allocations.

SILK frame decoding now retains typed decoder, entropy, PCM and output-count
pointers behind the public Opus_silk_decode_frame uintptr escape adapter. The
control struct and shell-aligned pulse buffer are Go-owned; 120-sample frames
retain 128 pulse slots. There is no frame TLS allocation/cursor/save/restore.
Normal/FEC dispatch still follows the exact LBRR flag test; all other flags
conceal. History shift precedes copying live PCM, then PLC/CNG/glue, the lag store
and final output-count store remain ordered. The two history assertions retain
their separate source/error positions. Internal core and PLC calls are typed.

Actual renamed decode_frame.c fixtures compare full numeric decoder images,
PCM/guards, counts and all eleven entropy fields: rates 8/12/16, 2/4 subframes,
normal/FEC flags and independent/conditional coding across three consecutive
calls, plus loss/FEC-fallback/negative/other flags. Go embedded pointer fields
are cleared in a numeric temporary before exporting bytes; native tables are
rebound on the C stack and cleared before import. Original Go pointer fields are
restored via a typed struct assignment, not raw pointer-byte stores. Input entropy
buffers remain unchanged. Native history and lag/count alias fixtures complement
the whole-frame oracle; overlapping history memcpy cases are Go-only.

Focused checkptr now covers full active normal/FEC/loss frame paths with nil TLS,
GC/stack growth, heap codebooks, guards, shell padding, unused nil entropy on loss,
validation-before-output ordering and an untouched TLS sentinel. The original
loss-frame golden enters the typed frame with a Go-owned count, unchanged. Rounds
one/two checked typed history/finish helpers; round three covered active loss
with Go-owned control; round four enabled the entire frame after pulse migration.

Control allocation exposed an outer API stack-lifetime defect: the original
TestSilkDecodeLostFrameState returned count 0 rather than 80. The legacy public
Opus_silk_Decode ABI now declares uintptrescapes, and its frame-local count is
forwarded directly as a typed pointer. No test/golden was weakened. The outer
SILK API remains largely legacy, as do other decoder boundaries and opaque
byte-backed pointer-bearing storage. Host native/QEMU validation is not direct
macOS CI or a global GC-safety proof.

SILK normal decode-core ownership now retains typed decoder, control, PCM and pulse
pointers behind the public Opus_silk_decode_core uintptr escape adapter. The
k=2 output-history staging and final LPC-state destination use typed fields rather
than hard-coded state offsets. The PLC-to-unvoiced transition clears the live
five-tap control view, stores its center tap, then reloads/stores the decoder lag;
inactive branches do not consume control. Pulse excitation uses typed samples
and quantization-table indexing, unsigned RNG/shift wrapping and a second live
pulse load after excitation stores. The second load matters for Go-only
pulse/excitation aliases.

LPC/LTP coefficients use live fixed-array pointers. Only LPC_order entries are
copied into the LPC snapshot; unused tail entries remain untouched. Rewhitening
still receives the original live A, and LTP prediction reads the live five-tap B
in the original order. Go-only overlapping snapshots use copy and retain the
live source pointer. Sole typed pulse/coefficient interiors retain scanned
backing through GC/stack growth; decoder-history fixtures retain heap codebooks.

The actual renamed decode_core.c integration oracle checks complete numeric
state/control images, excitation, LPC history, output/guards and unchanged pulses
for rates 8/12/16, 2/4 subframes, signal types 0/1/2, interpolation 0/4, prior loss
0/1, gain changes, voiced rewhitening and PLC transition across three consecutive
core calls. Actual-source fixtures also check valid int16 PCM/pulse aliases and
PCM pointing into decoder outBuf. Embedded fixture
pointers remain nil; raw image copying is not a write-barrier-safe pointer import.
Native leaf fixtures cover history staging, transition stores, extreme/zero pulse
excitation and coefficient snapshots. Excitation includes seeds -128/-1/0/1/17/127,
lengths 0/1/17/80/160/320 and both quantization offsets. Alias/error cases that C
cannot define are Go-only.

The following four storage rounds migrate whitening, Q15 LTP history, residual/
PCM views, then LPC history. All four arrays are Go-owned: ltp_mem_length int16
samples, ltp_mem_length + frame_length Q15 words, subfr_length residual words and
subfr_length + MAX_LPC_ORDER LPC words. Reverse loads and prediction/state stores
use numeric indices, not integer addresses. Gain scaling and each LPC/LTP MAC
still narrow individually; wrapping residual/state shifts, PCM product narrowing,
rounding and saturation are unchanged. LPC retains the ten-tap plus optional
six-tap order, per-sample order assertion and PCM-before-history-copy ordering.
Go copy handles state shifts, including zero-length unused cursors safely.

Whitening compares native LPC analysis at k=0/2, multiple start offsets, lengths
160/240/320 and orders 10/16, with untouched prefixes/guards. Native SILK macro
fixtures compare reverse whitening/scale stores, five-tap prediction and PCM
narrowing. Residual signed-overflow cases are Go-only; native residual fixtures
stay within defined signed ranges.

After the last round, focused checkptr covers the full active normal core with
nil TLS, GC/stack growth, heap codebooks, zero-length unused PCM/pulses, guards and
an untouched TLS sentinel. Order-failure fixtures preserve excitation/gain stores
before the assertion and leave PCM untouched. Original normal-core tests now
enter the typed driver without pseudostack setup and retain their exact goldens.
There is no core TLS allocation/cursor/save/restore; only the public uintptr ABI
adapter remains. Earlier storage rounds checked helpers until the final legacy
LPC boundary was gone. Original frame goldens and decode/encode baselines remain
unchanged. Outer SILK API, CELT paths and opaque byte-backed decoder storage
are not made globally GC-safe. Full amd64/386 and ARM64/QEMU checks do not imply
direct macOS coverage.

SILK PLC dispatch now holds typed decoder, control and PCM pointers. The public
Opus_silk_PLC uintptr ABI is an explicit escape adapter that converts all three
arguments before forwarding. The rate mismatch/reset helper uses typed state;
reset/fs update precede dispatch, nonzero (including negative) loss flags select
concealment, lossCnt increments only after concealment returns, and nonloss
updates never consume PCM (nil and numeric aliases are valid on that branch).
The concealment entry now retains typed decoder/control/PCM and PLC-state
owners. LTP coefficients use a live fixed-array pointer and random excitation a
live 128-word decoder-owned view. Five-tap attenuation retains int16 gain
narrowing; random mixing retains per-MAC int32 narrowing and wrapping shifts.
The LPC history/PCM phase is a typed standalone kernel: copy the 16-word state,
check order >=10, perform the first ten MACs plus remaining order taps, narrow
each MAC, saturate the prediction shift and excitation sum, scale/round/saturate
PCM, then save history after all PCM stores. Modeled frame/order reads remain
live. The redundant nested SAT16 is collapsed without changing its value;
SMULWW still narrows before RSHIFT_ROUND (the max-int32 product fixture yields
-256, not a clamp of the wide product).

Whitening samples and Q14 prediction/history now use typed slices and numeric
indices, including reverse-order five-tap loads. The int16 whitening buffer has
ltp_mem_length samples; the int32 synthesis buffer has ltp_mem_length +
frame_length words. Both are Go-owned. There is no concealment TLS allocation,
cursor addressing or save/restore. Previous-LPC reset uses clear and coefficient
snapshots use copy. Whitening still runs before any PCM stores, the five MACs
narrow individually, subframe attenuation/pitch drift remain ordered, and LPC
uses the same live synthesis tail. The obsolete private uintptr concealment
adapter is removed; only the public Opus_silk_PLC ABI adapter remains.

Focused checkptr now covers complete active typed concealment/dispatch with
nil TLS, forced GC/stack growth, heap codebooks, PCM/guards, reset/type/loss
matrices and an untouched TLS sentinel. Original voiced/unvoiced C-reference
goldens now enter the typed driver without fixture pseudostack setup and pass
unchanged. Opaque byte-backed decoder allocations, outer SILK API,
CELT concealment and other legacy boundaries are not made globally GC-safe.

Actual PLC.c dispatcher fixtures compare the full decoder/control structs and
PCM/guards: rates 8/12/16, 2/4 subframes, orders 10/16 on update, all signal types,
matching/mismatching rates, prior loss counts 0/1/3, flags 1/-1/7, first-frame
reset, and voiced/unvoiced concealment. Native byte-image fixture states leave
embedded codebook/coefficient pointers nil; they do not import Go pointers via
raw stores. Native loss fixtures now call Go with nil TLS, and actual PLC.c
fixtures also check valid int16 decoder-history/PCM aliases (whitening precedes
output). Native comparisons remain host-only. Focused tests cover rate ownership,
full nonloss/loss dispatch with GC/stack growth, nil/unused PCM and Go-only invalid-
control reset/store ordering. A sole typed PLC interior retains a scanned decoder
and heap codebook through forced GC. Round one checked the rate helper while
control remained integer-addressed; the first dispatcher rounds checked full
nonloss dispatch, not active concealment. The coefficient/random/PCM/history rounds
also check active typed LPC buffers with nil TLS, GC/stack growth, lengths
0/1/10/16/17/80/320, orders 10/16, guard words, saturation and copy-before-order-
assertion behavior. Decoder-backed coefficient/random interiors retain heap
codebooks through weak-owner tests. Go-only PCM/state alias fixtures verify
history save after PCM; they do not claim C effective-type alias parity.
Macro-based native leaf fixtures add coefficient decay, random mixing, narrowed
PCM scaling, and the full LPC kernel/working history (zero/extreme coefficients,
several gains, all listed lengths). Actual PLC.c lost-dispatch fixtures remain
the complete native integration reference. The following whitening/Q14/scratch
rounds check typed whitening and five-tap helpers first, then Go synthesis storage,
and finally the whole active path after the int16 scratch migration. Whitening
fixtures cover lengths 160/240/320, orders 10/16, multiple consumed offsets,
untouched prefixes and guards. Go-only failure fixtures retain rate reset/bandwidth
expansion, do not increment lossCnt and do not touch PCM/control when rewhitening
asserts. All rounds retain the original baselines/goldens and cross-architecture
checks; QEMU is not direct macOS CI.

CELT allocation interpolation now takes typed mode/entropy, four band inputs,
three band outputs and balance/intensity/dual-stereo output pointers. The unused
TLS save/restore is removed, the interpolation integer adapter is gone, and the
complete leaf accepts nil TLS. Six-step bisection, int32 multiply/shift wrapping,
unsigned celt_udiv, backward skip decisions, entropy flags, reservation refunds,
N=1/N=2 fine-energy cases, caps/rebalancing, assertions and sequential alias stores
are preserved. Views remain live: no snapshots of aliased band/scalar values.

The allocation driver now takes typed mode/entropy, offset/cap inputs, scalar
outputs and pulse/energy/priority arrays. All four mode-band-length scratch arrays
(bits1, bits2, threshold, trim) are Go-owned; there is no TLS initialization,
allocation, cursor arithmetic or stack restoration in the core driver. The
public Opus_clt_compute_allocation ABI still has an explicit uintptr escape
wrapper, converting every pointer argument before allocation/stack growth.
Vector accesses retain typed mode-owned backing and the separately cached band
stride. Reservation/refund order, negative-total clamping, do-while vector search,
per-multiply int32 wrapping, threshold/tilt shifts, single-coefficient correction,
positive-only tilt application, dynalloc skip start and interpolation store order
are preserved. Wide-trim wrapping fixtures are Go-only, not C signed-overflow
parity claims.

Actual rate.c interpolation/driver fixtures compare returned coded bands, all
arrays/guards/scalars, eleven entropy fields and full byte buffers for C=1/2,
LM=0–3, multiple starts, budgets, encode/decode and stereo reservations. Driver
fixtures add trim 0/5/10 and dynalloc boosts. Leaf fixtures also cover input/output
aliases, shared intensity/dual outputs, balance/energy aliases, fine/priority
aliases and intensity/pulse aliases. Driver alias fixtures additionally cover
cap/pulse and cap/fine-energy sharing; native driver fixtures now use nil TLS.
Heap-owner/GC/stack-growth tests check mode bands/logN/vectors, entropy backing,
live inputs, curve scratch and the full active driver/interpolation chain. They
cover exact-sized one-band outputs, negative total, nil unused entropy, shared
scalar outputs and an untouched TLS sentinel. Focused checkptr passes on
amd64/386 and ARM64/QEMU; native fixtures remain host-only. The first three
scratch-migration rounds checked helpers/leaf paths while the driver still had
TLS scratch; the final round checks the complete active typed driver. Outer
quantization/decode/PLC boundaries, opaque allocation scans and extension EOF
are not made globally safe by this batch.

Decoder CTL dispatch now has typed CELT custom, Opus, multistream and projection
entries, using OpusDecoderCtlArgs for integer values and GC-visible scalar, range,
mode and decoder output slots. Legacy vararg entries delegate; internal forwarding
uses no TLS vararg/range scratch. Mode/decoder outputs use Go pointer stores and
write barriers. Typed aligned traversal preserves live layout reads, range-clear
before XOR, child-call order/early returns and first-stream getters. Numeric cursor
offsets avoid unused one-past pointers, including a dependent multistream-init fix
found by exact-sized composite checkptr tests; extension iterator EOF remains a
separate unresolved boundary.
Actual decoder C sources compare all supported requests, invalid/unknown requests,
boundary values, nil outputs, reset images, SILK/CELT pitch modes, one/three streams
and coupled layouts, decoder-output offsets and numeric output/layout aliases.
CELT custom CTL's C state parameter is restrict-qualified: its state/output aliases
are tested for Go store order, not claimed as valid C-oracle inputs. Focused checkptr
and GC/stack-growth tests cover the active forwarding chain and pointer-output sole
owners (including a heap mode retained through a returned component decoder) on
amd64, 386 and ARM64/QEMU. Native comparisons are host-only; raw vararg boundaries
and unscanned opaque byte-backed decoder storage still prevent a global safety claim.

SILK CNG now takes typed decoder/control/PCM pointers, shifts excitation with copy,
and clears only the active LPC history. Its length+16 synthesis buffer is Go-owned;
no TLS scratch allocation, restore, integer reconstruction or pinning remains in
this leaf. The fixed coefficient array and per-MAC int32 narrowing, approximate
integer square root, shifts, rounding, saturation, seed and final-history store
order match actual silk/CNG.c fixtures (including control/excitation and PCM/history
aliases). Native fixtures cover orders 10/16, rates 8/12/16, two/four subframes,
reset/no-reset, signal types, zero/short/full frames, loss counts, gain branches and
extreme signed samples. Focused checkptr includes active loss synthesis, nil TLS,
GC/stack growth and guards on amd64, 386 and ARM64/QEMU; original loss-path expected
outputs are unchanged. Earlier rounds checked non-loss paths only while TLS scratch
was still legacy. The enclosing SILK frame/PLC APIs and their pseudostack remain
legacy; this is not a global decoder safety claim.

Opaque byte-backed allocations, architecture FFT headers, outer integer APIs and
some outer band-table locals remain legacy. Earlier ARM64 quant-partition and stereo
fixture lifetime failures required pinning. With the internal spectral chain typed,
those pins and unused pseudostack setup are removed: the unchanged expected outputs
now pass with forced GC/stack growth and checkptr on all three tested architectures.
This does not migrate the outer quantizer.
MDCT lookups retain typed FFT-state and trig-table pointers; FFT states in turn
retain typed bit-reversal and twiddle pointers. Grouped FFT tests compare both C
layouts and force GC/stack growth with a heap lookup as the only table owner.
Forward/inverse MDCT cases use actual mdct.c plus scalar kiss_fft.c and compare
input/output bits and guards across shifts 0–3, overlap 0/4/120, signed/zero/strided
spectra, standard FFT sizes and odd N/4. Forward folding and rotation use Go scratch,
not TLS pseudostack storage. Inverse de-shuffling preserves both-end capture and
the double-processed odd middle pair before TDAC. Go also checks forward fold-before-
output aliases; native cases keep restrict-qualified buffers distinct. The obsolete
FFT integer adapter is removed. Synthesis now uses the typed MDCT chain;
other outer decode/concealment MDCT boundaries remain legacy.
Architecture headers and opaque decoder allocations are still not globally GC-safe.
FFT driver cases share the butterfly tests and use native-generated factors,
bit-reversal and twiddles for sizes 4–480, including shared-table shifts -1–2.
Forward FFT cases add native bit-reversal/scaling, separate buffers and both
partial-overlap directions, retaining explicit typed tables throughout the driver.
Inverse FFT reuses the grouped transform cases and checks both conjugation
passes and an exact four-point forward/inverse round trip.
Band normalization tests live beside denormalization tests and compare partial
bands/channel strides, empty bands/M=0, C's channel-zero visit, exceptional energies,
untouched tails and input preservation with the native bands implementation.
Mini-FFT radix-2 shares the FFT butterfly test files and compares actual static
mini_kfft.c helpers, arbitrary positive widths, repeated strides and raw float bits.
Mini radix-4 additionally compares all nonzero inverse flags, inverse sign/store
order and m=1 twiddle use (unlike the CELT twiddle-free special case).
Mini radix-3 adds typed epi3 selection and scalar rounding checks, updating the
existing C-reference fixture without adding another test file.
Mini radix-5 tests preserve left-associated sums and the distinct negated-product
expressions, including cancellation-heavy inputs, impulse output and guards.
Mini-FFT recursive work compares native-generated complete states and twiddles,
all four radices, inverse flags, repeated/strided reads, and unchanged inputs/state.
Each subsequent pointer round also runs full ARM64 tests and focused checkptr
under QEMU, in addition to amd64/386, native comparisons and GC stress.
Typed mini-FFT stride entry tests reuse these complete-state fixtures, including
zero stride in native comparisons and stack growth/GC with Go-owned state/input.
The typed unit-stride mini-FFT entry also compares actual C entry dispatch and
checks its unnormalized forward/inverse round trip; the real-FFT adapter remains legacy.
Band energies share the normalization/denormalization test files and compare
scalar C square accumulation/sqrt, empty bands, channel gaps, LM=0–3, exceptional
inputs and untouched energy tails; Go checks explicit product rounding on ARM64.
Decoder state validation tests are grouped in decoder_validation test files.
Opus checks compare actual static C assertions (captured by a test-only longjmp),
valid rate/channel combinations, rejected fields and unchanged state.
CELT state validation reuses those files, comparing actual native mode, band,
channel, pitch/period, tapset and architecture assertions without mutating state.
Multistream validation compares actual static C dispatch and layout results,
including invalid mappings: its ignored layout return is intentionally preserved.
Custom CELT decoder sizes reuse the validator fixtures, comparing native header
size, channels, overlaps and band counts; Go separately checks int32 size narrowing.
Time/frequency decode reuses entropy-bit tests and the actual static C helper,
comparing all LM/transient choices, budget exhaustion, selected bands and entropy state.
Extension payload skipping compares actual extensions.c helpers across all ID
bytes, lacing boundaries, trailing-short reservations and failure output preservation;
Go tests also check ownership through returned interior pointers.
Whole-extension skipping reuses those fixtures to compare ID/header consumption,
empty/negative lengths, all ID bytes and failure cursor/header behavior.
Payload writing shares extension fixtures and compares actual C lacing for
short/long IDs, 255-byte boundaries, final payloads, sizing-only calls, capacity
failures and untouched buffers. Packet generation now takes typed descriptor and
output pointers, uses fixed 48-frame scratch and typed record helpers, reloads
numeric descriptor fields after output stores, and pads with overlap-safe copy.
The integer generator writer adapters are removed. Actual extensions.c fixtures
cover repeats/interleaved frame order, short/long IDs and lacing, capacity/errors,
NULL-output sizing, padding, guards and numeric descriptor/header aliases. Original
generation goldens now run on Go-owned descriptors/payload/output storage under
checkptr on amd64, 386 and ARM64/QEMU.
Whole-extension writing additionally compares ID-byte narrowing and partial writes
before payload errors, using the same extension fixtures and actual static C helper.
Iterator initialization, repeat/next and find use typed state/output arguments and
GC-visible packet pointers. The same grouped fixtures compare every cursor, frame,
length and output against actual extensions.c, including captured assertions,
repeat L=0 handling, trailing short payloads, dynamic frame limits, nil outputs,
malformed packets and unchanged failed/not-found outputs. Random next/find fixtures
supplement structured cases; Go tests force GC/stack growth and retain returned
nonempty payloads after discarding the iterator.
Packet-extension count/count_ext/parse/parse_ext use typed input/output pointers,
direct clears, barrier-aware extension stores and fixed prefix-sum scratch. Grouped
actual-C comparisons cover random/structured packets, partial counts/errors,
capacity failures with unchanged capacity, frame-order output gaps, nil/assertion
ordering and capacity/count/output aliases. Count ignores malformed tails as C does;
parse reports errors after committing prior outputs. Capacity is read live, and
frame prefix sums are snapshotted before any output writes.
C-style end cursors still have an unresolved boundary for exactly sized Go
allocations: advancing one past an allocation fails checkptr (observed on 386).
Focused fixtures keep logical EOF inside guarded backing storage and verify guards;
this is not a fix for that general cursor-representation limitation. These passes
do not establish global GC safety for iterator endpoints or outer
decoder boundaries.
CELT FIR comparisons compile the actual celt_lpc.c scalar helper beside existing
LPC fixtures: reversed coefficients/history, 4-lane/tail arithmetic, odd orders,
signed zeros, subnormals, guards and partial overlap are checked bitwise.
CELT IIR shares those scalar C fixtures and checks recurrence patch/store order,
positive tail scratch, output-derived memory (including N<order history), guards,
in-place/partial overlap, signed zeros and subnormals bitwise.
CELT autocorrelation shares LPC fixtures and explicitly links the scalar pitch.c
bridge (the library's SIMD dispatch can change accumulation). Window copying,
overlapping window ends, prefix/tail grouping, lag bounds, input/output aliasing,
input/window preservation and guarded results are compared bitwise.
The internal packet parser compares native opus_packet_parse_impl outputs,
including untouched/error outputs, all framing branches, optional outputs,
self-delimited streams, duration/size limits, padding and signed size narrowing.
Frames/padding are typed packet interiors; outer integer APIs use an explicit adapter.
The public packet parser shares these fixtures and additionally checks native
public-API optional outputs and frame ownership across GC/stack growth on Go;
architecture guards check both consumed and untouched array entries.
LBRR detection uses that typed parser and compares native results for every TOC
and first payload byte, SILK mono/stereo/durations, CELT's pre-parse short circuit,
malformed packets and zero-size frames, without changing input bytes.
Multistream packet validation now uses stack-owned typed parser outputs instead
of TLS allocation. Native static-helper comparisons cover concatenated/self-delimited
streams, duration mismatches, missing/malformed streams and all supported rates;
existing C-reference fixtures also run directly on Go-owned packet buffers.
Pitch downsampling shares scalar pitch/LPC fixtures: mono/stereo and unusual
channel counts, factors 1–4, lag windowing, LPC/FIR stages, input preservation
and guarded output are compared bitwise; the decoder channel-table adapter remains legacy.
Pitch search shares those fixtures, comparing coarse/fine search, odd lengths,
lag limits, silence/periodic/random inputs and guarded pitch output; decimated
input and correlation scratch are now Go-owned rather than TLS allocations.
Pitch-doubling removal shares scalar pitch fixtures and compares period/gain
bits, continuity thresholds, rolling energy lookup, minimum/clamped/odd periods,
short windows and input/output guards; energy lookup is now Go-owned scratch.
SILK decoder rate setup reuses reset/resampler fixtures and compiles the actual
silk_decoder_set_fs helper. Internal/API/frame-only/no-change transitions compare
full Go state, resampler history, table identities and C untouched-state checks
for both subframe counts and all supported rates; fixed-offset clears are gone.
Mini-FFT allocation reuses FFT fixtures to compare native size queries, null or
undersized caller storage, full initialized bytes, factor/twiddle tables and guards;
typed owning returns keep the full flexible-array allocation alive across GC.
Real-FFT allocation shares these tests, comparing relative interior pointers,
architecture-sized headers, substate/super-twiddle bytes and query/capacity behavior;
the owning return and its three interior fields are typed.
Real-FFT transforms reuse native mini_kfft.c fixtures for all supported radix
combinations, odd/even complex halves, impulse/signed-zero inputs, exact spectrum
and scratch bits, immutable state/twiddles, guards and in-place time/frequency output.
CELT PLC pitch search now owns a fixed Go low-pass buffer and calls the fully
typed pitch chain. The actual static C driver is linked to scalar pitch fixtures;
silence, periodic/random mono/stereo data, input guards and shared channels are covered.
The surrounding concealment state/scratch boundary remains legacy.
Custom-mode lookup compares the actual static modes.c path over rates/frame sizes;
Go-only tests retain int32 shift wrapping without treating C signed-overflow UB as
an oracle. Comb transitions use renamed scalar celt.c for tapsets, gain changes,
history, zero-length/memmove paths, unchanged filters and overlapping buffers.
PVQ search/quantization reuse renamed vq.c: exact pulses, energy, fallback/sign bits,
reconstruction, collapse masks, entropy state and finalized bytes, including tiny
encoder capacities. Search/pulse scratch is Go-owned. Subsequent migrations
converted the partition/synthesis internal pointer chains and mode table fields;
outer decode/concealment adapters remain legacy.
One-bin quantization uses typed context/entropy/sample pointers and preserves
cached encode mode, mono/stereo aliases, lowband store order and insufficient-bit
behavior. Grouped bands fixtures compare sign bits, all entropy fields and finalized
buffers (including tiny/zero capacities) with actual bands.c.
Anti-collapse and spreading use explicit typed band tables and buffer/output
pointers. Anti-collapse preserves the float32 exp2 Horner/bit reconstruction, seed
order, mono decoder's second-channel history and normalization. Its actual bands.c
reference enables FLOAT_APPROX as the generated build does and binds scalar
normalization rather than linked-library SIMD. Spreading fixtures compare decisions,
recursive/HF state, shared output pointers, early exits and captured assertions.
The coarse-energy encoding leaf uses typed arrays/encoder and fixed predictor
scratch; actual quant_bands.c fixtures compare badness, error/reconstructed energy
store aliases, all LM/coding/LFE choices, budget branches, band-20 probability
clipping and full finalized entropy buffers. Float32 rounding and C store order
are preserved. The outer coarse-energy driver now uses typed mode, arrays, delayed history and
encoder snapshots, plus Go candidate-energy/error and saved-byte scratch. Actual
quant_bands.c cases compare one-/two-pass and forced/automatic intra behavior,
low budgets, pre-existing prefix/tail bits, tiny buffers, LFE, loss bias and output/
delayed-history aliases. Full-band cases initialize every C candidate-error slot;
partial-band intra copies can expose indeterminate C scratch and are not a defined
native oracle. Go scratch is zero-initialized. Outer quantizers and band-context
fields still cross legacy adapters; these checks are not a global GC-safety proof.
Typed single-stream, multistream and projection destroy entries pair with the typed
factories. Each releases only its allocation base. Native free spies verify exactly
one base free, including nil; Go weak-reference tests keep TLS live while verifying
that destruction releases the registry owner. Legacy destroy ABIs retain integer
registry-key frees without reconstructing heap pointers. These APIs do not change
the C no-use-after-destroy contract or make opaque byte allocations GC-scanned.
Decoder reset/init fixtures compare complete state images and guards using actual
celt_decoder.c, opus_decoder.c, SILK init_decoder.c and dec_API.c scalar builds;
only mode-pointer addresses are normalized. CELT reset uses offsetof(rng), retains
configuration, clears its complete flexible tail and seeds both log histories.
Custom/rate-aware CELT initialization retains typed mode ownership, zero-channel
behavior, unchanged argument failures and initialized state before unsupported-rate
assertions (captured in C). Typed Opus initialization removes its TLS vararg scratch;
Go-owned buffers, reinitialization, stack growth and GC are exercised. Creation,
multistream/control and decode boundaries still have explicit legacy adapters.
SILK parameter/index decoding reuses decoder-reset fixtures and actual
silk/decode_parameters.c and decode_indices.c for all rates/subframe counts,
voicing, interpolation/reset/loss cases, coding modes and entropy fields. Typed
state/control/entropy arguments, direct LTP tables and copy/clear operations keep
local scratch visible; codebook and table address fields remain legacy.
Multistream/projection init fixtures compare complete state images (mode addresses
only normalized), invalid/partial writes, guards and aliased mappings/matrices.
Projection coefficient and identity-map scratch is Go-owned, and the complete
initialization chain no longer needs TLS scratch. Creation/decode/control adapters
and broader allocation ownership remain legacy.
Typed registered allocation (`libcshim.XmallocPointer`/`XfreePointer`) derives
aligned pointers with unsafe.Add from the original Go backing slice, rather than
reconstructing them from integer addresses. Alignment, zero/oversized/nil-TLS
behavior, registry/free compatibility and backing lifetime are checked under
checkptr on amd64/386/ARM64. The three decoder factories have `_typed` entry points
with typed returns and mapping/matrix inputs; old exported integer-address APIs
remain explicit adapters. Creation calls typed initialization directly, and the
three unused initialization adapters were removed. Grouped native fixtures call
actual C factories with an allocator matching the Go shim's zeroed storage and
inject allocation failures, comparing error ordering and full state images
(normalizing mode addresses only), including invalid layouts, rates and matrices.
The registry and underlying byte allocations remain: a typed pointer retains the
backing object, but does not make pointer fields inside an opaque byte allocation
GC-scanned. This is not a global ownership/GC-safety proof, nor a migration of the
outer decode/control interfaces or extension end-cursor representation.
Float-to-PCM conversion, VAD initialization,
Laroia weights, sum-of-squares, bandwidth expansion (16/32-bit), 2:1 downsampling,
analysis filter bank, high-quality 2× upsampling, mono/stereo biquads, low-pass
cutoff control, and pitch decoding tests cover both `opuscc` and `opusccenc`.
The encoder float-to-PCM test also checks C's NaN-to-−32768 behavior.
The cutoff test compares PCM and state at every transition position in both
directions; its static tap helper also has exhaustive Q16 interpolation tests
in `opuscc`.

Integer helper comparisons are exact. Renormalization must exactly match a C
implementation of the scalar formula from `vq.c`/`pitch.h`. The linked native
build presumes SSE even with `arch=0`, so its different accumulation order is
checked separately with a relative bound of eight float32 machine epsilons
(about 9.54e-7). The initial maximum observed relative difference was 4.56e-7.
Run that test with `-v` to see the measured difference:

```sh
go test -tags compareopus ./tools/compareopus
go test -tags compareopus ./tools/compareopus -run TestRenormaliseAgainstC -v
```
