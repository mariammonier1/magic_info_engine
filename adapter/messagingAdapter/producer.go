package messageDistributor

import (
	"magic_info_engine/adapter/messagingAdapter/producers"
	"magic_info_engine/port/outputPort/messagingOutputPort"
)

func NewProducerAdapter() messagingOutputPort.IMessagingProducerPort {
	producer := producers.NewNatsProducer()
	return producer
}
