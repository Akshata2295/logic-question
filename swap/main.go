package main

import "fmt"

func main() {
	var a,b int 
	fmt.Println("Enter value for a and b")
	fmt.Scanln(&a, &b)
	a,b = b,a
	fmt.Println("Swap of two numbers :", a, b)
}