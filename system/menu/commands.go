package menu

// Command represents a menu navigation intent, independent of input device.
type Command int

const (
	CmdNone Command = iota
	CmdUp
	CmdDown
	CmdSelect
	CmdBack
)
