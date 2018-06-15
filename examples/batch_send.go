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

	// connect
	cli := centcom.New(addr, key)
	err := centcom.Connect(cli)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer centcom.Disconnect(cli)

	err = cli.CheckHttp()
	if err != nil {
		fmt.Println(err)
		return
	}

	channel := "public:data"
	p := []int{1}
	payload, _ := json.Marshal(p)
	started := time.Now()
	// publish 1000 messages with http
	for i := 0; i < 1000; i++ {
		_ = cli.Http.AddPublish(channel, payload)
	}
	//time.Sleep(1*time.Second)
	result, err := cli.Http.Send()
	fmt.Println("Sent", len(result), "messages in one request")
	fmt.Println(time.Since(started))
}
