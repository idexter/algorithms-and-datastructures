package linked_list

// removeNthFromEnd removes Nth node from end of Linked List.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/603/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && n == 1 {
		return nil
	}

	last := head
	elementToDelete := head
	beforeLast := head
	k := 0
	for {
		if k >= n {
			elementToDelete = elementToDelete.Next
		}
		if last.Next != nil && last.Next.Next == nil {
			beforeLast = last
		}
		if last.Next == nil {
			break
		}
		last = last.Next
		k++
	}

	if elementToDelete == head {
		return elementToDelete.Next
	}

	if n == 1 && k == 1 {
		head.Next = nil
		return head
	}

	if elementToDelete == last {
		beforeLast.Next = nil
		return head
	}

	*elementToDelete = *elementToDelete.Next
	return head
}
