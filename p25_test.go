package main

import (
	"reflect"
	"testing"
)

func TestP25(t *testing.T) {
	table := []Case{
		{
			name:     "should generate a random permutation of the elements of a list",
			arg:      []string{"a", "b", "c", "d", "e", "f", "g", "h"},
			expected: []string{"h", "f", "g", "d", "c", "e", "a", "b"},
		},
	}
	for _, tc := range table {
		out := p25(tc.arg.([]string))
		unexpected := len(out) != len(tc.arg.([]string)) ||
			reflect.DeepEqual(out, tc.arg.([]string)) ||
			!isShuffled(out, tc.arg.([]string))
		if unexpected {
			t.Errorf(tc.errorInfo(out))
		}
	}
}

func isShuffled(s1 []string, s2 []string) bool {
	used := map[string]int{}
	for _, char := range s1 {
		used[char]++
	}
	shuffled := true
	for _, char := range s2 {
		if val, ok := used[char]; !ok || val == 0 {
			shuffled = false
			break
		}
		used[char]--
	}
	return shuffled
}
