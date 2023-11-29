package main

import (
	"reflect"
	"testing"
)

func TestP23(t *testing.T) {
	type args struct {
		n    int
		list []string
	}
	table := []Case{
		{
			name: "should extract a given number of randomly selected elements from a list",
			arg: args{
				3,
				[]string{"a", "b", "c", "d", "e", "f", "g", "h"},
			},
			expected: []string{"g", "c", "b"},
		},
	}
	for _, tc := range table {
		out := p23(tc.arg.(args).n, tc.arg.(args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
