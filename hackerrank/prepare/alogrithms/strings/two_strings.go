package strings

// twoStrings Problem: Two Strings.
//
// HackerRank: https://www.hackerrank.com/challenges/two-strings/problem
func twoStrings(s1 string, s2 string) string {

	strMap := make(map[rune]struct{})
	for _, v := range s2 {
		strMap[v] = struct{}{}
	}

	for _, v := range s1 {
		if _, ok := strMap[v]; ok {
			return "YES"
		}
	}
	return "NO"
}
