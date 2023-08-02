package messagingInputPort

import "github.com/nats-io/nats.go"

type IMessagingRespondPort interface {
	Respond(topic string, cb func(message *nats.Msg))
}
