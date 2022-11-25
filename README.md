# openvpn-mgmt-go
openvpn-mgmt-go provides an API for connecting to, receiving messages from, and sending commands to an OpenVPN management socket.

## Example
```go
package main

import (
	openvpn "github.com/cedws/openvpn-mgmt-go"
	"github.com/cedws/openvpn-mgmt-go/command"
	"github.com/cedws/openvpn-mgmt-go/message"
)

func main() {
	socket, err := openvpn.DialUnix("/run/openvpn/server/management.sock")
	if err != nil {
		// ...
	}

	plugin := AuthPlugin{socket}
	plugin.Init()
}

type AuthPlugin struct {
	socket *openvpn.Socket
}

func (a *AuthPlugin) Init() {
	a.socket.OnClientConnect(a.onClientConnect)
	a.socket.Start()
}

func (a *AuthPlugin) onClientConnect(c message.ClientConnect) {
	// allow the connecting client
	a.socket.Dispatch(command.ClientAuthNt{c.ClientID, c.KeyID})
}
```
