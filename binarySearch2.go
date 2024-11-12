package main

import (
	"fmt"
	"sort"
)

func binarySearch2(arr []int, target int) (int, error) {
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] > target {
			right = mid - 1
		}

		if arr[mid] < target {
			left = mid + 1
		}

		if arr[mid] == target {
			return mid, nil
		}
	}
	return -1, fmt.Errorf("Not found")
}
