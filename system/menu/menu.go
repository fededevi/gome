package menu

import (
	"gome/system/events"
	"gome/system/input"
)

// Menu holds the current menu and selected index as reactive properties
type Menu struct {
	Current  *events.Property[MenuItem]
	Selected *events.Property[int]
}

// New creates a new Menu and subscribes to input events
func New(root MenuItem) *Menu {
	m := &Menu{
		Current:  events.NewProperty(root),
		Selected: events.NewProperty(0),
	}

	return m
}

// handleCommand routes input commands to the correct menu item logic
func (m *Menu) HandleCommand(cmd input.Command) {
	if m.Current == nil {
		return
	}

	children := m.Current.Get().Children()
	if len(children) == 0 {
		return
	}

	selectedIndex := m.Selected.Get()
	selectedItem := children[selectedIndex]

	switch cmd {
	case input.CmdUp:
		if selectedIndex > 0 {
			m.Selected.Set(selectedIndex - 1)
		}
	case input.CmdDown:
		if selectedIndex < len(children)-1 {
			m.Selected.Set(selectedIndex + 1)
		}
	case input.CmdLeft:
		m.applyLeftRight(selectedItem, -1)
	case input.CmdRight:
		m.applyLeftRight(selectedItem, 1)
	case input.CmdSelect:
		m.applySelect(selectedItem)
	case input.CmdBack:
		if parent := m.Current.Get().Parent(); parent != nil {
			m.Current.Set(parent)
			m.Selected.Set(0)
		}
	case input.CmdMenu:
		// Optional: handle main menu toggle
	}
}

// applyLeftRight handles left/right commands for sliders and enums
func (m *Menu) applyLeftRight(item MenuItem, delta int) {
	switch it := item.(type) {
	case *Slider:
		newValue := it.Value() + delta
		if newValue < it.min {
			newValue = it.min
		} else if newValue > it.max {
			newValue = it.max
		}
		it.value = newValue
		if it.OnChange != nil {
			it.OnChange(newValue)
		}
	case *EnumItem:
		newIndex := it.index + delta
		if newIndex < 0 {
			newIndex = 0
		} else if newIndex >= len(it.values) {
			newIndex = len(it.values) - 1
		}
		it.index = newIndex
		if it.OnChange != nil {
			it.OnChange(it.values[it.index])
		}
	}
}

// applySelect executes the action for buttons or opens submenus
func (m *Menu) applySelect(item MenuItem) {
	switch it := item.(type) {
	case *Button:
		if it.Action != nil {
			it.Action()
		}
	case *Submenu:
		if len(it.Children()) > 0 {
			m.Current.Set(it)
			m.Selected.Set(0)
		}
	}
}

// LinkParents recursively sets parent pointers
func LinkParents(root MenuItem) MenuItem {
	for _, child := range root.Children() {
		child.SetParent(root)
		LinkParents(child)
	}
	return root
}
