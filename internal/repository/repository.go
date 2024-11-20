package repository

import "Expense-Tracker-API/internal/model"

type Repository interface {
	UserRepository
	ExpenseRepository
}

type UserRepository interface {
	AddUser(data model.User) error
	GetUser(login string) (model.User, error)
}

type ExpenseRepository interface {
	GetAllExpenses(userId int) ([]model.Expense, error)
	AddExpense(data model.Expense) error
	UpdateExpense(data model.Expense) error
	DeleteExpense(expenseId int) error
}
