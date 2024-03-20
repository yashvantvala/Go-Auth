package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yashvantvala/Go-Auth/db"
	"github.com/yashvantvala/Go-Auth/routes"
)

func main() {
	fmt.Println("new request came")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("could not load env file")
		panic(err)
	}
	db.ConnectDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
