package app

import (
	"app/db"
	"app/service"
	"app/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	router := echo.New()

	// ===== Middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// ===== Initial resource from MongoDB
	dbConnection := db.Connection{"user", "password", "demo", "localhost"}
	resource := dbConnection.CreateConnection()
	defer resource.Close()

	// ===== Prefix of all routes
	publicRoute := router.Group("/api/v1")
	privateRoute := router.Group("/api/v1/users")
	user.SetupRouter(publicRoute, resource)

	config := middleware.JWTConfig{
		Claims:     &service.JWTClaimsModel{},
		SigningKey: []byte("AllYourBase"),
	}
	privateRoute.Use(middleware.JWTWithConfig(config))

	user.SetupPrivateRouter(privateRoute, resource)

	// ===== Start server
	router.Start(":8080")
}
