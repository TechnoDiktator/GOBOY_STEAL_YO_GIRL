//ATOMIC PACKAGE IN GO
/*

The atomic package in Go provides low-level atomic memory operations useful for synchronizing access to shared variables among multiple goroutines. These operations ensure that certain operations (like reading, writing, or modifying a variable) are performed atomically, meaning they are not interrupted by other goroutines.

Here are some common functions provided by the atomic package:

Load and Store Operations:

Load*: Functions like LoadInt32, LoadInt64, LoadUint32, LoadUint64, LoadPointer, etc., are used to atomically load the value of a variable.
Store*: Functions like StoreInt32, StoreInt64, StoreUint32, StoreUint64, StorePointer, etc., are used to atomically store a value into a variable.
Additive and Subtractive Operations:

Add*: Functions like AddInt32, AddInt64, AddUint32, AddUint64, etc., are used to atomically add a value to a variable.
Sub*: Functions like SubInt32, SubInt64, SubUint32, SubUint64, etc., are used to atomically subtract a value from a variable.
Bitwise Operations:

And*, Or*, Xor*: Functions like AndInt32, OrInt32, XorInt32, etc., are used for bitwise operations on integers.
Compare and Swap:

CompareAndSwap*: Functions like CompareAndSwapInt32, CompareAndSwapInt64, CompareAndSwapUint32, CompareAndSwapUint64, etc., are used to atomically compare and swap a value in a variable if it matches an expected value.
These functions ensure that operations on shared variables are performed atomically, preventing race conditions and ensuring consistency in concurrent programs.

Here's a simple example demonstrating the usage of the atomic package:

*/
package main
import (
	"fmt"
	"sync/atomic"
	"sync"
)

func main () {


	//with atomic
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum , 1)
	fmt.Println(sum)

	fmt.Println("-------------------------------------------")

	//without atomic
	var mu sync.Mutex
	fmt.Println(sum)
	mu.Lock()
	sum += 1
	mu.Unlock()
	fmt.Println(sum)

	fmt.Println("-------------------------------------------")

	var diffsum int64
	fmt.Println(atomic.LoadInt64(&diffsum))
	atomic.StoreInt64(&diffsum , 1)
	fmt.Println(diffsum)



	//atomic.value
	//NOTE that in this when we pass an object to store we sont pass its address instead we pass the object(copy)
	//whereas above we are passing addresses like in AddInt64 , and StoreInt64
	//so there is a distinction
	tarang := ninja{name:"tarang"}
	var v atomic.Value
	v.Store(tarang)
	v.Load()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		w := v.Load().(ninja)
		w.name = "Not tarang"
		v.Store(w)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(v.Load().(ninja).name)

}


type  ninja struct {
	name string
}