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
