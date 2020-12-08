package main

import "fmt"

type Node struct {
	Data interface{} // any Go' object as value of the node
	Next *Node       // link to the next node
	Prev *Node       // link to the previous node
}

//
// List can contain arbitrary data (belongs to any GoLang type)
//
type List struct {
	Head *Node
}

// insert new node to the start of the list
func (l *List) Insert(data interface{}) {
	newNode := new(Node)
	newNode.Data = data
	newNode.Next = l.Head
	l.Head = newNode
}

// Create new list with single element
func newList(data interface{}) *List {
	newL := new(List)
	newL.Insert(data)
	return newL
}

// Returns last element of the list
func (l List) Last() interface{} {
	n := l.Head
	for n.Next != nil {
		n = n.Next
	}
	return n.Data
}

func (l *List) append(data interface{}) {
	n := l.Head
	for n.Next != nil {
		n = n.Next
	}
	newNode := new(Node)
	newNode.Data = data
	n.Next = newNode

}

// Returns first element of the list
func (l List) First() interface{} {
	n := l.Head
	return n.Data
}

// Counts amount of elements - simple implementation without
// any caching
func (l List) Count() int {
	n := l.Head
	cnt := 0
	for n != nil {
		cnt++
		n = n.Next
	}
	return cnt
}

// Returns tail of the list - List without first element
func (l *List) tail() *List {
	newL := new(List)
	node := l.Head.Next
	for node != nil {
		newL.Insert(node.Data)
		node = node.Next
	}
	ll := newL.reverse()
	return ll
}

// Reverse order od elements in the list
func (l *List) reverse() *List {
	node := l.Head
	newList := new(List)
	for node != nil {
		newList.Insert(node.Data)
		node = node.Next
	}
	return newList
}

func (L List) PrintAll() {
	node := L.Head
	for node != nil {
		m := node.Data
		fmt.Printf("%v", m)
		if node.Next != nil {
			fmt.Printf(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

func main() {
	list := newList("Start")
	list.Insert(22)
	list.Insert("Second")
	list.Insert(44)
	list.Insert("float:")
	list.Insert(66.33)

	fmt.Println("Initial list")
	list.PrintAll()
	fmt.Println()
	fmt.Println("List properties:")
	fmt.Println("First element:: ", list.First())
	fmt.Println("Last element:: ", list.Last())
	fmt.Println("Count of elements:: ", list.Count())
	fmt.Println()

	reversedList := list.reverse()
	fmt.Println("Reversed list:")
	reversedList.PrintAll()
	fmt.Println()

	fmt.Println("Tail of the reversed list:")
	tail := reversedList.tail()
	tail.PrintAll()
	fmt.Println()

	fmt.Println("Append array  to the end of reversedList::")
	arr := []int{1, 2, 3, 4, 5}
	reversedList.append(arr)
	reversedList.PrintAll()
	fmt.Println()

}
