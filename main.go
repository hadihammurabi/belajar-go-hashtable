package main

import (
	"fmt"

	"github.com/hadihammurabi/belajar-go-hashtable/linkedlist"
)

func main() {
	// fmt.Println("Not implemented, run command bellow:")
	// fmt.Println("go test ./...")

	ll := linkedlist.NewLinkedList()

	size := 5
	for size > 0 {
		ll.Add(&linkedlist.Node{Name: fmt.Sprintf("%d", size), Data: size + 1})
		size--
	}

	ll.Delete("4")
	// fmt.Println(ll.Exists("2"))
	ll.Show()
}
