package router

import (
	"Expense-Tracker-API/internal/controller"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo, controller controller.Controller) {
	e.GET("/api/expense/get", controller.HandleGetExpenses)
	e.POST("/api/expense/add", controller.HandleAddExpense)
	e.PUT("/api/expense/update", controller.HandleUpdateExpense)
	e.DELETE("/api/expense/delete", controller.HandleDeleteExpense)
	e.POST("/api/user/login", controller.HandleLogin)
	e.POST("/api/user/register", controller.HandleRegister)
}
