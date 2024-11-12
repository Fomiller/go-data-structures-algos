package main

func mergeSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	// Recursively sort each half
	return merge(mergeSort1(left), mergeSort1(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, ii := 0, 0

	for i < len(left) && ii < len(right) {
		if left[i] < right[ii] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[ii])
			ii++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[ii:]...)

	return result
}
