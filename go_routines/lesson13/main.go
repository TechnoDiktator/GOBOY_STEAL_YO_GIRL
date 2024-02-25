package main

import (
	"fmt"
	"sync"

)
/*


so the waitgroup will tell for how many goriutines the function whree the waitgroup is declared will wait

ChatGPT
Exactly. The sync.WaitGroup is used to wait for a specified number of Goroutines to complete their tasks. 
The number you pass to Add method indicates how many Goroutines the WaitGroup should wait for.

When a Goroutine is started, you increment the WaitGroup counter using Add(1). When the Goroutine completes its task, 
it decrements the counter using Done(). The main Goroutine (or any other Goroutine that needs to wait) calls Wait() 
to block until the counter becomes zero, meaning that all the associated Goroutines have completed.

Here's a breakdown:

wg.Add(1): Indicates that you are starting a Goroutine, and the WaitGroup should wait for this Goroutine to finish.

wg.Done(): Indicates that a Goroutine has completed its task. You call this at the end of your Goroutine.

wg.Wait(): Blocks the calling Goroutine until the counter becomes zero, 
meaning that all Goroutines added to the WaitGroup have completed.

It's a useful tool for coordinating the execution of multiple Goroutines and ensuring that t
he main program doesn't exit before they are all done.

*/



//since the
func main() {	
	var beeper sync.WaitGroup
	
	evil_ninjas := []string{"tommy" , "johhny" , "bobby"}

	beeper.Add( len( evil_ninjas ))
	
	for _ , ninja := range evil_ninjas{
		go attack(ninja , &beeper)	
	}
	 // in the waitgroup we are waiting for the goroutines to complete
	beeper.Wait()
	fmt.Println("Mission Complete")


}

func attack(evilNinja string , beeper *sync.WaitGroup) {
	fmt.Println("Attacked the evil ninja"  , evilNinja)
	beeper.Done()

}







