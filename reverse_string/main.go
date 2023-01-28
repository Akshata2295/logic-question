package main
import "fmt"

func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _ , v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	str := "Reverse"

	fmt.Println(str)
	// invoke reverseString
	fmt.Println(reverseString(str))
}