package game

import "gome/system/audio"

// These will hold the loaded sounds
var (
	ClickSound     *audio.Sound
	ExplosionSound *audio.Sound
	BgmSound       *audio.Sound
	AllSounds      []*audio.Sound
)

// InitializeSounds sets up all sounds with proper channels
func InitializeSounds(as *audio.AudioSystem) *audio.AudioSystem {
	sounds := CreateSounds(as.Effects, as.Music)

	// Assign to globals for easy access
	ClickSound = sounds[0]
	ExplosionSound = sounds[1]
	BgmSound = sounds[2]
	AllSounds = sounds

	return as
}

// CreateSounds creates the list of sounds with channels assigned
func CreateSounds(effects, music *audio.AudioChannel) []*audio.Sound {
	return []*audio.Sound{
		{
			Data:    audio.MustLoadAudio("game/assets/audio/menu_navigation.mp3"),
			Channel: effects,
			Volume:  1.0,
		},
		{
			Data:    audio.MustLoadAudio("game/assets/audio/menu_navigation.mp3"),
			Channel: effects,
			Volume:  1.0,
		},
		{
			Data:    audio.MustLoadAudio("game/assets/audio/menu_navigation.mp3"),
			Channel: music,
			Volume:  0.5,
		},
	}
}
