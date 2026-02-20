// OpusPlayer is a high level API for decoding Ogg Opus audio streams.
// The main method is Read, which reads decoded PCM data into a byte slice.
package player

import (
    "fmt"
    "io"
    "os"
    "bytes"
    "bufio"
    "time"
    "math"
    // "log"

    "github.com/kazzmir/opus-go/ogg"
    "github.com/kazzmir/opus-go/opus"
)

type DataType interface {
    int16 | float32
}

func dataSize[T DataType](zero T) int {
    switch any(zero).(type) {
        case int16:
            return 2
        case float32:
            return 4
        default:
            panic("unsupported data type")
    }
}

type OpusPlayer[SampleT DataType] struct {
    reader *ogg.OpusReader
    decoder *opus.Decoder
    bufferInt16 []int16
    bufferFloat32 []float32
    preskipRemaining int64
    position int
    bytesPerSample int
    finished bool
    // how many samples have been read so far
    totalSamples int64
    lastTimestamp time.Duration
}

func newPlayerFromReader[T DataType](reader io.Reader) (*OpusPlayer[T], error) {
    opusReader, err := ogg.NewOpusReader(reader)
    if err != nil {
        return nil, err
    }

    decoder, err := opus.NewDecoderFromHead(opusReader.Head)
    if err != nil {
        return nil, err
    }

    var zero T

    return &OpusPlayer[T]{
        reader: opusReader,
        decoder: decoder,
        preskipRemaining: int64(opusReader.Head.PreSkip),
        position: 0,
        bytesPerSample: dataSize(zero),
        // buffer does not need to be initialized here because it will be allocated on first read
    }, nil
}

// Create a new OpusPlayer from an io.Reader. If the reader is seekable,
// then seeking within the stream will be supported. The Read() and ReadPacket() methods
// will produce int16 PCM data
//
// Internally the reader is wrapped in a bufio.Reader, so you do not have to
// wrap it yourself.
func NewPlayerFromReader(reader io.Reader) (*OpusPlayer[int16], error) {
    return newPlayerFromReader[int16](reader)
}

// Same as NewPlayerFromReader but the Read() and ReadPacket() methods will produce float32 PCM data
func NewPlayerF32FromReader(reader io.Reader) (*OpusPlayer[float32], error) {
    return newPlayerFromReader[float32](reader)
}

func newPlayerFromFile[T DataType](path string, stream bool) (*OpusPlayer[T], error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    if stream {
        // dont need bufio here because OggOpusReader already uses bufio internally
        // note that we rely on the garbage collector to close the file when the player is done with it
        return newPlayerFromReader[T](file)
    } else {
        defer file.Close()
        var data bytes.Buffer
        _, err := io.Copy(&data, bufio.NewReader(file))
        if err != nil {
            return nil, err
        }

        return newPlayerFromReader[T](bytes.NewReader(data.Bytes()))
    }
}

// Create a new OpusPlayer from a file path that produces int16 PCM data.
// If stream is true, the file will be kept open and 
// otherwise if stream is false then the entirety of the file will be read into memory.
func NewPlayerFromFile(path string, stream bool) (*OpusPlayer[int16], error) {
    return newPlayerFromFile[int16](path, stream)
}

// Same as NewPlayerFromFile but produces float32 PCM data
func NewPlayerF32FromFile(path string, stream bool) (*OpusPlayer[float32], error) {
    return newPlayerFromFile[float32](path, stream)
}

// Returns true when the stream has finished and all bytes decoded
// have been read, meaning when the internal buffer is empty.
func (player *OpusPlayer[SampleT]) IsFinished() bool {
    var zero SampleT
    switch any(zero).(type) {
        case int16:
            return player.finished && player.position >= len(player.bufferInt16)
        case float32:
            return player.finished && player.position >= len(player.bufferFloat32)
    }

    return player.finished
}

// Read as much data as possible into p. Returns the number of bytes read, and io.EOF
// if there is no more output available.
func (player *OpusPlayer[SampleT]) Read(p []byte) (int, error) {
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

// Read at most one opus packet's worth of audio into p.
// The length of p should not be less than 4 if the opus stream is mono,
// and should not be less than 2 if the opus stream is stereo.
func (player *OpusPlayer[SampleT]) ReadPacket(p []byte) (int, error) {
    // we have len(p) bytes to fill up, which is len(p)/2 int16 samples
    // we are going to produce stereo audio, so the number of samples read per channel will be len(p)/4

    // log.Printf("OpusPlayer read: p=%d position=%d buffer=%d finished=%v", len(p), player.position, len(player.buffer), player.finished)

    // fmt.Printf("Current sample: %v\n", player.totalSamples)

    var zero SampleT
    switch any(zero).(type) {
        case int16:
            return player.readPacketInt16(p)
        case float32:
            return player.readPacketFloat32(p)
    }

    return 0, fmt.Errorf("unsupported data type")
}

func (player *OpusPlayer[float32]) readPacketFloat32(p []byte) (int, error) {
    if player.position >= len(player.bufferFloat32) {
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

        player.updateTimestamp(packet.GranulePos)

        if packet.EOS {
            player.finished = true
        }

        player.bufferFloat32 = player.bufferFloat32[:cap(player.bufferFloat32)]
        player.position = 0

        // DecodePacket will re-allocate the buffer if necessary
        decoded, n, err := player.decoder.DecodePacketF32(packet, player.bufferFloat32)
        if err != nil {
            return 0, err
        }

        player.bufferFloat32 = decoded

        newSize := player.handleSkip(n, packet, len(player.bufferFloat32))
        player.bufferFloat32 = player.bufferFloat32[:newSize]
    }

    switch player.reader.Head.Channels {
        case 1:
            // we have to produce stereo audio, so each input sample becomes two output samples
            atMost := min(len(p) / 8, len(player.bufferFloat32) - player.position)
            // log.Printf("Rendering opus: p=%d buffer=%d atMost=%d position=%d", len(p), len(player.buffer), atMost, player.position)
            count := 0
            for count < atMost {
                sample := player.bufferFloat32[player.position + count]

                v := math.Float32bits(sample)
                p[count*8+0] = byte(v)
                p[count*8+1] = byte(v >> 8)
                p[count*8+2] = byte(v >> 16)
                p[count*8+3] = byte(v >> 24)

                p[count*8+4] = byte(v)
                p[count*8+5] = byte(v >> 8)
                p[count*8+6] = byte(v >> 16)
                p[count*8+7] = byte(v >> 24)

                count += 1
            }
            player.position += count
            player.totalSamples += int64(count)

            return count * 8, nil

        case 2:
            atMost := min(len(p) / 4, len(player.bufferFloat32) - player.position)

            // log.Printf("Rendering opus: p=%d buffer=%d atMost=%d position=%d", len(p), len(player.buffer), atMost, player.position)

            count := 0
            for count < atMost {
                sample := player.bufferFloat32[player.position + count]

                v := math.Float32bits(sample)
                p[count*4+0] = byte(v)
                p[count*4+1] = byte(v >> 8)
                p[count*4+2] = byte(v >> 16)
                p[count*4+3] = byte(v >> 24)

                count += 1
            }
            player.position += count
            player.totalSamples += int64(count / 2)

            return count * 4, nil
    }

    return 0, fmt.Errorf("unsupported number of channels: %d", player.reader.Head.Channels)
}

func (player *OpusPlayer[T]) handleSkip(n int, packet *ogg.OpusAudioPacket, maxLength int) int {
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
            // fmt.Printf("Dropping %d samples to match granule position at %d\n", actual - maxSamples, packet.GranulePos)

            excessSamples := actual - maxSamples
            upper := excessSamples * uint64(player.reader.Head.Channels)
            if upper < uint64(maxLength) {
                return maxLength - int(upper)
            }
        }
    }

    return maxLength
}

func (player *OpusPlayer[int16]) readPacketInt16(p []byte) (int, error) {
    if player.position >= len(player.bufferInt16) {
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

        player.updateTimestamp(packet.GranulePos)

        if packet.EOS {
            player.finished = true
        }

        player.bufferInt16 = player.bufferInt16[:cap(player.bufferInt16)]
        player.position = 0

        // DecodePacket will re-allocate the buffer if necessary
        decoded, n, err := player.decoder.DecodePacket(packet, player.bufferInt16)
        if err != nil {
            return 0, err
        }

        player.bufferInt16 = decoded

        newSize := player.handleSkip(n, packet, len(player.bufferInt16))
        player.bufferInt16 = player.bufferInt16[:newSize]
    }

    switch player.reader.Head.Channels {
        case 1:
            // we have to produce stereo audio, so each input sample becomes two output samples
            atMost := min(len(p) / 4, len(player.bufferInt16) - player.position)
            // log.Printf("Rendering opus: p=%d buffer=%d atMost=%d position=%d", len(p), len(player.buffer), atMost, player.position)
            count := 0
            for count < atMost {
                low := byte(player.bufferInt16[player.position + count] & 0xFF)
                high := byte((player.bufferInt16[player.position + count] >> 8) & 0xFF)

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
            atMost := min(len(p) / 2, len(player.bufferInt16) - player.position)

            // log.Printf("Rendering opus: p=%d buffer=%d atMost=%d position=%d", len(p), len(player.buffer), atMost, player.position)

            count := 0
            for count < atMost {
                p[count*2] = byte(player.bufferInt16[player.position + count] & 0xFF)
                p[count*2+1] = byte((player.bufferInt16[player.position + count] >> 8) & 0xFF)
                count += 1
            }
            player.position += count
            player.totalSamples += int64(count / 2)

            return count * 2, nil
    }

    return 0, fmt.Errorf("unsupported number of channels: %d", player.reader.Head.Channels)
}

// always 2 channels (stereo)
func (player *OpusPlayer[T]) Channels() int {
    return 2 // always stereo output
}

// The sample rate of the decoded PCM stream, which is always 48000 Hz for Opus
func (player *OpusPlayer[T]) SampleRate() int {
    return ogg.OpusSampleRateHz
}

// offset is a byte position within the decoded PCM stream
//
// whence is one of io.SeekStart, io.SeekCurrent, io.SeekEnd
//
// Returns the new offset in bytes from the start of the stream.
func (player *OpusPlayer[T]) Seek(offset int64, whence int) (int64, error) {
    bytesPerSample := int64(player.bytesPerSample * 2)
    // 1 sample = 2 bytes per channel, where the decoded stream is always stereo
    byteToSample := func(b int64) int64 {
        return b / bytesPerSample
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

    return player.totalSamples * bytesPerSample, err
}

// Total length in bytes of the decoded stream (not samples).
//
// note that this method reads packets to determine the end of the stream
// if the underlying reader is seekable, you may want to seek back to the start after calling this method.
// the length is cached, however, so it is safe and efficient to call multiple times on the same stream.
func (player *OpusPlayer[T]) Length() int64 {
    total, err := player.reader.TotalSamples()
    if err != nil {
        return 0
    }
    return total * 2 * int64(player.bytesPerSample)
}

// position is a number of samples (not bytes) from the start of the stream.
//
// e.g., 0 is the start of the stream (after preskip), and the last available position is
// the total samples - 1 (which is the same as the last granule position - preskip)
func (player *OpusPlayer[T]) SeekSample(position uint64) error {
    // granule positions must take preskip into account
    position += uint64(player.reader.Head.PreSkip)

    // force reader to go back to the page that contains the desired position
    granule, err := player.reader.SeekToPage(position)
    if err != nil {
        return err
    }

    player.updateTimestamp(granule)

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

    preskip := int64(player.reader.Head.PreSkip)
    // position already takes preskip into account
    player.preskipRemaining = 0
    player.totalSamples = int64(granule) - preskip

    // start a fresh buffer
    player.position = 0
    player.bufferInt16 = player.bufferInt16[:0]
    player.bufferFloat32 = player.bufferFloat32[:0]

    _, err = io.CopyN(io.Discard, player, int64(skipSamples * uint64(player.bytesPerSample) * 2))
    return err
}

// Seek to the position specified by the argument in terms of time.
func (player *OpusPlayer[T]) SeekTime(when time.Duration) error {
    samples := uint64(when * time.Duration(ogg.OpusSampleRateHz) / time.Second)
    return player.SeekSample(samples)
}

// Current position in terms of how many samples have been rendered. This is independent of the number of channels the
// underlying opus stream has. Basically this is the number of stereo samples in the
// decoded PCM stream. Seeking to the beginning of the stream will reset this to zero.
func (player *OpusPlayer[T]) CurrentSample() int64 {
    return player.totalSamples
}

// Current position in terms of time from when the player started.
func (player *OpusPlayer[T]) CurrentTime() time.Duration {
    return time.Duration(player.totalSamples) * time.Second / time.Duration(ogg.OpusSampleRateHz)
}

// Return the total number of samples in the stream per channel. This is a destructive operation,
// so you should seek back to the start if you need to read the stream again.
func (player *OpusPlayer[T]) TotalSamples() (int64, error) {
    return player.reader.TotalSamples()
}

// Return the total duration of the stream. This is a destructive operation, similar to TotalSamples.
func (player *OpusPlayer[T]) TotalDuration() (time.Duration, error) {
    return player.reader.TotalDuration()
}

// Returns the time stamp of the most recently decoded audio in the stream.
// This is the granule position of the last packet in opus terms. This can be different from CurrentTime()
// when the stream being played is a network stream that started in the past.
func (player *OpusPlayer[T]) CurrentStreamTimestamp() time.Duration {
    return player.lastTimestamp
}

func (player *OpusPlayer[T]) updateTimestamp(granule uint64) {
    player.lastTimestamp = time.Duration(granule) / time.Duration(ogg.OpusSampleRateHz) * time.Second
}
