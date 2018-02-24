package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":           "Golanding",
			"subtitle":        "Your favorite landing page",
			"sub_subtitle":    "In bootstrap :D",
			"action_button":   "Let do it",
			"about_title":     "About",
			"services_title":  "Services",
			"portfolio_title": "Portfolio",
			"contact_title":   "Contact",
			"contact_call":    "Let's Get In Touch!",
			"contact_text":    "Ready to start",
			"contact_phone":   "+55 11 1234 1234",
			"contact_email":   "eduardonunesp@gmail.com",
		})
	})
	router.Run(":3000")
}
