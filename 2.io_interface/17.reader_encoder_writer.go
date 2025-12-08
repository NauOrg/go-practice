package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {

	// file, _ := os.Create("data.gz")
	// w := gzip.NewWriter(file)
	// w.Write([]byte("hello naushad"))
	// w.Close()

	// 1. open .gz file
	file, err := os.Open("data.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. gunzip reader (decompress while reading)
	gzipR, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gzipR.Close()

	// 3. base64 writer (encode while writing)
	b64W := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer b64W.Close()

	// 4. pipeline: (gzip → base64 → stdout)
	n, err := io.Copy(b64W, gzipR)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\nProcessed %d bytes\n", n)
}
