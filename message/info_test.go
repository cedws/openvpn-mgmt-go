package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidInfo(t *testing.T) {
	t.Parallel()
	var n Info

	err := n.Parse(">INFO:rmation")
	assert.Nil(t, err)
	assert.Equal(t, "rmation", n.Message)
}
