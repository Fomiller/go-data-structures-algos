package main

import (
	"fmt"
	"testing"
)

func TestLeet1(t *testing.T) {
	nums := []int{1, 10, 23, 4}
	target := 33

	result := leet1(nums, target)

	if result[0] != 1 && result[1] != 2 {
		t.Errorf("Incorrect %v:%v", result[0], result[1])
	}
	fmt.Printf("Leet1: %v:%v\n", result[0], result[1])
}

func TestLeet2(t *testing.T) {
	nums := []int{1, 10, 23, 4, 43, 8}
	target := 44

	res := leet2(nums, target)

	if res[0] != 0 && res[1] != 4 {
		t.Errorf("Incorrect %v:%v", res[0], res[1])
	}

	fmt.Printf("Leet2: %v:%v\n", res[0], res[1])
}

func TestLeet3(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 26

	res := leet3(nums, target)

	if nums[res[0]]+nums[res[1]] != target {
		t.Errorf("Incorrect %v:%v", res[0], res[1])
	}

	fmt.Printf("Leet3: %v:%v\n", res[0], res[1])
}

func TestLeet4(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	res := leet4(nums)

	if res != expected {
		t.Errorf("Incorrect: %v", res)
	}

	fmt.Printf("Leet3: %v\n", res)
}
