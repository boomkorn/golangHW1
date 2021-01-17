package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUserHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		users, _ := repo.getAllUser()
		code := http.StatusOK
		return c.JSON(code, users)
	}
}
func getUserByNameHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := ReqGetUserByName{}
		// if err := c.Bind(&body); err != nil {
		// return c.JSON(http.StatusBadRequest, echo.H{"error": err.Error()})
		// }
		users, _ := repo.getByName(body)
		code := http.StatusOK
		return c.JSON(code, users)
	}
}

func postUserHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := ReqPostUser{}
		// if err := c.Bind(&body); err != nil {
		// return c.JSON(http.StatusBadRequest, echo.H{"error": err.Error()})
		// }

		// if body.Name == "" || body.Age == 0 {
		// 	return c.JSONP(http.StatusBadRequest, Users{})
		// } else {
		users, _ := repo.postUser(body)
		code := http.StatusOK
		return c.JSON(code, users)
		// }
	}
}
