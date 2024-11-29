package main

import "fmt"

func main(){
	ch := make(chan int,10)

	go func() {
		ch <- 42
		ch <- 52
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}