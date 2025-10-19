package input

import (
	"gome/system/events"

	"github.com/hajimehoshi/ebiten/v2"
)

var keyBindings = map[Command]ebiten.Key{
	CmdUp:     ebiten.KeyArrowUp,
	CmdDown:   ebiten.KeyArrowDown,
	CmdLeft:   ebiten.KeyArrowLeft,
	CmdRight:  ebiten.KeyArrowRight,
	CmdSelect: ebiten.KeyEnter,
	CmdBack:   ebiten.KeyEscape,
	CmdMenu:   ebiten.KeyM,
}

type KeyboardInput struct {
	OnDown    events.Event[Command]
	OnUp      events.Event[Command]
	keyStates map[Command]bool
}

func NewKeyboardInput() *KeyboardInput {
	return &KeyboardInput{
		OnDown:    events.Event[Command]{},
		OnUp:      events.Event[Command]{},
		keyStates: make(map[Command]bool),
	}
}

func (gi *KeyboardInput) Update() {
	for cmd, key := range keyBindings {
		pressed := ebiten.IsKeyPressed(key)
		wasPressed := gi.keyStates[cmd]

		if pressed && !wasPressed {
			gi.OnDown.Emit(cmd)
		} else if !pressed && wasPressed {
			gi.OnUp.Emit(cmd)
		}

		gi.keyStates[cmd] = pressed
	}
}

func (gi *KeyboardInput) IsDown(cmd Command) bool {
	return gi.keyStates[cmd]
}

func (gi *KeyboardInput) OnKeyDown() *events.Event[Command] { return &gi.OnDown }
func (gi *KeyboardInput) OnKeyUp() *events.Event[Command]   { return &gi.OnUp }
