package main

import (
	"reflect"
	"testing"
)

func TestP14(t *testing.T) {
	table := []Case{
		{
			name: "should duplicate the elements of a list",
			arg:  []string{"a", "b", "c", "d"},
			expected: []string{
				"a",
				"a",
				"b",
				"b",
				"c",
				"c",
				"d",
				"d",
			},
		},
	}
	for _, tc := range table {
		out := p14(tc.arg.([]string))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
