package main

import "fmt"

type QueueEntry struct {
	Data interface{}
}

type Queue struct {
	queue *[]QueueEntry
}

func NewQueue() *Queue {
	ss := new(Queue)
	ss.queue = new([]QueueEntry)
	return ss
}

func (s Queue) Depth() int {
	ss := s.queue
	return len(*ss)
}

func (s *Queue) Head() interface{} {
	ss := *s.queue
	if len(ss) > 0 {
		// pick up top element
		ste := ss[0]
		// return data field
		return ste.Data
	}
	return nil
}

func (s *Queue) In(data interface{}) {
	ste := QueueEntry{data}
	sst := *s.queue
	sst = append(sst, ste)
	s.queue = &sst
}

func (s *Queue) Out() interface{} {
	data := s.Head()
	sst := *s.queue
	sst = sst[1:len(sst)]
	s.queue = &sst
	return data
}

func (s Queue) printQueue() {
	sst := *s.queue
	for i, v := range sst {
		fmt.Printf("%2d: %+v\n", i, v.Data)
	}
}

func main() {

	queue := NewQueue()

	queue.In("First")
	fmt.Println("Head of the Queue - ", queue.Head())
	queue.In("Second")
	queue.In("Third")

	arr := []int{1, 2, 3, 4, 5}
	queue.In(arr)

	fmt.Println("Queue head - ", queue.Head())
	fmt.Println("Depth - ", queue.Depth())

	fmt.Println("Filled Queue:")
	queue.printQueue()

	element := queue.Out()
	fmt.Println("Extracted element - ", element, "depth - ", queue.Depth())

	element = queue.Out()
	fmt.Println("Extracted element - ", element, "depth - ", queue.Depth())

	fmt.Println("Remained queue:")
	queue.printQueue()

}
