package main

import "fmt"

func main() {
	input := "00010110"
	value := 0

	for _, char := range input {
		if char == '0' {
			value += 1
		} else if char == '1' {
			value += 2
		}
	}

	fmt.Println("Output:", value)

}