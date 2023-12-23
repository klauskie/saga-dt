package server

import (
	"github.com/gin-gonic/gin"
	"github.com/klauskie/saga-dt/orders/config"
	"strings"
)

type App struct {
	router *gin.Engine
}

func NewApp(env config.Env) App {
	app := App{
		router: gin.Default(),
	}

	// Routes
	app.router.GET("/health", health)

	v1Group := app.router.Group("/api/v1")
	v1Group.POST("/orders", submitOrder(env))

	return app
}

func (a App) Run(port string) error {
	// Listen and Serve
	port = strings.TrimPrefix(port, ":")
	if port == "" {
		port = "8080"
	}
	return a.router.Run(":" + port)
}
