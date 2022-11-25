package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidHold(t *testing.T) {
	t.Parallel()
	var n Hold

	err := n.Parse(">HOLD:Waiting for hold release:10")
	assert.Nil(t, err)
	assert.Equal(t, "Waiting for hold release", n.Message)
	assert.Equal(t, 10, n.Wait)
}
