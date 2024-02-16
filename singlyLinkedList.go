package main

type Node struct {
	Value string
	Next  *Node
}
type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func (s *SinglyLinkedList) Append(value string) {
	newNode := &Node{Value: value}
	if s.Head == nil {
		s.Head = newNode
		s.Tail = newNode
		return
	}
	s.Tail.Next = newNode
	s.Tail = newNode
}
