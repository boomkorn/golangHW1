package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserHandler(repo repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		users, _ := repo.getAllUser()
		code := http.StatusOK
		c.JSON(code, users)
	}
}
func getUserByNameHandler(repo repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := ReqGetUserByName{}
		if err := c.Bind(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users, _ := repo.getByName(body)
		code := http.StatusOK
		c.JSON(code, users)
	}
}

func postUserHandler(repo repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := ReqPostUser{}
		if err := c.Bind(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if body.Name == "" || body.Age == 0 {
			c.JSON(http.StatusBadRequest, Users{})
		} else {
			users, _ := repo.postUser(body)
			code := http.StatusOK
			c.JSON(code, users)
		}
	}
}
