package main

import (
	"reflect"
	"testing"
)

func TestP12(t *testing.T) {
	table := []Case{
		{
			name: "should decode a run-length encoded list",
			arg: []EncodedListP12{
				{4, "a"},
				{1, "b"},
				{2, "c"},
				{2, "a"},
				{1, "d"},
				{4, "e"},
			},
			expected: []string{
				"a", "a", "a", "a",
				"b",
				"c", "c",
				"a", "a",
				"d",
				"e", "e", "e", "e",
			},
		},
	}
	for _, tc := range table {
		out := p12(tc.arg.([]EncodedListP12))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
