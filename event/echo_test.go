package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidEcho(t *testing.T) {
	var n Echo

	err := n.Parse(">ECHO:1101519562,forget-passwords")
	assert.Nil(t, err)
	assert.Equal(t, int64(1101519562), n.Timestamp)
	assert.Equal(t, "forget-passwords", n.Message)
}
