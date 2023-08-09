package main

import (
	"api/controllers"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //init the gin router

	models.ConnectDatabase() // new

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/thema", controllers.FindThema)
	r.POST("/thema", controllers.PostThema)

	r.Run()
}
