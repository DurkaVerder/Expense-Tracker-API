package service

import "Expense-Tracker-API/internal/repository"

type Service interface {
	Login(data model.User) error
	Register(data model.User) error
	GetExpenses(userId int) ([]model.Expense, error)
	AddExpense(data model.Expense) error
	UpdateExpense(data model.Expense) error
	DeleteExpense(expenseId int) error	
}

type ServiceManager struct {
	repo *repository.Repository
}

func NewServiceManager(repo *repository.Repository) *ServiceManager {
	return &ServiceManager{repo}
}
