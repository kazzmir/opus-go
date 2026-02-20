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
func NewPlayerFromReader(reader io.Reader) (*OpusPlayer[int16], error) {
    return player.NewPlayerFromReader(reader)
}

// Create a new player from a file path. If stream is true, the file will be streamed instead of fully loaded into memory.
// Note that internally the file object is closed when the garbage collector collects the player.
func NewPlayerFromFile(path string, stream bool) (*OpusPlayer[int16], error) {
    return player.NewPlayerFromFile(path, stream)
}
