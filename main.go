package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tannineo/buffeed/route"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// 中间件
	e.Use(middleware.Logger())

	// 路由
	route.Route(e)

	e.Logger.Fatal(e.Start(":1323"))
}
