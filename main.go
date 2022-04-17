package main

import (
	"sample_gin_app/controllers"
	"sample_gin_app/db"
	_ "sample_gin_app/entity"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.Close()

	r := gin.Default()
	r.GET("/", healthCheckHandler)
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.PostUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
