package linked_list

// hasCycle checks if list has cycle.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/773/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	existsMap := make(map[*ListNode]struct{})
	for head.Next != nil {
		if _, ok := existsMap[head]; ok {
			return true
		}
		existsMap[head] = struct{}{}
		head = head.Next
	}
	return false
}
