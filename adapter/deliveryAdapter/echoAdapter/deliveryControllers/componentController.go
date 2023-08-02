package deliveryControllers

import (
	"magic_info_engine/adapter/deliveryAdapter/echoAdapter/deliveryControllers/componentController"
	"magic_info_engine/port/domainPort/componentPort"
	"magic_info_engine/port/inputPort/deliveryPort"

	"github.com/labstack/echo/v4"
)

type ComponentController struct {
	router     *echo.Echo
	domainPort componentPort.IComponentDomainPort
}

func NewComponentController(router *echo.Echo, domainPort componentPort.IComponentDomainPort) deliveryPort.IController {
	return &ComponentController{
		router,
		domainPort,
	}

}
func (thisPC *ComponentController) Control() {
	componentController.ComponentCrudsRouter(thisPC.router, thisPC.domainPort)
}
