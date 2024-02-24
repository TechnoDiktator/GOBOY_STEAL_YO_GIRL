package main

import (
	"fmt"

)

func main () {
	ninja1 , ninja2 := make(chan string) , make(chan string)
	go captainElect(ninja1 , "NINJA 1")
	go captainElect(ninja2 , "NINJA 2")
	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
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


