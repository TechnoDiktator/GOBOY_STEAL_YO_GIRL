package main

/*
what is a cond in goland

ChatGPT
In Go, a sync.Cond is a condition variable, 
used for thread synchronization. 
It allows one Goroutine to signal another Goroutine that a particular condition has been met. Condition variables are commonly used in concurrent programming to coordinate the execution of multiple Goroutines.

Here's a brief overview of how sync.Cond works:

Initialization: To use a sync.Cond, you typically create it with the sync.NewCond function, 
passing it a locker (typically a mutex) that it will use to coordinate access to the condition variable.

Waiting: Goroutines can wait on a sync.Cond by acquiring 
the associated locker and then calling the Wait method on the condition variable. This releases the locker and blocks the Goroutine until another Goroutine calls Signal or Broadcast on the same condition variable.

Signaling: When a Goroutine wants to signal that a condition has been met, 
it acquires the associated locker and then calls either Signal to wake up one waiting Goroutine or Broadcast to wake up all waiting Goroutines. After signaling, the locker is released.

Handling Spurious Wakeups: It's important to note that Goroutines waiting 
on a condition variable may sometimes wake up without the condition being true (spurious wakeups). Therefore, it's common to wrap the call to Wait in a loop that rechecks the condition after waking up.

Here's a simple example of how sync.Cond can be used:


*/


import (
    "fmt"
    "sync"
    "time"
)

var (
    ready bool
    readyMutex sync.Mutex
)

func main() {
    cond := sync.NewCond(&sync.Mutex{})
    go worker(cond)

    time.Sleep(2 * time.Second)

    readyMutex.Lock()
    ready = true
    readyMutex.Unlock()

    cond.Broadcast()

    time.Sleep(2 * time.Second)
}

func worker(cond *sync.Cond) {
    cond.L.Lock()
    defer cond.L.Unlock()

    for !ready {
        fmt.Println("Worker waiting...")
        cond.Wait()
    }

    fmt.Println("Worker done!")
}
/*
In this example:

We create a condition variable cond associated with a mutex.
We start a worker Goroutine that waits on the condition variable.
After 2 seconds, we signal the condition variable and set the ready flag to true.
The worker Goroutine wakes up, observes that the condition is true, and proceeds with its task.
*/
