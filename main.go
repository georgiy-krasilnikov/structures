package main

import (
	"fmt"

	"structures/types"
)

func main() {
	// l := types.NewOneLinkedList()

	// l.PushFront("hello")
	// l.PushBack("bye")
	// l.Erase(1)

	// n := l.GetAttr(1)
	// var d string

	// if n != nil {
	// 	d = n.Data
	// } else {
	// 	d = "no"
	// }
	
	l := types.NewTwoLinkedList()

	l.PushBack("hi")
	l.PushBack("bye")
	l.PushBack("bob")

	for n := l.Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	var stack types.Stack

	stack.Push("good")
	stack.Push("morning")

	for len(stack) > 0 {
		s, ok := stack.Pop()
		if ok {
			fmt.Println(s)
		}
	}

	d := []int{10, 5, 7, 16, 13, 2}
	t := types.NewTree()
	for _, v := range(d) {
		t.Add(types.NewNode(v))
	}

	t.Show(t.Root)
}
