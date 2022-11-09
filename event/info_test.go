package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidInfo(t *testing.T) {
	var n Info

	err := n.Parse(">INFO:rmation")
	assert.Nil(t, err)
	assert.Equal(t, "rmation", n.Message)
}
