package types

type OneNode struct {
	Data string
	Next *OneNode
}

type OneLinkedList struct {
	Head *OneNode
	Tail *OneNode
}

func NewOneNode(data string) *OneNode {
	return &OneNode{
		Data: data,
		Next: nil,
	}
}

func NewOneLinkedList() *OneLinkedList {
	return &OneLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *OneLinkedList) GetAttr(i int) *OneNode {
	if i < 0 {
		return nil
	}

	node := l.Head
	n := 0

	for node != nil && n != i && node.Next != l.Tail {
		node = node.Next
		n++
	}

	if n == i {
		return node
	}

	return nil
}

func (l *OneLinkedList) PopFront() {
	switch true {
	case l.Head == nil:
		return
	case l.Head == l.Tail:
		l.Head = nil
		l.Tail = nil
	default:
		node := l.Head
		l.Head = node.Next
	}
}

func (l *OneLinkedList) PushBack(data string) {
	node := NewOneNode(data)

	switch true {
	case l.Head == nil:
		l.Head = node
	case l.Tail != nil:
		l.Tail.Next = node
	}

	l.Tail = node
}

func (l *OneLinkedList) PushFront(data string) {
	node := NewOneNode(data)
	node.Next = l.Head
	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}
}

func (l *OneLinkedList) PopBack() {
	switch true {
	case l.Tail == nil:
		return
	case l.Tail == l.Head:
		l.Head = nil
		l.Tail = nil
	}

	node := l.Head

	for ; node != l.Tail; node = node.Next {
		node.Next = nil
		l.Tail = node
	}
}

func (l *OneLinkedList) Insert(i int, data string) {
	left := l.GetAttr(i)

	if left == nil {
		return
	}

	right := left.Next
	node := NewOneNode(data)

	left.Next = node
	node.Next = right

	if right == nil {
		l.Tail = node
	}
}

func (l *OneLinkedList) Erase(i int) {
	switch true {
	case i < 0:
		return
	case i == 0:
		l.PopFront()
	}

	left := l.GetAttr(i - 1)
	node := left.Next

	if node == nil {
		return
	}

	right := node.Next
	left.Next = right

	if node == l.Tail {
		l.Tail = left
	}

}
