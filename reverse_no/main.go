package main
import "fmt"
func reverseNumber() int {
	var num int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&num)

   res := 0
   for num>0 {
      remainder := num % 10
      res = (res * 10) + remainder
      num /= 10
   }
   return res
}

func main(){
   fmt.Println(reverseNumber())
   
}