package routes

import (
	"net/http"

	"github.com/brycehayden/resume/internal/database"
	"github.com/labstack/echo/v4"
)
type Router struct {
	server *echo.Echo
	db *database.DBStore
}

func Setup (server *echo.Echo, db *database.DBStore) {
	// server.GET("/health", s.healthHandler)
	r := Router{
		server: server,
		db: db,
	}

	apiGroup := server.Group("/v1")

	apiGroup.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "API Router")
	})
	
	login := apiGroup.Group("/login")
	login.POST("", r.Login)


	// TODO add middleware to verify JWT && on some routes role is admin
	search := apiGroup.Group("/search")
	search.POST("", r.searchRestaurant)

	restaurant := apiGroup.Group("/restaurants")
	restaurant.GET("", r.getRestaurants)
	restaurant.GET("/:id", r.getRestaurantDetails)
	


	admin := apiGroup.Group("/admin")
	admin.POST("/restaurants", r.createRestaurants)

	
}
