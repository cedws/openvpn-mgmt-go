package openvpn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeSCRV1Valid(t *testing.T) {
	t.Parallel()

	var dec SCRV1
	err := dec.DecodeString("SCRV1:Zm9v:ODY3NTMwOQ==")
	assert.Nil(t, err)
	assert.Equal(t, "foo", dec.Password)
	assert.Equal(t, "8675309", dec.ChallengeResponse)
}

func TestDecodeSCRV1Invalid(t *testing.T) {
	t.Parallel()

	var dec SCRV1
	err := dec.DecodeString("SCRV2:Zm9v:ODY3NTMwOQ==")
	assert.Equal(t, ErrSCRV1Prefix, err)
}
