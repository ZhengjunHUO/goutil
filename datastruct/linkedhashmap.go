package datastruct

// Doubly Linked Node
type DoublyLinkedNode struct {
	Key, Val	interface{}
	Prev, Next	*DoublyLinkedNode
}

// NewDoublyLinkedNode initializes a DoublyLinkedNode
func NewDoublyLinkedNode(k, v interface{}) *DoublyLinkedNode {
	return &DoublyLinkedNode{k, v, nil, nil}
}


// Doubly Linked List
type DoublyLinkedList struct {
	Head, Tail	*DoublyLinkedNode
	Len		int
}

// NewDoublyLinkedList initializes a DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	Head, Tail := NewDoublyLinkedNode(nil, nil), NewDoublyLinkedNode(nil, nil)
	Head.Next, Tail.Prev = Tail, Head

	return &DoublyLinkedList{Head, Tail, 0}
}

// Push appendes a node at the list's end
func (dl *DoublyLinkedList) Push(node *DoublyLinkedNode) {
	node.Prev, node.Next = dl.Tail.Prev, dl.Tail
	dl.Tail.Prev.Next, dl.Tail.Prev = node, node
	dl.Len++
}

// Pop moves out the node from the beginning of the list
func (dl *DoublyLinkedList) Pop() *DoublyLinkedNode {
	if dl.Head.Next == dl.Tail {
		return nil
	}

	target := dl.Head.Next
	dl.Delete(target)
	return target
}

// Delete removes the chosen node from the list
func (dl *DoublyLinkedList) Delete(node *DoublyLinkedNode) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	dl.Len--
}

// Linked Hash Map
type LinkedHashmap struct {
	Map	map[interface{}]*DoublyLinkedNode
	List	*DoublyLinkedList
}

// NewLinkedHashmap initializes a linked hashmap
func NewLinkedHashmap() *LinkedHashmap {
	return &LinkedHashmap{make(map[interface{}]*DoublyLinkedNode), NewDoublyLinkedList()}
}

// NewLinkedHashmapFromKV initializes a linked hashmap from a list of key-value tuples
func NewLinkedHashmapFromKV(kvList [][2]interface{}) *LinkedHashmap {
	rslt := NewLinkedHashmap()

	for _, v := range kvList {
		rslt.Put(v[0], v[1])
	}

	return rslt
}

// Put adds a node to the list and registers to the dictionary if not exist yet.
// Update the value of exist one.
func (lm *LinkedHashmap) Put(k, v interface{}) {
	if _, ok := lm.Map[k]; ok {
		lm.Map[k].Val = v
		return
	}

	node := NewDoublyLinkedNode(k, v)
	lm.List.Push(node)
	lm.Map[k] = node
}

// Get returns the node's value if exist.
func (lm *LinkedHashmap) Get(k interface{}) interface{} {
	if _, ok := lm.Map[k]; !ok {
		return nil
	}

	return lm.Map[k].Val
}

// Delete removes the node if exist.
func (lm *LinkedHashmap) Delete(k interface{}) {
	if _, ok := lm.Map[k]; !ok {
		return
	}

	lm.List.Delete(lm.Map[k])
	delete(lm.Map, k)
}

// PopEldest moves out the node from the beginning of the list
func (lm *LinkedHashmap) PopEldest() *DoublyLinkedNode {
	if lm.List.Len == 0 {
		return nil
	}

	node := lm.List.Pop()
	delete(lm.Map, node.Key)
	return node
}

// BecomeNewest moves the choosen node to the end of the list
func (lm *LinkedHashmap) BecomeNewest(k interface{}) {
	if _, ok := lm.Map[k]; !ok {
		return
	}

	node := lm.Map[k]
	lm.List.Delete(node)
	lm.List.Push(node)
}

// Contains checks if the linked hashmap contains a certain value
func (lm *LinkedHashmap) Contains(k interface{}) bool {
	if _, ok := lm.Map[k]; !ok {
		return false
	}

	return true
}

// Size returns the linked hashmap's size
func (lm *LinkedHashmap) Size() int {
	return lm.List.Len
}

// IntoIter returns an iterator to read through the ordered linked map
func (lm *LinkedHashmap) IntoIter() <-chan *DoublyLinkedNode {
	next := make(chan *DoublyLinkedNode)

	go func() {
		defer close(next)
		current := lm.List.Head.Next
		for current.Key != nil {
			next <- current
			current = current.Next
		}
	}()

	return next
}
