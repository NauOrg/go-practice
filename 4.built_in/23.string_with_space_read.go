package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println("bufio.NewReader", line)
	fmt.Println("bufio.NewReader.Split", strings.Split(line, " "))

	var first, last string
	fmt.Scanln(&first, &last)
	fmt.Println("first, last", first, last)

}
