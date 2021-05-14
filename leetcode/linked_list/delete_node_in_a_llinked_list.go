package linked_list

// deleteNode removes node from LinkedList
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/553/
func deleteNode(node *ListNode) {
	*node = *node.Next
}
