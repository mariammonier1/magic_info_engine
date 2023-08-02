package deliveryControllers

import (
	"magic_info_engine/port/inputPort/deliveryPort"

	"github.com/labstack/echo/v4"
)

type PingController struct {
	router *echo.Echo
	//grant  gocloakecho.AuthenticationMiddleWare
}

func NewPingController(router *echo.Echo) deliveryPort.IController {
	return &PingController{
		router,
		//grant,
	}
}

func (thisPC *PingController) Control() {
	thisPC.ping()
}

func (thisPC *PingController) ping() {
	thisPC.router.GET("v1/ping", func(c echo.Context) (err error) {
		return c.JSON(200, map[string]interface{}{
			"message": "pong",
		})
	})
}
