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
	sm "gome/system/statemachine"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	audioSys *audio.AudioSystem
	menu     *menu.Menu
	fsm      *game.GomeFSM
}

func (g *Game) Update() error {
	// Update menu input
	cmd := input.GetMenuCommand()
	g.menu.HandleCommand(cmd)

	// Automatically transition from Initial -> OptionsMenu
	if g.fsm.FSM.Current() == g.fsm.Initial {
		g.fsm.FSM.TryTransition()
	}

	if game.QuitRequested {
		return errors.New("quit requested")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw menu only if OptionsMenu state is active
	if g.fsm.FSM.Current() == g.fsm.OptionsMenu {
		render.DrawMenu(screen, g.menu)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	audioSys := audio.NewAudioSystem()
	game.InitializeSounds(audioSys)

	// Create menu
	menuTree := game.CreateGameMenu(audioSys)
	menu.LinkParents(menuTree)
	gameMenu := menu.New(menuTree)

	// Create FSM
	fsm := game.NewGomeFSM()
	fsm.FSM.OnStateChange.Do(func(s *sm.State) {
		fmt.Println("Entered Options Menu")
		_ = audioSys.Play(game.ClickSound)
	})

	fsm.OptionsToGame.OnActivate.Do(func(t *sm.Transition) {
		fmt.Println("Starting game!")
	})

	g := &Game{
		menu:     gameMenu,
		audioSys: audioSys,
		fsm:      fsm,
	}

	// Start background music
	_ = audioSys.Play(game.BgmSound)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gome FSM Menu System")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
