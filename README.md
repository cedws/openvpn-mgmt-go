# openvpn-mgmt-go
openvpn-mgmt-go provides an API for connecting to, receiving messages from, and sending commands to an OpenVPN management socket.

## Example
```go
package main

import (
	"log"

	mgmt "github.com/cedws/openvpn-mgmt-go"
	"github.com/cedws/openvpn-mgmt-go/command"
	"github.com/cedws/openvpn-mgmt-go/event"
)

func main() {
	s, _ := mgmt.DialUnix("/run/openvpn/server/management.sock")

	s.HandleFunc(handleFunc)
	s.ErrorFunc(errorFunc)

	s.Dispatch(command.Bytecount{60})

	s.Start()
}

func handleFunc(v any) {
	switch message := v.(type) {
	case *event.ClientConnect:
		log.Printf("client %v is connecting with key %v", message.ClientID, message.KeyID)
	}
}

func errorFunc(v error) {
	log.Println(v)
}

```
