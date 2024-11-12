package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func jsonDecode(data []byte) Person {
	var person Person

	err := json.Unmarshal(data, &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return person
	}
	return person
}

func jsonEncode(data interface{}) []byte {

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return []byte{}
	}
	return bytes
}
