package events

type Property[T comparable] struct {
	value    T
	OnChange Event[T]
}

// NewProperty creates a property with an initial value
func NewProperty[T comparable](initial T) *Property[T] {
	return &Property[T]{value: initial}
}

// Set assigns a new value and triggers OnChange if changed
func (p *Property[T]) Set(newValue T) {
	if p.value != newValue {
		p.value = newValue
		p.OnChange.Emit(newValue)
	}
}

// Get returns the current value
func (p *Property[T]) Get() T {
	return p.value
}
