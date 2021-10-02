package main

import (
	"testing"
)

func TestMultipleOf3(t *testing.T) {
	testCases := []struct {
		desc    string
		input   string
		want    bool
		wantErr error
	}{
		{
			desc:    "should be true for 0",
			input:   "0",
			want:    true,
			wantErr: nil,
		},
		{
			desc:    "should be false for 1",
			input:   "1",
			want:    false,
			wantErr: nil,
		},
		{
			desc:    "should be true for 9",
			input:   "1001",
			want:    true,
			wantErr: nil,
		},
		{
			desc:    "should be false for 19",
			input:   "10011",
			want:    false,
			wantErr: nil,
		},
		{
			desc:    "should throw error for invalid input",
			input:   "abc",
			want:    false,
			wantErr: &BadTransitionInputError{},
		},
	}

	dfa, err := MakeDFA(
		[]int{0, 1, 2},   // all states
		[]rune{'0', '1'}, // all symbols
		0,                // initial state
		[]int{0},         // final states
	)
	if err != nil {
		t.Fail()
	}
	err = dfa.AddTransitions([]*Transitions{
		{0, '0', 0},
		{0, '1', 1},
		{1, '0', 2},
		{1, '1', 0},
		{2, '0', 1},
		{2, '1', 2},
	})
	if err != nil {
		t.Fail()
	}
	dfa.PrintTransitionTable()

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := dfa.Run(tC.input)
			// TODO: here I want to compare specific error objects
			// perferably using errors.As(error,interface{}) but I don't know how
			if tC.wantErr != nil && err == nil || got != tC.want {
				t.Fail()
			}
		})
	}
}
