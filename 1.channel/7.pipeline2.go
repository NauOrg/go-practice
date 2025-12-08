package main

import (
	"fmt"
	"sync"
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

	RawRowCh := make(chan RawRow)
	CleanRowCh := make(chan CleanRow)

	var extractWg sync.WaitGroup
	var transformWg sync.WaitGroup

	data := []RawRow{
		{1, "name1", 30},
		{2, "name2", 40},
		{3, "name3", 50},
		{4, "name4", 60},
		{5, "name5", 70},
	}

	//-----------------------
	// STAGE 1 — EXTRACTOR
	//-----------------------
	extractWg.Add(1)
	go func() {
		defer extractWg.Done()
		for _, rawRow := range data {
			RawRowCh <- rawRow
		}
		close(RawRowCh)
	}()

	//-----------------------
	// STAGE 2 — TRANSFORM (3 workers)
	//-----------------------
	for w := 0; w < 3; w++ {
		transformWg.Add(1)
		go func(w int) {
			defer transformWg.Done()
			for rawRow := range RawRowCh {
				CleanRowCh <- CleanRow{
					ID:     rawRow.ID,
					Name:   rawRow.Name,
					Score:  rawRow.Score,
					Passed: rawRow.Score >= 50,
				}
			}
		}(w)
	}

	go func() {
		transformWg.Wait()
		close(CleanRowCh)
	}()

	//-----------------------
	// STAGE 3 — LOADER
	//-----------------------
	for cleanRow := range CleanRowCh {
		fmt.Println(cleanRow)
	}

	extractWg.Wait()

}
