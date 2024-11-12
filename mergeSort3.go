package main

func mergeSort3(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	return merge3(mergeSort3(left), mergeSort3(right))
}

func merge3(left, right []int) []int {
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
