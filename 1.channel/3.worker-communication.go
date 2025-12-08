package main

import (
	"fmt"
	"time"
)

func worker(work int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d \n", work, j)
		time.Sleep(time.Microsecond * 200)
		results <- j * 2
	}
}
func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 0; w < 5; w++ {
		go worker(w, jobs, results)
	}

	for w := 0; w < 5; w++ {
		jobs <- w
	}

	close(jobs)

	for w := 0; w < 5; w++ {
		fmt.Println("results", <-results)
	}
}
