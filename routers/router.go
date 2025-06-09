package routers

import (
	"coinpe/controllers"
	"coinpe/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(app config.App, ctrl controllers.BaseController) {
	app.Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	app.Router.GET("/health", func(ctx *gin.Context) {
		// Send a ping to make sure the database connection is alive.
		db, err := app.DB.DB()
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"live": "not ok"})
			return
		}
		err = db.PingContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{"live": "not ok"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"live": "ok"})
	})

	// Register All routes
	v1Routes(app, ctrl)
}
