package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	target := 6

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := binarySearch(arr, target)

	if res != target-1 {
		t.Errorf("Incorrect %v", res)
	}

	fmt.Printf("Binary1: %v\n", res)
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{1, 6, 3, 9, 4, 5, 2, 7, 8, 10}
	target := 6
	res, _ := binarySearch2(arr, target)

	if res != target-1 {
		t.Errorf("Incorrect %v", res)
	}

	fmt.Printf("Binary2: %v\n", res)
}

func TestBinarySearch3(t *testing.T) {
	arr := []int{1, 6, 3, 9, 4, 5, 2, 7, 8, 10}
	target := 6
	res, _ := binarySearch2(arr, target)

	if res != target-1 {
		t.Errorf("Incorrect %v", res)
	}

	fmt.Printf("Binary3: %v\n", res)
}

func TestBinarySearch4(t *testing.T) {
	target := 6

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := binarySearch4(arr, target)

	if arr[result] != target {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Binary4: %v\n", result)
	}
}

func TestBinarySearch5(t *testing.T) {
	target := 6

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := binarySearch5(arr, target)

	if arr[result] != target {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Binary5: %v\n", result)
	}
}

func TestBinarySearch6(t *testing.T) {
	target := 6

	arr := []int{1, 2, 3, 4, 6, 6, 7, 8, 9, 10}

	result := binarySearch6(arr, target)

	if arr[result] != target {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Binary6: %v\n", result)
	}
}

func TestBinarySearch7(t *testing.T) {
	target := 3

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := binarySearch7(arr, target)

	if arr[result] != target {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Binary7: %v\n", result)
	}
}
