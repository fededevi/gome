package game

import (
	sm "gome/system/statemachine"
)

type GomeFSM struct {
	FSM *sm.StateMachine

	// States
	Initial     *sm.State
	OptionsMenu *sm.State
	GameStart   *sm.State
	Gameplay    *sm.State
	Paused      *sm.State
	FinishGame  *sm.State

	// Transitions
	InitialToOptions *sm.Transition
	OptionsToGame    *sm.Transition
	GameStartToPlay  *sm.Transition
	PlayToPause      *sm.Transition
	PauseToPlay      *sm.Transition
	PlayToFinish     *sm.Transition
	FinishToInitial  *sm.Transition
}

func NewGomeFSM() *GomeFSM {
	fsm := sm.NewStateMachine()

	// Define states
	initial := sm.NewState("Initial")
	options := sm.NewState("OptionsMenu")
	gameStart := sm.NewState("GameStart")
	gameplay := sm.NewState("Gameplay")
	paused := sm.NewState("Paused")
	finish := sm.NewState("FinishGame")

	// Define transitions
	initialToOptions := sm.NewTransition(initial, options, nil, "InitToOptions")
	optionsToGame := sm.NewTransition(options, gameStart, nil, "StartGame")
	gameStartToPlay := sm.NewTransition(gameStart, gameplay, nil, "StartGameplay")
	playToPause := sm.NewTransition(gameplay, paused, nil, "PauseGame")
	pauseToPlay := sm.NewTransition(paused, gameplay, nil, "ResumeGame")
	playToFinish := sm.NewTransition(gameplay, finish, nil, "FinishGame")
	finishToInitial := sm.NewTransition(finish, initial, nil, "ReturnToMenu")

	// Register transitions
	fsm.AddTransition(initialToOptions)
	fsm.AddTransition(optionsToGame)
	fsm.AddTransition(gameStartToPlay)
	fsm.AddTransition(playToPause)
	fsm.AddTransition(pauseToPlay)
	fsm.AddTransition(playToFinish)
	fsm.AddTransition(finishToInitial)

	// Set initial state
	fsm.SetState(initial)

	return &GomeFSM{
		FSM:              fsm,
		Initial:          initial,
		OptionsMenu:      options,
		GameStart:        gameStart,
		Gameplay:         gameplay,
		Paused:           paused,
		FinishGame:       finish,
		InitialToOptions: initialToOptions,
		OptionsToGame:    optionsToGame,
		GameStartToPlay:  gameStartToPlay,
		PlayToPause:      playToPause,
		PauseToPlay:      pauseToPlay,
		PlayToFinish:     playToFinish,
		FinishToInitial:  finishToInitial,
	}
}
