package player

import (
    "testing"
    "time"
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

    if absTime(player.CurrentTime() - 1*time.Second) > 1*time.Millisecond {
        test.Fatalf("Expected current time to be approximately 1 second after seeking, got %v", player.CurrentTime())
    }

    err = player.SeekSample(1 << 40)
    if err == nil {
        test.Fatalf("Expected error when seeking beyond file length, but got none")
    }
}
