package main_test
import (
	"sync"
	"sync/atomic"
	"testing"	
)

const iterations = 1000

func BenchmarkMutexAdd (t *testing.B) {

	var wg sync.WaitGroup
	wg.Add(iterations)
	var sum int64
	var mutex sync.Mutex
	for i:=0; i<iterations ; i++{
		go func(){
			defer wg.Done()
			for i:=0 ; i <iterations ; i++ {
				mutex.Lock()
				sum++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
}


func BenchmarkAtomicAdd (t *testing.B) {
	var wg sync.WaitGroup

	wg.Add(iterations)
	var sum int64

	for i := 0; i < iterations; i++ {
		go func() {
			defer  wg.Done()
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&sum , 1)
			}
		}()
	}

	wg.Wait()

}




