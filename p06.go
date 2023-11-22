package main

func p06(list []int) bool {
	isPalindrome := true
	for i, el := range list {
		if el != list[len(list)-i-1] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}
