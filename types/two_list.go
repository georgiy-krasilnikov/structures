package types

type TwoNode struct {
	Data string
	Prev *TwoNode
	Next *TwoNode
}

type TwoLinkedList struct {
	Head *TwoNode
	Tail *TwoNode
}

func NewTwoNode(data string) *TwoNode {
	return &TwoNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}

func NewTwoLinkedList() *TwoLinkedList {
	return &TwoLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *TwoLinkedList) GetAttr(i int) *TwoNode {
	if i < 0 {
		return nil
	}

	node := l.Head
	n := 0

	for n != i {
		if node == nil {
			return node
		}
		node = node.Next
		n++
	}

	if n == i {
		return node
	}

	return nil
}

func (l *TwoLinkedList) PopFront() {
	if l.Head == nil {
		return
	}

	node := l.Head.Next

	switch true {
	case node != nil:
		node.Prev = nil
	default:
		l.Tail = nil
	}

	l.Head = node
}

func (l *TwoLinkedList) PushBack(data string) *TwoNode {
	node := NewTwoNode(data)
	node.Prev = l.Tail

	switch true {
	case l.Tail != nil:
		l.Tail.Next = node
	case l.Head == nil:
		l.Head = node
	}

	l.Tail = node

	return node
}

func (l *TwoLinkedList) PushFront(data string) *TwoNode {
	node := NewTwoNode(data)
	node.Next = l.Head

	switch true {
	case l.Head != nil:
		l.Head.Prev = node
	case l.Tail == nil:
		l.Tail = node
	}

	l.Head = l.Tail

	return node
}

func (l *TwoLinkedList) PopBack() {
	if l.Tail == nil {
		return
	}

	node := l.Tail.Prev

	switch true {
	case node != nil:
		node.Next = nil
	default:
		l.Head = nil
	}

	l.Tail = node
}

func (l *TwoLinkedList) Insert(i int, data string) *TwoNode {
	right := l.GetAttr(i)

	if right == nil {
		return l.PushBack(data)
	}

	left := right.Prev
	if left == nil {
		return l.PushFront(data)
	}
	node := NewTwoNode(data)
	node.Prev = left
	node.Next = right
	left.Next = node
	right.Prev = node

	return node

}

func (l *TwoLinkedList) Erase(i int) {
	node := l.GetAttr(i)
	if node == nil {
		return
	}

	switch true {
	case node.Prev == nil:
		l.PopFront()
		return
	case node.Next == nil:
		l.PopBack()
		return
	}

	left := node.Prev
	right := node.Next
	left.Next = right
	right.Prev = left
}
