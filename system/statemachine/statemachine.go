package sm

import (
	"fmt"

	"gome/system/events"
)

// ---------- State ----------
type State struct {
	Name    string
	OnEnter events.Event[*State]
	OnExit  events.Event[*State]
}

func NewState(name string) *State {
	return &State{
		Name:    name,
		OnEnter: events.Event[*State]{},
		OnExit:  events.Event[*State]{},
	}
}

// ---------- Transition ----------
type Transition struct {
	Name       string
	Source     *State
	Target     *State
	OnActivate events.Event[*Transition]
	Guard      func() bool
}

func NewTransition(source, target *State, guard func() bool, name string) *Transition {
	t := &Transition{
		Source:     source,
		Target:     target,
		OnActivate: events.Event[*Transition]{},
		Name:       name,
	}
	if guard != nil {
		t.Guard = guard
	} else {
		t.Guard = func() bool { return true }
	}
	return t
}

// ---------- State Machine ----------
type StateMachine struct {
	initial       *State
	current       *events.Property[*State]
	transitions   []*Transition
	OnTransition  events.Event[*Transition]
	OnStateChange events.Event[*State]
}

func NewStateMachine() *StateMachine {
	initial := &State{Name: "InitialState"}
	return &StateMachine{
		initial:     initial,
		current:     events.NewProperty[*State](initial),
		transitions: []*Transition{},
	}
}

func (sm *StateMachine) ActivateTransition(t *Transition) *State {
	if sm.Current() != t.Source {
		fmt.Printf("Cannot activate transition '%s': current state is '%s', expected '%s'\n",
			t.Name, sm.Current().Name, t.Source.Name)
		return sm.Current()
	}
	t.OnActivate.Emit(t)
	sm.OnTransition.Emit(t)
	sm.SetState(t.Target)
	return t.Target
}

func (sm *StateMachine) AddTransition(t *Transition) {
	sm.transitions = append(sm.transitions, t)
}

func (sm *StateMachine) Current() *State {
	return sm.current.Get()
}

func (sm *StateMachine) Reset() {
	sm.SetState(sm.initial)
}

func (sm *StateMachine) SetState(s *State) {
	if sm.current.Get() != nil {
		sm.current.Get().OnExit.Emit(sm.current.Get())
	}
	sm.current.Set(s)
	s.OnEnter.Emit(s)
	sm.OnStateChange.Emit(s)
}

func (sm *StateMachine) TryTransition() bool {
	current := sm.current.Get()
	for _, t := range sm.transitions {
		if t.Source == current && t.Guard() {
			if current != nil {
				current.OnExit.Emit(current)
			}
			t.OnActivate.Emit(t)
			sm.OnTransition.Emit(t)
			sm.current.Set(t.Target)
			t.Target.OnEnter.Emit(t.Target)
			sm.OnStateChange.Emit(t.Target)
			return true
		}
	}
	return false
}
