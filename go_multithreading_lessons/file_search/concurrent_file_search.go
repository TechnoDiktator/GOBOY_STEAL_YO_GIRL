package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)


var (

	matches []string
	waitgroup  = sync.WaitGroup{}
	lock = sync.Mutex{}
)

func fileSearch(root string , filename string){
	fmt.Println("Searching in ", root)
	files , _ := ioutil.ReadDir(root)
	for _ , file :=range files {
		if strings.Contains(file.Name() , filename){
			//why do we need to get a lock here ....
			//i think
			lock.Lock()
			matches = append(matches, filepath.Join(root , file.Name()))
			lock.Unlock()
		}else if (file.IsDir()){
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root , file.Name())  , filename)	
		}
	}
	//so this is a recursive code that creates a waitgroup everytime it detects that the current poath is a folder
	waitgroup.Done()
}











