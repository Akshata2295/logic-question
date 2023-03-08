package main

import (
    "fmt"
   // "sort"
)

func findMissingElements(arr []int) []int {
    // sort.Ints(arr) // Sort the array in ascending order

    var missing []int

    for i := 1; i < len(arr); i++ {
        for j := arr[i-1] + 1; j < arr[i]; j++ {
            missing = append(missing, j)
        }
    }

    return missing
}

func main() {
    arr := []int{9,0}
    missing := findMissingElements(arr)

    fmt.Printf("Missing elements: %v\n", missing) 
}


