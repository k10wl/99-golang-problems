package main

import "math/rand"

func p25(list []string) []string {
	res := make([]string, len(list))
	copy(res, list)
	for i := range list {
		k := rand.Intn(len(list) - 1)
		res[i], res[k] = res[k], res[i]
	}
	return res
}
