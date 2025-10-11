package game

import (
	sm "gome/system/statemachine"
)

type GomeFSM struct {
	FSM         *sm.StateMachine
	Initial     *sm.State
	OptionsMenu *sm.State
	GameStart   *sm.State

	// Exposed transitions
	InitialToOptions *sm.Transition
	OptionsToGame    *sm.Transition
}

func NewGomeFSM() *GomeFSM {
	fsm := sm.NewStateMachine()

	initial := sm.NewState("Initial")
	options := sm.NewState("OptionsMenu")
	gameStart := sm.NewState("GameStart")

	initialToOptions := sm.NewTransition(initial, options, nil, "InitToOptions")
	optionsToGame := sm.NewTransition(options, gameStart, nil, "StartGame")

	fsm.AddTransition(initialToOptions)
	fsm.AddTransition(optionsToGame)
	fsm.SetState(initial)

	return &GomeFSM{
		FSM:              fsm,
		Initial:          initial,
		OptionsMenu:      options,
		GameStart:        gameStart,
		InitialToOptions: initialToOptions,
		OptionsToGame:    optionsToGame,
	}
}
