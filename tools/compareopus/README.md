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
