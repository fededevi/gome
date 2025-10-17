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
	gameMap  *game.Map
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

		// Apply brightness shader

	}

	if g.fsm.FSM.Current() == g.fsm.GameStart {
		if g.gameMap != nil {
			g.gameMap.Draw(screen)
		}

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {

	// 1️⃣ Create audio system first
	audioSys := audio.NewAudioSystem()
	game.InitializeSounds(audioSys)

	// 3️⃣ Create FSM
	fsm := game.NewGomeFSM()
	fsm.FSM.OnStateChange.Do(func(s *sm.State) {
		fmt.Println("Entered Options Menu")
		_ = audioSys.Play(game.ClickSound)
	})

	g := &Game{
		menu:     menu.New(menu.LinkParents(game.CreateGameMenu(audioSys))),
		audioSys: audioSys,
		fsm:      fsm,
		gameMap:  game.NewMap("game/assets/audio/maps/lush.png"),
	}

	game.OnStartGame.Do(func(_ any) {
		//Log
		fmt.Println("OnStartGame event received in main.go")
		fsm.FSM.ActivateTransition(fsm.OptionsToGame)
	})

	mermaidDiagram := sm.ToMermaidLiveURL(sm.GenerateMermaidDiagram(fsm.FSM))
	fmt.Println("Game State Machine:", mermaidDiagram)

	// Start background music
	_ = audioSys.Play(game.BgmSound)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gome FSM Menu System")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
