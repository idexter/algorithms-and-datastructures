package linked_list

// reverseList return reversed linked list.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/560/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	// Go through whole list
	for head.Next != nil {
		next := head.Next // Save next
		head.Next = prev  // Set next to previous
		prev = head       // set new previous to current
		head = next       // set head to next
	}
	head.Next = prev // for last case set next to previous
	return head
}
