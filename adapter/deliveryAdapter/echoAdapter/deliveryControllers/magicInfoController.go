package deliveryControllers

import (
	"magic_info_engine/adapter/deliveryAdapter/echoAdapter/deliveryControllers/magicInfoController"
	"magic_info_engine/port/domainPort/magicInfoPort"
	"magic_info_engine/port/inputPort/deliveryPort"

	"github.com/labstack/echo/v4"
)

type MagicInfoController struct {
	router     *echo.Echo
	domainPort magicInfoPort.IMagicInfoDomainPort
}

func NewMagicInfoController(router *echo.Echo, domainPort magicInfoPort.IMagicInfoDomainPort) deliveryPort.IController {
	return &MagicInfoController{
		router,
		domainPort,
	}

}
func (thisPC *MagicInfoController) Control() {
	magicInfoController.MagicInfoCrudsRouter(thisPC.router, thisPC.domainPort)
}
