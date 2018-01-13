package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.tmpl.html", nil)
	})

	router.GET("/blog/cellphones", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cellphones.tmpl.html", nil)
	})

	router.GET("/music", func(c *gin.Context) {
		c.HTML(http.StatusOK, "music.tmpl.html", nil)
	})

	router.Run(":" + port)
}
