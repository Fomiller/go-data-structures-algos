package main

import (
	"sort"
)

func binarySearch5(arr []int, target int) int {
	sort.Ints(arr)

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
