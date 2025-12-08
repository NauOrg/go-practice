package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("User(%s, %d years)", u.Name, u.Age)
}

func main() {
	user := User{"Naushad", 25}
	fmt.Println(user) // calls .String()
}
