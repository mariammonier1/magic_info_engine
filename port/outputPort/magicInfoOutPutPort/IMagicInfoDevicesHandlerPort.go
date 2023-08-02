package magicInfoOutputPort

import (
	"magic_info_engine/domain/domainModel"
)

type IMagicInfoDevicesHandlerPort interface {
	ListAllDevices() (*domainModel.MagicInfoResponseMessage, error) 
}
