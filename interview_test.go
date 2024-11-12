package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceSum(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{9, 8, 7, 6}
	result := sliceSum(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect: %v", result)

	} else {
		fmt.Printf("SliceSum: %v\n", result)
	}
}
