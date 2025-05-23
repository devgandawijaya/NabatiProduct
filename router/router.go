package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, userController *controller.UserController) {
	e.GET("/product", userController.GetAllProduct)
}
