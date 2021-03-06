package main

import (
	"fmt"
	"github.com/synw/centcom"
	"time"
)

func main() {
	addr := "localhost:8001"
	key := "secret_key"
	// set options
	centcom.SetVerbosity(2)

	started := time.Now()
	// connect
	cli := centcom.New(addr, key)
	err := centcom.Connect(cli)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer centcom.Disconnect(cli)
	/* suscribe. Note: this namespace must be set to public=true in the Centrifugo's config
	in order to suscribe to the channel */
	err = cli.Subscribe("public:data")
	if err != nil {
		fmt.Println(err)
	}
	// listen
	go func() {
		fmt.Println("Listening ...")
		for msg := range cli.Channels {
			if msg.Channel == "public:data" {
				fmt.Println("PAYLOAD", msg.Payload, msg.UID)
			}
		}
	}()

	// publish
	payload := []int{1}
	err = cli.Publish("public:data", payload)
	if err != nil {
		fmt.Println(err)
	}

	payload2 := make(map[string]string)
	payload2["hello"] = "world"
	_ = cli.Publish("public:data", payload2)

	// unsuscribe
	err = cli.Unsubscribe("public:data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Since(started))
}
