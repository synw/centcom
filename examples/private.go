package main

import (
	"fmt"
	"time"
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
	
	channel := "cmd:$ch1"
	/* suscribe. Note: this namespace must be set to public=true in the Centrifugo's config 
	in order to suscribe to the channel */
	cli, err = cli.Suscribe(channel)
	if err != nil {
		fmt.Println(err)
	}
	// listen
	go func() {
	fmt.Println("Listening ...")
	for msg := range(cli.Channels) {
		if msg.Channel == channel {
			fmt.Println("PAYLOAD", msg.Payload)
		}
	}
	}()
	
	// publish
	payload := []int{1}
	err = cli.Publish(channel, payload)
	if err != nil {
		fmt.Println(err)
	}
	
	payload2 := make(map[string]string)
	payload2["hello"] = "world"
	_ = cli.Publish(channel, payload2)
	
	// unsuscribe
	cli, err = cli.Unsuscribe(channel)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(time.Since(started))
}
