package main

import (
	"reflect"
	"testing"
)

func TestP21(t *testing.T) {
	type args struct {
		new  string
		at   int
		list []string
	}
	table := []Case{
		{
			name: "should insert an element at a given position into a list",
			arg: args{
				"new",
				1,
				[]string{"a", "b", "c", "d"},
			},
			expected: []string{"a", "new", "b", "c", "d"},
		},
	}
	for _, tc := range table {
		out := p21(tc.arg.(args).new, tc.arg.(args).at, tc.arg.(args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
