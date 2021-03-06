package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func receive(event cloudevents.Event) {
	// do something with event.
	fmt.Printf("%s", event)
	log.Printf("%s", event)
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("start main")
	// The default client is HTTP.
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), receive))
}
