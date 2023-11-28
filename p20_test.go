package main

import (
	"reflect"
	"testing"
)

func TestP20(t *testing.T) {
	type args struct {
		k    int
		list []string
	}
	table := []Case{
		{
			name: "should remove the Kth element from a list",
			arg: args{
				1,
				[]string{
					"a", "b", "c", "d",
				},
			},
			expected: KindaTuppleP20{
				[]string{
					"a", "c", "d",
				},
				"b",
			},
		},
	}
	for _, tc := range table {
		out := p20(tc.arg.(args).k, tc.arg.(args).list)
		if !reflect.DeepEqual(out, tc.expected) {
			t.Errorf(tc.errorInfo(out))
		}
	}
}
