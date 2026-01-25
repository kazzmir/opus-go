package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/ebitengine/oto/v3"
	opusgo "github.com/kazzmir/opus-go"
)

func playURL(ctx context.Context, url string) error {
	log.Printf("Fetching: %s", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "opus-go/examples-net")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 8*1024))
		return fmt.Errorf("http %s: %s", resp.Status, string(body))
	}

	opusPlayer, err := opusgo.NewPlayerFromReader(resp.Body)
	if err != nil {
		return err
	}

	var options oto.NewContextOptions
	options.SampleRate = opusgo.OpusSampleRateHz
	options.ChannelCount = 2
	options.Format = oto.FormatSignedInt16LE

	audioContext, ready, err := oto.NewContext(&options)
	if err != nil {
		return err
	}
	// defer audioContext.Close()

	log.Printf("Waiting for audio context to be ready")
	<-ready

	log.Printf("Playing")
	otoPlayer := audioContext.NewPlayer(opusPlayer)

	// 250ms buffer: sampleRate * channels * 2 bytes/sample / 4
	otoPlayer.SetBufferSize(options.SampleRate * options.ChannelCount * 2 / 4)
	otoPlayer.Play()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Stop the network stream; this will eventually propagate to the decoder.
			_ = resp.Body.Close()
			return ctx.Err()

		case <-ticker.C:
			if otoPlayer.Err() != nil {
				return otoPlayer.Err()
			}

			if opusPlayer.IsFinished() {
				// Wait just a bit longer for the last audio packet to finish playing.
				time.Sleep(500 * time.Millisecond)
				runtime.KeepAlive(otoPlayer)
				return nil
			}

			runtime.KeepAlive(otoPlayer)
		}
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <http(s)://...>\n", os.Args[0])
		fmt.Println("The URL should be an Ogg Opus stream/file (.opus/.ogg containing Opus).")
		os.Exit(2)
	}

	url := os.Args[1]

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := playURL(ctx, url); err != nil {
		// Treat Ctrl+C as a normal exit.
		if err == context.Canceled {
			log.Printf("Stopped")
			return
		}
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
