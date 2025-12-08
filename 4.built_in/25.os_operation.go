package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//📌 Create file
	file, err := os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello from Go!")
	file.Close()

	//📌 Open + Read file
	readfile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := io.ReadAll(readfile)
	fmt.Println(string(data))
	readfile.Close()

	//🔥 5️⃣ File / Folder Metadata
	info, _ := os.Stat("demo.txt")
	fmt.Println("info.Name()", info.Name())
	fmt.Println("info.Size()", info.Size())
	fmt.Println("info.IsDir()", info.IsDir())
	fmt.Println("info.ModTime()", info.ModTime())

	//📌 Remove a file
	os.Remove("demo.txt")

	//🔥 2️⃣ Directory Operations
	os.Mkdir("myfolder", 0755)
	// os.Remove("myfolder")

	//List files in folder
	entries, _ := os.ReadDir(".")
	for _, e := range entries {
		fmt.Println("e.Type(), e.Name()", e.Type(), e.Name())
	}

	//🔥 3️⃣ Environment Variables
	fmt.Println(os.Getenv("PATH"))

	os.Setenv("MY_NAME", "naushad")
	fmt.Println(os.Getenv("MY_NAME"))

}
