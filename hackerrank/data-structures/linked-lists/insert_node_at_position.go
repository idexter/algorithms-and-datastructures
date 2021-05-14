package linkedlists

// insertNodeAtPosition inserts node at position.
// See: https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem
func insertNodeAtPosition(list *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	first := list

	if position == 0 {
		node := &SinglyLinkedListNode{
			data: data,
			next: list,
		}
		return node
	}

	var at int32 = 0
	for list.next != nil {
		if position == at+1 {
			node := &SinglyLinkedListNode{
				data: data,
				next: list.next,
			}
			list.next = node
			break
		}
		list = list.next
		at++
	}

	return first

}
