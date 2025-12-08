/*
package main

import (
"bufio"
"bytes"
"os"
)

func main() {
file1, _ := os.Create("byte.txt")
file1.Write([]byte("byte text"))

file2, _ := os.Create("byteBuffer.txt")
var buffered bytes.Buffer
buffered.WriteString("buffer string \n")
buffered.Write([]byte("buffer byte \n"))
file2.Write(buffered.Bytes())

file3, _ := os.Create("bufio.txt")
buffio := bufio.NewWriter(file3)
buffio.WriteString("bufio string \n")
buffio.WriteString("bufio string2 \n")
buffio.Flush()
}
*/

/*
package main

import (
"bufio"
"bytes"
"compress/gzip"
"os"
)

func main() {
file1, _ := os.Create("byte.gz")
gzipWriter1 := gzip.NewWriter(file1)
gzipWriter1.Write([]byte("[]byte gzip"))

var buffered bytes.Buffer
gzipWriter2 := gzip.NewWriter(&buffered)
gzipWriter2.Write([]byte("buffer string \n"))
gzipWriter2.Write([]byte("buffer byte \n"))
gzipWriter2.Close()
gzipWriter2File, _ := os.Create("byteBuffer.gz")
gzipWriter2File.Write(buffered.Bytes())
defer gzipWriter2File.Close()

file3, _ := os.Create("bufio.gz")
buffio := bufio.NewWriter(file3)
gzipWriter3 := gzip.NewWriter(buffio)
gzipWriter3.Write([]byte("bufio string1"))
gzipWriter3.Write([]byte("bufio string2"))
buffio.Flush()
}
*/

package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	file1, _ := os.Create("byte.zip")
	defer file1.Close()
	zipWriter1 := zip.NewWriter(file1)
	defer zipWriter1.Close()
	entry, _ := zipWriter1.Create("msg.txt") // file inside ZIP
	entry.Write([]byte("Hello Naushad (zip + []byte)\n"))

	var buffered bytes.Buffer
	zipWriter2 := zip.NewWriter(&buffered)
	entry2, _ := zipWriter2.Create("msg.txt")
	entry2.Write([]byte("Hello Naushad (zip + bytes.Buffer)\n"))
	zipWriter2.Close()
	zipWriter2File, _ := os.Create("byteBuffer.zip")
	zipWriter2File.Write(buffered.Bytes())
	defer zipWriter2File.Close()

	file3, _ := os.Create("bufio.zip")
	buffio := bufio.NewWriter(file3)
	zipWriter3 := zip.NewWriter(buffio)
	entry3, _ := zipWriter3.Create("msg.txt")
	entry3.Write([]byte("Hello Naushad (zip + bufio.Writer)\n"))
	entry3.Write([]byte("Another line...\n"))

	buffio.Flush()

	http.HandleFunc("/byte", func(w http.ResponseWriter, r *http.Request) {
		data := []byte("Hello Naushad (HTTP + []byte)\n")
		w.Write(data)
	})

	http.HandleFunc("/byteBuffer", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer

		buf.WriteString("Hello Naushad (HTTP + bytes.Buffer)\n")
		buf.WriteString("Appending more data into the buffer...\n")

		w.Write(buf.Bytes()) // write once
	})

	http.HandleFunc("/bufio", func(w http.ResponseWriter, r *http.Request) {
		buffered := bufio.NewWriter(w)

		buffered.WriteString("Hello Naushad (HTTP + bufio)\n")
		buffered.WriteString("Line 2\n")
		buffered.WriteString("Line 3\n")

		buffered.Flush() // VERY IMPORTANT!
	})

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
