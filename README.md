# opusgo (vAudience fork)

> **Fork notes.** `github.com/vAudience/opus-go` is a fork of
> [kazzmir/opus-go](https://github.com/kazzmir/opus-go) at `v1.6.0`, kept for its
> **encoder**, which go-vaicap uses to compress TTS audio. It differs from upstream in
> one place, the C-runtime shim (`libcshim`):
>
> - **The pseudo-stack never moves.** Upstream grows `TLS.Alloc`'s stack by copying it
>   into a larger slice, which leaves every outstanding `uintptr` into the old one
>   pointing at memory nothing keeps alive (a use-after-free). Here it is a fixed 1 MiB
>   reservation, and running out of it panics instead of moving.
> - **C memory lives outside the Go heap.** On Linux, macOS and the BSDs, `Xmalloc` and
>   the stack come from anonymous `mmap` regions (other platforms fall back to Go
>   memory). The collector never scans, moves or frees them, so a `uintptr` into them is
>   an ordinary address. `TLS.Close` releases them, with a finalizer as a backstop.
> - **The transpiled packages and the shim are marked `//go:nocheckptr`.** They take
>   addresses into package-level tables by design, which checkptr (enabled by `-race`)
>   rejects even though the data segment never moves. Data-race detection stays on, and
>   so does checkptr for the hand-written wrappers.
> - ⚠ **The decoder (`opuscc`) is not held to the encoder's bar.** It passes
>   `uintptr(unsafe.Pointer(&local))` into deep call chains. When the goroutine stack
>   moves mid-call, the callee reads the abandoned copy, and two decoder tests return
>   wrong values under `-race` for exactly that reason (they are skipped there, with the
>   reason). `internal/unsafeaudit` finds the pattern, and a guard test holds `opusccenc`
>   at zero. **Use the decoder in production only after that is fixed.**
>
> `go test -race ./...` is green.


[![Go Reference](https://pkg.go.dev/badge/github.com/kazzmir/opus-go.svg)](https://pkg.go.dev/github.com/kazzmir/opus-go)

This is a pure go implementation of an Ogg/Opus parser, decoder, and encoder. It was produced by transpiling the libopus C sources to Go using [ccgo](https://pkg.go.dev/modernc.org/ccgo/v4), as well as using GPT 5.2 help. No cgo is used.

Directory structure:
- `cmd/oggopusdump`: command-line tool to dump Ogg/Opus headers and packet sizes
- `cmd/oggopusextract`: command-line tool to extract Opus audio packets
- `cmd/oggopus2wav`: command-line tool to decode Ogg/Opus
- `cmd/wav2oggopus`: command-line tool to encode WAV to Ogg/Opus
- `ogg` - Parser for the Ogg format and Opus packet reader
- `opus` - Encoder and decoder opus API
- `opuscc` - Transpiled libopus C source of the decoder logic
- `opusccenc` - Transpiled libopus C source of the encoder logic
- `libcshim` - Small libc shim for the transpiled C code, replaces some modernc.org/libc functionality
- `examples` - Example programs using the library

## Minimal high level API example
Decoding an opus file to get PCM samples. Note the sample rate of the PCM data is always 48000 Hz.
```go
player, _ := opusgo.NewPlayerFromFile("file.opus", true) // true means stream from disk
data := make([]byte, 48000*2*2) // 1 second buffer for stereo s16le
io.ReadFull(player, data)
// data now contains PCM samples
```

## Minimal low level API example

Decoding an opus file to get PCM samples. Note the sample rate of the PCM data is always 48000 Hz.
```go
// error handling omitted for brevity
input, _ := os.Open("file.opus")
reader, _ := ogg.NewOpusReader(input) // input can be any io.Reader
decoder, _ := opus.NewDecoderFromHead(reader.Head)
for {
  packet, err := reader.ReadAudioPacket()
  if err != nil {
    // no more audio packets
    break
  }
  decoded, n, _ := decoder.DecodePacket(packet, nil) // nil can instead be a pre-allocated []int16, otherwise new memory is allocated
  // decoded is an []int16 PCM audio buffer with n samples per channel, so total samples = n * reader.Head.Channels
  // use decoded PCM samples...
}
```

## Example program usage

From this folder:

- Dump headers + packet sizes:

```sh
go run ./cmd/oggopusdump --no-crc path/to/file.ogg
```

- Extract length-prefixed audio packets:

```sh
go run ./cmd/oggopusextract --out packets.bin path/to/file.ogg
```

`packets.bin` format: repeating `(u32le length, length bytes packet)`.

- Decode `.opus` to `.wav` (48kHz s16le):

```sh
go run ./cmd/oggopus2wav --out out.wav path/to/file.opus
```

- Profile decoding:

```sh
go run ./cmd/oggopus2wav --cpuprofile cpu.pprof --out out.wav path/to/file.opus
go tool pprof -top cpu.pprof
```

- Encoding `.wav` to `.opus`:

```sh
go run ./cmd/wav2oggopus --bitrate 64000 --out out.opus path/to/file.wav
```

  Input is 16-bit PCM WAV, mono or stereo, at any sample rate: 8/12/16/24/48 kHz
  are encoded directly and anything else (e.g. 44.1 kHz) is resampled to 48 kHz
  first (package `resample`). Pass `--cpuprofile cpu.pprof` to profile it.

## Local CI gate

GitHub Actions are disabled on this repository. `make ci-local` is THE gate and must pass before every merge and every
release tag: gitleaks (full history), gofmt, `go mod tidy` drift, govulncheck, vet, golangci-lint, `go test -race`,
`GOARCH=386 go test`, and cross-compile vets for darwin/arm64 and linux/arm64 (the closest local substitute for the
retired arm64 and macOS runners).
