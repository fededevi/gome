package menu

type Item struct {
	Label    string
	Action   func()
	Children []*Item
	Parent   *Item
}

type Menu struct {
	Current  *Item
	Selected int
}

// New creates a new menu from a root node.
func New(root *Item) *Menu {
	return &Menu{Current: root}
}

// HandleCommand updates the menu state based on an abstract command.
func (m *Menu) HandleCommand(cmd Command) {
	switch cmd {
	case CmdUp:
		if m.Selected > 0 {
			m.Selected--
		}
	case CmdDown:
		if m.Selected < len(m.Current.Children)-1 {
			m.Selected++
		}
	case CmdSelect:
		if len(m.Current.Children) == 0 {
			return
		}
		selected := m.Current.Children[m.Selected]
		if len(selected.Children) > 0 {
			m.Current = selected
			m.Selected = 0
		} else if selected.Action != nil {
			selected.Action()
		}
	case CmdBack:
		if m.Current.Parent != nil {
			m.Current = m.Current.Parent
			m.Selected = 0
		}
	}
}

// Utility to link parent pointers recursively.
func LinkParents(root *Item) {
	for _, child := range root.Children {
		child.Parent = root
		LinkParents(child)
	}
}
