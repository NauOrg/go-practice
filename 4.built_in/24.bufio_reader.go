package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()

	outFile, _ := os.Create("output.txt")
	defer outFile.Close()

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	for {
		line, err := reader.ReadString('\n')
		writer.WriteString(line)

		if err != nil {
			break
		}
	}
	writer.Flush()
	fmt.Println("file copied")
}

/*
func main() {
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()

	outFile, _ := os.Create("output.txt")
	defer outFile.Close()

	reader := bufio.NewReader(inFile)

	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)

		if err != nil {
			break
		}
	}
}
*/
