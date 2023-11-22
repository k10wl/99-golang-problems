package main

import "testing"

func TestP06(t *testing.T) {
	table := []Case{
		{
			name:     "should find out whether a list is a palindrome",
			arg:      []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "should return false if list is not a palindrome",
			arg:      []int{1, 2, 3, 1},
			expected: false,
		},
	}

	for _, tc := range table {
		out := p06(tc.arg.([]int))
		if out != tc.expected {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
