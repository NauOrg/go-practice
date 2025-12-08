package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	// Stage 1
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
	}()

	// Stage 2 (3 parallel workers)
	for w := 1; w <= 3; w++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			for v := range ch1 {
				ch2 <- v * v // square
			}
		}(w)
	}

	// Close ch2 when all workers are done
	go func() {
		wg2.Wait()
		close(ch2)
	}()

	// Stage 3
	for result := range ch2 {
		fmt.Println("Result:", result)
	}
}
