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
	var n int
	fmt.Println("Enter Number")
	fmt.Scan(&n)
	for i:=2;i<=n;i++ {
	if Prime(i) {
		fmt.Println(i, "It is prime")
	} else {
		fmt.Println(i, "It is not prime")
	}
}
}

