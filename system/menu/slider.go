package menu

type Slider struct {
	BaseItem
	label    string
	value    int
	min, max int
	OnChange func(newValue int)
}

func NewSlider(label string, min, max, value int, onChange func(int)) *Slider {
	return &Slider{label: label, min: min, max: max, value: value, OnChange: onChange}
}

func (s *Slider) Label() string { return s.label }

func (s *Slider) Update(cmd Command) {
	old := s.value
	switch cmd {
	case CmdLeft:
		if s.value > s.min {
			s.value--
		}
	case CmdRight:
		if s.value < s.max {
			s.value++
		}
	}
	if old != s.value && s.OnChange != nil {
		s.OnChange(s.value)
	}
}

// Value returns the current value of the slider
func (s *Slider) Value() int {
	return s.value
}
