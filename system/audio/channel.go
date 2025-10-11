package audio

type AudioChannel struct {
	Volume float64
}

func (c *AudioChannel) SetVolume(v float64) {
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	c.Volume = v
}

func (c *AudioChannel) GetVolume() float64 {
	return c.Volume
}
