package main

import (
	"fmt"
	"time"
)

//without channels the functionality acheived is same so why do we need achannels 

func main() {
	now := time.Now()
	defer func(){
		fmt.Println(time.Since(now))
	}()


	evilNinja := "Tommy"
	go attack(evilNinja)
	time.Sleep(time.Second*2)

}



func attack (target string){

	time.Sleep(time.Second)
	fmt.Println("Throwing  ninja starts as " , target)
}