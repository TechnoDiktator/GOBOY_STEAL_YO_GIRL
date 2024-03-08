package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)


//COND


//SIGNAL AND BROADCAST
/*
In Go's sync package, the sync.Cond type represents a condition variable, 
which is a synchronization primitive used to coordinate the execution of Goroutines
based on some condition.

When you call cond.Signal() or cond.Broadcast() on a sync.Cond variable, 
you're signaling waiting Goroutines that the condition they are waiting for may have been met. 
Here's a breakdown:


SIGNAL
cond.Signal(): This method wakes up one Goroutine waiting on the condition variable cond. 
If there are multiple Goroutines waiting, it's not specified which one will be woken up. Typically, you use Signal() when you know that only one Goroutine needs to be woken up.


BROADCAST
cond.Broadcast(): This method wakes up all Goroutines waiting on the condition variable cond. 
This is useful when multiple Goroutines may need to be woken up because the condition has been met.

*/
// Cond package 
//Cond implements a condition variable  , a rendevouz point
//for goroutines waiting for or announcing the occurance of an event



//Each Cond has an associatesd Locker L (often a *Mutex or a *RWMutex) 
//which must be held when changing the condition and 
//when callling thewait method

// A Cond must not be copied after first use


// THSI IS WHAT A COND LOOKS LIKE
// type Cond struct{
// 	noCopy noCopy

// 	//l is held while observing or changing the condition
// 	L Locker

// 	notify notifyList
	
// 	checker copyChecker
// }

// func NewCond(l Locker) *Cond{
// 	return &Cond{L:l}

//}




//CODE BELOW


/*
Yes, in the provided code, cond.Wait() is blocking the execution of the gettingreadyForMisiionWithCond() function until the condition variable ready becomes true.

Let's break down the flow:

The main Goroutine calls gettingreadyForMisiionWithCond().

Inside gettingreadyForMisiionWithCond(), a new sync.Cond variable cond is created with a mutex.

A Goroutine is launched with go gettingReadyWithCond(cond).

The main Goroutine enters a loop where it waits until the ready variable becomes true.

Within gettingReadyWithCond(), after a random sleep period, the ready variable is set to true and cond.Signal() 
is called to wake up one Goroutine waiting on the condition variable.

Once woken up, the main Goroutine checks if ready is true and then prints the message indicating that the mission is ready.

So yes, cond.Wait() is indeed blocking the main Goroutine until the ready condition is met and it receives a signal from another Goroutine. 
Once the condition is met and the signal is received, the main Goroutine continues its execution.


*/



var ready bool

func main() {
	gettingreadyForMisiionWithCond()
}


func gettingreadyForMisiionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0
	cond.L.Lock()
	for !ready {
		workIntervals++ 
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("We are now ready! After %d work intervals. \n" , workIntervals)

}


func gettingReadyWithCond (cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()

}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5))*time.Second
	time.Sleep(someTime)
}

// func gettingreadyForMisiion() {
// 	go gettingReady()
// 	workIntervals := 0
// 	for !ready {
// 		workIntervals++ 
// 	}

// 	fmt.Printf("We are now ready! After %d work intervals. \n" , workIntervals)

// }


// func gettingReady() {
// 	sleep() // this will make the main theread sleep
// 	ready  = true
// }








