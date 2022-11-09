package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidFatal(t *testing.T) {
	var n Fatal

	err := n.Parse(">FATAL:ity")
	assert.Nil(t, err)
	assert.Equal(t, "ity", n.Message)
}
