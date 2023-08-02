package magicInfo

import (
	"magic_info_engine/domain/domainModel"
	"magic_info_engine/port/domainPort/magicInfoPort"
	magicInfoOutputPort "magic_info_engine/port/outputPort/magicInfoOutPutPort"
)

type listMagicInfoDevicesUseCase struct {
	magicInfoOutputPort.IMagicInfoHandlerPort
}

func NewListMagicInfoDevicesUseCase(magicInfoPort magicInfoOutputPort.IMagicInfoHandlerPort) magicInfoPort.IMagicInfoDevicesPort {
	return &listMagicInfoDevicesUseCase{
		magicInfoPort,
	}
}
func (thisLMID *listMagicInfoDevicesUseCase) ListMagicInfoDevices() (*domainModel.MagicInfoResponseMessage, error) {
	devices, err := thisLMID.IMagicInfoHandlerPort.ListAllDevices()
	if err != nil {
		return nil, err
	}
	return devices, nil
}
