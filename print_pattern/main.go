package main

import "fmt"

func main() {
	var num, i int

    fmt.Print("\n Enter the Maximum Natural Number Limit = ")
    fmt.Scanln(&num)

    fmt.Println("\n Numbers from 1 to ", num, " are = ")
    for i = 1; i <= num; i++ {
		if i<=5 {
        fmt.Print(i," ")
		} else if i<= 10 {
			fmt.Print("*", " ")
		} else if i<=15 {
			fmt.Print(i, " ")
		} else if i<=20 {
			fmt.Print("*", " ")
		}
    }
}