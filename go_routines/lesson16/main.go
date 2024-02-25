package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	lock sync.Mutex   // this is the mutex lock that we will aquie
	rwlock sync.RWMutex

)




func readAndWrite() {
	go read()
	go write()
	time.Sleep(5*time.Second)
	fmt.Println("Done")
}



// when we talk anout reading 
// the lock can be acquired by an arbitrary number of readers or a Single 
//Writer
// the zero value of an rwmutex is an unlocked mtex
func read() {
	rwlock.RLock()
	defer rwlock.RUnlock()
	fmt.Println("Read locking")
	time.Sleep(1*time.Second)
	fmt.Println("reading unlocking")
}

func write() {
	rwlock.Lock()
	defer rwlock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1*time.Second)
	fmt.Println("Write unlockeing")


}



func basics() {
	iterations:= 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(2*time.Second) 
	fmt.Println("result count is :" , count)

}

func main() {
	readAndWrite()
}



func increment () {
	lock.Lock()
	count++
	lock.Unlock()

}
