package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main () {
	rdb := redis.NewClient(&redis.Options{

		Addr: "localhost:6379",
		Password: "",
		DB: 0,

	})

	//define a context with a timeout of 10 seconds
	ctx :=  context.Background()

	ctx , cancel := context.WithTimeout(ctx , time.Second*10)
	defer cancel()


	pubsub := rdb.PSubscribe(ctx  , "mychannel")
	ch := pubsub.Channel()


	go func() {
		for msg := range ch {
			fmt.Print("received message: %s \n" , msg.Payload)
		}
	}()
	
	err := rdb.Publish(ctx , "mychannel1.1" , "Hello worrls" ).Err()

	if err != nil {
		log.Fatalf("Failed to publish the message" , err)
	
	}
	
	<-ctx.Done()
}




