package main_test

import (
	
	"sync"
	"testing"
)

var iterations int = 1000

func BenchmarkRegularMap(t *testing.B) {
	var mutex sync.Mutex
	regularMap := make(map[int]int , iterations)
	var wg sync.WaitGroup
	wg.Add(iterations - 1)
	regularMap[0] = 0

	for i:=1 ; i< iterations ; i++ {
		go func(){
			defer wg.Done()

			//j := 1
			for j:= 1 ; j< iterations ; j++ {
				go func () {
					mutex.Lock()
					defer mutex.Unlock()
					regularMap[i] = j
					//fmt.Println("writing value "  , j , " to key " , i)
				}()
			}
		}()
	}
	wg.Wait()

}

func BenchmarkSyncmap(b *testing.B) {
	var syncmap sync.Map
	var wg sync.WaitGroup
	wg.Add(iterations - 1)


	for i:=1 ; i< iterations ; i++ {
		go func(){
			defer wg.Done()

			for j := 1 ; j< iterations ; j++ {
				go func(){
					syncmap.Store(i , j)
				}()
			}
		}()
	}

	wg.Wait()


}

























