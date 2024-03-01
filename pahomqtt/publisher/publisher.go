package publisher

import (
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func Publisher(wg *sync.WaitGroup) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("Publisher")
	client:= mqtt.NewClient(opts)
	if token:= client.Connect();token.Wait()&& token.Error() != nil{
		panic("PHAT GYI RE BABA")
	}

	text := "Hello from publisher"

	

	//token := client.Publish("text/topic" , 0 , false , 	text)

	for{
		client.Publish("text/topic" , 0 , false , 	text)
		time.Sleep(time.Second*2)
	}
	//token.Wait()
	fmt.Println("Published \n" , text)


	wg.Done()
}













