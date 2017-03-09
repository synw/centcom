# Centcom

Go API for the Centrifugo websockets server.

## Quick example

Initialize connection:

   ```go
cli := centcom.New("locahost", 8001, "secret_key")
cli, err := centcom.Connect(cli)
if err != nil {
	fmt.Println(err)
	return
}
defer centcom.Disconnect(cli)
   ```
   
Suscribe and listen to a channel:

   ```go
cli, err = cli.Suscribe("public:data")
if err != nil {
	fmt.Println(err)
}
// listen
go func() {
fmt.Println("Listening ...")
for msg := range(cli.Channels) {
	if msg.Channel == "public:data" {
		fmt.Println("PAYLOAD", msg.Payload, msg.UID)
	}
}
}()
   ```
   
Publish into a channel using client drivers (channel has to be set `public=true` in Centrifugo's config):

   ```go
payload := []int{1}
err = cli.Publish("public:data", payload)
if err != nil {
	fmt.Println(err)
}
   ```
   
Publish to a channel using the server drivers (no restrictions):

   ```go
d := []int{1, 2}
dataBytes, err := json.Marshal(d)
if err != nil {
	fmt.Println(err)
}
ok, err := cli.Http.Publish(channel, dataBytes)
if err != nil {
	fmt.Println(err, ok)
}
   ```

Check the [examples](https://github.com/synw/centcom/tree/master/examples)

## API

#### Centcom methods:

`centcom.New(host string, port int, key string)`: initialize client

`centcom.Connect(cli *Cli)`: connect the client drivers

`centcom.Disconnect(cli *Cli)`: disconnect the client drivers

#### Cli methods:

`cli.CheckHttp()`: verify the server side drivers connection

`cli.Suscribe(channel string)`: subscribe to a channel

`cli.Unsuscribe(channel string)`: unsubscribe to a channel

`cli.Publish(channel string, payload interface{})`: publish into a channel using client drivers4

`cli.Http` is a *gocent.Client with all its method: check the [Gocent API](https://godoc.org/github.com/centrifugal/gocent)
