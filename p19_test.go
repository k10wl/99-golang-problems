package main

import (
	"reflect"
	"testing"
)

func TestP19(t *testing.T) {
	type args struct {
		i    int
		list []string
	}
	table := []Case{
		{
			name: "should rotate a list 3 places to the left",
			arg: args{
				3,
				[]string{
					"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
				},
			},
			expected: []string{
				"d", "e", "f", "g", "h", "i", "j", "k", "a", "b", "c",
			},
		},
		{
			name: "should rotate a list 2 places to the right",
			arg: args{
				-2,
				[]string{
					"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
				},
			},
			expected: []string{
				"j", "k", "a", "b", "c", "d", "e", "f", "g", "h", "i",
			},
		},
	}
	for _, tc := range table {
		out := p19(tc.arg.(args).i, tc.arg.(args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
