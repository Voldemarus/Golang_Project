package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	NumberOfRandomValues int = 16 // total amount of random values
	ChunkSize            int = 4  // size of chunk, generated on each step
)

// Create top level wait group to join chunks
var topWg sync.WaitGroup

func main() {

	// Create array to store random data.
	randomData := make([]int, 0, NumberOfRandomValues)

	numThreads := runtime.NumCPU()
	runtime.GOMAXPROCS(numThreads)
	fmt.Println("Max amount of threads = ", numThreads)
	// Option 1. Create single channel with size enough to fit all generated data
	ch := make(chan int, ChunkSize*numThreads)

	// Spawn goroutines to generate numbers
	for i := 0; i < NumberOfRandomValues/ChunkSize; i++ {
		topWg.Add(2)
		// Prepare channel to supply main function with data
		// Option 2. Create separate channel for each pair of routines
		//		ch := make(chan int, ChunkSize)
		time.Sleep(500 * time.Millisecond)
		go makeRandomNumbers(ch) // spawn random generaotr
		go receiveRandomData(i, ch, &randomData)
	}

	topWg.Wait() // join
}

// Generate chunk of random numbers and put them into channel
func makeRandomNumbers(ch chan int) {
	source := rand.NewSource(time.Now().UnixNano()) // set up random generator seed
	generator := rand.New(source)                   // run random generator
	for i := 0; i < ChunkSize; i++ {
		ch <- generator.Intn(100) //  pick and scale to range [0..99]
	}
	topWg.Done() // release mutex
}

// Receive data from the channel and return as a slice of int
func receiveRandomData(index int, ch chan int, data *[]int) {
	for i := 0; i < ChunkSize; i++ {
		*data = append(*data, (<-ch))
	}
	fmt.Println("Genertor: ", index, " data = ", data)
	topWg.Done()
}
