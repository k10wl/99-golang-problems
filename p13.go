package main

type EncodedListP13 struct {
	n  int
	el string
}

func p13(list []string) []EncodedListP13 {
	encoded := []EncodedListP13{}
	var prev string
	for _, char := range list {
		if prev == char {
			encoded[len(encoded)-1].n++
			continue
		}
		encoded = append(encoded, EncodedListP13{1, char})
		prev = char
	}
	return encoded
}
