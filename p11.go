package main

import (
	"log"
)

type PackedDuplicateP11 struct {
	n int
	e string
}

type OutP11 interface{}

func p11(list []string) []interface{} {
	d := []interface{}{}
	prev := ""
	for _, e := range list {
		if e == prev {
			i := len(d) - 1
			switch d[i].(type) {
			case string:
				{
					d[i] = PackedDuplicateP11{2, prev}
				}
			case PackedDuplicateP11:
				{
					d[i] = PackedDuplicateP11{d[i].(PackedDuplicateP11).n + 1, prev}
				}
			default:
				{
					log.Fatalf("received unexpected value: %v\n", d[i])
				}
			}
			continue
		}
		d = append(d, e)
		prev = e
	}
	return d
}
