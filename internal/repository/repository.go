package repository

import "Expense-Tracker-API/internal/model"

type Repository interface {
	UserRepository
	ExpenseRepository
}

type UserRepository interface {
	AddUser(data model.User) error
	GetUser(userId int) (model.User, error)
}

type ExpenseRepository interface {
	GetAllExpenses(userId int) ([]model.Expense, error)
	AddExpense(data model.Expense, userId int) error
	UpdateExpense(data model.Expense, currentExpenseId int) error
	DeleteExpense(expenseId int) error
}
