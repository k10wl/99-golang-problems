package main

import (
	"testing"
)

func TestP03(t *testing.T) {
	type half struct {
		arg1 int
		arg2 []int
	}
	table := []Case{
		{
			name: "should find the Kth element of a list",
			arg: half{
				2,
				[]int{1, 1, 2, 3, 5, 8},
			},
			expected: 2,
		},
		{
			name: "should return -1 if Kth element not found",
			arg: half{
				-1,
				[]int{1, 1, 2, 3, 5, 8},
			},
			expected: -1,
		},
		{
			name: "should return -1 if Kth element not found",
			arg: half{
				6,
				[]int{1, 1, 2, 3, 5, 8},
			},
			expected: -1,
		},
	}

	for _, tc := range table {
		out := p03(tc.arg.(half).arg1, tc.arg.(half).arg2)
		if out != tc.expected {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
