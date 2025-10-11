package audio

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type AudioData struct {
	Data       []byte
	SampleRate int
	Channels   int
	BitDepth   int
}

// LoadWav loads a WAV file into memory as AudioData
func LoadWav(path string) (*AudioData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stream, err := wav.DecodeWithoutResampling(f)
	if err != nil {
		return nil, err
	}

	data := make([]byte, stream.Length())
	_, err = io.ReadFull(stream, data)
	if err != nil {
		return nil, err
	}

	return &AudioData{
		Data:       data,
		SampleRate: stream.SampleRate(),
		Channels:   2,
		BitDepth:   16,
	}, nil
}

// LoadMP3 loads an MP3 file into memory as AudioData
func LoadMP3(path string) (*AudioData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stream, err := mp3.DecodeWithoutResampling(f)
	if err != nil {
		return nil, err
	}

	data := make([]byte, stream.Length())
	_, err = io.ReadFull(stream, data)
	if err != nil {
		return nil, err
	}

	return &AudioData{
		Data:       data,
		SampleRate: stream.SampleRate(),
		Channels:   2,
		BitDepth:   16,
	}, nil
}

// MustLoadAudio decides which loader to use based on file extension
func MustLoadAudio(path string) *AudioData {
	var data *AudioData
	var err error

	switch ext := path[len(path)-4:]; ext {
	case ".wav", ".WAV":
		data, err = LoadWav(path)
	case ".mp3", ".MP3":
		data, err = LoadMP3(path)
	default:
		panic("unsupported audio format: " + path)
	}

	if err != nil {
		panic(fmt.Sprintf("failed to load audio %s: %v", path, err))
	}

	return data
}
