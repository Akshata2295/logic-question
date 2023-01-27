package main

import "fmt"

func main() {
	var n int

	var fact int = 1

	fmt.Print("Enter the Number: ")

	fmt.Scan(&n)

	if n < 0 {

		fmt.Print("\nFactorial of a negative number is not possible")

	}

	for i := 1; i <= n; i++ {

		fact = fact * i

	}

	fmt.Printf("\nFactorial is %d", fact)

}