package main
import "fmt"

func Palindrome(n int) string{
   num := n
   var remainder int
   res := 0
   for n>0 {
      remainder = n % 10
      res = (res * 10) + remainder
      n = n / 10
   }
   if num == res {
      return "Palindrome"
   } else {
      return "Not a Palindrome"
   }
}

func main(){
	n := 123
   fmt.Println(Palindrome(n))
}
