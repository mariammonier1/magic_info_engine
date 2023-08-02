package messageDistributor

import (
	"magic_info_engine/adapter/messagingAdapter/consumers"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/inputPort/messagingInputPort"
)

func NewConsumerAdapter(domainPort domainPort.IDomainPort) messagingInputPort.IMessagingConsumerPort {
	consumer := consumers.NewNatsConsumerAdapter(domainPort)
	return consumer
}
