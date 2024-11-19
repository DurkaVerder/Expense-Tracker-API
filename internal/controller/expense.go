package controller

import (
	"Expense-Tracker-API/internal/model"

	"github.com/labstack/echo"
)

func (c *ControllerManager) HandleGetExpenses(ctx echo.Context) error {
	userId, err := c.cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(401, "Invalid data")
	}

	expenses, err := c.service.GetExpenses(userId)
	if err != nil {
		return ctx.JSON(500, "Error get expenses")
	}

	return ctx.JSON(200, expenses)
}

func (c *ControllerManager) HandleAddExpense(ctx echo.Context) error {
	userId, err := c.cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(401, "Invalid data")
	}

	enterData := model.Expense{}
	if err := ctx.Bind(&enterData); err != nil {
		return ctx.JSON(400, "Invalid data")
	}

	enterData.CreatorId = userId
	if err := c.service.AddExpense(enterData); err != nil {
		return ctx.JSON(500, "Error add expense")
	}

	return ctx.JSON(200, "Add expense")
}

func (c *ControllerManager) HandleUpdateExpense(ctx echo.Context) error {
	newData := model.Expense{}
	if err := ctx.Bind(&newData); err != nil {
		return ctx.JSON(400, "Invalid data")
	}

	if err := c.service.UpdateExpense(newData); err != nil {
		return ctx.JSON(500, "Error update expense")
	}

	return ctx.JSON(200, "Update expense")
}

func (c *ControllerManager) HandleDeleteExpense(ctx echo.Context) error {
	expense := model.Expense{}
	if err := ctx.Bind(&expense); err != nil {
		return ctx.JSON(400, "Invalid data")
	}

	if err := c.service.DeleteExpense(expense.Id); err != nil {
		return ctx.JSON(500, "Error delete expense")
	}

	return ctx.JSON(200, "Delete expense")
}
