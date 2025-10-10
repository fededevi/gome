package input

import (
	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevKeys = map[ebiten.Key]bool{}

// GetMenuCommand returns a command only once per key press
func GetMenuCommand() menu.Command {
	keys := []struct {
		key     ebiten.Key
		command menu.Command
	}{
		{ebiten.KeyArrowUp, menu.CmdUp},
		{ebiten.KeyArrowDown, menu.CmdDown},
		{ebiten.KeyArrowLeft, menu.CmdLeft},
		{ebiten.KeyArrowRight, menu.CmdRight},
		{ebiten.KeyEnter, menu.CmdSelect},
		{ebiten.KeyEscape, menu.CmdBack},
	}

	for _, k := range keys {
		pressed := ebiten.IsKeyPressed(k.key)
		justPressed := pressed && !prevKeys[k.key] // edge detection
		prevKeys[k.key] = pressed
		if justPressed {
			return k.command
		}
	}

	return menu.CmdNone
}
