package menu

type Button struct {
	BaseItem // embeds parent/children logic
	label    string
	action   func()
}

func NewButton(label string, action func()) *Button {
	return &Button{label: label, action: action}
}

func (b *Button) Label() string { return b.label }
func (b *Button) Update(cmd Command) {
	if cmd == CmdSelect && b.action != nil {
		b.action()
	}
}
