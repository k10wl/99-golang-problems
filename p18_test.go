package main

import (
	"reflect"
	"testing"
)

func TestP18(t *testing.T) {
	type args struct {
		i    int
		k    int
		list []string
	}
	table := []Case{
		{
			name: "should extract a slice from a list",
			arg: args{
				3,
				7,
				[]string{
					"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
				},
			},
			expected: []string{
				"d", "e", "f", "g",
			},
		},
	}
	for _, tc := range table {
		out := p18(tc.arg.(args).i, tc.arg.(args).k, tc.arg.(args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
