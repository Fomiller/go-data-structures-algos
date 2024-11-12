package main

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// while left is less than or equal to right bounds
	for left <= right {
		// set middle index
		mid := left + (right-left)/2

		// if the arr index of mid is equal to the target return the index
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			// if the value at mid is less than the target increase the left bound
			left = mid + 1
		} else {
			// if the value at mid is greater than the target increase the right bound
			right = mid - 1
		}
	}
	// return -1 if not found
	return -1
}
