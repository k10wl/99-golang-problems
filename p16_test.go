package main

import (
	"reflect"
	"testing"
)

func TestP16(t *testing.T) {
	type Args struct {
		n    int
		list []string
	}
	table := []Case{
		{
			name: "should drop every Nth element from a list",
			arg: Args{
				3,
				[]string{
					"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
				},
			},
			expected: []string{
				"a", "b", "d", "e", "g", "h", "j", "k",
			},
		},
	}
	for _, tc := range table {
		out := p16(tc.arg.(Args).n, tc.arg.(Args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
