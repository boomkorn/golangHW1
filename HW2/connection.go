package app

import (
	"app/db"
	"app/user"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// ===== Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// ===== Initial resource from MongoDB
	dbConnection := db.Connection{"user", "password", "demo", "localhost"}
	resource := dbConnection.CreateConnection()
	defer resource.Close()

	// ===== Prefix of all routes
	publicRoute := router.Group("/api/v1")
	user.SetupRouter(publicRoute, resource)

	// ===== Start server
	router.Run() // listen and serve on 0.0.0.0:8080
	// can use router.Run({port number}) ex. router.Run(:9999)
}
