package main

import (
	"fmt"
	"sync"
	"time"
)

var topWg sync.WaitGroup

func main() {

	s, g := factorial(7)
	fmt.Println("Factorial ", s, "= ", g)

	result := factorialMT(7)
	fmt.Println("Factorial multithreaded", result)

}

//
//  "Classic" factorial implementation. Simple recursion
//
func factorial(m int) (int, int) {
	result := 1
	if m > 1 {
		_, result := factorial(m - 1)
		return m, result * m
	}
	return m, result
}

//
//  Multithreade Factorial caluclation
//

// Wrapper for multithreaded recursion function
func factorialMT(m int) int {
	channel := make(chan int)
	topWg.Add(1)
	go factorial1(1, m, channel)
	time.Sleep(time.Second * 3)

	topWg.Wait()
	return <-channel
}

func factorial1(from int, to int, channel chan int) {

	if from >= to {
		channel <- from
	} else if (to - from) == 1 {
		channel <- from * to
	} else {
		middle := (to + from) / 2
		middle2 := middle + 1
		if middle2 == from {
			middle2++
		}
		//	var left, right int

		topWg.Add(1)
		chanLeft := make(chan int, 1)
		go factorial1(from, middle, chanLeft)
		chanRight := make(chan int, 1)
		topWg.Add(1)
		go factorial1(middle2, to, chanRight)
		time.Sleep(time.Millisecond * 600)
		left := <-chanLeft
		right := <-chanRight
		res := left * right
		channel <- res
		fmt.Println("from = ", from, " to = ", to, "result = ", res)
	}
	topWg.Done()

}
