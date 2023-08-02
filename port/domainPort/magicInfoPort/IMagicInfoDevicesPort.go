package magicInfoPort

import "magic_info_engine/domain/domainModel"

type IMagicInfoDevicesPort interface {
	ListMagicInfoDevices() (*domainModel.MagicInfoResponseMessage, error)
}
