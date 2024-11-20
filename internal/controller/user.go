package controller

import (
	"Expense-Tracker-API/internal/model"

	"github.com/labstack/echo"
)

func (c *ControllerManager) HandleLogin(ctx echo.Context) error {
	enterData := model.User{}
	if err := ctx.Bind(&enterData); err != nil {
		ctx.JSON(400, "Invalid data")
	}

	if err := c.service.Login(enterData); err != nil {
		return ctx.JSON(400, "Invalid data")
	}

	if err := c.cookie.SaveJWTInCookie(ctx, enterData.Id); err != nil {
		return ctx.JSON(500, "Error save cookie")
	}

	return ctx.JSON(200, "Login")
}

func (c *ControllerManager) HandleRegister(ctx echo.Context) error {
	enterData := model.User{}

	if err := ctx.Bind(&enterData); err != nil {
		return ctx.JSON(400, "Invalid data")
	}

	if err := c.service.Register(enterData); err != nil {
		return ctx.JSON(500, "Error register user")
	}

	return ctx.JSON(200, "Register")
}



