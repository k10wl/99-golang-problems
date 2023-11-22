package main

type PackedDuplicatesP10 []struct {
	count int
	char  string
}

func p10(list []string) PackedDuplicatesP10 {
	pack := PackedDuplicatesP10{}
	last := ""
	for _, char := range list {
		if char == last {
			pack[len(pack)-1].count++
			continue
		}
		pack = append(pack, struct {
			count int
			char  string
		}{
			1,
			char,
		})
		last = char
	}
	return pack
}

