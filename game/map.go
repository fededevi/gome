package game

import (
	"image"
	_ "image/png" // PNG decoder
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Map represents a game map
type Map struct {
	background *ebiten.Image
	width      int
	height     int
}

// NewMap loads the background image and returns a Map
func NewMap(backgroundPath string) *Map {
	file, err := os.Open(backgroundPath)
	if err != nil {
		log.Fatalf("failed to open background image: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode background image: %v", err)
	}

	bg := ebiten.NewImageFromImage(img)
	return &Map{
		background: bg,
		width:      bg.Bounds().Dx(),
		height:     bg.Bounds().Dy(),
	}
}

func (m *Map) Draw(screen *ebiten.Image) {
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := m.width, m.height

	// Compute scale factor to fit screen while maintaining aspect ratio
	scaleX := float64(sw) / float64(bw)
	scaleY := float64(sh) / float64(bh)
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	// Center the image
	op.GeoM.Translate(float64(sw)/2-float64(bw)*scale/2, float64(sh)/2-float64(bh)*scale/2)

	screen.DrawImage(m.background, op)
}
