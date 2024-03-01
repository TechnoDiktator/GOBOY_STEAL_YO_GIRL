package main

import (
	"mqttgo/publisher"
	"mqttgo/subscriber"
	"sync"
)

func main(){

	var wg sync.WaitGroup
	wg.Add(2)
	
	go subscriber.Subscriber(&wg)
	go publisher.Publisher(&wg)

	wg.Wait()



}

