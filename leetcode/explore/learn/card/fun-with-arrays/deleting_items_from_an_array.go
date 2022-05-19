package fun_with_arrays

// removeElement Problem: Remove Element.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3247/
//
// Hint 1:
//	The problem statement clearly asks us to modify the array in-place and it also says that the element beyond
//	the new length of the array can be anything. Given an element, we need to remove all the occurrences of
//	it from the array. We don't technically need to remove that element per-say, right?
// Hint 2:
//	We can move all the occurrences of this element to the end of the array. Use two pointers!
// Hint 3:
//	Yet another direction of thought is to consider the elements to be removed as non-existent.
//	In a single pass, if we keep copying the visible elements in-place, that should also solve this problem for us.
func removeElement(nums []int, val int) int {
	initialLength := len(nums)
	endPosition := len(nums) - 1
	removed := 0

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	if len(nums) == 1 && nums[0] != val {
		return 1
	}

	for currentPosition := 0; currentPosition < endPosition+1; currentPosition++ {
		if nums[currentPosition] == val {
			removed++

			nums[currentPosition] = nums[endPosition]
			nums[endPosition] = 0
			endPosition--

			if nums[currentPosition] == val {
				currentPosition--
				continue
			}
		}
	}

	return initialLength - removed
}

// removeDuplicates Problem: Remove Duplicates from Sorted Array.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
//
// Hint 1:
//	In this problem, the key point to focus on is the input array being sorted.
//	As far as duplicate elements are concerned, what is their positioning in the array when the given array is sorted?
//	Look at the image above for the answer. If we know the position of one of the elements,
//	do we also know the positioning of all the duplicate elements?
// Hint 2:
//	We need to modify the array in-place and the size of the final array would potentially be smaller
//	than the size of the input array. So, we ought to use a two-pointer approach here.
//	One, that would keep track of the current element in the original array and another one for just the unique elements.
// Hint 3:
//	Essentially, once an element is encountered, you simply need to **bypass**
//	its duplicates and move on to the next unique element.
func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	writePointer := 1
	for readPointer := 1; readPointer < len(nums); readPointer++ {
		if nums[readPointer] != nums[readPointer-1] {
			nums[writePointer] = nums[readPointer]
			writePointer++
		}
	}

	return writePointer
}
