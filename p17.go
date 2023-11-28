package main

func p17(n int, list []string) [][]string {
	if len(list) <= n {
		return [][]string{list}
	}
	return [][]string{
		list[:n],
		list[n:],
	}
}
