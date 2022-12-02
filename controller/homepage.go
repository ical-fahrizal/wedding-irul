package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", index)
	r.GET("/about", about)
	r.GET("/contact", contact)
	r.GET("/gallery", gallery)
	// r.GET("/services", services)

}

func index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Geeksbeginner",
		},
	)
}

func about(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"about.html",
		gin.H{},
	)
}

func contact(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"contact.html",
		gin.H{
			"title": "Contact",
		},
	)
}

func gallery(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"gallery.html",
		gin.H{
			"title": "Gallery",
		},
	)
}

func services(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"services.html",
		gin.H{
			"title": "Services",
		},
	)
}
