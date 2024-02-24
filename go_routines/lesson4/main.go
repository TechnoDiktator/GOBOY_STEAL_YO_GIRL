package main

import (
	"fmt"
	
)


func main() {
	channel := make(chan string ,  2)


	channel <- "FIRST MESSAGE"			
	channel <- "SECOND MESSAGE"

	fmt.Println("SEND DATA TO THE CHANNEL WE JUST CRETAED ABOVE")
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}


