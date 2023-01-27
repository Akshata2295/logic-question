package main

import "fmt"

func main() {
    for i := 3; i > 0; i-- {
        for j := i; j < 3; j++ {
            fmt.Print("  ")
        }
        for k := i; k > 0; k-- {
            fmt.Print(k, " ")
        }
        fmt.Println()
    }
}
