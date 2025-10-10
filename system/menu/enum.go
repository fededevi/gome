package menu

// EnumItem is a menu item that cycles through a list of values
type EnumItem struct {
	BaseItem
	label    string
	values   []string     // list of possible options
	index    int          // current selected index
	OnChange func(string) // callback when selection changes
}

// NewEnumItem creates a new EnumItem
func NewEnumItem(label string, values []string, initial int, onChange func(string)) *EnumItem {
	if initial < 0 || initial >= len(values) {
		initial = 0
	}
	return &EnumItem{
		label:    label,
		values:   values,
		index:    initial,
		OnChange: onChange,
	}
}

// Label returns the label
func (e *EnumItem) Label() string { return e.label }

// Update handles left/right commands to cycle the enum
func (e *EnumItem) Update(cmd Command) {
	oldIndex := e.index
	switch cmd {
	case CmdLeft:
		if e.index > 0 {
			e.index--
		}
	case CmdRight:
		if e.index < len(e.values)-1 {
			e.index++
		}
	}
	if e.index != oldIndex && e.OnChange != nil {
		e.OnChange(e.values[e.index])
	}
}

// Value returns the currently selected value
func (e *EnumItem) Value() string {
	return e.values[e.index]
}
