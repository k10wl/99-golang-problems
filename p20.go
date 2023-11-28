package main

type KindaTuppleP20 struct {
	list    []string
	removed string
}

func p20(k int, list []string) KindaTuppleP20 {
	tmp := list[k]
	return KindaTuppleP20{append(list[:k], list[k+1:]...), tmp}
}
