package main

import "fmt"

func main() {
	n := 4
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := i; k < n; k++ {
			if i%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print("", k)
			}
		}
		fmt.Println()
	}
}
