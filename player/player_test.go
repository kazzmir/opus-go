package player

import (
    "testing"
    "time"
    "bytes"
    "io"
    "math"
    _ "fmt"
)

const testFilePath = "../test/music_64kbps.opus"

func TestBasic(test *testing.T) {
    player, err := NewPlayerFromFile(testFilePath, true)
    if err != nil {
        test.Fatalf("Failed to create player: %v", err)
    }

    if player.CurrentSample() != 0 {
        test.Fatalf("Expected current sample to be 0, got %d", player.CurrentSample())
    }
}

func absTime(d time.Duration) time.Duration {
    if d < 0 {
        return -d
    }
    return d
}

func TestSeek1(test *testing.T) {
    player, err := NewPlayerFromFile(testFilePath, true)
    if err != nil {
        test.Fatalf("Failed to create player: %v", err)
    }

    if player.CurrentSample() != 0 {
        test.Fatalf("Expected current sample to be 0, got %d", player.CurrentSample())
    }

    err = player.SeekSample(1)
    if err != nil {
        test.Fatalf("Failed to seek to sample 1: %v", err)
    }

    if player.CurrentSample() != 1 {
        test.Fatalf("Expected current sample to be 1 after seeking, got %d", player.CurrentSample())
    }

    err = player.SeekTime(1 * time.Second)
    if err != nil {
        test.Fatalf("Failed to seek to 1 second: %v", err)
    }

    if player.CurrentSample() != int64(player.SampleRate()) {
        test.Fatalf("Expected current sample to be %d after seeking to 1 second, got %d", player.SampleRate(), player.CurrentSample())
    }

    if absTime(player.CurrentTime() - 1*time.Second) > 1*time.Millisecond {
        test.Fatalf("Expected current time to be approximately 1 second after seeking, got %v", player.CurrentTime())
    }

    err = player.SeekSample(1 << 40)
    if err == nil {
        test.Fatalf("Expected error when seeking beyond file length, but got none")
    }
}

// this test decodes the entire stream into memory, then seeks to a position and reads data.
// the read data should exactly equal the corresponding data in the full stream at the same position
func TestSeek2(test *testing.T) {
    player, err := NewPlayerFromFile(testFilePath, true)
    if err != nil {
        test.Fatalf("Failed to create player: %v", err)
    }

    player2, err := NewPlayerFromFile(testFilePath, true)
    if err != nil {
        test.Fatalf("Failed to create second player: %v", err)
    }
    // decode the entire stream into memory
    var fullStream bytes.Buffer
    n, err := io.Copy(&fullStream, player2)
    if err != nil {
        test.Fatalf("Failed to decode full stream: %v", err)
    }

    if n != player.Length() {
        test.Fatalf("Expected full stream length %d, got %d", player.Length(), n)
    }

    position := int64(5000 * 4)
    where, err := player.Seek(position, io.SeekStart)
    if err != nil {
        test.Fatalf("Failed to seek to position %d: %v", position, err)
    }

    if where != position {
        test.Fatalf("Expected seek position %d, got %d", position, where)
    }

    decoded := make([]byte, 800 * 4)
    decodedLength, err := player.Read(decoded)
    if err != nil {
        test.Fatalf("Failed to read after seeking: %v", err)
    }

    if decodedLength != len(decoded) {
        test.Fatalf("Expected to read %d bytes, got %d", len(decoded), decodedLength)
    }

    /*
    check := 80
    fmt.Printf("First %v bytes\n", check)
    fmt.Printf("Decoded:    ")
    for i := range check {
        fmt.Printf("%02x ", decoded[i])
    }
    fmt.Println()

    fmt.Printf("Fullstream: ")
    for i := range check {
        fmt.Printf("%02x ", fullStream.Bytes()[position+int64(i)])
    }
    fmt.Println()
    */

    index := bytes.Index(fullStream.Bytes(), decoded)
    if index != int(position) {

        b := fullStream.Bytes()[position:]
        for i := range decoded {
            if decoded[i] != b[i] {
                test.Fatalf("Data mismatch at byte %d after seeking, expected %02x, got %02x", i, b[i], decoded[i])
            }
        }

        test.Fatalf("Decoded data does not match expected data after seeking, expected index %d, got %d", position, index)
    }

    for i := range 6 {
        position := int(math.Pow(7, float64(i + 1))) * 4
        player.Seek(int64(position), io.SeekStart)

        decodedLength, err := player.Read(decoded)
        if err != nil {
            test.Fatalf("Failed to read after seeking to position %d: %v", position, err)
        }
        if decodedLength != len(decoded) {
            test.Fatalf("Expected to read %d bytes after seeking to position %d, got %d", len(decoded), position, decodedLength)
        }

        index := bytes.Index(fullStream.Bytes(), decoded)
        if index != int(position) {

            b := fullStream.Bytes()[position:]
            for i := range decoded {
                if decoded[i] != b[i] {
                    test.Fatalf("Data mismatch at byte %d after seeking, expected %02x, got %02x", i, b[i], decoded[i])
                }
            }

            test.Fatalf("Decoded data does not match expected data after seeking to position %d, expected index %d, got %d", position, position, index)
        }
    }

}
