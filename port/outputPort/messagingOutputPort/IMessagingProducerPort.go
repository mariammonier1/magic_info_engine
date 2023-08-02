package messagingOutputPort

type IMessagingProducerPort interface {
	Produce(topic string, message interface{}) error
}
