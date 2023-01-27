package main

import "fmt"

func main() {
EvenOdd()
}

func EvenOdd() {
	var n int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n)
	for x := 1; x <= n; x++ {
	if(x%2 != 0) {
		fmt.Println("It is even number", x)
	} else {
		fmt.Println("It is odd number", x)
	}
}
}