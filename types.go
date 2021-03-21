package main

import "fmt"

// WaveFormat holds information regarding the Codec/Format a wav chunk holds.
type WaveFormat uint16

// ChannelType holds the number of channels to use
type ChannelType uint16

func (wf WaveFormat) String() string {
	var format string
	switch wf {
	case PCM:
		format = "PCM "
	case IEEEFloat:
		format = "IEEE Float "
	case ALaw:
		format = "A-Law "
	case MuLaw:
		format = "µ-law "
	case Extensible:
		format = "Extension "
	}

	return fmt.Sprintf("%s0x%x", format, int16(wf))
}

func (ct ChannelType) String() string {
	switch ct {
	case Mono:
		return "mono"
	case Stereo:
		return "stereo"
	default:
		return fmt.Sprintf("%d", ct)
	}
}
