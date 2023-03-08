package main

import (

	"sync"
	"fmt"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)
    for i:=1;i<=10;i++ {
		go func(n int){
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

