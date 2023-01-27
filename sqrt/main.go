package main

import (
    "fmt"
    "math"
)

func main() {

    var x float64 
	fmt.Println("Enter the given number :")
	fmt.Scan(&x)
    result := math.Sqrt(x)
    fmt.Println("Square root of given number is :", result)
}