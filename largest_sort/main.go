package main

import "fmt"

func main() {
	arr := []int{3, 6, 8, 1, 9, 2, 5}
	largest := arr[0]
	secondLargest := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Second Largest:", secondLargest)
}
