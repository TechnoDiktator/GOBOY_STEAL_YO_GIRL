package subscriber

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func onMessageReceived (client mqtt.Client , message mqtt.Message){
	fmt.Print("SUBSCRIBER")
	fmt.Println("HELLo \n" , message.Topic() , message.Payload())

}	


func Subscriber(wg *sync.WaitGroup){
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("subscriber")
	opts.SetDefaultPublishHandler(onMessageReceived)
	client := mqtt.NewClient(opts)

	if token := client.Connect() ; token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token  := client.Subscribe("text/topic" , 0 , nil)
	token.Wait()
	fmt.Println("Subscribed to the text topic") 
	
	for{
		
	}
	defer wg.Done()

}		



