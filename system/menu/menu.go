package menu

import (
	"gome/system/events"
)

// Menu holds the current menu and selected index as reactive properties
type Menu struct {
	Current  *events.Property[MenuItem]
	Selected *events.Property[int]
}

// New creates a new Menu
func New(root MenuItem) *Menu {
	return &Menu{
		Current:  events.NewProperty(root),
		Selected: events.NewProperty(0),
	}
}

// HandleCommand updates the menu according to input
func (m *Menu) HandleCommand(cmd Command) {
	if m.Current == nil {
		return
	}

	children := m.Current.Get().Children()
	if len(children) == 0 {
		return
	}

	selected := m.Selected.Get()

	switch cmd {
	case CmdUp:
		if selected > 0 {
			m.Selected.Set(selected - 1)
		}
	case CmdDown:
		if selected < len(children)-1 {
			m.Selected.Set(selected + 1)
		}
	case CmdLeft, CmdRight, CmdSelect:
		selectedItem := children[selected]
		selectedItem.Update(cmd)
		if len(selectedItem.Children()) > 0 && cmd == CmdSelect {
			m.Current.Set(selectedItem)
			m.Selected.Set(0)
		}
	case CmdBack:
		if parent := m.Current.Get().Parent(); parent != nil {
			m.Current.Set(parent)
			m.Selected.Set(0)
		}
	}
}

// LinkParents recursively sets parent pointers
func LinkParents(root MenuItem) {
	for _, child := range root.Children() {
		child.SetParent(root)
		LinkParents(child)
	}
}
