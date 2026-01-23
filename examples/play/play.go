package main

import (
    "os"
    "fmt"
    "time"
    "runtime"
    "log"
    "io"

    "github.com/kazzmir/opus-go"
    "github.com/ebitengine/oto/v3"
)

func play(filename string) error {

    opusPlayer, err := opusgo.NewPlayerFromFile(filename, true)
    if err != nil {
        return err
    }

    var options oto.NewContextOptions
    options.SampleRate = 48000
    options.ChannelCount = 2
    options.Format = oto.FormatSignedInt16LE

    context, ready, err := oto.NewContext(&options)
    if err != nil {
        return err
    }

    log.Printf("Waiting for audio context to be ready")
    <-ready

    log.Printf("Playing")

    log.Printf("Length: %v", opusPlayer.Length())

    // opusPlayer.Seek(80000)
    // err = opusPlayer.SeekTime(42382 * time.Millisecond)
    n, err := opusPlayer.Seek(-5, io.SeekEnd)
    if err != nil {
        log.Printf("Error seeking: %v", err)
    } else {
        log.Printf("Seeked to byte position: %d", n)
    }

    log.Printf("Current position: %v", opusPlayer.CurrentPosition())
    log.Printf("Current time: %v", opusPlayer.CurrentTime())

    otoPlayer := context.NewPlayer(opusPlayer)
    otoPlayer.SetBufferSize(options.SampleRate * 2 * 2 / 4) // 250ms buffer
    otoPlayer.Play()

    if otoPlayer.Err() != nil {
        return otoPlayer.Err()
    }

    for !opusPlayer.IsFinished() {
        time.Sleep(1 * time.Second)
        runtime.KeepAlive(otoPlayer)
    }

    // wait just a bit longer for the last audio packet to finish playing
    time.Sleep(500 * time.Millisecond)
    runtime.KeepAlive(otoPlayer)

    return nil
}

func main() {
    log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
    if len(os.Args) < 2 {
        fmt.Println("Provide a .ogg or .opus file to play")
        return
    }

    filename := os.Args[1]

    err := play(filename)
    if err != nil {
        fmt.Printf("Error playing file: %v\n", err)
    }

    log.Println("Exiting")
}
