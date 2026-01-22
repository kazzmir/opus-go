package player

import (
    "fmt"
    "io"
    "os"
    "bytes"
    "bufio"

    "github.com/kazzmir/opus-go/ogg"
    "github.com/kazzmir/opus-go/opus"
)

type OpusPlayer struct {
    reader *ogg.OpusReader
    decoder *opus.Decoder
    buffer []int16
    preskipRemaining int
    position int
    finished bool
}

func NewPlayerFromReader(reader io.Reader) (*OpusPlayer, error) {
    opusReader, err := ogg.NewOpusReader(reader)
    if err != nil {
        return nil, err
    }

    decoder, err := opus.NewDecoderFromHead(opusReader.Head)
    if err != nil {
        return nil, err
    }

    return &OpusPlayer{
        reader: opusReader,
        decoder: decoder,
        preskipRemaining: int(opusReader.Head.PreSkip),
        position: 0,
        // buffer does not need to be initialized here because it will be allocated on first read
    }, nil
}

// Create a new OpusPlayer from a file path
// If stream is true, the file will be kept open and 
// otherwise if stream is false then the entirety of the file will be read into memory
func NewPlayerFromFile(path string, stream bool) (*OpusPlayer, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    if stream {
        // dont need bufio here because OggOpusReader already uses bufio internally
        return NewPlayerFromReader(file)
    } else {
        defer file.Close()
        var data bytes.Buffer
        _, err := io.Copy(&data, bufio.NewReader(file))
        if err != nil {
            return nil, err
        }

        return NewPlayerFromReader(bytes.NewReader(data.Bytes()))
    }
}

func (player *OpusPlayer) IsFinished() bool {
    return player.finished && player.position >= len(player.buffer)
}

func (player *OpusPlayer) Read(p []byte) (int, error) {
    // we have len(p) bytes to fill up, which is len(p)/2 int16 samples
    // we are going to produce stereo audio, so the number of samples read per channel will be len(p)/4

    if player.position >= len(player.buffer) {
        packet, err := player.reader.ReadAudioPacket()
        if err != nil {
            // fill with silence
            for i := range p {
                p[i] = 0
            }

            return len(p), err
        }

        if packet.EOS {
            player.finished = true
        }

        player.buffer = player.buffer[:cap(player.buffer)]
        player.position = 0

        // DecodePacket will re-allocate the buffer if necessary
        decoded, n, err := player.decoder.DecodePacket(packet, player.buffer)
        if err != nil {
            return 0, err
        }

        player.buffer = decoded
        if player.preskipRemaining > 0 {
            skip := min(n, player.preskipRemaining)
            player.preskipRemaining -= skip
            player.position += skip * int(player.reader.Head.Channels)
        }
    }

    switch player.reader.Head.Channels {
        case 1:
            // we have to produce stereo audio, so each input sample becomes two output samples
            atMost := min(len(p) / 4, len(player.buffer) - player.position)
            // log.Printf("Rendering opus: p=%d buffer=%d atMost=%d position=%d", len(p), len(player.buffer), atMost, player.position)
            count := 0
            for count < atMost {
                low := byte(player.buffer[player.position + count] & 0xFF)
                high := byte((player.buffer[player.position + count] >> 8) & 0xFF)

                p[count*4] = low
                p[count*4+1] = high
                p[count*4+2] = low
                p[count*4+3] = high
                count += 1
            }
            player.position += count

            return count * 4, nil

        case 2:
            atMost := min(len(p) / 2, len(player.buffer) - player.position)

            // log.Printf("Rendering opus: p=%d buffer=%d atMost=%d position=%d", len(p), len(player.buffer), atMost, player.position)

            count := 0
            for count < atMost {
                p[count*2] = byte(player.buffer[player.position + count] & 0xFF)
                p[count*2+1] = byte((player.buffer[player.position + count] >> 8) & 0xFF)
                count += 1
            }
            player.position += count

            return count * 2, nil
    }

    return 0, fmt.Errorf("unsupported number of channels: %d", player.reader.Head.Channels)
}
