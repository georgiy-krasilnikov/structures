package main

import (
	"fmt"

	"structures/one_list"
)

func main() {
	l := one_list.NewOneLinkedList()
	l.PushFront("hello")
	l.PushBack("bye")
	l.Erase(1)
	n := l.GetAttr(1)
	var d string

	if n != nil {
		d = n.Data
	} else {
		d = "no"
	}
	fmt.Print(d)
}
