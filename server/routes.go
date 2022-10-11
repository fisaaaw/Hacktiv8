package server

import (
	"ass3/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	routes := gin.Default()
	routes.GET("/weather", controller.GetWeather)
	routes.POST("/weather", controller.SetWeather)

	routes.Run(":8080")
}
