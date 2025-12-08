package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Raw struct {
	ID    int
	Value string
}

type Clean struct {
	ID  int
	Num int
}

type Batch struct {
	No    int
	Items []Clean
}

type LoadResult struct {
	BatchNo int
	Msg     string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rawCh := make(chan Raw, 20)
	cleanCh := make(chan Clean, 20)
	batchCh := make(chan Batch, 10)

	resultCh := make(chan LoadResult, 20)
	dlqCh := make(chan Batch, 10) // ❗ Dead-letter queue
	errCh := make(chan error, 10)

	var wgExtract, wgTransform, wgBatch, wgLoad sync.WaitGroup

	// ------------------------------------
	// 1️⃣ Extract
	// ------------------------------------
	wgExtract.Add(1)
	go func() {
		defer wgExtract.Done()
		defer close(rawCh)

		rawData := []Raw{
			{1, "10"}, {2, "xyz"}, {3, "25"}, {4, "hello"},
			{5, "40"}, {6, "50"}, {7, "60"}, {8, "oops"}, {9, "70"},
		}
		for _, r := range rawData {
			rawCh <- r
		}
	}()

	// ------------------------------------
	// 2️⃣ Transform (skip bad rows)
	// ------------------------------------
	for w := 0; w < 3; w++ {
		wgTransform.Add(1)
		go func() {
			defer wgTransform.Done()
			for raw := range rawCh {
				var num int
				_, err := fmt.Sscan(raw.Value, &num)
				if err != nil {
					errCh <- fmt.Errorf("skipping row %d: %v", raw.ID, err)
					continue
				}
				cleanCh <- Clean{raw.ID, num}
			}
		}()
	}

	go func() {
		wgTransform.Wait()
		close(cleanCh)
	}()

	// ------------------------------------
	// 3️⃣ Batch Stage
	// ------------------------------------
	wgBatch.Add(1)
	go func() {
		defer wgBatch.Done()
		defer close(batchCh)

		batchSize := 3
		batch := make([]Clean, 0, batchSize)
		batchNo := 1

		for clean := range cleanCh {
			batch = append(batch, clean)
			if len(batch) == batchSize {
				batchCh <- Batch{batchNo, batch}
				batch = batch[:0]
				batchNo++
			}
		}
		if len(batch) > 0 {
			batchCh <- Batch{batchNo, batch}
		}
	}()

	// ------------------------------------
	// 4️⃣ Load Stage (Fan-out + DLQ)
	// ------------------------------------
	wgLoad.Add(1)
	go func() {
		defer wgLoad.Done()
		defer close(resultCh)
		defer close(dlqCh)
		defer close(errCh)

		for batch := range batchCh {
			// Fan-out loading
			errDB := loadToDB(batch)
			errWH := loadToWarehouse(batch)

			// If both failed → send to DLQ
			if errDB != nil && errWH != nil {
				dlqCh <- batch
				errCh <- fmt.Errorf("Batch %d FAILED → DLQ", batch.No)
				continue
			}

			// Success
			resultCh <- LoadResult{batch.No, "Batch loaded to all destinations"}
		}
	}()

	// ------------------------------------
	// Output
	// ------------------------------------

	// Successful loads
	for r := range resultCh {
		fmt.Println("✔", r.Msg, "| Batch", r.BatchNo)
	}

	// DLQ items
	for dlq := range dlqCh {
		fmt.Println("❌ DLQ → Batch", dlq.No, dlq.Items)
	}

	// Processing errors
	for e := range errCh {
		fmt.Println("⚠", e)
	}
}

// ------------------------------------
// Mock Loader Functions (Fan-out)
// Random failures to demonstrate DLQ
// ------------------------------------

func loadToDB(batch Batch) error {
	if rand.Float64() < 0.30 { // 30% fail probability
		return fmt.Errorf("DB failed on batch %d", batch.No)
	}
	fmt.Printf("📌 DB loaded batch %d\n", batch.No)
	return nil
}

func loadToWarehouse(batch Batch) error {
	if rand.Float64() < 0.30 { // 30% fail probability
		return fmt.Errorf("Warehouse failed on batch %d", batch.No)
	}
	fmt.Printf("📦 Warehouse loaded batch %d\n", batch.No)
	return nil
}
