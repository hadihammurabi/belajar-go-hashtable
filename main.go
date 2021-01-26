package main

import (
	"fmt"
	"time"

	"github.com/hadihammurabi/belajar-go-hashtable/linkedlist"

	"github.com/hadihammurabi/belajar-go-hashtable/hashtable"
)

func main() {
	table := hashtable.NewHashTable()
	count := 1000
	for i := 0; i < count; i++ {
		table.Add(&linkedlist.Node{Name: fmt.Sprint(i), Data: i})
	}

	start := time.Now()
	fmt.Println(table.Exists("78"))
	fmt.Println(table.Exists("68"))
	fmt.Println(table.Exists("4778"))
	fmt.Println(table.Exists("89"))
	fmt.Println(table.Exists("2929"))

	fmt.Printf("%s \n\n", time.Since(start))
}
