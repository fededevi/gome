package main

import (
	"errors"
	"fmt"
	"log"

	"gome/game"
	"gome/system/audio"
	"gome/system/input"
	sm "gome/system/statemachine"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	audioSys  *audio.AudioSystem
	menu      *game.GomeMenu  // original menu
	startMenu *game.StartMenu // new start menu
	fsm       *game.GomeFSM
	gameMap   *game.Map
	hub       *input.InputHub
}

func (g *Game) Update() error {
	current := g.fsm.FSM.Current()
	g.hub.Update()

	// Show StartMenu only in GameStart
	if current == g.fsm.GameStart {
		g.startMenu.Activate()
	} else {
		g.startMenu.Deactivate()
	}

	// Show original menu in OptionsMenu or Paused
	switch current {
	case g.fsm.OptionsMenu, g.fsm.Paused:
		g.menu.Activate()
	default:
		g.menu.Deactivate()
	}

	// Automatically move from Initial -> OptionsMenu (example)
	if current == g.fsm.Initial {
		g.fsm.FSM.ActivateTransition(g.fsm.InitialToOptions)
	}

	// Check if quit requested from the original menu
	if g.menu.QuitRequested {
		return errors.New("quit requested")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	current := g.fsm.FSM.Current()

	// Draw StartMenu only in GameStart
	g.startMenu.Draw(screen)

	// Draw original menu in OptionsMenu or Paused
	g.menu.Draw(screen)

	// Draw gameplay map only in Gameplay state
	if current == g.fsm.Gameplay {
		if g.gameMap != nil {
			g.gameMap.Draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	// Initialize audio system
	audioSys := audio.NewAudioSystem()
	game.InitializeSounds(audioSys)

	// Create FSM
	fsm := game.NewGomeFSM()
	fsm.FSM.OnStateChange.Do(func(s *sm.State) {
		fmt.Println("Entered state:", s.Name)
		_ = audioSys.Play(game.ClickSound)
	})

	// Create InputHub
	hub := input.NewInputHub(input.NewControllerInput(0), input.NewKeyboardInput())

	// Create original menu
	menu := game.NewGomeMenu(audioSys, hub)

	// Keep original OnStartGame event for GomeMenu
	menu.OnStartGame.Do(func(_ any) {
		fmt.Println("GomeMenu OnStartGame event received")
		fsm.FSM.ActivateTransition(fsm.OptionsToGame)
	})

	// Create StartMenu
	startMenu := game.NewStartMenu(audioSys, hub)

	// Listen to OnBegin event from StartMenu to transition to Gameplay
	startMenu.OnBegin.Do(func(difficulty string) {
		fmt.Println("StartMenu OnBegin event received with difficulty:", difficulty)
		fsm.FSM.ActivateTransition(fsm.GameStartToPlay)
	})

	// Print FSM Mermaid diagram
	fmt.Println("FSM Mermaid diagram:", sm.ToMermaidLiveURL(sm.GenerateMermaidDiagram(fsm.FSM)))

	// Start background music
	_ = audioSys.Play(game.BgmSound)

	// Initialize Game struct
	g := &Game{
		audioSys:  audioSys,
		menu:      menu,
		startMenu: startMenu,
		fsm:       fsm,
		gameMap:   game.NewMap("game/assets/audio/maps/lush.png"),
		hub:       hub,
	}

	// Set window properties
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gome FSM Menu System")

	// Run the game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
