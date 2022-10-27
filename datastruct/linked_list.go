package datastruct

import (
	"fmt"
	"sync"
)

// LinkedList contains two pointers indicating 
// the linked list's head and tail position
type LinkedList[V any] struct {
	head, tail *LinkedListNode[V]
	lock  sync.RWMutex
}

// LinkedListNode is the basic element in LinkedList
// contains a value and a pointer to the next element 
type LinkedListNode[V any] struct {
        val  V
        next *LinkedListNode[V]
}

// NewLinkedList creates a linked list from a slice of any type
func NewLinkedList[V any](slice []V) *LinkedList[V] {
	l := &LinkedList[V]{}

	for i := 0; i < len(slice); i++ {
		l.Push(slice[i])
	}

	return l
}

// Push appends a new element to the linked list's tail
func (l *LinkedList[V]) Push(v V) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.tail == nil {
		l.head = &LinkedListNode[V]{v, nil}
		l.tail = l.head
		return
	}

	l.tail.next = &LinkedListNode[V]{v, nil}
	l.tail = l.tail.next
}

// ToSlice converts the linked list back to a slice 
// of the same element value's type
func (l *LinkedList[V]) ToSlice() []V {
	l.lock.RLock()
	defer l.lock.RUnlock()

	var slice []V

	for n := l.head; n != nil; n = n.next {
		slice = append(slice, n.val)
	}

	return slice
}

// PrintAll prints all elements' value in the linked list, comma separated.
func (l *LinkedList[V]) PrintAll() {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if l.head != nil {
	        curr := l.head
	        for {
			fmt.Printf("%v", curr.val)

	                if curr.next == nil {
				fmt.Println()
	                        break
	                }
	                curr = curr.next
			fmt.Printf(", ")
	        }
	}
}
