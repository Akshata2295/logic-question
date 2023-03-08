package main

import "fmt"

// wap to find out squareroot of given number

func square(n int) int {
    if n == 1 || n == 0 {
		return n
	  } else {
		  return n * n
	  }
}

func main() {
    fmt.Println(square(12))
}