package main

import (
	"reflect"
	"testing"
)

func TestP22(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	table := []Case{
		{
			name:     "should create a list containing all integers within a given range",
			arg:      args{4, 9},
			expected: []int{4, 5, 6, 7, 8, 9},
		},
	}
	for _, tc := range table {
		out := p22(tc.arg.(args).start, tc.arg.(args).end)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
