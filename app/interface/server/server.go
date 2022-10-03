package server

import (
	"ass-david/app/interface/container"
	"ass-david/app/shared/config"
	"ass-david/app/transport"

	"github.com/gin-gonic/gin"
)

func SetupServer(container container.Container) {
	app := gin.Default()
	transport := transport.SetupTransport(container)
	SetupRouter(transport, app)

	app.Run(config.Server.PORTHTTP)
}
