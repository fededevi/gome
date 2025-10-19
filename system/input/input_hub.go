package input

import "gome/system/events"

type InputHub struct {
	providers []InputProvider

	OnKeyDown events.Event[Command]
	OnKeyUp   events.Event[Command]
}

func NewInputHub(providers ...InputProvider) *InputHub {
	hub := &InputHub{
		providers: providers,
		OnKeyDown: events.Event[Command]{},
		OnKeyUp:   events.Event[Command]{},
	}

	for _, p := range providers {
		p.OnKeyDown().Do(func(cmd Command) { hub.OnKeyDown.Emit(cmd) })
		p.OnKeyUp().Do(func(cmd Command) { hub.OnKeyUp.Emit(cmd) })
	}

	return hub
}

func (h *InputHub) Update() {
	for _, p := range h.providers {
		p.Update()
	}
}

func (h *InputHub) IsDown(cmd Command) bool {
	for _, p := range h.providers {
		if p.IsDown(cmd) {
			return true
		}
	}
	return false
}
