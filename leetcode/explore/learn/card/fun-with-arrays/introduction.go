package fun_with_arrays

// findMaxConsecutiveOnes Problem: Max Consecutive Ones.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/
//
// Hint 1:
//	You need to think about two things as far as any window is concerned.
//	One is the starting point for the window. How do you detect that a new window of 1s has started?
//	The next part is detecting the ending point for this window.
//	How do you detect the ending point for an existing window?
//	If you figure these two things out, you will be able to detect the windows of consecutive ones.
//	All that remains afterward is to find the longest such window and return the size.
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	current := 0

	for i := 0; i < len(nums); i++ {
		if current > max {
			max = current
		}

		if nums[i] == 1 {
			current++
			continue
		}

		if nums[i] == 0 {
			current = 0
		}
	}

	if current > max {
		max = current
		current = 0
	}

	return max
}

// findNumbers Problem: Find Numbers with Even Number of Digits.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3237/
//
// Hint 1:
//	How to compute the number of digits of a number ?
// Hint 2:
//	Divide the number by 10 again and again to get the number of digits.
func findNumbers(nums []int) int {
	even := 0
	for _, v := range nums {
		digit := v
		numberOfDigits := 0
		for !(digit < 1) {
			numberOfDigits++
			digit = digit / 10
		}

		if numberOfDigits%2 == 0 {
			even++
		}

	}

	return even
}

// sortedSquares Problem: Squares of a Sorted Array.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3240/
func sortedSquares(nums []int) []int {
	newArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newArr[i] = nums[i] * nums[i]
	}

	if len(newArr) == 1 {
		return newArr
	}

	// Buble sort implementation
	for i := 1; i < len(newArr); i++ {
		f := 0
		for k := 0; k < len(newArr)-i; k++ {
			if newArr[k] > newArr[k+1] {
				tmp := newArr[k+1]
				newArr[k+1] = newArr[k]
				newArr[k] = tmp
				f = 1
			}
		}
		if f == 0 {
			break
		}
	}

	return newArr

}
