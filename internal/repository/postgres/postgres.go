package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo() *PostgresRepo {
	return PostgresRepo(initDB)
}

func initDB() *sql.DB {
	open := "user=postgres password=durka dbname=Expense-Tracker sslmode=disable"
	db, err := sql.Open("postgres", open)
	if err != nil {
		log.Fatal("Error open db: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connect db: ", err)
	}

	return db
}
