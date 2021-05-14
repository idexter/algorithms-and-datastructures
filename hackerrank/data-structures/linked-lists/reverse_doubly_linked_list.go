package linkedlists

// reverse reverses DoublyLinkedListNode.
// See: https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list/problem
func reverse(head *DoublyLinkedListNode) *DoublyLinkedListNode {

	if head.next == nil {
		return head
	}

	for {
		if head.prev == nil {
			head.prev = head.next
			head.next = nil
			head = head.prev
			continue
		}

		if head.next == nil {
			head.next = head.prev
			head.prev = nil
			break
		}

		head.next, head.prev = head.prev, head.next
		head = head.prev
	}
	return head
}
