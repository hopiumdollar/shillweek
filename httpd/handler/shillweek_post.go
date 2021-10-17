package handler

import (
	"net/http"

	"shillweek/platform/shillweek"

	"github.com/gin-gonic/gin"
)

type shillweekPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func ShillWeekPost(feed shillweek.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := shillweekPostRequest{}
		c.Bind(&requestBody)

		item := shillweek.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
