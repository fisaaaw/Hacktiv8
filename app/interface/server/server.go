package server

import (
	"fmt"
	"swag-tutor/app/interface/container"
	"swag-tutor/app/transport"
	"swag-tutor/app/shared/config"

	"github.com/gin-gonic/gin"
)

func StartServer(container container.Container) {

	app := gin.Default()

	transport  := transport.SetupTransport(container)
	SetUpRouter(transport, app)

	fmt.Println(app.Run(config.Server.PORTHTTP))
}