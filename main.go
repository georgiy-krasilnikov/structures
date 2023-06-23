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
}
