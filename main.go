package centcom

import (
	"fmt"
	"strconv"
	"github.com/synw/centcom/ws"
	"github.com/synw/centcom/state"
)


func New(host string, port int, key string) *ws.Cli {
	return ws.NewClient(host, port, key)
}

func Connect(cli *ws.Cli) (*ws.Cli, error) {
	return ws.Connect(cli)
}

func Disconnect(cli *ws.Cli) {
	cli.Conn.Close()
	if state.Verbosity > 0 {
		msg := "Disconnected from "+cli.Host
		fmt.Println(msg)
	}
	close(cli.Channels)	
}

func SetVerbosity(v int) {
	state.Verbosity = v
}

func State() string {
	v := strconv.Itoa(state.Verbosity)
	msg := "- Verbosity is set to "+v
	return msg
}

func PrintState() {
	fmt.Println(State())
}
