package route

import (
	"github.com/labstack/echo"
	"github.com/tannineo/buffeed/control"
)

// Route 来人给echo上route
func Route(e *echo.Echo) {
	// 静态资源
	e.Static("/public", "public")

	// 大概RESTful
	// 首页
	// 用户
	e.POST("/sginup", control.UserCreateUser)
	e.GET("/user/:name", control.UserGetUserByName)
	e.POST("/user/:name", control.ModifyUserAccess)
	e.GET("/users", control.UserGetAll)
	e.GET("/users/count", control.UserGetUserCount)
	// feed
}
