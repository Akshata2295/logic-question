package main

import (
    "fmt"
)

func main() {
    str1 := "listen"
    str2 := "silent"

    if isAnagram(str1, str2) {
        fmt.Println("The strings are anagrams.")
    } else {
        fmt.Println("The strings are not anagrams.")
    }
}

func isAnagram(str1, str2 string) bool {
    if len(str1) != len(str2) {
        return false
    }

    // Create a map to store the frequency of characters
    freq := make(map[rune]int)

    // Iterate over each character in the first string
    // and increment the frequency in the map
    for _, r := range str1 {
        freq[r]++
    }

    // Iterate over each character in the second string
    // and decrement the frequency in the map
    for _, r := range str2 {
        if freq[r] == 0 {
            return false
        }
        freq[r]--
    }

    return true
}
