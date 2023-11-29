package main

import "math/rand"

func p24(n int, m int) []int {
	pool := make([]int, n)
	for i := range pool {
		pool[i] = rand.Intn(m)
	}
	return pool
}
