package connections_test

import (
	"github.com/NeowayLabs/wabbit/amqptest/server"
	"github.com/streadway/amqp"
	"gotest.tools/assert"
	"testing"
)

func TestDial(t *testing.T) {
	url := "amqp://localhost:5672/%2f"
	amqpServer := server.NewServer(url)
	assert.NilError(t, amqpServer.Start())
	defer stopServer(t, amqpServer)

	connection, err := amqp.Dial(url)

	assert.NilError(t, err)
	defer closeConnection(t, connection)
	assert.Check(t, connection != nil)
}

func stopServer(t *testing.T, amqpServer *server.AMQPServer) {
	assert.NilError(t, amqpServer.Stop())
}

func closeConnection(t *testing.T, connection *amqp.Connection) {
	assert.NilError(t, connection.Close())
}
