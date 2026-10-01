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
comparisons. Poisoned reset fixtures skip GC pointer slots and use real sentinels.
CELT window, allocation-vector and cap-table fields are typed; native fixtures compare
mode/cache layout and table payloads, and heap-owner cases exercise caps and MDCT.
Band/log and pulse-cache index/bit fields, opaque byte-backed allocations and outer
integer APIs remain legacy. ARM64 exposed a quant-partition fixture stack-lifetime
failure; its crossing objects are now pinned without changing expected outputs.
This is fixture ownership repair, not a migration of the outer quantizer.
MDCT lookups retain typed FFT-state and trig-table pointers; FFT states in turn
retain typed bit-reversal and twiddle pointers. Grouped FFT tests compare both C
layouts and force GC/stack growth with a heap lookup as the only table owner.
Forward/inverse MDCT cases use actual mdct.c plus scalar kiss_fft.c and compare
input/output bits and guards across shifts 0–3, overlap 0/4/120, signed/zero/strided
spectra, standard FFT sizes and odd N/4. Forward folding and rotation use Go scratch,
not TLS pseudostack storage. Inverse de-shuffling preserves both-end capture and
the double-processed odd middle pair before TDAC. Go also checks forward fold-before-
output aliases; native cases keep restrict-qualified buffers distinct. The obsolete
FFT integer adapter is removed; the outer synthesis MDCT boundary remains legacy.
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
failures and untouched buffers. The generator API remains explicitly legacy.
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
do not establish global GC safety for iterator endpoints, the generator or outer
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
encoder capacities. Search/pulse scratch is Go-owned; outer partition/synthesis
adapters and mode pointer fields remain legacy.
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
