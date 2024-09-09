package server

import (
	"strconv"

	"github.com/brycehayden/resume/internal/database"
	"github.com/brycehayden/resume/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initialize(port int, db *database.DBStore) {
	echoApp := echo.New()

	echoApp.Use(middleware.Logger())
	echoApp.Use(middleware.Recover())

	routes.Setup(echoApp, db)	

	echoApp.Logger.Fatal(echoApp.Start(":" + strconv.Itoa(port)))
}

// func Initialize() *http.Server { 
// 	port, _ := strconv.Atoi(os.Getenv("API_PORT"))
// 	NS := &Server{
// 		port: port,
// 		db:   database.New(),
// 	}

// 	// Declare Server config
// 	server := &http.Server{
// 		Addr:         fmt.Sprintf(":%d", NS.port), 
// 		Handler:      NS.RegisterRoutes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 	}

// 	return server
// }
