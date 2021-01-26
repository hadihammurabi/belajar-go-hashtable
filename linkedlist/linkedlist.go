package linkedlist

import (
	"fmt"
)

// LinkedList structure
type LinkedList struct {
	Head   *Node
	length uint
}

// NewLinkedList create a new linked list container
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add func
func (l *LinkedList) Add(newNode *Node) {
	if l.length == 0 {
		l.Head = newNode
		l.length++
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.length++
}

// Exists func
func (l LinkedList) Exists(name string) bool {
	size := l.Size()
	for size > 0 {
		if l.Head.Name == name {
			return true
		}

		l.Head = l.Head.Next
		size--
	}

	return false
}

// Delete func
func (l *LinkedList) Delete(name string) bool {
	if !l.Exists(name) {
		return false
	}

	position := uint(l.GetPosition(name))
	if position <= 0 {
		l.Head = l.GetAt(0).Next
	} else {
		l.GetAt(position - 1).Next = l.GetAt(position).Next
	}

	l.length--
	return false
}

// GetPosition func
func (l LinkedList) GetPosition(name string) int {
	for i := 0; i < int(l.length); i++ {
		if l.Head.Name == name {
			return i
		}

		l.Head = l.Head.Next
	}

	return -1
}

// GetAt func
func (l LinkedList) GetAt(position uint) *Node {
	ptr := l.Head
	if position < 0 {
		return ptr
	}
	if position > (l.length - 1) {
		return nil
	}

	for i := 0; i < int(position); i++ {
		ptr = ptr.Next
	}

	return ptr
}

// Size func
func (l *LinkedList) Size() uint {
	return l.length
}

// Show func
func (l LinkedList) Show() {
	for l.Head != nil {
		fmt.Println(l.Head)
		l.Head = l.Head.Next
	}
}
