# Go Ogg/Opus packet extractor (no cgo)

This folder contains a small Go module that can:

- Parse Ogg pages and reassemble laced packets
- Parse `OpusHead` and `OpusTags`
- Stream Opus **audio** packets (`[]byte`) that you can pass to an Opus decoder
- Decode Ogg Opus to signed 16-bit PCM via a pure-Go libopus translation (no cgo)

The Opus decoder is backed by a ccgo-transpiled build of the libopus C sources.

The generated bindings use a small stdlib-only runtime (`opusgo/libcshim`) that provides the minimal TLS/allocator/varargs functionality needed by the translated C.

Concurrency: the ccgo build uses `NONTHREADSAFE_PSEUDOSTACK` for speed, but patched so its scratch state is stored per-decoder/TLS (no shared global variables). The Go `opus.Decoder` also serializes `Decode`/`Close` with a mutex so using a single decoder from multiple goroutines won’t race (but it will decode sequentially).

## Usage

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
