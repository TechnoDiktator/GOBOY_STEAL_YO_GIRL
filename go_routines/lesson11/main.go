package main

import (
	"fmt"
)

func main () {
	ninja1 , ninja2 := make(chan string) , make(chan string)
	go captainElect(ninja1 , "NINJA 1")
	go captainElect(ninja2 , "NINJA 2")
	
	
	
	// tHe select statement will block the code from executing further
	//until and unless one of its cases are met 
	//if both cases are available ...it eill select one at random and execute 


	//If you give a default value then it can become nonblocking
	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	default:
		fmt.Println("Non blocking !!!!")
	}
	roughlyFair()

} 



func captainElect(ninja chan string , message string) {
	ninja <- message
}


// I AHVE NO LCUE G+HOW THE BELOW CODE IS WORKING
func roughlyFair(){
	ninja1 := make(chan interface{}) ; close(ninja1)
	ninja2 := make(chan interface{}) ; close(ninja2)

	var ninja1Count ,  ninja2Count int 
	for i := 0; i < 10000; i++ {
		select {
			case <- ninja1:
				ninja1Count++

			case <- ninja2:
				ninja2Count++

		}	
	}

	fmt.Printf("ninja1Count %d , ninja2Count: %d \n" , ninja1Count , ninja2Count)



}


