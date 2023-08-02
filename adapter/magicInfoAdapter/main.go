package magicInfoAdapter

import (
	"magic_info_engine/adapter/magicInfoAdapter/handler"
	magicInfoOutputPort "magic_info_engine/port/outputPort/magicInfoOutPutPort"
)

func NewMagicInfoAdapter() magicInfoOutputPort.IMagicInfoHandlerPort {
	return handler.NewMagicInfoConnection()
}
