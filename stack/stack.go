package stack

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	Top  *Node[T]
	Size int
	Tail *Node[T]
}

// Append Stack is a basic LIFO stack that resizes as needed.
// The complexity of this implementation is O(1)
func (s *Stack[T]) Append(value T) {
	if s.Top == nil {
		s.Top = &Node[T]{Value: value}
		s.Tail = s.Top
	} else {
		s.Tail.Next = &Node[T]{Value: value}
		s.Tail = s.Tail.Next
	}
}

func (s *Stack[T]) Pop() *Node[T] {
	if s.Top == nil {
		return nil
	}
	node := s.Top
	s.Top = s.Top.Next
	return node
}

func (n *Stack[T]) Peek() *Node[T] {
	return n.Top
}
