package main

import (
	"fmt"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	input := "abcabcdabcc"
	expected := "abcd"
	result := longestSubString(input)
	if result != expected {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("LongestSubString: %v\n", result)
	}
}

func TestLongestSubString2(t *testing.T) {
	input := "abcabcdabcc"
	expected := "abcd"
	result := longestSubString2(input)
	if result != expected {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("LongestSubString2: %v\n", result)
	}
}

func TestLongestSubString3(t *testing.T) {
	input := "abcabcdabcc"
	expected := "abcd"
	result := longestSubString3(input)
	if result != expected {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("LongestSubString3: %v\n", result)
	}
}
