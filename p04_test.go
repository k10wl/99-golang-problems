package main

import "testing"

func TestP04(t *testing.T) {
	table := []Case{
		{
			name:     "should find the number of elements of a list.",
			arg:      []int{1, 1, 2, 3, 5, 8},
			expected: 6,
		},
	}

	for _, tc := range table {
		out := p04(tc.arg.([]int))
		if out != tc.expected {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
