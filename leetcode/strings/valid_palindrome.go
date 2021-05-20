package strings

import (
	"strings"
	"unicode"
)

// isPalindrome checks is string is valid palindrome.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/883/
func isPalindrome(s string) bool {
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
