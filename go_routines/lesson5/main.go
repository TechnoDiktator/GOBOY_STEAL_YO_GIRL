package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main (){
	fmt.Println("Welcome to golang")
	channel :=  make(chan string)
	numRounds := 10
	go throwingNinjaStar(channel  , numRounds)
	
	for i := 0; i < numRounds; i++ {
		fmt.Println(<-channel)	
	}

}
func throwingNinjaStar(channel  chan string  , numRounds int) {
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < numRounds; i++ {		
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored : " , score)
	}
}
