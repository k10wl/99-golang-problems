package main

func p01(list []int) int {
	if len(list) == 0 {
		return -1
	}
	return list[len(list)-1]
}
