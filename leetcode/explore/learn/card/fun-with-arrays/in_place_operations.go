package fun_with_arrays

// replaceElements Problem: Replace Elements with Greatest Element on Right Side.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/
//
// Hint 1:
//	Loop through the array starting from the end.
// Hint 2:
//	Keep the maximum value seen so far.
func replaceElements(arr []int) []int {
	max := 0

	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j != i; j-- {
			if arr[j] > max {
				max = arr[j]
				arr[i] = max
			}
		}
		max = 0
	}

	arr[len(arr)-1] = -1

	return arr
}
