package main

import (
	"fmt"
	"time"
)

type Operation1 struct {
	worker int
	job    int
	data   int
}

type Operation2 struct {
	worker int
	job    int
	data   string
}

const (
	kNumJobsOp1 int = 10 // amount of jobs for operation 1
	kNumJobsOp2 int = 9  // amount of jobs for operation 2
	kWorkers1   int = 4  // amount of workers in Pool 1
	kWorkers2   int = 6  // amount of workers in Pool 2
)

//
//  Worker 1. Performs first operation
//
func worker1(workerNum int, jobs <-chan int, results chan<- Operation1) {
	for j := range jobs {
		fmt.Println("Worker ", workerNum, "Operation 1 started as job #", j)
		time.Sleep(time.Millisecond * 500) // hard work imitation
		fmt.Println("Worker ", workerNum, "Operation 1 finished as job #", j)
		// send result on iteration
		result := 100*workerNum + j
		results <- Operation1{workerNum, j, result}
	}
}

func worker2(workerNum int, jobs <-chan int, results chan<- Operation2) {
	for j := range jobs {
		fmt.Println("Worker ", workerNum, "Operation 2 started as job #", j)
		time.Sleep(time.Millisecond * 300) // hard work imitation
		fmt.Println("Worker ", workerNum, "Operation 2 finished as job #", j)
		// send result on iteration
		result := fmt.Sprintf("Worker #%d :: job : %d result = %d", workerNum, j, j*j)
		results <- Operation2{workerNum, j, result}
	}
}

func main() {

	channelJobs := make(chan int, kNumJobsOp1)       // channel for requests 1
	channelOp1 := make(chan Operation1, kNumJobsOp1) // channel to get results from worker1
	channelJobs2 := make(chan int, kNumJobsOp2)      // channel for requests to worker 2
	channelOp2 := make(chan Operation2, kNumJobsOp2) // channels to get results from worker 2

	// spawn workers
	for w1 := 1; w1 <= kWorkers1; w1++ {
		go worker1(w1, channelJobs, channelOp1)
	}
	for w2 := 1; w2 <= kWorkers2; w2++ {
		go worker2(w2, channelJobs2, channelOp2)
	}
	// All workers are waiting for incoming requests

	// Start requests, which will activate workers
	// Activate all operations at once
	for j := 1; j <= kNumJobsOp1; j++ {
		channelJobs <- j
	}
	for j := 1; j <= kNumJobsOp2; j++ {
		channelJobs2 <- j
	}
	// and close channels with requests
	close(channelJobs)
	close(channelJobs2)

	// now get results
	arrivedData := 0
	for arrivedData < (kNumJobsOp2 + kNumJobsOp1) {
		select {
		case res := <-channelOp1:
			fmt.Println("Operation 1 - ", res)
			arrivedData++
		case res2 := <-channelOp2:
			fmt.Println("Operation 2 - ", res2)
			arrivedData++
		}
	}

}
