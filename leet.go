package main

func leet1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for ii := i + 1; ii < n-1; ii++ {
			if nums[i]+nums[ii] == target {
				return []int{i, ii}
			}
		}
	}
	return nil
}

func leet2(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, found := hashMap[complement]; found {
			return []int{idx, i}
		}

		hashMap[num] = i
	}
	return nil
}

func leet3(arr []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range arr {
		// calculate a compliment
		complement := target - v
		// check if the complementing value exists in map
		if idx, ok := hashMap[complement]; ok {
			return []int{i, idx}
		}
		// if it does not exist create an entry for the current value and index
		hashMap[v] = i
	}
	return nil
}

func leet4(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	min := arr[0]
	max := 0

	for _, v := range arr[1:] {
		if v < min {
			min = v
		} else {
			profit := v - min
			if profit > max {
				max = profit
			}
		}
	}

	return max
}
