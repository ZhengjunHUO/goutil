package datastruct

import (
	"fmt"
)

// LinkedList contains two pointers indicating 
// the linked list's head and tail position
type LinkedList[V any] struct {
	Head, Tail *LinkedListNode[V]
}

// LinkedListNode is the basic element in LinkedList
// contains a value and a pointer to the next element 
type LinkedListNode[V any] struct {
        Val  V
        Next *LinkedListNode[V]
}

// NewLinkedList creates a linked list from a slice of any type
func NewLinkedList[V any](slice []V) *LinkedList[V] {
	l := &LinkedList[V]{}

	n := len(slice)
	/*
	if n == 0 {
		return l
	}
	*/

	for i := 0; i < n; i++ {
		l.Push(slice[i])
	}

	return l
}

// Push appends a new element to the linked list's tail
func (l *LinkedList[V]) Push(v V) {
	if l.Tail == nil {
		l.Head = &LinkedListNode[V]{v, nil}
		l.Tail = l.Head
		return
	}

	l.Tail.Next = &LinkedListNode[V]{v, nil}
	l.Tail = l.Tail.Next
}

// ToSlice converts the linked list back to a slice 
// of the same element value's type
func (l *LinkedList[V]) ToSlice() []V {
	var slice []V

	for n := l.Head; n != nil; n = n.Next {
		slice = append(slice, n.Val)
	}

	return slice
}

// PrintAll prints all elements' value in the linked list, comma separated.
func (l *LinkedList[V]) PrintAll() {
	if l.Head != nil {
	        curr := l.Head
	        for {
			fmt.Printf("%v", curr.Val)

	                if curr.Next == nil {
				fmt.Println()
	                        break
	                }
	                curr = curr.Next
			fmt.Printf(", ")
	        }
	}
}
