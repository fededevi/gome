package main

import (
	"errors"
	"fmt"
	"log"

	"gome/game"
	"gome/system/audio"
	"gome/system/input"
	"gome/system/menu"
	"gome/system/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	menu     *menu.Menu
	audioSys *audio.AudioSystem
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
	// Create audio system
	audioSys := audio.NewAudioSystem()
	game.InitializeSounds(audioSys)

	menuTree := game.CreateGameMenu(audioSys)
	menu.LinkParents(menuTree)

	// Play background music
	_ = audioSys.Play(game.BgmSound)

	g := &Game{
		menu:     menu.New(menuTree),
		audioSys: audioSys,
	}

	g.menu.Current.OnChange.Do(func(_ menu.MenuItem) {
		fmt.Println("Menu changed to:")
		_ = audioSys.Play(game.ClickSound)
	})

	
	g.menu.Current.OnChange.Do(func(_ menu.MenuItem) {
		fmt.Println("Menu changed to:")
		_ = audioSys.Play(game.ClickSound)
	})

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gome Menu System")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
