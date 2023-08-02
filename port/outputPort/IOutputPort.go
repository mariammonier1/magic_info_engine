package outputPort

import (
	"magic_info_engine/port/outputPort/messagingOutputPort"
)

type IOutputPort struct {
	Producer messagingOutputPort.IMessagingProducerPort
}
