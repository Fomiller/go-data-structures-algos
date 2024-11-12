package main

func mergeSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	return merge2(mergeSort2(left), mergeSort2(right))
}

func merge2(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
