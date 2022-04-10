package main

import (
	_ "sample_gin_app/entity"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", healthCheckHandler)

	r.Run()
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
