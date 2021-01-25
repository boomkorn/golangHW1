package user

import (
	"app/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func registerHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(ReqRegister)
		code := http.StatusOK
		if err := c.Bind(body); err != nil {
			// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			return c.JSON(code, ResponseDefault{400, "Input not match, please try again"})
		} else if err := repo.register(body); err != nil {
			// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			return c.JSON(code, ResponseDefault{400, "Something wrong, please try again"})
		}
		return c.JSON(http.StatusOK, ResponseDefault{200, "Registeration success"})
	}
}

func loginHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(ReqLogin)
		code := http.StatusOK
		if err := c.Bind(body); err != nil {
			// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			return c.JSON(code, ResponseDefault{400, "Input not match, please try again"})
		}
		user, err := repo.login(body)
		if err != nil {
			// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			return c.JSON(code, ResponseDefault{400, fmt.Sprintf("Something wrong, please try again : %v", err)})
		}
		jwt, err := service.GenerateJWT(user.Id.Hex())
		return c.JSON(code, ResponseLogin{200, "Login success", ReponseWithToken{user, jwt}})
	}
}

func saveProfileHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(ReqSaveProfile)
		token := c.Get("user").(*jwt.Token)

		code := http.StatusOK
		if err := c.Bind(body); err != nil {
			return c.JSON(code, ResponseDefault{400, "Input not match, please try again"})
		}

		claims := token.Claims.(*service.JWTClaimsModel)
		user, err := repo.saveProfile(claims.UserId, body)

		if err != nil {
			return c.JSON(code, ResponseDefault{400, fmt.Sprintf("Something wrong, please try again : %v", err)})
		}
		return c.JSON(code, ResponseWithUser{200, "Save profile success", user})
	}
}

func changePasswordHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(ReqChangePassword)
		token := c.Get("user").(*jwt.Token)

		code := http.StatusOK
		if err := c.Bind(body); err != nil {
			return c.JSON(code, ResponseDefault{400, "Input not match, please try again"})
		}

		claims := token.Claims.(*service.JWTClaimsModel)
		err := repo.updatePassword(claims.UserId, body)

		if err != nil {
			return c.JSON(code, ResponseDefault{400, fmt.Sprintf("Something wrong, please try again : %v", err)})
		}
		return c.JSON(code, ResponseDefault{200, "Change password success"})
	}
}

func getProfileHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		code := http.StatusOK

		claims := token.Claims.(*service.JWTClaimsModel)
		user, err := repo.getProfile(claims.UserId)

		if err != nil {
			return c.JSON(code, ResponseDefault{400, fmt.Sprintf("Something wrong, please try again : %v", err)})
		}
		return c.JSON(code, ResponseWithUser{200, "Get profile success", user})
	}
}

func getUserHandler(repo repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		users, _ := repo.getAllUser()
		code := http.StatusOK
		return c.JSON(code, users)
	}
}
