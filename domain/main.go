package domain

import (
	"magic_info_engine/domain/domianUsecase/component"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/domainPort/componentPort"
	"magic_info_engine/port/outputPort"
)

type domainUseCase struct {
	componentPort.IComponentDomainPort
}

func NewDomain(outputPort outputPort.IOutputPort) domainPort.IDomainPort {

	componentUseCases := component.NewComponentUseCases()

	return &domainUseCase{
		componentUseCases,
	}
}
