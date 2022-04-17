package controllers

import (
	"net/http"
	"sample_gin_app/db"
	"sample_gin_app/entity"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func GetUsers(c *gin.Context) {
	var users []entity.User

	result := db.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, result)
}

func FindUser(c *gin.Context) {
	var user entity.User

	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{Name: input.Name, Age: input.Age, Email: input.Email}

	db.DB.Create(&user)

	c.IndentedJSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user entity.User

	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&user).Updates(input)

	c.IndentedJSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user entity.User

	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&user)

	c.IndentedJSON(http.StatusOK, user)
}
