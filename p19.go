package main

func p19(i int, list []string) []string {
	if i < 0 {
		return append(list[len(list)+i:], list[:len(list)+i]...)
	}
	return append(list[i:], list[:i]...)
}
