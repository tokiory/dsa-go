package linkedList

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head  *Node[T]
	Value T
}

// Iterate goes through all elements in linked list
func (ll *LinkedList[T]) Iterate(callback func(idx int, item *Node[T])) {
	idx := 0
	for current := ll.Head; current != nil; current = current.Next {
		callback(idx, current)
		idx++
	}
}

// Last node of the linked list
func (ll *LinkedList[T]) Last() (last *Node[T]) {
	ll.Iterate(func(_ int, node *Node[T]) {
		last = node
	})
	return
}

func (ll *LinkedList[T]) Append(node Node[T]) {
	last := ll.Last()
	if last != nil {
		last.Next = &node
	} else {
		ll.Head = &node
	}
}

func (ll *LinkedList[T]) Remove() *Node[T] {
	var secondFromEnd *Node[T]
	for current := ll.Head; current.Next != nil; current = current.Next {
		secondFromEnd = current
	}

	lastNode := secondFromEnd.Next
	secondFromEnd.Next = nil
	return lastNode
}

func (ll *LinkedList[T]) Find(value T) (bool, *Node[T]) {
	var foundNode *Node[T]
	for current := ll.Head; current != nil; current = current.Next {
		if current.Value == value {
			foundNode = current
			break
		}
	}

	return foundNode != nil, foundNode
}

func New[T comparable](slice []T) LinkedList[T] {
	ll := LinkedList[T]{}
	var currentNode *Node[T]

	for _, item := range slice {
		if ll.Head == nil {
			n := Node[T]{Value: item}
			ll.Head = &n
			currentNode = ll.Head
		} else {
			n := Node[T]{Value: item}
			currentNode.Next = &n
			currentNode = currentNode.Next
		}
	}
	return ll
}
