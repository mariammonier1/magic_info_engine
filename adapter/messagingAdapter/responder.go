package messageDistributor

import (
	"magic_info_engine/adapter/messagingAdapter/responders"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/inputPort/messagingInputPort"
)

func NewResponderAdapter(domainPort domainPort.IDomainPort) messagingInputPort.IMessagingRespondPort {
	responder := responders.NewNatsResponderAdapter(domainPort)
	return responder
}
