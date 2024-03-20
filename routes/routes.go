package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/yashvantvala/Go-Auth/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok!"})
	})
	server.POST("/signup", user.SignUp)
	server.POST("/login", user.Login)
}
