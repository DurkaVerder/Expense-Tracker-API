package repository_test

import (
	"Expense-Tracker-API/internal/model"
	"Expense-Tracker-API/internal/repository/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresRepo(t *testing.T) {
	connect := "user=postgres password=durka dbname=Expense-Tracker-Test sslmode=disable"
	testRepo := postgres.NewPostgresRepo(connect)

	err := testRepo.CheckConnect()

	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	connect := "user=postgres password=durka dbname=Expense-Tracker-Test sslmode=disable"
	testRepo := postgres.NewPostgresRepo(connect)
	realUser := model.User{Id: 1, Name: "Dev", Login: "develop", Password: "developQ1"}
	user, err := testRepo.GetUser("develop")

	assert.NoError(t, err)
	assert.Equal(t, realUser, user)
}
