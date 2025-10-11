package menu

import (
	"gome/system/events"
)

// Menu holds the current menu and selection
type Menu struct {
	Current  *events.Property[MenuItem]
	Selected int
}

// New creates a new Menu
func New(root MenuItem) *Menu {
	return &Menu{Current: events.NewProperty(root)}
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

	selectedItem := children[m.Selected]

	switch cmd {
	case CmdUp:
		if m.Selected > 0 {
			m.Selected--
		}
	case CmdDown:
		if m.Selected < len(children)-1 {
			m.Selected++
		}
	case CmdLeft, CmdRight, CmdSelect:
		selectedItem.Update(cmd)
		if len(selectedItem.Children()) > 0 && cmd == CmdSelect {
			m.Current.Set(selectedItem)
			m.Selected = 0
		}
	case CmdBack:
		if m.Current.Get().Parent() != nil {
			m.Current.Set(m.Current.Get().Parent())
			m.Selected = 0
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
