package main
//wE WILL TALK ABOUT Sync.MAP


import (
	"sync"
	"fmt"
)

func main() {
	regularMap := make(map[int]interface{})

	// for i:=1 ; i <100 ; i++ {

	// 	//Now this will fuck up the map 
	// 	//because it is a regular map that is the functionality of writing a vaue to a key in the mao is not atomic
	// 	//and writing in the map via spawned go routines 
	// 	go func(){
	// 		regularMap[0] = i
	// 	}()

	// }

	// THIS is why we will use sync map
	syncMap := sync.Map{}

	regularMap[0] = 0
	regularMap[1] = 1
	regularMap[2] = 2


	syncMap.Store(0 , 0)
	syncMap.Store(0 , 0)
	syncMap.Store(0 , 0)

	//get 
	regularValue , regularOk := regularMap[0]
	fmt.Println(regularValue , regularOk)
	
	
	mu := sync.Mutex{}
	//delete
	mu.Lock()   //to avoid error in a goroutine you will have to use a lock with a regular mao
	delete(regularMap , 2)
	mu.Unlock()

	//put  with goroutines we will have to do this !!!
	mu.Lock()
	regularValue , regularOk = regularMap[1]
	if regularOk {
		regularMap[1] = 3
		regularValue = regularMap[1]
		fmt.Println()
	}
	mu.Unlock()

	//iteration
	for key , value := range regularMap{
		fmt.Print(key , value , "|")
	}
	fmt.Println()




	
	//get in the sync map
	syncValue , syncOk := syncMap.Load(0)
	fmt.Println(syncValue , syncOk)
	
	//delete just do this and the wrok is done
	syncMap.LoadAndDelete(2)
		

	//get and put
	syncValue , _  = syncMap.LoadOrStore(1 , 12)
	fmt.Println(syncValue)

	//iteration 
	//looping fucking ugly
	syncMap.Range(func(key , value interface{})bool {
		fmt.Print("key : " , key , " value : " ,  value , " | ")
		return true
	})

}





