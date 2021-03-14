package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file.pcm> <file.wav>", os.Args[0])
		os.Exit(-1)
		return
	}
	data, err := os.ReadFile(os.Args[1])
	s := uint32(unsafe.Sizeof(FormatChunk{}))
	l := uint32(len(data))

	riff := RIFFChunk{
		RIFFID:     RIFF,
		ChunkSize:  l + s - 8,
		WaveHeader: WAVE,
	}

	chunk := FormatChunk{
		FormatID:      FMT,
		Size:          16,
		Format:        PCM,
		ChannelsNum:   Mono,
		SamplesPerSec: 8000,
		BitsPerSample: 16,
		DataID:        DATA,
		DataSize:      l + s - 44,
	}
	chunk.BytesPerSec = chunk.SamplesPerSec * uint32(chunk.ChannelsNum) * uint32(chunk.BitsPerSample) / 8
	chunk.BlockAlign = uint16(chunk.ChannelsNum) * chunk.BitsPerSample / 8

	err = WriteWavFile(os.Args[2], riff, chunk, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
