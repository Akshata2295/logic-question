package main

import "fmt"

func Prime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i<n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 25
	for i:=2; i<=n;i++ {
	if Prime(i) {

		fmt.Println("Prime numbers", i, " ")
	}
}
}
