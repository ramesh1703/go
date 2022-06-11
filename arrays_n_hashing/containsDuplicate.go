package main

import (
	"fmt"
)

func main() {
	input := []int{1, 3, 2, 4}
	fmt.Println(containsDuplicate(input))
}

func containsDuplicate(nums []int) bool {
	uniq := make(map[int]int)
	for _, num := range nums {
		//fmt.Println("v ", v)
		if _, ok := uniq[num]; ok {
			return true
		}
		uniq[num] = num
	}
	return false
}
