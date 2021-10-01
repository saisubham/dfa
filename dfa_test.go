package main

import (
	"testing"
)

func TestMultipleOf3(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "should be true for 0",
			input:    "0",
			expected: true,
		},
		{
			desc:     "should be false for 1",
			input:    "1",
			expected: false,
		},
		{
			desc:     "should be true for 9",
			input:    "1001",
			expected: true,
		},
		{
			desc:     "should be false for 19",
			input:    "10011",
			expected: false,
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
			res, err := dfa.Run(tC.input)
			if err != nil {
				t.Fail()
			}
			if res != tC.expected {
				t.Fail()
			}
		})
	}
}
