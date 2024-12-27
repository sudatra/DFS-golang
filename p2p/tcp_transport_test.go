package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"

	assert.Equal(t, tr.listenAddress, listenAddr)
	assert.Nil(t, tr.ListenAndAccept())
	tr := NewTCPTransport(listenAddr)

	select {}
}