package linkedlists

// SinglyLinkedListNode describes LinkedList element.
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// DoublyLinkedListNode describes doubly linked list node.
type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}
