package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// TmplDir templete dir
var TmplDir = filepath.Join(os.Getenv("GOPATH"), "src/github.com/andmorefine/golang_copilot/templates/*")

func main() {
	// success
	log.Println("success")

	router := gin.Default()
	router.LoadHTMLGlob(TmplDir)

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

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":  "Main website",
			"body":   "body",
			"status": http.StatusText(http.StatusOK),
		})
	})

	router.Run(":80")
}
