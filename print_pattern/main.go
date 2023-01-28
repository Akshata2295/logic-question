package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n)
    for i := 1; i <= n; i++ {
        if i <= 5 {
            fmt.Print(i, " ")
        } else if i > 5 && i <= 10 {
            fmt.Print("* ")
        } else if i > 10 && i <= 15 {
            fmt.Print(i, " ")
        } else {
            fmt.Print("* ")
        }
    }
}
