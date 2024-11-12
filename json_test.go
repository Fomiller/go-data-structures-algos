package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	data := []byte(`{"name": "Alice", "age": 30, "email": "alice@example.com"}`)
	expected := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	result := jsonDecode(data)
	if result != expected {
		t.Errorf("Incorrect: %v", result)
	}

	fmt.Printf("JsonDecode: %v\n", result)
}

func TestJsonEncode(t *testing.T) {
	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	expected := []byte(`{"name": "Alice", "age": 30, "email": "alice@example.com"}`)
	result := jsonEncode(p)

	if bytes.Equal(result, expected) {
		t.Errorf("Incorrect JsonEncode")
	}

	fmt.Printf("JsonEncode: Pass\n")
}
