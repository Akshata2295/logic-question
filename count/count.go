
package main

import (
	"fmt"
)

func main() {
	input := "aabbcdddededfg"
	charCount := map[rune]int{}
	for _, char := range input {
		charCount[char]++
	}
	output := make(map[string]int)
	for key, value := range charCount {
		output[string(key)] = value
	}
	fmt.Println(output)
}
