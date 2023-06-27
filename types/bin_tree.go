package types

import "fmt"

type LeafNode struct {
	Data  int
	Left  *LeafNode
	Right *LeafNode
}

type BinTree struct {
	Root *LeafNode
}


func NewNode(data int) *LeafNode {
	return &LeafNode{data, nil, nil}
}

func NewTree() *BinTree {
	return &BinTree{nil}
}

func (t *BinTree) Add(l *LeafNode) *LeafNode {
	if t.Root == nil {
		t.Root = l 
		return l
	}

	n, _, ok := t.Find(t.Root, nil, l.Data)
	
	if !ok && n != nil {
		if l.Data < n.Data {
			n.Left = l
		} else {
			n.Right = l
		}
	}

	return l
}

func (t *BinTree) Find(l *LeafNode, p *LeafNode, d int) (*LeafNode, *LeafNode, bool) {
	switch true {
	case d < l.Data:
		if l.Left != nil {
			return t.Find(l.Left, l, d)
		}
	case d < l.Data:
		if l.Right != nil {
			return t.Find(l.Right, l, d)
		}
	case d == l.Data:
		return l, p, true
	case l == nil:
		return nil, p, false
	}

	return l, p, false
}

func (t *BinTree) Show(l *LeafNode) {
	if l == nil {
		return
	}

	t.Show(l.Left)
	fmt.Println(l.Data)
	t.Show(l.Right)
}


