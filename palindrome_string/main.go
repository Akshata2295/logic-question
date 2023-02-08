package main

import (
	"fmt"
)

func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _ , v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	str := "Hello"

	fmt.Println(str)
	// invoke reverseString
	fmt.Println(reverseString(str))

	Palindrome := reverseString(str)

	if str == Palindrome {
		fmt.Println("IS Palindrome")
	} else {
		fmt.Println("Is not Palindrome")
	}

}