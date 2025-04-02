package controllers

import (
	"gin-api-project/config"
	"gin-api-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateItem(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateItemとDeleteItemも同様に実装
