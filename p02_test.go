package main

import (
	"testing"
)

func TestP02(t *testing.T) {
	table := []Case{
		{
			name:     "should find the last but one element of a list",
			arg:      []int{1, 1, 2, 3, 5, 8},
			expected: 5,
		},
		{
			name:     "should return -1 if list length is less than 2",
			arg:      []int{1},
			expected: -1,
		},
		{
			name:     "should return -1 if list length is less than 2",
			arg:      []int{},
			expected: -1,
		},
	}

	for _, tc := range table {
		out := p02(tc.arg.([]int))
		if out != tc.expected {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
