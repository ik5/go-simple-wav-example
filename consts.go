package main

// The following constants represents WAV codecs
// The constants does not represents all available Codecs
const (
	PCM        WaveFormat = 0x0001
	IEEEFloat  WaveFormat = 0x0003
	ALaw       WaveFormat = 0x0006
	MuLaw      WaveFormat = 0x0008
	Extensible WaveFormat = 0xFFFE
)

// Header constants
var (
	RIFF = [4]byte{'R', 'I', 'F', 'F'}
	WAVE = [4]byte{'W', 'A', 'V', 'E'}
	FMT  = [4]byte{'f', 'm', 't', ' '}
	DATA = [4]byte{'d', 'a', 't', 'a'}
)

// Type of channels
const (
	Mono   ChannelType = iota + 1 // single channel (mono)
	Stereo                        // dual channels (stereo)
)
