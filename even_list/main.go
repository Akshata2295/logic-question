package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evenNumbers := []int{}

    for i:=0; i<len(arr); i++ {
        if arr[i]%2 == 0 {
            evenNumbers = append(evenNumbers, arr[i])
        }
    }

    fmt.Println("Even numbers:", evenNumbers)
}
