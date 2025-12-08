package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int)

	go func() {
		fmt.Println("sending starts")
		jobs <- 5
		fmt.Println("sending done")
	}()

	fmt.Println("RECEVING STARTS")
	C := <-jobs
	fmt.Println("RECEVED", C)
}
