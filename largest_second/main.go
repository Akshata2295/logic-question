package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7}
	large1 := arr[0]
	large2 := arr[0]

	for i := 1; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {
			large2 = arr[i]
		}
	}
	fmt.Println("Second largest element is: ", large2)
	fmt.Println("Largest element:", large1)
}