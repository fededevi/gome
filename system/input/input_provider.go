package input

import "gome/system/events"

type Command int

const (
	CmdUp Command = iota
	CmdDown
	CmdLeft
	CmdRight
	CmdSelect
	CmdBack
	CmdMenu
)

type InputProvider interface {
	Update()
	IsDown(cmd Command) bool
	OnKeyDown() *events.Event[Command]
	OnKeyUp() *events.Event[Command]
}
