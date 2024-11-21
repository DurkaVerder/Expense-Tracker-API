package postgres

import (
	"Expense-Tracker-API/internal/model"
	"errors"
)

func (r *PostgresRepo) GetUser(login string) (model.User, error) {
	req := "SELECT * FROM users WHERE login = $1"
	row := r.db.QueryRow(req, login)

	user := model.User{}
	if err := row.Scan(&user.Id, &user.Login, &user.Name, user.Password); err != nil {
		return user, errors.New("Error scan row")
	}

	return user, nil
}

func (r *PostgresRepo) AddUser(data model.User) error {
	req := "INSERT INTO users (login, name, password) VALUES ($1, $2, $3)"
	if _, err := r.db.Exec(req, data.Login, data.Name, data.Password); err != nil {
		return errors.New("Error add user in data base")
	}
	return nil
}
