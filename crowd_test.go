package main

import (
	"fmt"
	"testing"
)

func TestCrowd1(t *testing.T) {
	s := "loveleetcode"
	result := crowd1(s)

	if result != 2 {
		t.Errorf("Incorrect")
	} else {
		fmt.Printf("Crowd1: %v\n", result)
	}
}

func TestCrowd2(t *testing.T) {
	s := "[{})()]"

	result := crowd2(s)

	if result == false {
		fmt.Printf("Crowd2: %v\n", result)
	} else {
		t.Errorf("Incorrect")
	}

}

func TestCrowd3(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 10, 11, 12, 13, 14, 15}
	expected := 6
	res := crowd3(arr)
	if res != expected {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Crowd3: %v\n", res)
	}
}
