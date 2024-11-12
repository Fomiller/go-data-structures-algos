package main

import (
	"slices"
	"strings"
)

func longestSubString(str string) string {
	tmp := make([]string, 0, len(str))
	longest := make([]string, 0, len(str))
	letters := strings.Split(str, "")

	for _, v := range letters {
		if !slices.Contains(tmp, v) {
			tmp = append(tmp, v)
		} else {
			if len(tmp) > len(longest) {
				longest = tmp
			}
			tmp = []string{v}
		}
	}

	if len(tmp) > len(longest) {
		return strings.Join(tmp, "")
	} else {
		return strings.Join(longest, "")
	}
}
