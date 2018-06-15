package main

import (
	"encoding/json"
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

	// verify the connection
	err = cli.CheckHttp()
	if err != nil {
		fmt.Println(err)
		return
	}

	channel := "public:data"
	// suscribe
	err = cli.Subscribe(channel)
	if err != nil {
		fmt.Println(err)
	}
	// listen
	go func() {
		fmt.Println("Listening ...")
		for msg := range cli.Channels {
			if msg.Channel == channel {
				fmt.Println("PAYLOAD", msg.Payload, msg.UID)
			}
		}
	}()

	// publish http
	d := []int{1, 2}
	dataBytes, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	ok, err := cli.Http.Publish(channel, dataBytes)
	if err != nil {
		fmt.Println(err, ok)
	}
	// and all the other methods of *gocent.Client
	presence, _ := cli.Http.Presence(channel)
	fmt.Printf("Presense for channel %s: %v\n", channel, presence)
	history, _ := cli.Http.History(channel)
	fmt.Printf("History for channel %s, %d messages: %v\n", channel, len(history), history)

	// unsuscribe
	err = cli.Unsubscribe(channel)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Since(started))
}
