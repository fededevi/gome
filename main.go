package main

import (
	"errors"
	"log"

	"gome/system/input"
	"gome/system/menu"
	"gome/system/render"

	"gome/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	menu *menu.Menu
}

func (g *Game) Update() error {
	cmd := input.GetMenuCommand()
	g.menu.HandleCommand(cmd)

	if game.QuitRequested {
		return errors.New("quit requested")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DrawMenu(screen, g.menu)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	menu.LinkParents(game.GameMenuTree)
	g := &Game{menu: menu.New(game.GameMenuTree)}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gome Menu System")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
