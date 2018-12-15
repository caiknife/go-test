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

func home(g *gin.Context) {
	g.String(http.StatusOK, "Hello, world!")
}
