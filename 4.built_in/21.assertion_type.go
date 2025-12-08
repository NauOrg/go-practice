package main

import "fmt"

type PersonType struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var value1 interface{} = "hello"
	var value2 interface{} = 32
	var value3 interface{} = PersonType{"naushad", 30}

	checktype(value1)
	checktype(value2)
	checktype(value3)

	checktypeSwitch(value1)
	checktypeSwitch(value2)
	checktypeSwitch(value3)
}

func checktype(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Println("value is string", str)
	}
	if num, ok := value.(int); ok {
		fmt.Println("value is int", num)
	}
	if person, ok := value.(PersonType); ok {
		fmt.Println("value is person", person)
	}
}

func checktypeSwitch(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("value is string in switch", v)
	case int:
		fmt.Println("value is int in switch", v)
	case PersonType:
		fmt.Println("value is person in switch", v)

	}
}
