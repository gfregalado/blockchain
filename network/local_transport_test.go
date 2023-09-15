package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	errA := tra.Connect(trb)
	errB := trb.Connect(tra)

	assert.Nil(t, errA)
	assert.Nil(t, errB)
	assert.Equal(t, tra.GetPeer(trb.Addr()), trb)
	assert.Equal(t, trb.GetPeer(tra.Addr()), tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	errA := tra.Connect(trb)
	errB := trb.Connect(tra)

	assert.Nil(t, errA)
	assert.Nil(t, errB)

	msg := []byte("Hello World")

	err := tra.SendMessage(trb.Addr(), msg)
	assert.Nil(t, err)

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.Addr())
}
