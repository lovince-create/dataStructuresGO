package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (h *Node) printLinkedList() {
	for h.next != nil {
		fmt.Printf("%d ", h.next.value)
		h = h.next
	}
	fmt.Println()
}

func (h *Node) add(val int) {
	for h.next != nil {
		h = h.next
	}
	h.next = new(Node)
	h.next.value = val
}

func (h *Node) getIndex(val int) int {
	i := 0
	for h.next != nil {
		if h.next.value == val {
			break
		}
		i++
		h = h.next
	}
	return i
}

func (h *Node) getValue(index int) int {
	i := 0
	for h.next != nil {
		if i == index {
			break
		}
		i++
		h = h.next
	}
	return h.next.value
}

func (h *Node) removeIndex(index int) {
	for i := 0; i < index; i++ {
		h = h.next
	}
	h.next = h.next.next
}

func (h *Node) remove(val int) {
	for h.next != nil {
		if h.next.value == val {
			break
		}
		h = h.next
	}
	h.next = h.next.next
}

func main() {
	ll := new(Node)
	ll.add(3)
	ll.add(2)
	ll.add(1)
	ll.printLinkedList()
}
