package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var atomicCounter int64
	var mutexCounter int64
	var m sync.Mutex

	start := time.Now()
	var wg1 sync.WaitGroup
	for i := 0; i < 1_000_000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	wg1.Wait()
	fmt.Println("Atomic duration:", time.Since(start))

	start = time.Now()
	var wg2 sync.WaitGroup
	for i := 0; i < 1_000_000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			m.Lock()
			mutexCounter++
			m.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("Mutex duration:", time.Since(start))
}
