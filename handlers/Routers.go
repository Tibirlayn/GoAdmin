package handlers

import (
	"net/http"

	"github.com/Tibirlayn/GoAdmin/config"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {

	config.AccountConfiguration()

	// Обработчик GET запроса на корневой URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Привет, это API сервер на Golang с использованием Gin"})
	})

}