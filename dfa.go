package main

import (
	"errors"
	"fmt"
	"strings"
)

const deadState = -1

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
	numStates    uint
	inputSymbols []rune
	transitions  map[TransitionInput]int
	initState    uint
	finalStates  []int
}

func MakeDFA(numStates uint, inputSymbols []rune, finalStates []int) (*DFA, error) {
	if numStates <= 0 {
		return nil, errors.New("bad number of states")
	}

	dfa := &DFA{
		numStates:    numStates,
		inputSymbols: inputSymbols,
		transitions:  make(map[TransitionInput]int),
		initState:    uint(0),
		finalStates:  finalStates,
	}

	// Transitions for dead state
	// for _, c := range inputSymbols {
	// 	dfa.AddTransitions([]*Transitions{
	// 		{deadState, c, deadState},
	// 	})
	// }
	return dfa, nil
}

func (dfa *DFA) AddTransition(src int, input rune, dst int) error {
	// check if valid src & dst state
	if src < deadState || src >= int(dfa.numStates) || dst < 0 || dst >= int(dfa.numStates) {
		return errors.New("bad state")
	}

	// check if valid input symbol
	if pos := strings.IndexRune(string(dfa.inputSymbols), input); pos == -1 {
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
	for _, c := range dfa.inputSymbols {
		fmt.Printf("%3c", c)
	}
	fmt.Println()
	for i := 0; i < len(dfa.inputSymbols)+3; i++ {
		fmt.Printf("---")
	}
	fmt.Println()

	for i := 0; i < int(dfa.numStates); i++ {
		fmt.Printf("%3d |", i)
		for _, j := range dfa.inputSymbols {
			in := TransitionInput{int(i), j}
			fmt.Printf("%3d", dfa.transitions[in])
		}
		fmt.Println()
	}
}

func (dfa *DFA) Run(s string) (bool, error) {
	curState := int(dfa.initState)
	for _, c := range s {
		tIn := TransitionInput{curState, c}
		nextState, prs := dfa.transitions[tIn]
		if !prs {
			return false, &BadTransitionInputError{tIn}
		}
		curState = nextState
	}

	// is current state a goal state?
	for f := range dfa.finalStates {
		if curState == f {
			return true, nil
		}
	}
	return false, nil
}
