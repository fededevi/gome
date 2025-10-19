package input

import (
	"gome/system/events"

	"github.com/hajimehoshi/ebiten/v2"
)

type ControllerInput struct {
	OnDown    events.Event[Command]
	OnUp      events.Event[Command]
	keyStates map[Command]bool
	playerID  int // Gamepad ID
}

func NewControllerInput(playerID int) *ControllerInput {
	return &ControllerInput{
		OnDown:    events.Event[Command]{},
		OnUp:      events.Event[Command]{},
		keyStates: make(map[Command]bool),
		playerID:  playerID,
	}
}

func (ci *ControllerInput) Update() {
	// Map buttons to commands (example, adjust according to gamepad mapping)
	buttonMap := map[Command]ebiten.GamepadButton{
		CmdUp:     ebiten.GamepadButton10, // D-Pad Up
		CmdDown:   ebiten.GamepadButton12, // D-Pad Down
		CmdLeft:   ebiten.GamepadButton13, // D-Pad Left
		CmdRight:  ebiten.GamepadButton11, // D-Pad Right
		CmdSelect: ebiten.GamepadButton0,  // A
		CmdBack:   ebiten.GamepadButton1,  // B
		CmdMenu:   ebiten.GamepadButton7,  // Start
	}

	for cmd, button := range buttonMap {
		pressed := ebiten.IsGamepadButtonPressed(ebiten.GamepadID(ci.playerID), button)
		wasPressed := ci.keyStates[cmd]

		if pressed && !wasPressed {
			ci.OnDown.Emit(cmd)
		} else if !pressed && wasPressed {
			ci.OnUp.Emit(cmd)
		}

		ci.keyStates[cmd] = pressed
	}
}

func (ci *ControllerInput) IsDown(cmd Command) bool {
	return ci.keyStates[cmd]
}

func (ci *ControllerInput) OnKeyDown() *events.Event[Command] { return &ci.OnDown }
func (ci *ControllerInput) OnKeyUp() *events.Event[Command]   { return &ci.OnUp }
