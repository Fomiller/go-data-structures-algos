package main

func longestSubString3(input string) string {
	chars := make(map[rune]int)
	start := 0
	output := ""

	for i, char := range input {
		if found_idx, found := chars[char]; found && found_idx >= start {
			start = found_idx + 1
		}

		chars[char] = i

		if i-start+1 > len(output) {
			output = input[start : i+1]
		}
	}
	return output
}
