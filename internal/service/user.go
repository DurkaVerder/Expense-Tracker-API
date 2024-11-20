package service

import "Expense-Tracker-API/internal/model"

func (s *ServiceManager) Login(data model.User) error

func (s *ServiceManager) Register(data model.User) error

func (s *ServiceManager) ValidUserData(data model.User) bool
