package main

import (
	"reflect"
	"testing"
)

func TestP05(t *testing.T) {
	table := []Case{
		{
			name:     "should reverse list",
			arg:      []int{1, 1, 2, 3, 5, 8},
			expected: []int{8, 5, 3, 2, 1, 1},
		},
	}

	for _, tc := range table {
		out := p05(tc.arg.([]int))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
