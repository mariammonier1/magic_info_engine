package consumers

import (
	"magic_info_engine/adapter/messagingAdapter/natsAdapter"
	"magic_info_engine/common"
	"magic_info_engine/domain/domainCommon/domainLogs"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/inputPort/messagingInputPort"
)

var config = common.GetEnvConfig()

type mdConsumer struct {
	DomainPort domainPort.IDomainPort
}

func NewNatsConsumerAdapter(domainPort domainPort.IDomainPort) messagingInputPort.IMessagingConsumerPort {
	consumer := &mdConsumer{
		DomainPort: domainPort,
	}
	consumer.Start()
	return consumer
}

func (thisMdC *mdConsumer) Consume(Topic string, cb func(message []byte)) {
	connection, connectionErr := natsAdapter.NewConnection()
	if connectionErr != nil {
		domainLogs.PrintError(connectionErr)
	}
	_, subscriptionErr := connection.QueueSubscribe(Topic, "magic_info_enginewafik", cb)
	if subscriptionErr != nil {
		domainLogs.PrintError(subscriptionErr)
	}
}

func (thisMdC *mdConsumer) Start() {

}
