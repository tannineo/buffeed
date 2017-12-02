package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tannineo/buffeed/route"
	"github.com/tannineo/buffeed/setting"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// 中间件
	e.Use(middleware.Logger())

	// 路由
	route.Route(e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(setting.Config.Port))))
}
