package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 9}
    var duplicates []int

    // Create a map to store the count of each element
    m := make(map[int]int)
    for _, v := range arr {
        m[v]++
    }

    // Iterate through the map and append elements with count greater than 1 to the duplicates slice
    for key, value := range m {
        if value > 1 {
            duplicates = append(duplicates, key)
        }
    }

    fmt.Println("Duplicate numbers in the array:", duplicates)
}
