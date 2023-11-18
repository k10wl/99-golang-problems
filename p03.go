package main

func p03(k int, list []int) int {
	if k >= len(list) || k < 0 {
		return -1
	}
	return list[k]
}
