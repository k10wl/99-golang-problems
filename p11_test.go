package main

import (
	"reflect"
	"testing"
)

func TestP11(t *testing.T) {
	table := []Case{
		{
			name: "should modified run-length encoding",
			arg: []string{
				"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e",
			},
			expected: []interface{}{
				PackedDuplicateP11{4, "a"},
				"b",
				PackedDuplicateP11{2, "c"},
				PackedDuplicateP11{2, "a"},
				"d",
				PackedDuplicateP11{4, "e"},
			},
		},
	}
	for _, tc := range table {
		out := p11(tc.arg.([]string))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
