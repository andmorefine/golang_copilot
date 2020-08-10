package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// success
	log.Println("success")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run(":80")
}
