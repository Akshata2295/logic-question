package main

import "fmt"

func main() {
    n := 25
    for i := 2; i <= n; i++ {
        if isPrime(i) {
            fmt.Print(i, " ")
        }
    }
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
