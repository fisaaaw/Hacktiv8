package server

import (
	"swag-tutor/app/transport"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "swag-tutor/docs"
	// "../../../docs"
)


// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
// @schemes http
func SetUpRouter(tp *transport.Tp, app *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	route := app.Group("/v1")
	{
		route.GET("/person", tp.Transport.GetAllPersons)
		route.GET("/person/:id", tp.Transport.GetOnePerson)
		route.POST("/person", tp.Transport.CreatePerson)
		route.PUT("/person/:id", tp.Transport.UpdatePersonRequest)
		route.DELETE("/person/:id", tp.Transport.DeletePerson)
	}
}