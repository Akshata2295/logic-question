// Golang program to
// print the Fibonacci series using recursion

package main

import "fmt"

func printFibonacii(n int) {

	
	var n3, n1, n2 int = 0, 0, 1
 
  for i := 1; i <= n; i++ {
 
    fmt.Println(n1)
 
    n3 = n1 + n2
 
    n1 = n2
 
    n2 = n3
 
  }
}

func main() {
	var n int 
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n)
	printFibonacii(n)
}
