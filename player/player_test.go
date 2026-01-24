package player

import (
    "testing"
)

func TestBasic(test *testing.T) {
    player, err := NewPlayerFromFile("../test/music_64kbps.opus", true)
    if err != nil {
        test.Fatalf("Failed to create player: %v", err)
    }

    if player.CurrentSample() != 0 {
        test.Fatalf("Expected current sample to be 0, got %d", player.CurrentSample())
    }
}
