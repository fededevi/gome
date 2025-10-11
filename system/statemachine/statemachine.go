package sm

import (
	"fmt"
	"strings"

	"gome/system/events"
)

// ---------- State ----------

// State represents a node in the FSM
type State struct {
	Name    string // optional name for logging/documentation
	OnEnter events.Event[*State]
	OnExit  events.Event[*State]
}

// NewState creates a new state with the given name
func NewState(name string) *State {
	return &State{
		Name:    name,
		OnEnter: events.Event[*State]{},
		OnExit:  events.Event[*State]{},
	}
}

// ---------- Transition ----------

// Transition connects a Source state to a Target state
type Transition struct {
	Name       string // optional name for logging/documentation
	Source     *State
	Target     *State
	OnActivate events.Event[*Transition]
	Guard      func() bool // optional guard function
}

// NewTransition creates a new transition with optional guard and name
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
		t.Guard = func() bool { return true } // default always allowed
	}
	return t
}

// Activate triggers the transition
func (t *Transition) Activate() *State {
	t.OnActivate.Emit(t)
	return t.Target
}

// ---------- State Machine ----------

// StateMachine manages the current state and transitions
type StateMachine struct {
	initial       *State
	current       *events.Property[*State]
	transitions   []*Transition
	OnTransition  events.Event[*Transition]
	OnStateChange events.Event[*State]
}

// NewStateMachine creates a new FSM with a default initial state
func NewStateMachine() *StateMachine {
	initial := &State{Name: "InitialState"} // basic initial state
	return &StateMachine{
		initial:     initial,
		current:     events.NewProperty[*State](initial),
		transitions: []*Transition{},
	}
}

// AddTransition registers a transition
func (sm *StateMachine) AddTransition(t *Transition) {
	sm.transitions = append(sm.transitions, t)
}

// Current returns the active state
func (sm *StateMachine) Current() *State {
	return sm.current.Get()
}

// Reset sets the FSM back to the initial state
func (sm *StateMachine) Reset() {
	sm.SetState(sm.initial)
}

// SetState forces a state change
func (sm *StateMachine) SetState(s *State) {
	if sm.current.Get() != nil {
		sm.current.Get().OnExit.Emit(sm.current.Get())
	}
	sm.current.Set(s)
	s.OnEnter.Emit(s)
	sm.OnStateChange.Emit(s)
}

// TryTransition attempts a valid transition from the current state
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

// ToMermaid generates a Mermaid flowchart string of the FSM
func (sm *StateMachine) ToMermaid() string {
	var sb strings.Builder
	sb.WriteString("flowchart TD\n")

	// Gather all states from transitions
	states := make(map[*State]bool)
	for _, t := range sm.transitions {
		states[t.Source] = true
		states[t.Target] = true
	}

	// Print each transition
	for _, t := range sm.transitions {
		srcName := t.Source.Name
		if srcName == "" {
			srcName = fmt.Sprintf("%p", t.Source)
		}
		tgtName := t.Target.Name
		if tgtName == "" {
			tgtName = fmt.Sprintf("%p", t.Target)
		}
		label := t.Name
		sb.WriteString(fmt.Sprintf("    %s -->|%s| %s\n", srcName, label, tgtName))
	}

	// Mark initial state
	if sm.initial != nil {
		initName := sm.initial.Name
		if initName == "" {
			initName = fmt.Sprintf("%p", sm.initial)
		}
		sb.WriteString(fmt.Sprintf("    %% Initial state: %s\n", initName))
	}

	return sb.String()
}
