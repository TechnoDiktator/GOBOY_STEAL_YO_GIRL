package main

import (
	"fmt"
	"time"
	"math/rand"
)

//SIGNAL AND BROADCAST


// Cond package 
//Cond implements a condition variable  , a rendevouz point
//for goroutines waiting for or announcing the occurance of an event



//Each Cond has an associatesd Locker L (often a *Mutex or a *RWMutex) 
//which must be held when changing the condition and 
//when callling thewait method

// A Cond must not be copied after first use


var ready bool
func main() {
	gettingreadyForMisiion()
}

func gettingreadyForMisiion() {
	go gettingReady()
	workIntervals := 0

	/*
	now in the for loop below
	since we have made the main thread sleep for 5 seconds 
	in that time the go routine will finish 
	and ready will be set to true 
	thus we will increase the workInterval variable very few times 
	//maybe only once

	//BUT if we commented the Sleep statement in the loop below
	ALL HELL BERAKS LOOSE IN THE MAIN THREAD
	if we didnt sleep for the 5 seconds ...the work interval count would hhave been very large
	That is we would be basically be running the loop millions of times before the
	ready variable is set to true by the go routine

	//BUT HERE IS THE PROBLEM
	even if we sleep the main thread for 5 seconds .....(with the expectation)
	that the goroutine will definitely finish in 5 sec 
	// it is kind of an overrach
	because what is the go routine only takes 2 sec
	or what is it only takes FUCKING 50 sec
	so how the fuck do we decide for how much time we shall waiit
	//This can in a way be acheived by using channels
	//BUT AN EVEN BETTER WAY IS THERE


	COND  ----> Discussed in the next lesson 22

	*/
	for !ready {
		time.Sleep(5*time.Second) 
		workIntervals++ 
	}

	fmt.Printf("We are now ready! After %d work intervals. \n" , workIntervals)

}


func gettingReady() {
	sleep() // this will make the main theread sleep
	ready  = true
}


func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5))*time.Second
	time.Sleep(someTime)
}





