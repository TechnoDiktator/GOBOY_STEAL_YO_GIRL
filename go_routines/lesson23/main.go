package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"

)
var readyformission bool

func broadCastStartOfMission() {
	
	// The newcond method returns a pointer to the Cond object 
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)

	standByForMission(func() {
		fmt.Println("Ninja 1 has started the mission ...but he is a bit dim ....will execute slowly !!!!")
		wg.Done()
	} , beeper)

	standByForMission(func() {
		fmt.Println("Ninja 2  has started the mission in a hurry without underwear !!! ...he works a bit too fast")
		wg.Done()
		} , beeper)

	standByForMission(func(){
		fmt.Println("Ninja three has shat his pants on being called for the mission ....but he will still join ....ewwwww!!!!")
		wg.Done()
	} , beeper)
	
	beeper.Broadcast()

	wg.Wait()
	fmt.Println("All Ninja's somehow have completed the mission ...and the y are still alive ")


}


func standByForMission(fn func() , beeper *sync.Cond){
	var wg sync.WaitGroup  // actually i dont think we need this wait group 
	wg.Add(1)
	go func() {
		wg.Done()  
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()    //because we are waiting for the 
		fn() //because this function contains the waitgroup oif the main thread ...and this is the one that makes the main thread wait
	}()
	wg.Wait()

}


func getReadyForMissionWithCond(){

	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	for !readyformission {
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work intervals. \n" , workIntervals)


}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	readyformission = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5))*time.Second
	time.Sleep(someTime)
}



func main() {

	getReadyForMissionWithCond()
	broadCastStartOfMission()

}


//BREAKDOWN

/*

Sure, let's break down the code you provided:

Imports: You're importing necessary packages including fmt, sync, time, and math/rand.

Global Variables:

readyForMission: A boolean variable indicating whether the mission is ready to start. It's initially set to false.
broadCastStartOfMission Function:

This function orchestrates the start of the mission.
It creates a new sync.Cond variable beeper with a mutex.
It adds 3 tasks (representing ninjas) to a sync.WaitGroup named wg. Each task is a call to standByForMission function.
It broadcasts a signal using beeper.Broadcast() to wake up all waiting Goroutines (ninjas).
It waits for all ninjas to complete their tasks by calling wg.Wait().
Finally, it prints a message indicating that all ninjas have completed the mission.
standByForMission Function:

This function represents a ninja waiting for the mission to start.
It takes a function fn as a parameter, which represents the task of the ninja.
It creates a Goroutine that waits for a signal on the beeper condition variable using beeper.Wait().
Once the signal is received, it executes the fn function representing the task of the ninja.
getReadyForMissionWithCond Function:

This function prepares for the mission.
It creates a new sync.Cond variable cond with a mutex.
It starts a Goroutine gettingReadyWithCond to simulate getting ready for the mission.
It enters a loop and waits until readyForMission becomes true.
Once readyForMission becomes true, it prints a message indicating that the mission is ready.
gettingReadyWithCond Function:

This function simulates the process of getting ready for the mission.
It sleeps for a random duration to simulate some preparation time.
After sleeping, it sets readyForMission to true and signals on the cond condition variable using cond.Signal().
sleep Function:

This function generates a random sleep duration to simulate some task taking time.
main Function:

The main function serves as the entry point of the program.
It first prepares for the mission by calling getReadyForMissionWithCond.
Then, it orchestrates the start of the mission by calling broadCastStartOfMission.
Overall, the code simulates a scenario where multiple Goroutines (ninjas) are waiting for a condition to start a mission. Once the condition is met, all ninjas start their tasks, and when they complete their tasks, the mission is considered complete.

*/
