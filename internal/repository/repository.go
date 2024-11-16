package repository

import "Expense-Tracker-API/internal/model"

type Repository interface {
	UserRepository
	ExpenseRepository
}

type UserRepository interface {
	AddUser(data model.User)
}

type ExpenseRepository interface {
	
}
