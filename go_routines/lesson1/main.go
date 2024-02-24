package main
import (
	"fmt"
	"time"
)
//lets go goriutuines baby
//WE ARE USING CHANNELS WE HAVENt ACTUALLLY EXPLOITED THEIR CAPABILITY
//THIS  code seems to weok correctly because of the sleep we have provided
func main(){
	channel := make(chan bool)
	start := time.Now()
	defer func(){
		fmt.Println(time.Since(start))
	}()
	evilNinjas:= []string{"Tommy" , "Jhonny" , "Bobby" , "Andy"}
	for _ , evil_ninja := range evilNinjas {
		go attack(evil_ninja ,channel)
	} 
	time.Sleep((time.Second*3 + time.Millisecond*500))
	fmt.Print("FUNCTION ENDS ALL NINJAs WERE KILLED")
}


func attack(target string , done chan bool){
	fmt.Println("Throwing the ninja star at :" , target)
	time.Sleep(time.Second)
	fmt.Printf("Evill ninja %v has died \n" , target )
	done <- true
}















