package componentController

import (
	"fmt"
	"magic_info_engine/domain/domainCommon/domainErrors"
	"magic_info_engine/domain/domainModel"
	"magic_info_engine/port/domainPort/componentPort"

	"github.com/labstack/echo/v4"
)

func ComponentCrudsRouter(router *echo.Echo, domainPort componentPort.IComponentDomainPort) {
	router.POST("/v1/projects/:projectId/pages/:pageId/components", func(context echo.Context) error {
		var component domainModel.Component
		if err := context.Bind(&component); err != nil {
			return context.JSON(400, map[string]interface{}{
				"ValidationError": err.Error(),
			})
		}
		component.ProjectId = context.Param("projectId")
		component.PageId = context.Param("pageId")
		componentId, err := domainPort.CreateComponent(&component)
		if err != nil {
			if fmt.Sprintf("%T", err) == domainErrors.IsDomainError() {
				return context.JSON(err.(*domainErrors.ErrorModel).HttpResp, err)
			}
			return context.JSON(500, domainErrors.CheckResponseStatus(500, err.Error()))
		}
		return context.JSON(200, map[string]interface{}{
			"code":    2001,
			"message": "SUCCESS",
			"result":  componentId,
		})
	})
	
	}
