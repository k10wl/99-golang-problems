package main

import (
	"reflect"
	"testing"
)

func TestP07(t *testing.T) {
	table := []Case{
		{
			name:     "should flatten a nested list structure",
			arg:      NestedListP07{NestedListP07{1, 1}, 2, NestedListP07{3, NestedListP07{5, 8}}},
			expected: []int{1, 1, 2, 3, 5, 8},
		},
	}
	for _, tc := range table {
		out := p07(tc.arg.(NestedListP07))
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
