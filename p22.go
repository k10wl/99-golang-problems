package main

func p22(start int, end int) []int {
	res := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		res[i-start] = i
	}
	return res
}
