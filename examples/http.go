package main

import (
	"fmt"
	"time"
	"encoding/json"	
	"github.com/synw/centcom"	
)


func main() {
	host := "localhost"
	port := 8001
	key := "238a3fd4-71c9-4eb1-9995-914b235efef1"
	// set options
	centcom.SetVerbosity(2)
	
	started := time.Now()	
	// connect
	cli := centcom.New(host, port, key)
	cli, err := centcom.Connect(cli)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer centcom.Disconnect(cli)
	
	// verify the connection
	cli, err = cli.CheckHttp()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	channel := "public:data"
	// suscribe
	cli, err = cli.Suscribe(channel)
	if err != nil {
		fmt.Println(err)
	}
	// listen
	go func() {
	fmt.Println("Listening ...")
	for msg := range(cli.Channels) {
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
	cli, err = cli.Unsuscribe(channel)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(time.Since(started))	
}
