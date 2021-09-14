package arrays

// containsDuplicate checks is array contains duplilcates.
// https://leetcode.com/problems/contains-duplicate
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
