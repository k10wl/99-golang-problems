package main

import (
	"reflect"
	"testing"
)

func TestP08(t *testing.T) {
	table := []Case{
		{
			name:     "should eliminate consecutive duplicates of list elements",
			arg:      []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"},
			expected: []string{"a", "b", "c", "a", "d", "e"},
		},
	}
	for _, tc := range table {
		out := p08(tc.arg.([]string))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}

