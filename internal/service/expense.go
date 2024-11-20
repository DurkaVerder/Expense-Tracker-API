package service

import "Expense-Tracker-API/internal/model"

func (s *ServiceManager) GetExpenses(userId int) ([]model.Expense, error) 

func (s *ServiceManager) AddExpense(data model.Expense) error 

func (s *ServiceManager) UpdateExpense(data model.Expense) error 

func (s *ServiceManager) DeleteExpense(expenseId int) error
