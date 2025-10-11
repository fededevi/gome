package render

import (
	"fmt"
	"image/color"

	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

func DrawMenu(screen *ebiten.Image, m *menu.Menu) {
	children := m.Current.Get().Children()
	for i, item := range children {
		y := 100 + i*20

		display := item.Label()

		// Handle sliders
		if s, ok := item.(*menu.Slider); ok {
			display = fmt.Sprintf("%s: %d", s.Label(), s.Value())
		}

		// Handle enum items
		if e, ok := item.(*menu.EnumItem); ok {
			display = fmt.Sprintf("%s: %s", e.Label(), e.Value())
		}

		// Handle submenus
		if _, ok := item.(*menu.Submenu); ok {
			display += " >"
		}

		// Choose color based on selection
		var col color.Color
		if i == m.Selected.Get() {
			col = color.RGBA{255, 200, 0, 255} // gold/yellow for selected
		} else {
			col = color.White
		}

		text.Draw(screen, display, basicfont.Face7x13, 100, y, col)
	}
}
