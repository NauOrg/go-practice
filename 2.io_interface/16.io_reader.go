package main

import "fmt"

type HelloReader struct{}

func (HelloReader) Read(p []byte) (int, error) {
	copy(p, "hello")
	return 5, nil
}

func main() {
	buff := make([]byte, 5)
	r := HelloReader{}
	r.Read(buff)
	fmt.Println(string(buff))

}
