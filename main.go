package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tannineo/buffeed/route"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 路由
	route.Route(e)

	e.Logger.Fatal(e.Start(":1323"))
}
