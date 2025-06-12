package routers

import (
	"coinpe/controllers"
	"coinpe/pkg/config"
)

func v1Routes(app config.App, ctrl controllers.BaseController) {
	v1 := app.Router.Group("/v1")

	accountGroup := v1.Group("/accounts")
	accountGroup.POST("", ctrl.CreateAccount)

	v1.POST("/authenticate", ctrl.Authenticate)
	v1.POST("/verify", ctrl.VerifyAuthenticate)
}
