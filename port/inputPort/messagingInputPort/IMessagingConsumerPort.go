package messagingInputPort

type IMessagingConsumerPort interface {
	Consume(Topic string, cb func(message []byte))
}
