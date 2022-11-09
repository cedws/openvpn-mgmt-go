package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidClientConnect(t *testing.T) {
	var n ClientConnect

	err := n.Parse(">CLIENT:CONNECT,1,0")
	assert.Nil(t, err)

	assert.Equal(t, int64(1), n.ClientID)
	assert.Equal(t, int64(0), n.KeyID)
}

func TestParseInvalidClientConnect(t *testing.T) {
	var n ClientConnect

	err := n.Parse(">CLIENT:CONNECT,1,0,1")
	assert.NotNil(t, err)
}

func TestParseValidClientReauth(t *testing.T) {
	var n ClientReauth

	err := n.Parse(">CLIENT:REAUTH,1,0")
	assert.Nil(t, err)

	assert.Equal(t, int64(1), n.ClientID)
	assert.Equal(t, int64(0), n.KeyID)
}

func TestParseInvalidClientReauth(t *testing.T) {
	var n ClientReauth

	err := n.Parse(">CLIENT:REAUTH,1,0,1")
	assert.NotNil(t, err)
}

func TestParseValidClientEnv(t *testing.T) {
	var n ClientEnvVar

	err := n.Parse(">CLIENT:ENV,go=good")
	assert.Nil(t, err)

	assert.Equal(t, false, n.End)
	assert.Equal(t, "go", n.Key)
	assert.Equal(t, "good", n.Value)
}

func TestParseValidClientEnvEnd(t *testing.T) {
	var n ClientEnvVar

	err := n.Parse(">CLIENT:ENV,END")
	assert.Nil(t, err)

	assert.Equal(t, true, n.End)
}

func TestParseInvalidClientEnv(t *testing.T) {
	var n ClientEnvVar

	err := n.Parse(">CLIENT:ENV,")
	assert.NotNil(t, err)
}

func TestParseValidClientEstablished(t *testing.T) {
	var n ClientEstablished

	err := n.Parse(">CLIENT:ESTABLISHED,1")
	assert.Nil(t, err)

	assert.Equal(t, int64(1), n.ClientID)
}

func TestParseInvalidClientEstablished(t *testing.T) {
	var n ClientEstablished

	err := n.Parse(">CLIENT:ESTABLISHED,")
	assert.NotNil(t, err)
}

func TestParseValidClientDisconnect(t *testing.T) {
	var n ClientDisconnect

	err := n.Parse(">CLIENT:DISCONNECT,1")
	assert.Nil(t, err)

	assert.Equal(t, int64(1), n.ClientID)
}

func TestParseInvalidClientDisconnect(t *testing.T) {
	var n ClientDisconnect

	err := n.Parse(">CLIENT:DISCONNECT,")
	assert.NotNil(t, err)
}

func TestParseValidClientAddress(t *testing.T) {
	var n ClientAddress

	err := n.Parse(">CLIENT:ADDRESS,1,address,1")
	assert.Nil(t, err)

	assert.Equal(t, int64(1), n.ClientID, 1)
	assert.Equal(t, "address", n.Address)
	assert.Equal(t, true, n.Primary)
}

func TestParseInvalidClientAddress(t *testing.T) {
	var n ClientAddress

	err := n.Parse(">CLIENT:ADDRESS,")
	assert.NotNil(t, err)
}

func TestParseValidClientChallengeResponse(t *testing.T) {
	var n ClientChallengeResponse

	err := n.Parse(">CLIENT:CR_RESPONSE,1,0,aGVsbG8K")
	assert.Nil(t, err)

	assert.Equal(t, int64(1), n.ClientID)
	assert.Equal(t, int64(0), n.KeyID)
	assert.Equal(t, "aGVsbG8K", n.ResponseBase64)
}

func TestParseInvalidClientChallengeResponse(t *testing.T) {
	var n ClientChallengeResponse

	err := n.Parse(">CLIENT:CR_RESPONSE,1,0")
	assert.NotNil(t, err)
}
