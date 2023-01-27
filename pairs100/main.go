package main

import "fmt"

func main() {
    arr := []int{80, 60, 10, 50, 30, 100, 0, 50}
    target := 100
    pairs := findPairs(arr, target)
    fmt.Println("Pairs whose sum is", target, ":", pairs)
}

func findPairs(arr []int, target int) [][]int {
    var pairs [][]int
    m := make(map[int]bool)
    for _, val := range arr {
        diff := target - val
        if m[diff] {
            pairs = append(pairs, []int{diff, val})
        }
        m[val] = true
    }
    return pairs
}
