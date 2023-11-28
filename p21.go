package main

func p21(new string, at int, list []string) []string {
	s := list[:at]
	e := list[at:]
	return append(s, append([]string{new}, e...)...)
}
