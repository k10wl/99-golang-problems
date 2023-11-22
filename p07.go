package main

import "fmt"

type NestedListP07 []interface{}

func p07(list NestedListP07) []int {
	res := []int{}
	for _, el := range list {
		switch e := el.(type) {
		case int:
			{
				res = append(res, e)
			}
		case NestedListP07:
			{
				res = append(res, p07(e)...)
			}
		default:
			{
				panic(fmt.Sprintf("Unexpected element type - %s", e))
			}
		}
	}
	return res
}
