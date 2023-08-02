package component

import (
	"magic_info_engine/port/domainPort/componentPort"
)

type componentUseCases struct {
	componentPort.ICreateComponentPort
}

func NewComponentUseCases() componentPort.IComponentDomainPort {
	return &componentUseCases{
		NewCreateComponentUseCase(),
	}
}
