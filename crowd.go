package main

func crowd1(value string) int {
	hashMap := make(map[rune]int)

	for _, v := range value {
		hashMap[v]++
	}

	for i, v := range value {
		if hashMap[v] == 1 {
			return i
		}
	}
	return -1
}

func crowd2(value string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range value {
		if open, found := brackets[v]; found {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}

func crowd3(arr []int) int {
	numSet := make(map[int]struct{})
	for _, num := range arr {
		numSet[num] = struct{}{}
	}

	maxLength := 0

	for num := range numSet {
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			length := 1

			for _, exists := numSet[currentNum+1]; exists; _, exists = numSet[currentNum+1] {
				currentNum++
				length++
			}

			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}
