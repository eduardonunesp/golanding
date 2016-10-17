package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":3000")
}
