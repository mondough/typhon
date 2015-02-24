package test

import (
	"testing"

	"github.com/b2aio/typhon/server"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/require"
)

var stubServer *StubServer

func InitStubServer(t *testing.T) *StubServer {
	if stubServer == nil {
		stubServer = NewStubServer(t)
	}
	return stubServer
}

// Helper method to call a handler function (the function being tested)
// directly with a `proto.Message`.
// Returns errors that were returned from the handler function directly.
// Marshalling errors cause the test to fail instantly
func CallEndpoint(t *testing.T, endpoint server.Endpoint, reqProto proto.Message, respProto proto.Message) error {
	// Call handler with amqp delivery
	reqBytes, err := proto.Marshal(reqProto)
	require.NoError(t, err)
	resp, err := endpoint.HandleRequest(server.NewAMQPRequest(&amqp.Delivery{
		Body: reqBytes,
	}))
	if err != nil {
		return err
	}
	respBytes, err := resp.Encode()
	require.NoError(t, err)
	err = proto.Unmarshal(respBytes, respProto)
	require.NoError(t, err)
	return nil
}
