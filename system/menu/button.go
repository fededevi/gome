package menu

type Button struct {
	BaseItem // embeds parent/children logic
	label    string
	Action   func()
}

// NewButton creates a new Button
func NewButton(label string, action func()) *Button {
	return &Button{
		label:  label,
		Action: action,
	}
}

func (b *Button) Label() string { return b.label }

// Update executes the action when selected
func (b *Button) Update(cmd Command) {
	if cmd == CmdSelect && b.Action != nil {
		b.Action()
	}
}
