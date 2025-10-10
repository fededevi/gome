package render

import (
	"fmt"

	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawMenu(screen *ebiten.Image, m *menu.Menu) {
	children := m.Current.Children()
	for i, item := range children {
		y := 100 + i*20
		prefix := "  "
		if i == m.Selected {
			prefix = "> "
		}

		display := item.Label()

		// Handle sliders
		if s, ok := item.(*menu.Slider); ok {
			display = fmt.Sprintf("%s: %d", s.Label(), s.Value())
		}

		// Handle enum items
		if e, ok := item.(*menu.EnumItem); ok {
			display = fmt.Sprintf("%s: %s", e.Label(), e.Value())
		}

		// Handle submenus (optional marker)
		if _, ok := item.(*menu.Submenu); ok {
			display += " ▶"
		}

		ebitenutil.DebugPrintAt(screen, prefix+display, 100, y)
	}
}
