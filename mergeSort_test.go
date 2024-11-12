package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	input := []int{1, 5, 3, 6, 8, 2, 3, 4}
	result := mergeSort1(input)
	if !sort.IntsAreSorted(result) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Merge1: %v\n", result)

	}
}

func TestMergeSort2(t *testing.T) {
	input := []int{1, 5, 3, 6, 8, 2, 3, 4}
	result := mergeSort2(input)
	if !sort.IntsAreSorted(result) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Merge2: %v\n", result)

	}
}

func TestMergeSort3(t *testing.T) {
	input := []int{1, 5, 3, 6, 8, 2, 3, 4}
	result := mergeSort3(input)
	if !sort.IntsAreSorted(result) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Merge3: %v\n", result)

	}
}

func TestMergeSort4(t *testing.T) {
	input := []int{1, 5, 3, 6, 8, 2, 3, 4}
	result := mergeSort4(input)
	if !sort.IntsAreSorted(result) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Merge4: %v\n", result)

	}
}
