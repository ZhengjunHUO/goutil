package datastruct

import (
	"fmt"
	"sync"
)

// LinkedList contains two pointers indicating 
// the linked list's Head and Tail position
type LinkedList[V any] struct {
	Head, Tail *LinkedListNode[V]
	lock  sync.RWMutex
}

// LinkedListNode is the basic element in LinkedList
// contains a Value and a pointer to the Next element 
type LinkedListNode[V any] struct {
        Val  V
        Next *LinkedListNode[V]
}

// NewLinkedList creates a linked list from a slice of any type
func NewLinkedList[V any](slice []V) *LinkedList[V] {
	l := &LinkedList[V]{}

	for i := 0; i < len(slice); i++ {
		l.Push(slice[i])
	}

	return l
}

// Push appends a new element to the linked list's Tail
func (l *LinkedList[V]) Push(v V) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.Tail == nil {
		l.Head = &LinkedListNode[V]{v, nil}
		l.Tail = l.Head
		return
	}

	l.Tail.Next = &LinkedListNode[V]{v, nil}
	l.Tail = l.Tail.Next
}

// ToSlice converts the linked list back to a slice 
// of the same element Value's type
func (l *LinkedList[V]) ToSlice() []V {
	l.lock.RLock()
	defer l.lock.RUnlock()

	var slice []V

	for n := l.Head; n != nil; n = n.Next {
		slice = append(slice, n.Val)
	}

	return slice
}

// PrintAll prints all elements' Value in the linked list, comma separated.
func (l *LinkedList[V]) PrintAll() {
	l.lock.RLock()
	defer l.lock.RUnlock()

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
