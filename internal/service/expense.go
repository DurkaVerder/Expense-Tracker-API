package service

import "Expense-Tracker-API/internal/model"

func (s *ServiceManager) GetExpenses(userId int) ([]model.Expense, error) {
	expenses, err := s.repo.GetAllExpenses(userId)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (s *ServiceManager) AddExpense(data model.Expense) error {
	err := s.repo.AddExpense(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) UpdateExpense(data model.Expense) error {
	err := s.repo.UpdateExpense(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) DeleteExpense(expenseId int) error {
	err := s.repo.DeleteExpense(expenseId)
	if err != nil {
		return err
	}

	return nil
}
