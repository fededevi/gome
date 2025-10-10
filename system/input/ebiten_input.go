package input

import (
	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetMenuCommand() menu.Command {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		return menu.CmdUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		return menu.CmdDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return menu.CmdSelect
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return menu.CmdBack
	}
	return menu.CmdNone
}
