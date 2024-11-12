package main

func bubbleSort2(arr []int) []int {
	n := len(arr)

	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			jj := j + 1
			if arr[j] > arr[jj] {
				arr[j], arr[jj] = arr[jj], arr[j]
			}
		}
	}
	return arr
}
