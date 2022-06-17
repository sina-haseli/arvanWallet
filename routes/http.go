package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"wallet/handler"

	_ "wallet/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func RegisterRoutes(e *echo.Echo, handler *handler.BaseHandler) {
	api := e.Group("/api")
	api.GET("/balance/:user_id", handler.Credit.HandleGetBalanceRequest())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
