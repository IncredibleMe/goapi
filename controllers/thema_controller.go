package controllers

import (
	"api/models"
	"models"
	"net/http"
	"models"
	"github.com/gin-gonic/gin"
)

type CreateThemaInput struct {
	Onoma   string            `json:"onoma" binding:"required"`
	Patriko string            `json:"patriko" binding:"required"`
	Status  models.status.StatusName `json:"status" binding:"required"`
}

func CreateThema(c *gin.Context) {
	// Validate input
	var input CreateThemaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Thema{Onoma: input.Onoma, Patriko: "patriko", Status: input.Status}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindThema(c *gin.Context) {
	var themata []models.Thema
	models.DB.Find(&themata)

	c.JSON(http.StatusOK, gin.H{"data": themata})
}
