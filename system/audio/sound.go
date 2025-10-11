package audio

type Sound struct {
	Data    *AudioData
	Channel *AudioChannel
	Volume  float64 // per-play volume 0.0-1.0
}
