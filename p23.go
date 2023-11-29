package main

import (
	"math/rand"
)

func p23(n int, list []string) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = list[rand.Intn(len(list)-1)]
	}
	return res
}
