/*
Your code looks well-structured, but there's an issue in the standByForMission function. Let's address it:

In the standByForMission function, you're adding a WaitGroup wg and then immediately calling wg.Done() without waiting for any Goroutines to complete. This will cause wg.Wait() in broadCastStartOfMission to return immediately, which is not the intended behavior. Instead, you should call wg.Done() after the Goroutine completes its execution. Also, you don't need the inner WaitGroup wg as it's not serving any purpose. You can simplify the function.

Here's the corrected standByForMission function:
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var readyForMission bool

func broadCastStartOfMission() {
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)

	standByForMission(func() {
		fmt.Println("Ninja 1 has started the mission ...but he is a bit dim ....will execute slowly !!!!")
		wg.Done()
	}, beeper)

	standByForMission(func() {
		fmt.Println("Ninja 2  has started the mission in a hurry without underwear !!! ...he works a bit too fast")
		wg.Done()
	}, beeper)

	standByForMission(func() {
		fmt.Println("Ninja three has shat his pants on being called for the mission ....but he will still join ....ewwwww!!!!")
		wg.Done()
	}, beeper)

	beeper.Broadcast()

	wg.Wait()
	fmt.Println("All Ninja's somehow have completed the mission ...and the y are still alive ")
}

func standByForMission(fn func(), beeper *sync.Cond) {
	go func() {
		defer beeper.L.Unlock()
		beeper.L.Lock()
		beeper.Wait()
		fn()
	}()
}

func getReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	for !readyForMission {
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work intervals. \n", workIntervals)
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	readyForMission = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}

func main() {
	getReadyForMissionWithCond()
	broadCastStartOfMission()
}
