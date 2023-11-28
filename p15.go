package main

func p15(n int, list []string) []string {
	dup := make([]string, len(list)*n)
	for i, el := range list {
		for j := i * n; j < i*n+3; j++ {
			dup[j] = el
		}
	}
	return dup
}
