package service

import (
	"Expense-Tracker-API/internal/model"
	"errors"
)

func (s *ServiceManager) Login(data model.User) error {
	if !s.existUser(data) {
		return errors.New("User not found or invalid password")
	}

	return nil
}

func (s *ServiceManager) Register(data model.User) error {
	if !s.validUserData(data) {
		return errors.New("invalid data")
	}

	err := s.repo.AddUser(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) validUserData(data model.User) bool {
	return s.checkPassword(data.Password) && data.Login != "" && data.Name != ""
}

func (s *ServiceManager) existUser(data model.User) bool {
	user, err := s.repo.GetUser(data.Login)
	if err != nil {
		return false
	}

	return user.Password == data.Password
}

func (s *ServiceManager) checkPassword(password string) bool {
	bigLetters := 0
	smallLetters := 0
	numbers := 0
	for _, symb := range password {
		if symb >= 'A' && symb <= 'Z' {
			bigLetters++
		}
		if symb >= 'a' && symb <= 'z' {
			smallLetters++
		}
		if symb >= '0' && symb <= '9' {
			numbers++
		}
	}

	return bigLetters > 0 && smallLetters > 0 && numbers > 0 && len(password) >= 6
}
