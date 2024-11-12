package main

func twoSum(input []int, target int) [2]int {
	nums := make(map[int]int)

	for i, v := range input {
		complement := target - v
		if idx, ok := nums[complement]; ok {
			return [2]int{idx, i}
		} else {
			nums[v] = i
		}
	}

	return [2]int{}
}
