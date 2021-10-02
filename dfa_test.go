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

	dfa, err := MakeDFA(3, []rune{'0', '1'}, []int{0})
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

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := dfa.Run(tC.input)
			if tC.wantErr != nil && err == nil || got != tC.want {
				t.Fail()
			}
		})
	}
}
