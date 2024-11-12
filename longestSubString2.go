package main

func longestSubString2(input string) string {
	charMap := make(map[rune]int)
	start := 0
	longest := ""

	for i, char := range input {
		if idx, found := charMap[char]; found && idx >= start {
			start = idx + 1
		}

		charMap[char] = i

		if i-start+1 > len(longest) {
			longest = input[start : i+1]
		}

	}
	return longest
}
