package deliveryPort

import (
	"magic_info_engine/port/domainPort"
)

type IDeliveryServicePort interface {
	InitAdapter(domainPort domainPort.IDomainPort)
}
