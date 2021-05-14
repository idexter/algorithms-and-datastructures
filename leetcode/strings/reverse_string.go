package strings

// reverseString reverses string.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/
func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
