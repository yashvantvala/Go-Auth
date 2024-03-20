package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashvantvala/Go-Auth/models"
	"github.com/yashvantvala/Go-Auth/utils"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	err = user.CreateUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User has been created successfully!"})
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	err = user.FindUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could find the user with this account"})
		return
	}
	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success!", "token": token})
}
