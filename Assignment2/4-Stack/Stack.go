package main

import "fmt"

type StackEntry struct {
	Data interface{}
}

type Stack struct {
	stack *[]StackEntry
}

func NewStack() *Stack {
	ss := new(Stack)
	ss.stack = new([]StackEntry)
	return ss
}

func (s Stack) Depth() int {
	ss := s.stack
	return len(*ss)
}

func (s *Stack) Top() interface{} {
	ss := *s.stack
	if len(ss) > 0 {
		// pick up top element
		ste := ss[len(ss)-1]
		// return data field
		return ste.Data
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	ste := StackEntry{data}
	sst := *s.stack
	sst = append(sst, ste)
	s.stack = &sst
}

func (s *Stack) Pop() interface{} {
	data := s.Top()
	sst := *s.stack
	sst = sst[0 : len(sst)-1]
	s.stack = &sst
	return data
}

func main() {

	stack := NewStack()

	stack.Push("First")
	fmt.Println("Top of the stack - ", stack.Top())
	stack.Push("Second")
	fmt.Println("New Top - ", stack.Top())
	fmt.Println("Depth - ", stack.Depth())

	element := stack.Pop()
	fmt.Println("Popped element - ", element, "depth - ", stack.Depth())

	element = stack.Pop()
	fmt.Println("Popped element - ", element, "depth - ", stack.Depth())

}
