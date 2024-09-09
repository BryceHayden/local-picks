package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/brycehayden/resume/internal/database"
	"github.com/brycehayden/resume/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

/*
	Main is the launching point of the API. It coordinates environment variables with initial arguments

*/
var (
	api_port	  = os.Getenv("API_PORT")
	db_name        = os.Getenv("POSTGRES_DB")
	db_password   = os.Getenv("POSTGRES_PASSWORD")
	db_user       = os.Getenv("POSTGRES_USER")
	db_port       = os.Getenv("POSTGRES_PORT")
	db_host       = os.Getenv("POSTGRES_HOST")	
	prod_env      = os.Getenv("CURRENT_ENV") == "production"
)

func main() {
	port, err := strconv.Atoi(api_port)
	if err != nil { 
		panic("Missing Environment Variable: API_PORT") 
	}
	
	db, err := database.Initialize(db_user, db_password, db_host, db_port, db_name, prod_env)
	if err != nil {
		panic(fmt.Sprintf("Failed to start db: %s", err))
	}

	server.Initialize(port, db)

	// err := s.ListenAndServe()
	// if err != nil {
	// 	panic(fmt.Sprintf("cannot start server: %s", err))
	// }
}
