package main

import "strings"

type EncodedListP12 struct {
	n int
	e string
}

func p12(list []EncodedListP12) []string {
	var sb strings.Builder
	for _, e := range list {
		sb.WriteString(strings.Repeat(e.e, e.n))
	}
	return strings.Split(sb.String(), "")
}
