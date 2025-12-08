package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 2️⃣ HTTP server that returns downloadable ZIP file
func handleZipDownload(w http.ResponseWriter, r *http.Request) {
	log.Println("➡ Client called /gzip-response to download ZIP")

	// Create ZIP in memory
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	file, err := zipWriter.Create("data.txt")
	if err != nil {
		http.Error(w, "Failed to create zip", http.StatusInternalServerError)
		return
	}
	file.Write([]byte("Hello from server! This file is inside the ZIP archive."))

	zipWriter.Close()

	// Headers for download
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=report.zip")
	w.Write(buf.Bytes())
}

// 3️⃣ HTTP server that accepts gzip request body
func handleGzipRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("➡ Client called /gzip-request")

	var body []byte
	var err error

	if r.Header.Get("Content-Encoding") == "gzip" {
		log.Println("📩 Client sent GZIP")
		gzReader, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, "Invalid gzip body", http.StatusBadRequest)
			return
		}
		defer gzReader.Close()

		body, err = io.ReadAll(gzReader)
	} else {
		log.Println("📩 Client sent PLAIN")
		body, err = io.ReadAll(r.Body)
	}

	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	log.Println("📝 Decompressed/received body:", string(body))
	w.Write([]byte("Received OK"))
}

func main() {
	http.HandleFunc("/gzip-response", handleZipDownload)
	http.HandleFunc("/gzip-request", handleGzipRequest)

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
