package menu

type Submenu struct {
	BaseItem
	label string
}

func NewSubmenu(label string, children []MenuItem) *Submenu {
	s := &Submenu{label: label}
	s.SetChildren(children)
	return s
}

func (s *Submenu) Label() string      { return s.label }
