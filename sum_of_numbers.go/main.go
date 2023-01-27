package main

import "fmt"

func main() {
	sum()

}

func sum() {
	var n int 
	fmt.Println("The sum of first n numbers")
	fmt.Scanln(&n)
	sum := (n*(n+1))/2
	fmt.Println(sum)
	

}