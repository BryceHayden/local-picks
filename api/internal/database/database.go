package database

import (
	"errors"

	"github.com/brycehayden/resume/internal/database/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
)


type DBStore struct {
	DB *postgres.PSQL
}

func Initialize(db_user, db_password, db_host, db_port, db_name string, prod_env bool) (*DBStore, error) {
	if db_user == "" || db_password == "" || db_host == "" || db_port == "" || db_name == "" {
		return nil, errors.New("Missing Environment Variable: DB_VARIABLES")
	}

	db, err := postgres.Initialize(db_user, db_password, db_host, db_port, db_name, prod_env)
	if err != nil {
		return nil, err
	}

	return &DBStore{ DB: db }, nil
}
