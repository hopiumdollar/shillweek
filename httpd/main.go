package main

import (
	"net/http"
	"shillweek/httpd/handler"

	"shillweek/platform/shillweek"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := shillweek.New()
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	// r.GET("/", handler.ShillWeekGet(feed))
	r.POST("/", handler.ShillWeekPost(feed))

	r.Run()
}
