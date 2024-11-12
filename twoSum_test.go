package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	target := 8
	result := twoSum(input, target)

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("TwoSum: %v\n", result)
	}
}
