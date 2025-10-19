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


// Value returns the current value of the slider
func (s *Slider) Value() int {
	return s.value
}
