package main

import (
	"errors"
	"fmt"
)

type TransitionInput struct {
	src   int
	input rune
}

type Transitions struct {
	src   int
	input rune
	dst   int
}

type DFA struct {
	allStates    map[int]struct{}
	finalStates  map[int]struct{}
	initState    int
	inputSymbols map[rune]struct{}
	transitions  map[TransitionInput]int
}

func MakeDFA(allStates []int, inputSymbols []rune, initState int, finalStates []int) (*DFA, error) {
	dfa := &DFA{
		allStates:    make(map[int]struct{}),
		finalStates:  make(map[int]struct{}),
		initState:    initState,
		inputSymbols: make(map[rune]struct{}),
		transitions:  make(map[TransitionInput]int),
	}
	for _, i := range allStates {
		dfa.allStates[i] = struct{}{}
	}
	for _, i := range finalStates {
		dfa.finalStates[i] = struct{}{}
	}
	for _, i := range inputSymbols {
		dfa.inputSymbols[i] = struct{}{}
	}
	return dfa, nil
}

func (dfa *DFA) AddTransition(src int, input rune, dst int) error {
	if _, ok := dfa.allStates[src]; !ok {
		return errors.New("bad src state")
	}
	if _, ok := dfa.allStates[dst]; !ok {
		return errors.New("bad dst state")
	}
	if _, ok := dfa.inputSymbols[input]; !ok {
		return &BadInputError{input}
	}
	dfa.transitions[TransitionInput{src, input}] = dst
	return nil
}

func (dfa *DFA) AddTransitions(transitions []*Transitions) error {
	for _, t := range transitions {
		if err := dfa.AddTransition(t.src, t.input, t.dst); err != nil {
			return errors.New("bad transition")
		}
	}
	return nil
}

func (dfa *DFA) PrintTransitionTable() {
	fmt.Printf("%5c", ' ')
	for c := range dfa.inputSymbols {
		fmt.Printf("%3c", c)
	}
	fmt.Println()
	for i := 0; i < len(dfa.inputSymbols)+3; i++ {
		fmt.Printf("---")
	}
	fmt.Println()

	for i := range dfa.allStates {
		fmt.Printf("%3d |", i)
		for j := range dfa.inputSymbols {
			fmt.Printf("%3d", dfa.transitions[TransitionInput{i, j}])
		}
		fmt.Println()
	}
}

func (dfa *DFA) Run(s string) (bool, error) {
	curState := dfa.initState
	for _, c := range s {
		tIn := TransitionInput{curState, c}
		nextState, ok := dfa.transitions[tIn]
		if !ok {
			return false, &BadTransitionInputError{tIn}
		}
		curState = nextState
	}

	// is current state a goal state?
	if _, ok := dfa.finalStates[curState]; ok {
		return true, nil
	}
	return false, nil
}
