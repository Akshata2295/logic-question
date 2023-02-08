package main

import "fmt"

func main() {
	var num int = 123456
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	fmt.Println("Reverse No:", res)
}