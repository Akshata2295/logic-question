package main

import "fmt"

func main() {
	first := []int{1,2,3,4}
	second := []int{2,4}

	result := intersection(first, second)
	fmt.Println(result)
}

func intersection(list1, list2 []int) []int {
	var intersection []int
	for _, elem1 := range list1 {
		for _, elem2 := range list2 {
			if elem1 == elem2 {
				intersection = append(intersection, elem1)
				break
			}
		}
	}
	return intersection
}