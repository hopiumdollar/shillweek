package handler

import (
	"net/http"

	"shillweek/platform/shillweek"

	"github.com/gin-gonic/gin"
)

func ShillWeekGet(feed shillweek.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
