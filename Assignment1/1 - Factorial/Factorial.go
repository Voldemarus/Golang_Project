package main

import (
	"fmt"
	"sync"
)

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
	channel := make(chan int, 1)
	topWg := new(sync.WaitGroup)
	topWg.Add(1)
	go factorial1(topWg, 1, m, channel)
	topWg.Wait()
	result := <-channel
	return result
}

// Debug version = single level goroutine
func factorial2(parentWG *sync.WaitGroup, from int, to int, channel chan int) {
	_, v := factorial(to)
	fmt.Println("fact2 ", v)
	channel <- v
	parentWG.Done()
}

// multithreaded version
func factorial1(parentWG *sync.WaitGroup, from int, to int, channel chan int) {
	defer parentWG.Done()
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
		localWG := new(sync.WaitGroup)
		localWG.Add(2)
		leftChan := make(chan int, 1)
		rightChan := make(chan int, 1)
		go factorial1(localWG, from, middle, leftChan)
		go factorial1(localWG, middle2, to, rightChan)
		localWG.Wait()
		left := <-leftChan
		right := <-rightChan
		result := left * right
		channel <- result
		// fmt.Println("from = ", from, " to = ", to, "result = ", result)
	}
}
