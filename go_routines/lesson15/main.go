package main

import (
	"fmt"
	"sync"
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
	lock sync.Mutex   // this is the mutex lock that we will aquie
	rwlock sync.RWMutex

)





func main() {
	iterations:= 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(2*time.Second) 
	fmt.Println("result count is :" , count)
}



func increment () {
	lock.Lock()
	count++
	lock.Unlock()

}

// // THIS Is not na atomic operation 
// func increment() {
// 	count++  // or count  =  count +!
	
// 	//STEPS :
// 	// temp :=  count 
// 	// temp   =  temp +!
// 	// count  = temp
// }