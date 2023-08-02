package producers

import (
	"magic_info_engine/adapter/messagingAdapter/natsAdapter"
	"magic_info_engine/port/outputPort/messagingOutputPort"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

type natsProducer struct {
	Connection *nats.EncodedConn
}

func NewNatsProducer() messagingOutputPort.IMessagingProducerPort {
	connection, connectionErr := natsAdapter.NewConnection()
	if connectionErr != nil {
		log.Error(connectionErr)
	}
	return &natsProducer{
		Connection: connection,
	}
}

func (thisNs *natsProducer) Produce(topic string, message interface{}) error {
	log.Debug("topic: ", topic)
	log.Debug("message: ", message)
	return thisNs.Connection.Publish(topic, message)
}
