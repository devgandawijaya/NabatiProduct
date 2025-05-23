package controller

import (
	"go-api/service"
	"go-api/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) GetAllProduct(ctx echo.Context) error {
	product, err := c.Service.GetAllProduct()
	log.Printf("Error GetAllProduct: %v", err)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve product")
	}
	return utils.Respond(ctx, http.StatusOK, product, "Success")
}
