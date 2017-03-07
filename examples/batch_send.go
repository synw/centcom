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
	
	// connect
	cli := centcom.New(host, port, key)
	cli, err := centcom.Connect(cli)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer centcom.Disconnect(cli)
	
	cli, err = cli.CheckHttp()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	channel := "public:data"
	p := []int{1}
	payload, _ := json.Marshal(p)
	started := time.Now()
	// publish 1000 messages with http
	for i:=0;i<1000; i++ {
		_ = cli.Http.AddPublish(channel, payload)
	}
	//time.Sleep(1*time.Second)
	result, err := cli.Http.Send()
	fmt.Println("Sent", len(result), "messages in one request")	
	fmt.Println(time.Since(started))
}
