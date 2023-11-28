package main

import "math"

func p16(n int, list []string) []string {
	filteredLength := len(list) - int(math.Floor(float64(len(list))/float64(n)))
	filtered := make([]string, filteredLength)
	skipped := 0
	for i, el := range list {
		if (i+1)%n == 0 {
			skipped++
			continue
		}
		filtered[i-skipped] = el
	}
	return filtered
}
