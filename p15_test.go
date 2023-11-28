package main

import (
	"reflect"
	"testing"
)

func TestP15(t *testing.T) {
	type Args struct {
		times int
		list  []string
	}
	table := []Case{
		{
			name: "should duplicate the elements of a list a given number of times",
			arg: Args{
				3,
				[]string{"a", "b", "c", "d"},
			},
			expected: []string{
				"a",
				"a",
				"a",
				"b",
				"b",
				"b",
				"c",
				"c",
				"c",
				"d",
				"d",
				"d",
			},
		},
	}
	for _, tc := range table {
		out := p15(tc.arg.(Args).times, tc.arg.(Args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
