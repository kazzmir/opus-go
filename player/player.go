package player

import (
    "fmt"
    "io"
    "os"
    "bytes"
    "bufio"
    "time"
    // "log"

    "github.com/kazzmir/opus-go/ogg"
    "github.com/kazzmir/opus-go/opus"
)

type OpusPlayer struct {
    reader *ogg.OpusReader
    decoder *opus.Decoder
    buffer []int16
    preskipRemaining int64
    position int
    finished bool
    // how many samples have been read so far
    totalSamples int64
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
        preskipRemaining: int64(opusReader.Head.PreSkip),
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
        // note that we rely on the garbage collector to close the file when the player is done with it
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

// returns true when the stream has finished and all bytes decoded
// have been read, meaning when the internal buffer is empty
func (player *OpusPlayer) IsFinished() bool {
    return player.finished && player.position >= len(player.buffer)
}

// read as much data as possible into p
func (player *OpusPlayer) Read(p []byte) (int, error) {
    total := 0
    for total < len(p) {
        n, err := player.ReadPacket(p[total:])
        total += n
        if err != nil {
            return total, err
        }
    }

    return total, nil
}

// read at most one opus packet's worth of audio into p
// the length of p should not be less than 4 if the opus stream is mono
// and should not be less than 2 if the opus stream is stereo
func (player *OpusPlayer) ReadPacket(p []byte) (int, error) {
    // we have len(p) bytes to fill up, which is len(p)/2 int16 samples
    // we are going to produce stereo audio, so the number of samples read per channel will be len(p)/4

    // log.Printf("OpusPlayer read: p=%d position=%d buffer=%d finished=%v", len(p), player.position, len(player.buffer), player.finished)

    // fmt.Printf("Current sample: %v\n", player.totalSamples)

    if player.position >= len(player.buffer) {
        packet, err := player.reader.ReadAudioPacket()
        if err != nil {
            player.finished = true
            return 0, err

            // fill with silence
            /*
            for i := range p {
                p[i] = 0
            }

            return len(p), err
            */
        }

        // fmt.Printf("Packet granule: %v valid: %v sequence: %v eos: %v\n", packet.GranulePos, packet.GranuleValid, packet.PageSequence, packet.EOS)

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
            skip := min(int64(n), player.preskipRemaining)
            player.preskipRemaining -= skip
            player.position += int(skip) * int(player.reader.Head.Channels)
        }

        // fmt.Printf("Decoded samples: %d buffer length: %d\n", n, len(player.buffer))

        // discard excess samples based on granule position
        if packet.GranuleValid {
            actual := uint64(player.totalSamples + int64(n + int(player.reader.Head.PreSkip)))
            maxSamples := packet.GranulePos

            // this page's granule position indicates that we should drop some of the decoded samples
            if actual > maxSamples {
                excessSamples := actual - maxSamples
                upper := excessSamples * uint64(player.reader.Head.Channels)
                if upper < uint64(len(player.buffer)) {
                    player.buffer = player.buffer[:len(player.buffer) - int(upper)]
                }
            }
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
            player.totalSamples += int64(count)

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
            player.totalSamples += int64(count / 2)

            return count * 2, nil
    }

    return 0, fmt.Errorf("unsupported number of channels: %d", player.reader.Head.Channels)
}

func (player *OpusPlayer) Channels() int {
    return 2 // always stereo output
}

func (player *OpusPlayer) SampleRate() int {
    return ogg.OpusSampleRateHz
}

// offset is a byte position within the decoded PCM stream
// whence is one of io.SeekStart, io.SeekCurrent, io.SeekEnd
// returns the new offset in bytes from the start of the stream
func (player *OpusPlayer) Seek(offset int64, whence int) (int64, error) {
    // 1 sample = 2 bytes per channel, where the decoded stream is always stereo
    byteToSample := func(b int64) int64 {
        return b / 4
    }

    offset = byteToSample(offset)

    var err error

    switch whence {
        case io.SeekStart:
            err = player.SeekSample(uint64(offset))
        case io.SeekCurrent:
            n := max(0, offset + player.totalSamples)
            err = player.SeekSample(uint64(n))
        case io.SeekEnd:
            length := byteToSample(player.Length())
            n := max(0, offset + length)
            err = player.SeekSample(uint64(n))
        default:
            err = fmt.Errorf("invalid whence: %d", whence)
    }

    return player.totalSamples * 4, err
}

// total length in bytes of the decoded stream(not samples)
// note that this method reads packets to determine the end of the stream
// if the underlying reader is seekable, you may want to seek back to the start after calling this method.
// the length is cached, however, so it is safe and efficient to call multiple times on the same stream
func (player *OpusPlayer) Length() int64 {
    total, err := player.reader.TotalSamples()
    if err != nil {
        return 0
    }
    return total * 4
}

// position is a number of samples (not bytes) from the start of the stream
// e.g., 0 is the start of the stream (after preskip), and the last available position is
// the total samples - 1 (which is the same as the last granule position - preskip)
func (player *OpusPlayer) SeekSample(position uint64) error {
    // granule positions must take preskip into account
    position += uint64(player.reader.Head.PreSkip)

    // force reader to go back to the page that contains the desired position
    granule, err := player.reader.SeekToPage(position)
    if err != nil {
        return err
    }

    // reset decoder state
    // in theory, the decoder state should start fresh from a new audio page
    // after the end of a previous valid granule page
    decoder, err := opus.NewDecoderFromHead(player.reader.Head)
    if err != nil {
        return err
    }
    player.decoder = decoder

    // how many samples to skip in this packet sequence
    // if preskip is larger than granule, we need to skip less
    // e.g., preskip = 2500, granule = 2000, that means that sample 2500 is the first 'real' sample
    // so position=0 should skip 500 samples
    skipSamples := position - granule
    player.totalSamples = int64(granule)

    // preskip := max(0, int64(player.reader.Head.PreSkip) - int64(granule))
    preskip := int64(player.reader.Head.PreSkip)
    player.preskipRemaining = preskip

    // the first page we read will contain samples starting from granule
    // the position variable is the sample we are moving towards
    // keep reading packets and decoding samples until we reach the desired position
    for skipSamples > 0 {
        packet, err := player.reader.ReadAudioPacket()
        if err != nil {
            return err
        }

        player.finished = packet.EOS
        player.position = 0
        player.buffer = player.buffer[:cap(player.buffer)]

        decoded, n, err := player.decoder.DecodePacket(packet, player.buffer)
        if err != nil {
            return err
        }

        // FIXME: handle excess samples based on granule position when packet.GranuleValid is true

        move := min(uint64(n), skipSamples)
        skipSamples -= move

        // if move is less than n then we need to keep the remaining samples in the buffer
        player.buffer = decoded
        player.position += int(move * uint64(player.reader.Head.Channels))

        drop := min(int64(move), preskip)
        preskip -= drop
        player.preskipRemaining -= drop

        player.totalSamples += int64(move) - drop
    }

    return nil
}

// Seek to the position specified by the argument in terms of time
func (player *OpusPlayer) SeekTime(when time.Duration) error {
    samples := uint64(when * time.Duration(ogg.OpusSampleRateHz) / time.Second)
    return player.SeekSample(samples)
}

// Current position in samples. This is independent of the number of channels the
// underlying opus stream has. Basically this is the number of stereo samples in the
// decoded PCM stream
func (player *OpusPlayer) CurrentSample() int64 {
    return player.totalSamples
}

// Current position in terms of time
func (player *OpusPlayer) CurrentTime() time.Duration {
    return time.Duration(player.totalSamples) * time.Second / time.Duration(ogg.OpusSampleRateHz)
}
