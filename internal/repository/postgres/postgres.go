package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(connect string) *PostgresRepo {
	return &PostgresRepo{initDB(connect)}
}

func initDB(connect string) *sql.DB {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("Error open db: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connect db: ", err)
	}

	return db
}

func (r *PostgresRepo) CheckConnect() error {
	if err := r.db.Ping(); err != nil {
		return err
	}
	return nil
}
