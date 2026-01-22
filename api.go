package opusgo

// this top level module re-exports nested packages

import (
    "io"

    "github.com/kazzmir/opus-go/player"
)

type OpusPlayer = player.OpusPlayer

func NewPlayerFromReader(reader io.Reader) (*OpusPlayer, error) {
    return player.NewPlayerFromReader(reader)
}

func NewPlayerFromFile(path string, stream bool) (*OpusPlayer, error) {
    return player.NewPlayerFromFile(path, stream)
}
