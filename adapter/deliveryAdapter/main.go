package deliveryAdapter

import (
	"magic_info_engine/adapter/deliveryAdapter/echoAdapter"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/inputPort/deliveryPort"
	"sync"
)

var (
	instance *deliveryAdapter
	once     sync.Once
)

type deliveryAdapter struct {
	Adapter deliveryPort.IDeliveryServicePort
}

func NewDeliveryAdapter(domainPort domainPort.IDomainPort) deliveryPort.IDeliveryServicePort {
	once.Do(func() {
		instance = &deliveryAdapter{}
		instance.Adapter = &echoAdapter.RestEchoImpl{}
		instance.Adapter.InitAdapter(domainPort)
	})
	return instance.Adapter
}
