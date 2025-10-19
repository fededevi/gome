package game

import (
	"fmt"
	"gome/system/audio"
	"gome/system/events"
	"gome/system/input"
	"gome/system/menu"
	"gome/system/render"

	"github.com/hajimehoshi/ebiten/v2"
)

type StartMenu struct {
	AudioSys *audio.AudioSystem
	RootMenu *menu.Submenu
	Menu     *menu.Menu
	Active   bool
	OnBegin  events.Event[string] // emits selected difficulty
}

func NewStartMenu(audioSys *audio.AudioSystem, hub *input.InputHub) *StartMenu {
	sm := &StartMenu{
		AudioSys: audioSys,
		Active:   false,
		OnBegin:  events.Event[string]{},
	}

	// Build menu structure
	sm.RootMenu = sm.createMenuStructure()
	sm.Menu = menu.New(menu.LinkParents(sm.RootMenu))

	// Subscribe to input events
	hub.OnKeyDown.Do(func(cmd input.Command) {
		if !sm.Active {
			return
		}
		sm.Menu.HandleCommand(cmd)
	})

	return sm
}

func (sm *StartMenu) Activate()   { sm.Active = true }
func (sm *StartMenu) Deactivate() { sm.Active = false }

func (sm *StartMenu) Draw(screen *ebiten.Image) {
	if !sm.Active {
		return
	}
	render.DrawMenu(screen, sm.Menu)
}

func (sm *StartMenu) SelectedItem() menu.MenuItem {
	return sm.Menu.Current.Get()
}

// --- Build menu ---
func (sm *StartMenu) createMenuStructure() *menu.Submenu {
	audioSys := sm.AudioSys

	// Default difficulty
	difficulty := "Easy"

	return menu.NewSubmenu("Start Menu", []menu.MenuItem{
		// Difficulty selection
		menu.NewEnumItem("Difficulty", []string{"Easy", "Medium", "Hard", "Impossible"}, 0, func(val string) {
			difficulty = val
			fmt.Println("Difficulty set to", val)
			audioSys.Play(ClickSound)
		}),
		// Start Game button
		menu.NewButton("Start Game", func() {
			fmt.Println("Starting game with difficulty:", difficulty)
			sm.OnBegin.Emit(difficulty)
			audioSys.Play(ClickSound)
		}),
	})
}
