package arrays

// reverseArray Problem: Arrays - DS.
//
// HackerRank: https://www.hackerrank.com/challenges/arrays-ds/problem
func reverseArray(a []int32) []int32 {
	b := []int32{}
	for i := len(a); i > 0; i-- {
		b = append(b, a[i-1])
	}
	return b
}
