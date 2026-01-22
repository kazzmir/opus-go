package main

import (
    "os"
    "io"
    "fmt"
    "time"

    "github.com/kazzmir/opus-go"
)

func decode(filename string) error {
    stat, err := os.Stat(filename)
    if err != nil {
        return err
    }

    opusPlayer, err := opusgo.NewPlayerFromFile(filename, true)
    if err != nil {
        return err
    }

    start := time.Now()
    n, err := io.Copy(io.Discard, opusPlayer)
    if err != nil {
        return err
    }
    elapsed := time.Since(start)
    
    fmt.Printf("Decoded %d bytes in %s (%.2f kilobytes/sec)\n", n, elapsed, float64(n)/1024/elapsed.Seconds())
    fmt.Printf("Opus decode rate: %.2f kilobytes/sec\n", float64(stat.Size())/1024/elapsed.Seconds())
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Give a .ogg or .opus file to decode")
        return
    }

    filename := os.Args[1]
    err := decode(filename)
    if err != nil {
        fmt.Printf("Error decoding file: %v\n", err)
    }
}
