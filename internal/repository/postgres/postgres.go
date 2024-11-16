package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}
