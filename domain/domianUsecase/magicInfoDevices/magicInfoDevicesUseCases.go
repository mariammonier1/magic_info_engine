package magicInfo

import (
	"magic_info_engine/port/domainPort/magicInfoPort"
	magicInfoOutputPort "magic_info_engine/port/outputPort/magicInfoOutPutPort"
)

type magicInfoUseCases struct {
	magicInfoPort.IMagicInfoDevicesPort
}

func NewMagicInfoUseCases(magicInfoPort magicInfoOutputPort.IMagicInfoDevicesHandlerPort) magicInfoPort.IMagicInfoDomainPort {
	return &magicInfoUseCases{
		NewListMagicInfoDevicesUseCase(magicInfoPort),
	}
}
