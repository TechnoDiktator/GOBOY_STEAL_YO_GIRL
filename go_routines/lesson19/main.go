package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


//USING Sync.once in this code



//Here we are properly lovking the variable before 
//any goroutine accesses it so thet a single goroutine has the lock of the variable 
// this ensures that the otherGoriutines do not take a wrong value


var  (
	missionComplete bool
	missionLock sync.Mutex
)


func markMissionComplete(){
	missionLock.Lock()
	defer missionLock.Unlock()
	missionComplete = true
}


type Once struct {

}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return (0 == rand.Intn(10))	
}


func checkMissionComplete() {
	missionLock.Lock()
	defer missionLock.Unlock()
	if missionComplete{
		fmt.Println("Mission is complete now")
	}else {
		fmt.Println("Mission was a failure !!!")
	}
}

func main () {
	var wg sync.WaitGroup
	wg.Add(100)

	var once sync.Once

	// NOW IN THIS CODE WE ARE running 100 goroutines and waiting for them to be completed 
	//all these goroutines have one object to mark the mission complete global flag as true
	//although the rand int value being 0 has 10%chance ...when we run 100 go rooutines 
	//we increase the probability that atleast one of then=m will have 0 rand
	//and thius mark the moission as complete
	
	for i := 0; i < 100; i++ {
		go func(){	

			defer wg.Done()
			if foundTreasure() {
				once.Do(markMissionComplete)
			}
		}()		
	}
	wg.Wait()
	fmt.Println("I hope the current ninja has not died!!!")
	checkMissionComplete()
}








