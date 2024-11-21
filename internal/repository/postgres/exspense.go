package postgres

import (
	"Expense-Tracker-API/internal/model"
	"errors"
)

func (r *PostgresRepo) GetAllExpenses(userId int) ([]model.Expense, error) {
	req := "SELECT * FROM expenses WHERE creator_id = $1"
	rows, err := r.db.Query(req, userId)
	if err != nil {
		return nil, errors.New("Error get expenses")
	}
	expenses := []model.Expense{}
	for rows.Next() {
		expense := model.Expense{}
		if err := rows.Scan(&expense.Id, &expense.Title, &expense.Cost, &expense.CreatedData, &expense.CreatorId); err != nil {
			return nil, errors.New("Error scan")
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func (r *PostgresRepo) AddExpense(data model.Expense) error {
	req := "INSERT INTO expenses (title, cost, creator_id) VALUES ($1, $2, $3)"
	if _, err := r.db.Exec(req, data.Title, data.Cost, data.CreatorId); err != nil {
		return errors.New("Error add expense")
	}

	return nil
}

func (r *PostgresRepo) UpdateExpense(data model.Expense) error {
	req := "UPDATE expenses SET title = $1 cost = $2"
	if _, err := r.db.Exec(req, data.Title, data.Cost); err != nil {
		return errors.New("Error update expense")
	}
	return nil
}

func (r *PostgresRepo) DeleteExpense(expenseId int) error {
	req := "DELETE FROM expenses WHERE id = $1"
	if _, err := r.db.Exec(req, expenseId); err != nil {
		return errors.New("Error delete expense")
	}
	return nil
}
