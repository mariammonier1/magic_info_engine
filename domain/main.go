package domain

import (
	magicInfo "magic_info_engine/domain/domianUsecase/magicInfoDevices"
	"magic_info_engine/port/domainPort"
	"magic_info_engine/port/domainPort/magicInfoPort"
	"magic_info_engine/port/outputPort"
)

type domainUseCase struct {
	magicInfoPort.IMagicInfoDevicesPort
}

func NewDomain(outputPort outputPort.IOutputPort) domainPort.IDomainPort {

	magicInfoUseCases := magicInfo.NewMagicInfoUseCases(outputPort.MagicInfoHandeler)

	return &domainUseCase{
		magicInfoUseCases,
	}
}
