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

## Minimal API example

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
