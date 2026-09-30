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
