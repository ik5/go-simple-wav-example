package main

import (
	"encoding/binary"
	"os"
)

// WriteWavFile Generate a new wav file based on data
func WriteWavFile(filename string, riff RIFFChunk, format FormatChunk, data []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	binary.Write(f, binary.LittleEndian, riff)
	binary.Write(f, binary.LittleEndian, format)
	binary.Write(f, binary.LittleEndian, data)
	f.Sync()
	return nil
}
