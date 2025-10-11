package events

type Event[T any] struct {
	listeners []func(T)
}

// Register adds a listener callback
func (e *Event[T]) Do(listener func(T)) {
	e.listeners = append(e.listeners, listener)
}

// Emit triggers all registered callbacks
func (e *Event[T]) Emit(value T) {
	for _, listener := range e.listeners {
		listener(value)
	}
}
