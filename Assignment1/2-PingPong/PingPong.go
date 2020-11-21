//
// PingPong.go
//
//

package main

import (
	"fmt"
	"sync"
	"time"
)

const FinalNum int = 10 // Amount of strikes in the set

// define two channels, which are used to send ball from left and from right players
var leftStrike = make(chan int, 5)
var rightStrike = make(chan int, 5)

// Create wait group to join threads after predefined amount of strijes is passed
var wg sync.WaitGroup

func player(name string, acceptStrikeCh <-chan int, sendStrikeCh chan<- int) {
	for {
		strikeCount := <-acceptStrikeCh // accept integere from typed channel
		strikeCount++                   // increment data
		if strikeCount > FinalNum {
			if strikeCount == FinalNum+1 {
				fmt.Println(name, " makes final strike and won!")
				// Now we should stop appliction
				// 1. Send strike to another player
				sendStrikeCh <- strikeCount
			}
			wg.Done()
			//
			break
		} else {
			fmt.Println("Player ", name, " strike #", strikeCount)
			sendStrikeCh <- strikeCount
		}
		time.Sleep(time.Second / 2.0)
	}
}

func main() {
	fmt.Println("Ping pong started.")
	// init players as async routines
	wg.Add(2)
	go player("Tom", leftStrike, rightStrike)
	go player("Jerry", rightStrike, leftStrike)
	// put ball into the game - send it to right player
	leftStrike <- 0

	// main thread is being put into sleep to allow background threads
	// to play ping pong
	wg.Wait()

	fmt.Println("PingPong is finished.")
}
