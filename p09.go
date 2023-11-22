package main

type PackedDuplicatesP09 [][]string

func p09(list []string) PackedDuplicatesP09 {
	pack := PackedDuplicatesP09{}
	last := ""
	for _, char := range list {
		if last == char {
			pack[len(pack)-1] = append(pack[len(pack)-1], char)
			continue
		}
		pack = append(pack, []string{char})
		last = char
	}
	return pack
}

