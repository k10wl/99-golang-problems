package main

import (
	"reflect"
	"testing"
)

func TestP10(t *testing.T) {
	table := []Case{
		{
			name: "should run-length encoding of a list",
			arg:  []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"},
			expected: PackedDuplicatesP10{
				{4, "a"},
				{1, "b"},
				{2, "c"},
				{2, "a"},
				{1, "d"},
				{4, "e"},
			},
		},
	}
	for _, tc := range table {
		out := p10(tc.arg.([]string))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}

