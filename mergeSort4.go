package main

func mergeSort4(input []int) []int {
	if len(input) < 2 {
		return input
	}
	mid := len(input) / 2
	left := input[:mid]
	right := input[mid:]

	return merge4(mergeSort4(left), mergeSort4(right))
}

func merge4(left []int, right []int) []int {
	output := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			output = append(output, left[i])
			i++
		} else {
			output = append(output, right[j])
			j++
		}
	}

	output = append(output, left[i:]...)
	output = append(output, right[j:]...)

	return output
}
