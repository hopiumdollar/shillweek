package main

import (
	"shillweek/httpd/handler"

	"shillweek/platform/shillweek"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := shillweek.New()
	r := gin.Default()

	r.GET("/shillweek", handler.ShillWeekGet(feed))
	r.POST("/shillweek", handler.ShillWeekPost(feed))

	r.Run()
}
