package main

type Node4 struct {
	Value int
	Next  *Node4
}

type LinkedList4 struct {
	Head *Node4
	Size int
}

func (ll *LinkedList4) InsertAtTail(value int) {
	newNode := &Node4{Value: value}

	if !ll.IsEmpty() {
		current := ll.Head

		for current.Next != nil {
			current = current.Next
		}

		current.Next = newNode

	} else {
		ll.Head = newNode
	}
	ll.Size++
}

func (ll *LinkedList4) InsertAtHead(value int) {
	newNode := &Node4{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList4) DeleteFromHead() {
	if ll.Head != nil {
		ll.Head = ll.Head.Next
		ll.Size--
	}
}

func (ll *LinkedList4) DeleteFromTail() {
	if ll.Head == nil {
		return
	} else if ll.Head.Next == nil {
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

func (ll *LinkedList4) DeleteFromPostion(idx int) {
	if idx < 0 || idx >= ll.Size {
		return
	}

	if idx == 0 {
		ll.DeleteFromHead()
	} else {
		current := ll.Head
		for i := 0; i < idx-1; i++ {
			current = current.Next
		}

		if current.Next != nil {
			current.Next = current.Next.Next
		}
	}

	ll.Size--
}

func (ll *LinkedList4) Length() int {
	return ll.Size
}

func (ll *LinkedList4) IsEmpty() bool {
	return ll.Size == 0
}
