package main

import (
	"reflect"
	"testing"
)

func TestP17(t *testing.T) {
	type args struct {
		n    int
		list []string
	}
	table := []Case{
		{
			name: "should split a list into two parts",
			arg: args{
				3,
				[]string{
					"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
				},
			},
			expected: [][]string{
				{
					"a", "b", "c",
				},
				{
					"d", "e", "f", "g", "h", "i", "j", "k",
				},
			},
		},
	}
	for _, tc := range table {
		out := p17(tc.arg.(args).n, tc.arg.(args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
