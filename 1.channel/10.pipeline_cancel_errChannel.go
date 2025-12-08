package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type RawRow struct {
	ID    int
	Name  string
	Score int
}

type CleanRow struct {
	ID     int
	Name   string
	Score  int
	Passed bool
}

func main() {

	data := []RawRow{
		{1, "alice", 80},
		{2, "BOB", 55},
		{3, "charlie", 90},
		{4, "dave", 40},
		{5, "ERR_ROW", 72}, // Simulate "bad" row
	}

	rawCh := make(chan RawRow)
	cleanCh := make(chan CleanRow)
	errCh := make(chan error)

	// For cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wgExtract sync.WaitGroup
	var wgTransform sync.WaitGroup

	//------------------------------
	// STAGE 1 — EXTRACTOR
	//------------------------------
	wgExtract.Add(1)
	go func() {
		defer wgExtract.Done()
		defer close(rawCh)

		for _, r := range data {

			// Simulate bad data detection
			if r.Name == "ERR_ROW" {
				errCh <- errors.New("extractor: invalid row encountered")
				cancel()
				return
			}

			select {
			case <-ctx.Done():
				return
			case rawCh <- r:
			}
		}
	}()

	//------------------------------
	// STAGE 2 — TRANSFORM (3 workers)
	//------------------------------
	workerCount := 3

	for w := 1; w <= workerCount; w++ {
		wgTransform.Add(1)
		go func(id int) {
			defer wgTransform.Done()

			for r := range rawCh {

				// Simulate worker error
				if r.Score < 0 {
					errCh <- fmt.Errorf("worker %d: negative score", id)
					cancel()
					return
				}

				clean := CleanRow{
					ID:     r.ID,
					Name:   strings.ToUpper(strings.TrimSpace(r.Name)),
					Score:  r.Score,
					Passed: r.Score >= 60,
				}

				select {
				case <-ctx.Done():
					return
				case cleanCh <- clean:
				}
			}
		}(w)
	}

	// Close cleanCh after all workers finish
	go func() {
		wgTransform.Wait()
		close(cleanCh)
	}()

	//------------------------------
	// ERROR WATCHER
	//------------------------------
	go func() {
		for err := range errCh {
			fmt.Println("ERROR:", err)
			cancel()
			return
		}
	}()

	//------------------------------
	// STAGE 3 — LOADER (Final consumer)
	//------------------------------
	fmt.Println("=== LOADED DATA ===")
	for c := range cleanCh {
		fmt.Printf("%+v\n", c)
	}

	wgExtract.Wait()
	time.Sleep(100 * time.Millisecond) // small delay for error printing
}
