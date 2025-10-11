package audio

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const DefaultSampleRate = 44100

type AudioSystem struct {
	Context *audio.Context
	Master  float64

	Music   *AudioChannel
	Effects *AudioChannel

	mu sync.Mutex
}

func NewAudioSystem() *AudioSystem {
	ctx := audio.NewContext(DefaultSampleRate)
	return &AudioSystem{
		Context: ctx,
		Master:  1.0,
		Music:   &AudioChannel{Volume: 1.0},
		Effects: &AudioChannel{Volume: 1.0},
	}
}

func (as *AudioSystem) Play(s *Sound) error {
	if s == nil || s.Data == nil {
		return fmt.Errorf("sound or audio data is nil")
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	// Use Context method (non-deprecated)
	player := as.Context.NewPlayerFromBytes(s.Data.Data)
	player.SetVolume(as.Master * s.Channel.GetVolume() * s.Volume)
	player.Play() // async

	return nil
}
