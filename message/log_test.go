package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValidLog(t *testing.T) {
	t.Parallel()
	var n Log

	err := n.Parse(">LOG:1101519562,IFNWD,message")
	assert.Nil(t, err)
	assert.Equal(t, int64(1101519562), n.Timestamp)
	assert.Equal(t, "IFNWD", n.Flags)
	assert.Equal(t, "message", n.Message)
}
