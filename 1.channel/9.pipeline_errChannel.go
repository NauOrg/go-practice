package main

import (
	"fmt"
	"sync"
)

// ----------------------
// Types
// ----------------------

type Raw struct {
	ID    int
	Value string
}

type Clean struct {
	ID  int
	Num int
}

type LoadResult struct {
	ID  int
	Msg string
}

// ----------------------
// Pipeline
// ----------------------

func main() {

	rawCh := make(chan Raw, 10)
	cleanCh := make(chan Clean, 10)
	resultCh := make(chan LoadResult, 10)
	errCh := make(chan error, 10)

	var wgExtract, wgTransform, wgLoad sync.WaitGroup

	// ----------------------
	// 1️⃣ Extract Stage
	// ----------------------
	wgExtract.Add(1)
	go func() {
		defer wgExtract.Done()
		defer close(rawCh)

		rawData := []Raw{
			{1, "10"},
			{2, "xyz"}, // ❌ bad row
			{3, "25"},
			{4, "hello"}, // ❌ bad row
			{5, "40"},
		}

		for _, r := range rawData {
			rawCh <- r
		}
	}()

	// ----------------------
	// 2️⃣ Transform Stage (skip bad rows)
	// ----------------------
	for w := 0; w < 3; w++ { // 3 workers
		wgTransform.Add(1)
		go func() {
			defer wgTransform.Done()

			for raw := range rawCh {

				// Try to convert Value → int
				var num int
				_, err := fmt.Sscan(raw.Value, &num)
				if err != nil {
					errCh <- fmt.Errorf("skipping row %d: %v", raw.ID, err)
					continue // ❗ skip bad row
				}

				cleanCh <- Clean{raw.ID, num}
			}
		}()
	}

	go func() {
		wgTransform.Wait()
		close(cleanCh)
	}()

	// ----------------------
	// 3️⃣ Load Stage
	// ----------------------
	for w := 0; w < 2; w++ { // 2 workers
		wgLoad.Add(1)
		go func() {
			defer wgLoad.Done()

			for clean := range cleanCh {
				msg := fmt.Sprintf("Loaded ID=%d Num=%d", clean.ID, clean.Num)
				resultCh <- LoadResult{clean.ID, msg}
			}
		}()
	}

	go func() {
		wgLoad.Wait()
		close(resultCh)
		close(errCh)
	}()

	// ----------------------
	// OUTPUT: Read results + errors
	// ----------------------
	for res := range resultCh {
		fmt.Println("✔", res.Msg)
	}

	for err := range errCh {
		fmt.Println("⚠", err)
	}

}
