package main
import (
	"fmt"
	"math/rand"
	"time"
)

// READ THE SUMMARY BELOW

func main (){
	fmt.Println("Welcome to golang")
	channel :=  make(chan string)
	go throwingNinjaStar(channel )
	
	for message := range channel{
		fmt.Println(message)	
	}

}
func throwingNinjaStar(channel  chan string ) {
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i <10; i++ {		
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored : " , score)
	}


	//this is necessary 
	/*
		1. The for loop above i sending messages to the channel continuously
		2. But the thing is that it will stop sending message sto the channel after 10 rounds 
		3. SEE THAT the for loop in main function will keep on running even after all 10 message are sent
		4. But we dont have any more messages
		5. Hence the program will not end ...IT WILL GIVE  DEADLOCK !!!!! and stop...

		6. TO PREVENT THIS FROM HAPPENING WE WILL IMMEDIATELY CLOSE THE CHANNEL . that is basiclally 
		//prompt the main function that the channel has stopped so the its loop can terninate    

	*/
	// HENCE WE ARE USING THE CLOSE METHOD TO CLOSE THE CHANNEL !!!!!
	close(channel)

}
