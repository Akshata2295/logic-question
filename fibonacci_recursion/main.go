package main

import (
    "fmt"
)

func main() {
	var n int 
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n)
    for i:=0;i<n;i++ {
    result := fibonacci(i)
    fmt.Println(result)
}
}

func fibonacci(n int) int {
    if n <= 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}
