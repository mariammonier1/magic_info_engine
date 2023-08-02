package responders

import (
	"magic_info_engine/adapter/messagingAdapter/natsAdapter"
	"magic_info_engine/common"
	"magic_info_engine/domain/domainCommon/domainLogs"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/inputPort/messagingInputPort"
	"github.com/nats-io/nats.go"
)

var config = common.GetEnvConfig()

type mdResponder struct {
	DomainPort domainPort.IDomainPort
}

func NewNatsResponderAdapter(domainPort domainPort.IDomainPort) messagingInputPort.IMessagingRespondPort {
	responder := &mdResponder{
		DomainPort: domainPort,
	}
	responder.Start()
	return responder
}

func (thisMdC *mdResponder) Respond(topic string, cb func(message *nats.Msg)) {
	connection, connectionErr := natsAdapter.NewConnection()
	if connectionErr != nil {
		domainLogs.PrintError(connectionErr)
	}
	_, subscriptionErr := connection.QueueSubscribe(topic, "core-engine-queue55", cb)
	if subscriptionErr != nil {
		domainLogs.PrintError(subscriptionErr)
	}
}

func (thisMdC *mdResponder) Start() {
	
}
