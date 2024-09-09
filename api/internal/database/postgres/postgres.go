package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PSQL struct {
	DB *sqlx.DB
}

func Initialize(db_user, db_password, db_host, db_port, db_name string, prod_env bool) (*PSQL, error) {
	var err error

	sslMode := "disable"
	if prod_env { sslMode = "verify-full" } 
	
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", db_user, db_password, db_host, db_port, db_name, sslMode)	
	
	db, err := sqlx.Open("postgres", uri)
	if err != nil { return nil, err }
	
	return &PSQL{ DB: db }, nil
}



func (psql PSQL) Close() error {
	log.Printf("Disconnected from database: %s", "TODO put db_name here")
	return psql.DB.Close()
}


