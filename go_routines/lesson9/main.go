package main

import (
	"fmt"
)


// SELECT

/*
The select statement lets a goroutine wait on multiple communication operations
A select blocks until one of its cases can run  , then is executes the case . It chooses on e
at randon if multiple cases/conditions are ready
*/


func main() {
	//fmt.Print("Hello")
	ninja1 , ninja2 := make(chan string) , make(chan string)

	go captainElect(ninja1 , "Ninja 1")
	go captainElect(ninja2 , "Ninja 2")


	select {
		case message_from_ninja1 := <-ninja1:
			fmt.Println(message_from_ninja1)
		case message_from_ninja2 := <-ninja2:
			fmt.Println(message_from_ninja2)
	}
}



func captainElect(ninja chan string ,  message string){
	ninja <- message
}

















