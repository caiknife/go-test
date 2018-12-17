package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.New()

	g.GET("/", home)

	g.Run(":8888")
}

type Person struct {
	Team   string `json:"team" xml:"team"`
	Member string `json:"member" xml:"member"`
}

func home(c *gin.Context) {
	team := c.Query("team")
	member := c.Query("member")

	result := &Person{
		team,
		member,
	}

	c.JSON(http.StatusOK, result)
}
