package main

func p05(list []int) []int {
	rev := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		rev[i] = list[len(list)-i-1]
	}
	return rev
}
