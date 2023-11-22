package main

func p08(list []string) []string {
	uniq := []string{}
	last := ""
	for _, char := range list {
		if char == last {
			continue
		}
		last = char
		uniq = append(uniq, char)
	}
	return uniq
}

