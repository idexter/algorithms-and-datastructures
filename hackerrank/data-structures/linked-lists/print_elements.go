package linkedlists

import "fmt"

// printLinkedList prints LinkedLists.
// See: https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem
func printLinkedList(head *SinglyLinkedListNode) {
	itm := head
	for {
		fmt.Println(itm.data)
		itm = itm.next
		if itm == nil {
			return
		}
	}
}
