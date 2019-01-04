package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-test/model"
)

func main() {
	g := gin.New()

	g.GET("/", home)

	g.GET("/json", showJson)
	g.GET("/xml", showXml)

	g.Run(":8888")
}

func home(c *gin.Context) {
	team := c.DefaultQuery("team", "caiknife")
	member := c.DefaultQuery("member", "caiknife")

	result := &model.Profile{
		Team:   team,
		Member: member,
	}

	c.JSON(http.StatusOK, result)
}

func showJson(c *gin.Context) {
	team := c.DefaultQuery("team", "caiknife")
	member := c.DefaultQuery("member", "caiknife")

	result := &model.Profile{
		Team:   team,
		Member: member,
	}

	c.JSON(http.StatusOK, result)
}

func showXml(c *gin.Context) {
	team := c.DefaultQuery("team", "caiknife")
	member := c.DefaultQuery("member", "caiknife")

	result := &model.Profile{
		Team:   team,
		Member: member,
	}

	c.XML(http.StatusOK, result)
}
