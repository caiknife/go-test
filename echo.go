package main

import (
	"github.com/labstack/echo"
	"net/http"
	"go-test/model"
)

func main() {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, world")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.GET("/json", json)
	e.GET("/xml", xml)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func json(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")

	result := &model.Profile{
		Team:   team,
		Member: member,
	}

	return c.JSON(http.StatusOK, result)
}

func xml(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")

	result := &model.Profile{
		Team:   team,
		Member: member,
	}

	return c.XML(http.StatusOK, result)
}
