package main

import (
	"reflect"
	"testing"
)

func TestP24(t *testing.T) {
	type args struct {
		n, m int
	}
	table := []Case{
		{
			name:     "should draw N different random numbers from the set 1..M1..M",
			arg:      args{6, 49},
			expected: []int{7, 27, 11, 45, 33, 11},
		},
	}
	for _, tc := range table {
		out := p24(tc.arg.(args).n, tc.arg.(args).m)
		if len(out) != tc.arg.(args).n || !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
