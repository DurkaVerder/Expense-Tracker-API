package controller

import (
	"Expense-Tracker-API/internal/service"

	"github.com/labstack/echo"
)

type Controller interface {
	HandleLogin(ctx echo.Context) error
	HandleRegister(ctx echo.Context) error
	HandleGetExpenses(ctx echo.Context) error
	HandleAddExpense(ctx echo.Context) error
	HandleUpdateExpense(ctx echo.Context) error
	HandleDeleteExpense(ctx echo.Context) error
}

type ControllerManager struct {
	s *service.Service
}

func NewControllerManager(s *service.Service) *ControllerManager {
	return &ControllerManager(s)
}
