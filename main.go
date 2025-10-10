package main

import (
	"fmt"
	"log"

	"gome/system/input"
	"gome/system/menu"
	"gome/system/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	menu *menu.Menu
}

func (g *Game) Update() error {
	cmd := input.GetMenuCommand()
	g.menu.HandleCommand(cmd)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DrawMenu(screen, g.menu)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func buildMenu() *menu.Item {
	root := &menu.Item{Label: "Main Menu"}

	start := &menu.Item{Label: "Start Game", Action: func() {
		fmt.Println("Starting game…")
	}}
	settings := &menu.Item{Label: "Settings"}
	audio := &menu.Item{Label: "Audio"}
	video := &menu.Item{Label: "Video"}
	settings.Children = []*menu.Item{audio, video}
	quit := &menu.Item{Label: "Quit", Action: func() {
		fmt.Println("Exiting game…")
	}}

	root.Children = []*menu.Item{start, settings, quit}
	menu.LinkParents(root)
	return root
}

func main() {
	root := buildMenu()
	g := &Game{menu: menu.New(root)}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gome Menu System")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
