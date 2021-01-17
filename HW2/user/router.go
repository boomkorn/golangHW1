package user

import (
	"app/db"

	"github.com/labstack/echo/v4"
)

func SetupRouter(app *echo.Group, resource *db.Resource) {
	repository := createUserRepository(resource)

	app.GET("/users", getUserHandler(repository))
	app.POST("/users/name", getUserByNameHandler(repository))
	app.POST("/users", postUserHandler(repository))
}
