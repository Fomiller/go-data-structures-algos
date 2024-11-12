package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 3, 2}
	result := bubbleSort(arr)
	if !sort.IntsAreSorted(result) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("Bubble1: %v\n", result)
	}
}

func TestBubbleSort2(t *testing.T) {
	arr := []int{100, 6, 3, 9, 4, 5, 1, 7, 8, 2}
	res := bubbleSort2(arr)

	if !sort.IntsAreSorted(res) {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Bubble2: %v\n", res)
	}
}

func TestBubbleSort3(t *testing.T) {
	arr := []int{100, 6, 3, 9, 4, 5, 1, 7, 8, 2}
	res := bubbleSort3(arr)

	if !sort.IntsAreSorted(res) {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Bubble3: %v\n", res)
	}
}

func TestBubbleSort4(t *testing.T) {
	arr := []int{100, 6, 3, 9, 4, 5, 1, 7, 8, 2}
	res := bubbleSort4(arr)

	if !sort.IntsAreSorted(res) {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Bubble4: %v\n", res)
	}
}

func TestBubbleSort5(t *testing.T) {
	arr := []int{100, 6, 3, 9, 4, 5, 1, 7, 8, 2}
	res := bubbleSort5(arr)

	if !sort.IntsAreSorted(res) {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Bubble5: %v\n", res)
	}
}

func TestBubbleSort6(t *testing.T) {
	arr := []int{100, 6, 3, 9, 4, 5, 1, 7, 8, 2}
	res := bubbleSort6(arr)

	if !sort.IntsAreSorted(res) {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Bubble6: %v\n", res)
	}
}

func TestBubbleSort7(t *testing.T) {
	arr := []int{100, 6, 3, 9, 4, 5, 1, 7, 8, 2}
	res := bubbleSort7(arr)

	if !sort.IntsAreSorted(res) {
		t.Errorf("Incorrect: %v", res)
	} else {
		fmt.Printf("Bubble7: %v\n", res)
	}
}
