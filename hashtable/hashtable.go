package hashtable

import (
	"github.com/hadihammurabi/belajar-go-hashtable/linkedlist"
)

// const
const (
	TableSize int = 25
)

// HashTable structure
type HashTable struct {
	container [TableSize]*linkedlist.LinkedList
}

// NewHashTable func
func NewHashTable() *HashTable {
	table := &HashTable{}
	for i := range table.container {
		table.container[i] = linkedlist.NewLinkedList()
	}

	return table
}

func hash(name string) int {
	hashResult := 0
	for _, c := range name {
		hashResult += int(c)
	}
	return hashResult
}

func index(hashValue int) int {
	return hashValue % TableSize
}

// Add func
func (h *HashTable) Add(node *linkedlist.Node) {
	idx := index(hash(node.Name))
	h.container[idx].Add(node)
}

// Get func
func (h *HashTable) Get(name string) interface{} {
	idx := index(hash(name))
	container := h.container[idx]
	position := uint(container.GetPosition(name))
	node := container.GetAt(position)

	if node != nil {
		return node.Data
	}

	return nil
}

// Exists func
func (h HashTable) Exists(name string) bool {
	idx := index(hash(name))
	return h.container[idx].Exists(name)
}

// Delete func
func (h HashTable) Delete(name string) bool {
	idx := index(hash(name))
	if h.Exists(name) {
		h.container[idx].Delete(name)
		return true
	}

	return false
}

// Show func
func (h HashTable) Show() {
	for i := range h.container {
		h.container[i].Show()
	}
}
