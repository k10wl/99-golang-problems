package main

func p26(k int, list []string) [][]string {
	if k == 0 {
		return [][]string{{}}
	}
	if len(list) == 0 {
		return [][]string{}
	}
	head := list[0]
	tail := list[1:]
	headCombination := p26(k-1, tail)
	for i := range headCombination {
		headCombination[i] = append([]string{head}, headCombination[i]...)
	}
	tailCombinations := p26(k, tail)
	return append(headCombination, tailCombinations...)
}
