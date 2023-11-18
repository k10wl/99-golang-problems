package main

import (
	"testing"
)

func TestP01(t *testing.T) {
	table := []Case{
		{
			name:     "should return last element",
			arg:      []int{1, 1, 2, 3, 5, 8},
			expected: 8,
		},
		{
			name:     "should return -1 if slice is empty",
			arg:      []int{},
			expected: -1,
		},
	}

	for _, tc := range table {
		out := p01(tc.arg.([]int))
		if out != tc.expected {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
