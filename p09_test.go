package main

import (
	"reflect"
	"testing"
)

func TestP09(t *testing.T) {
	table := []Case{
		{
			name: "should pack consecutive duplicates of list elements into sublists",
			arg:  []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"},
			expected: PackedDuplicatesP09{
				[]string{"a", "a", "a", "a"},
				[]string{"b"},
				[]string{"c", "c"},
				[]string{"a", "a"},
				[]string{"d"},
				[]string{"e", "e", "e", "e"},
			},
		},
	}
	for _, tc := range table {
		out := p09(tc.arg.([]string))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}

