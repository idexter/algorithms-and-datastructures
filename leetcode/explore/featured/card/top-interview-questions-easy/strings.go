package top_interview_questions_easy

import (
	"strings"
	"unicode"
)

// reverseString Problem: Reverse String.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/879/
//
// Hint 1:
//	The entire logic for reversing a string is based on using the opposite directional two-pointer approach!
func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

// isPalindromeString Problem: Valid Palindrome.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/883/
func isPalindromeString(s string) bool {
	str := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, s)
	str = strings.ToLower(str)

	if len(str) == 2 {
		if str[0] != str[1] {
			return false
		}
		return true
	}

	k := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[k] {
			return false
		}
		k--
	}
	return true
}
