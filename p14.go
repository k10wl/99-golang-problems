package main

func p14(list []string) []string {
	duplicated := []string{}
	for _, el := range list {
		duplicated = append(duplicated, el, el)
	}
	return duplicated
}
