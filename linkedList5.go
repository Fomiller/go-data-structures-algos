package main

type Node5 struct {
	Value int
	Next  *Node5
}

type LinkedList5 struct {
	Head *Node5
	Size int
}

func (ll *LinkedList5) InsertAtTail(value int) {
	node := &Node5{Value: value}
	if ll.Size == 0 {
		ll.Head = node
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	ll.Size++
}

func (ll *LinkedList5) InsertAtHead(value int) {
	if ll.Size == 0 {
		ll.Head = &Node5{Value: value}
	} else {
		ll.Head = &Node5{Value: value, Next: ll.Head}
	}
	ll.Size++
}

func (ll *LinkedList5) DeleteFromTail(value int) {
	if ll.Size == 0 {
		return
	} else if ll.Size == 1 {
		ll.Head = nil
	} else {
		current := ll.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
	}
	ll.Size--
}

func (ll *LinkedList5) DeleteFromHead(value int) {
	if ll.Size == 0 {
		return
	} else if ll.Size == 1 {
		ll.Head = nil
	} else {
		ll.Head = ll.Head.Next
	}
	ll.Size--
}
