package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)



//SEE that this global variable is not synchronised 
// that is if multiple go routines try to chage the value of 
// this variable ...it can break the system / because when one routine is writing a value to the system
// another routine might access the value and ...stick to the value 
//although the previus goroutine was busy
var ( 	
	missionComplete bool
	missionLock sync.Mutex
)
//Once is an object that will perform only on action
//a Once must not be copied after first use

func markMissionComplete(){
	missionComplete = true
}

type Once struct{

}


func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return (0 == rand.Intn(10))
}

func checkMissionComplete() {
	if missionComplete {
		fmt.Println("Mission is now completed")
	}else {
		fmt.Println("Mission was a failure !!")
	}
}



func main() {

	var wg  sync.WaitGroup
	wg.Add(1)
	
	go func() {
		if foundTreasure() {
			markMissionComplete()
		}

		wg.Done()
	}()

	wg.Wait()

	checkMissionComplete()

}


