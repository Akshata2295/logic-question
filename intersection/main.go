package main

import "fmt"

func main() {
	first := []string{"one", "two", "three", "four"}
	second := []string{"two", "four"}

	result := intersection(first, second)
	fmt.Println(result)
}

func intersection(list1, list2 []string) []string {
	var intersection []string
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