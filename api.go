package opusgo

// this top level module re-exports nested packages

import (
    "io"

    "github.com/kazzmir/opus-go/player"
    "github.com/kazzmir/opus-go/ogg"
)

type OpusPlayer[T player.DataType] = player.OpusPlayer[T]

// the decoded sample rate for Opus streams
const OpusSampleRateHz = ogg.OpusSampleRateHz

// Create a new player from an existing reader. If the reader is seekable (implements io.Seeker), seeking will be supported.
// This player produces int16 PCM samples.
func NewPlayerFromReader(reader io.Reader) (*OpusPlayer[int16], error) {
    return player.NewPlayerFromReader(reader)
}

// Similar to NewPlayerFromReader but produces float32 PCM samples in the range [-1.0, 1.0].
func NewPlayerF32FromReader(reader io.Reader) (*OpusPlayer[float32], error) {
    return player.NewPlayerF32FromReader(reader)
}

// Create a new player from a file path. If stream is true, the file will be streamed instead of fully loaded into memory.
// This player produces int16 PCM samples.
// Note that internally the file object is closed when the garbage collector collects the player.
func NewPlayerFromFile(path string, stream bool) (*OpusPlayer[int16], error) {
    return player.NewPlayerFromFile(path, stream)
}

// Similar to NewPlayerFromFile but produces float32 PCM samples in the range [-1.0, 1.0].
func NewPlayerF32FromFile(path string, stream bool) (*OpusPlayer[float32], error) {
    return player.NewPlayerF32FromFile(path, stream)
}
