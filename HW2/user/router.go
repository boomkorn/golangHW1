package user

import (
	"app/db"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.RouterGroup, resource *db.Resource) {
	repository := createUserRepository(resource)

	app.GET("/users", getUserHandler(repository))
	app.POST("/users/name", getUserByNameHandler(repository))
	app.POST("/users", postUserHandler(repository))
}
