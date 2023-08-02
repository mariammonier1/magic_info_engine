package echoAdapter

import (
	"fmt"
	"log"
	"magic_info_engine/adapter/deliveryAdapter/echoAdapter/deliveryCommon"
	"magic_info_engine/adapter/deliveryAdapter/echoAdapter/deliveryControllers"
	"magic_info_engine/common"
	"magic_info_engine/port/domainPort"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var once sync.Once

// var grant gocloakecho.AuthenticationMiddleWare
var config = common.GetEnvConfig()

type RestEchoImpl struct {
	router *echo.Echo
	domain domainPort.IDomainPort
	//grant  gocloakecho.AuthenticationMiddleWare
}

func (thisRE *RestEchoImpl) InitAdapter(domainPort domainPort.IDomainPort) {
	once.Do(func() {
		thisRE.domain = domainPort
		thisRE.router = echo.New()
		thisRE.router.Binder = &deliveryCommon.CustomBinder{}
		thisRE.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"*",
			},
			AllowCredentials: true,
		}))
		thisRE.router.Use(middleware.Logger())
		thisRE.router.Use(middleware.Gzip())
		thisRE.router.Use(middleware.Secure())

		//ctx := context.Background()
		//client := gocloak.NewClient(config.AuthHost)
		//grant = gocloakecho.NewDirectGrantMiddleware(ctx, client, config.AuthRealm, config.AuthClientId, config.AuthClientSecret, "*", nil)
		//thisRE.grant = grant
	})
	err := thisRE.start(config.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
}

func (thisRE *RestEchoImpl) start(port int) error {
	thisRE.init()
	return thisRE.router.Start(fmt.Sprintf(":%d", port))
}

func (thisRE *RestEchoImpl) init() {
	//customValidator := &deliveryCommon.CustomValidator{Validator: validator.New()}
	//customValidator.Init()
	//thisRE.router.Validator = customValidator

	pingController := deliveryControllers.NewPingController(thisRE.router)
	pingController.Control()

	componentController := deliveryControllers.NewComponentController(thisRE.router, thisRE.domain)
	componentController.Control()

}
