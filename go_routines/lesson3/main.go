package main

import (
	"fmt"
	"time"
)

//without channels the functionality acheived is same so why do we need achannels
//THIS is WHERE WE ACTUALLU EXPLOITED THE CHANNEL THING
func main() {
	now := time.Now()
	defer func(){
		fmt.Println(time.Since(now))
	}()
	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja , smokeSignal)
	fmt.Println(" the Evil Ninja Was Attacket Successfully !!!")
	
	// ---------------------------- THIS LINE WILL RESULT IN DEADLOCK IF UNCOMMENTED --------------------------
	// smokeSignal <- false
	
	//EXPLANATION :
	/*
	The main() function starts executing.
	It creates a channel smokeSignal.
	It spawns a goroutine to execute the attack function.
	It prints a message indicating the Evil Ninja was attacked successfully.
	It tries to send a value (false) on the smokeSignal channel.
	The attack function in the goroutine executes, sleeps for one second, and tries to send a value (true) on the same smokeSignal channel.
	At this point, a deadlock occurs:

	The main() function is blocked on sending false into the smokeSignal channel because 
	there's no goroutine ready to receive the value.
	The attack goroutine is blocked on sending true into the same smokeSignal channel because 
	the main() function is not ready to receive.
	Both goroutines are waiting for each other, resulting in a deadlock situation.

	To fix this, you need to ensure that either the main() function or the attack goroutine 
	is ready to receive from the channel before sending data. 
	For example, you could add a receive operation (<-smokeSignal) in the main() 
	function before sending false or modify the synchronization logic to ensure proper communication between the goroutines.	
	*/
	//-----------------------------------------------DEAD LOCK EXPLANATION ---------------------------------------------------


	fmt.Println(<-smokeSignal)
	
}

func attack (target string , attacked chan bool){
	time.Sleep(time.Second)
	fmt.Println("Throwing  ninja starts as " , target)
	attacked <- true
}



