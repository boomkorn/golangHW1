package user

import (
	"app/db"

	"github.com/labstack/echo/v4"
)

func SetupRouter(app *echo.Group, resource *db.Resource) {
	repository := createUserRepository(resource)
	// sessionRepo := session.CreateSessionRepository(resource)

	app.POST("/register", registerHandler(repository))
	app.POST("/login", loginHandler(repository))
}

func SetupPrivateRouter(app *echo.Group, resource *db.Resource) {
	repository := createUserRepository(resource)

	app.GET("", getUserHandler(repository))
	app.GET("/profile", getProfileHandler(repository))
	app.POST("/profile", saveProfileHandler(repository))
	app.POST("/changePassword", changePasswordHandler(repository))
}
