package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name    string  `json:"name" validate:"required"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address"`
}

type Address struct {
	Street string `json:"steet" validate:"required"`
}

func main() {
	a := Address{}
	p := Person{Age: 30, Address: a}

	fmt.Println("struct", p)
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err)
		}
	}

	m, _ := json.Marshal(p)
	fmt.Println("marshal", string(m))

	var um Person
	json.Unmarshal([]byte(`{"name":"ansari","address":{"steet":"stree1"}}`), &um)
	fmt.Println("unmarshal", um)
}

/*
jsonData := `{"name":"Alice","age":30,"email":"alice@example.com","address":{"city":"New York","zip":"10001"}}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonData), &result)
	fmt.Println(result["name"])
	address, ok := result["address"].(map[string]interface{})
	if ok {
		fmt.Println(address["city"])
	}
*/
