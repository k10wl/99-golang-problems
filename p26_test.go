package main

import (
	"reflect"
	"testing"
)

func TestP26(t *testing.T) {
	type args struct {
		n    int
		list []string
	}
	type edgeCase struct {
		i    int
		list []string
	}
	type expected struct {
		possibilities int
		edgeCases     []edgeCase
	}
	table := []Case{
		{
			name: "should generate the combinations of K distinct objects chosen from the N elements of a list",
			arg:  args{3, []string{"a", "b", "c", "d", "e", "f"}},
			expected: expected{possibilities: 20, edgeCases: []edgeCase{
				{0, []string{"a", "b", "c"}},
				{1, []string{"a", "b", "d"}},
				{2, []string{"a", "b", "e"}},
				{19, []string{"d", "e", "f"}},
			}},
		},
	}
	for _, tc := range table {
		out := p26(tc.arg.(args).n, tc.arg.(args).list)
		e := tc.expected.(expected)
		if len(out) != e.possibilities {
			t.Errorf("Failed to create expected combinations.\nExpected: %d\nActual:   %v\n", e.possibilities, len(out))
		}
		for _, el := range e.edgeCases {
			if !reflect.DeepEqual(out[el.i], el.list) {
				t.Errorf("Failed combination creation.\nExpected: %v\nActual:   %v\n", el.list, out[el.i])
			}
		}
	}
}
