package main

func p02(list []int) int {
	l := len(list)
	if l < 2 {
		return -1
	}

	return list[l-2]
}
