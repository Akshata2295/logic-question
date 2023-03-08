package main

import "fmt"

func miss(arr []int) []int {
    min := arr[0]
    max := arr[0]

    for _, val := range arr {
        if val < min {
            min = val
        } else if val > max {
            max = val
        }
    }

    boolslice := make([]bool, max-min+1)

    for _, val := range arr {
        boolslice[val-min] = true
    }

    missing := []int{}

    for i, val := range boolslice {
        if !val {
            missing = append(missing, min+i)
        }
    }

    return missing
}

func main() {
    arr := []int{-1, 5, 3, 9}
    c := miss(arr)
    fmt.Println(c)

}