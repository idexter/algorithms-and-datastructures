package top_interview_questions_easy

// containsDuplicate Problem: Contains Duplicate.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/578/
func containsDuplicate(nums []int) bool {
	exists := make(map[int]struct{})
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return false
	}

	for _, v := range nums {
		if _, ok := exists[v]; ok {
			return true
		}
		exists[v] = struct{}{}
	}

	return false
}

// removeDuplicates Problem: Remove Duplicates from Sorted Array.
// See: https://github.com/idexter/algorithms-and-datastructures/leetcode/explore/learn/card/fun-with-arrays/deleting_items_from_an_array.go
