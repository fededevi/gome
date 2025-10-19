package menu

// MenuItem is the interface all menu items implement
type MenuItem interface {
	Label() string
	Parent() MenuItem
	SetParent(MenuItem)
	Children() []MenuItem
}

// BaseItem holds common parent/children data for reuse
type BaseItem struct {
	parent   MenuItem
	children []MenuItem
}

func (b *BaseItem) Parent() MenuItem         { return b.parent }
func (b *BaseItem) SetParent(p MenuItem)     { b.parent = p }
func (b *BaseItem) Children() []MenuItem     { return b.children }
func (b *BaseItem) SetChildren(c []MenuItem) { b.children = c }
