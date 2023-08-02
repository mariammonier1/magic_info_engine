package magicInfoController

import (
	"fmt"
	"magic_info_engine/domain/domainCommon/domainErrors"
	"magic_info_engine/port/domainPort/magicInfoPort"

	"github.com/labstack/echo/v4"
)

func MagicInfoCrudsRouter(router *echo.Echo, domainPort magicInfoPort.IMagicInfoDomainPort) {
	router.GET("/v1/magicInfo-devices", func(context echo.Context) error {
	
		res, err := domainPort.ListMagicInfoDevices()
		if err != nil {
			if fmt.Sprintf("%T", err) == domainErrors.IsDomainError() {
				return context.JSON(err.(*domainErrors.ErrorModel).HttpResp, err)
			}
			return context.JSON(500, domainErrors.CheckResponseStatus(500, err.Error()))
		}
		return context.JSON(200, map[string]interface{}{
			"code":    2001,
			"message": "SUCCESS",
			"result":  res,
		})
	})
	
	}
