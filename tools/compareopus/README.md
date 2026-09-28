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
`Fbuf` remains a legacy `uintptr`, with its test buffer owner explicitly retained.
SILK decoder comparisons also cover NLSF unpacking/reconstruction/stabilization,
LPC coefficient fitting (including input updates), and shell pulse decoding with
all numeric entropy-state fields checked. Additional decoder tests cover 16-bit
ICDFs, both Laplace variants, and all 5,625 SILK stereo predictor-index combinations
plus mid-only flags. Signed Laplace fixtures avoid zero-probability symbols
(e.g. negative values when `p0=32767`).
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
