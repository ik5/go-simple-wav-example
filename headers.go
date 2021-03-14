package main

// RIFFChunk holds basic information about a given RIFF format.
type RIFFChunk struct {
	RIFFID     [4]byte // RIFF Header Magic header
	ChunkSize  uint32  // RIFF Chunk Size
	WaveHeader [4]byte // WAVE Header
}

// FormatChunk specifies the format of the data.
type FormatChunk struct {
	FormatID      [4]byte     // FMT header
	Size          uint32      // Size of the fmt chunk (16, 18 or 40)
	Format        WaveFormat  // Audio format
	ChannelsNum   ChannelType // Number of channels
	SamplesPerSec uint32      // Sampling Per seconds
	BytesPerSec   uint32      // average bytes per seconds
	BlockAlign    uint16      // Data block size
	BitsPerSample uint16      // Number of bits per sample
	DataID        [4]byte     // "data"  string
	DataSize      uint32      // Sampled data length
}
