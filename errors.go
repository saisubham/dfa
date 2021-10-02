package main

import "fmt"

type BadInputError struct {
	input rune
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input %q", e.input)
}

type BadTransitionInputError struct {
	t TransitionInput
}

func (e *BadTransitionInputError) Error() string {
	return fmt.Sprintf("bad transition input %q", e.t)
}
