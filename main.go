package main

import (
	"log"
	"web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./static/css")
	r.Static("/fonts", "./static/fonts")
	r.Static("/images", "./static/images")
	r.Static("/js", "./static/js")
	r.Static("/sass", "./static/sass")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	// r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLGlob("templates/**")
	controller.Router(r)

	log.Println("Server started")
	r.Run(":8090") // default => listen and serve on 0.0.0.0:8080 (for windows "localhost:8090")
}
