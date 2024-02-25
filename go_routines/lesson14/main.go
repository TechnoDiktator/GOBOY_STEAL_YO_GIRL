package main

import (
	"fmt"
	"time"
)

// rubber chicken mutex



// this code makes the  main thread sleep for the 
// 5 whole seconds 
/*
that shuld be enough for all the 
go routines to be spawned and finish their execution 
and make the value of the globa; count variable 
1000

But no matter how many times you try is 
it will not show 1000
WHY :::

BECAUSE THE SIMPLE OPERATION of count++
is not atomic !!!!!!!!!!!
WHAT THE ACTUAL FUCK !!!!!
SO all the go routines could be simultaneously accessing different values of the count variable
 instead of one following the other 
// THIS  FUCKS up the last value that is returned

MAA KI AANKH MAZA AA GAYA

*/
var (
	count int
)

func main() {
	iterations:= 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(5*time.Second) 
	fmt.Println("result count is :" , count)
}


// THIS Is not na atomic operation 
func increment() {
	count++  // or count  =  count +!
	
	//STEPS :
	// temp :=  count 
	// temp   =  temp +!
	// count  = temp
}





















