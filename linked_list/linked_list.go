package main

import "fmt"

type Node struct {
	element interface{}
	next    *Node
}

type LinkedLists struct {
	head *Node
}

func NewLinkedLists() *LinkedLists {
	return &LinkedLists{head: &Node{
		element: "head",
		next:    nil,
	}}
}

func (l *LinkedLists) Find(element interface{}) *Node {
	current := l.head

	for current.element != element {
		current = current.next
	}

	return current
}

func (l *LinkedLists) Insert(newElement interface{}, element interface{}) {
	newNode := &Node{element: newElement}
	current := l.Find(element)

	if current == nil {
		fmt.Println("Element not found in the list")
		return
	}

	newNode.next = current.next
	current.next = newNode
}

func (l *LinkedLists) Print() {
	current := l.head

	for {
		fmt.Printf("%v", current.element)

		if current.next == nil {
			fmt.Println()
			break
		}

		current = current.next

		fmt.Println(" -> ")
	}
}

func (l *LinkedLists) FindPrevious(element interface{}) *Node {
	current := l.head
	for current.next != nil && current.next.element != element {
		current = current.next
	}
	return current
}

func (l *LinkedLists) Remove(element interface{}) {
	prevNode := l.FindPrevious(element)
	if prevNode.next != nil {
		current := prevNode.next
		prevNode.next = current.next
		current = nil
	}
}

func main() {
	l := NewLinkedLists()
	l.Insert("Hien", "head")
	l.Insert("Pham", "Hien")
	l.Insert("Minh", "Pham")
	l.Print()
	l.Remove("Pham")
	l.Print()
}
