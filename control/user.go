package control

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/tannineo/buffeed/model"
	"github.com/tannineo/buffeed/util"
)

// UserCreateUser 用户注册
func UserCreateUser(c echo.Context) (err error) {
	// dto
	newU := struct { // 验证密码是否输入两边是客户端的动作
		Name  string
		Pwd   string
		Email string
	}{}

	// bind
	if err = c.Bind(&newU); err != nil {
		return
	}
	c.Logger().Infoj(log.JSON{
		"name":  newU.Name,
		"email": newU.Email,
	})

	// validate
	switch {
	case !util.IsNickname(newU.Name):
		return errors.New("invalid nickname")
	case !util.IsPwd(newU.Pwd):
		return errors.New("invalid password")
	case !util.IsEmail(newU.Email):
		return errors.New("invalid email")
	}

	// 查找重复
	dupUsers := &[]model.User{}
	dupUsers, err = (&model.User{
		Name:  newU.Name,
		Email: newU.Email,
	}).FindUserByNameOrEmail()
	if err != nil {
		return
	} else if len(*dupUsers) > 0 {
		return c.String(http.StatusOK, "user duplicate")
	}

	// write
	userModel := model.User{
		Name:  newU.Name,
		Pwd:   newU.Pwd,
		Email: newU.Email,
	}
	if err = (&userModel).InsertIn(); err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

// UserGetUserCount 获取用户总数
func UserGetUserCount(c echo.Context) (err error) {
	var count = 0
	if count, err = model.CountUsers(); err != nil {
		return
	}
	return c.String(http.StatusOK, strconv.Itoa(count))
}

// UserGetUserByName 获取用户简介
func UserGetUserByName(c echo.Context) (err error) {
	name := c.Param("name")
	userInfo := &model.User{
		Name: name,
	}
	if has, err := userInfo.GetUser(); err == nil && has {
		// 存在用户
		return c.JSON(http.StatusOK, &map[string]string{
			"name":    userInfo.Name,
			"email":   userInfo.Email,
			"access":  string(userInfo.Access),
			"created": fmt.Sprintf("%d", userInfo.Created.Unix()),
		})
	}
	// 不存在用户
	return c.String(http.StatusNotFound, "")
}

// UserGetAll 获取所有用户
// 鉴于不太可能有很多用户
func UserGetAll(c echo.Context) (err error) {
	var users *[]model.User
	if users, err = model.AllUsers(); err != nil {
		return
	}
	// 包装
	var userInfos = make([]map[string]string, 0)
	for _, v := range *users {
		c.Logger().Info("name: " + v.Name)
		userInfos = append(userInfos, map[string]string{
			"name":    v.Name,
			"email":   v.Email,
			"access":  string(v.Access),
			"created": fmt.Sprintf("%d", v.Created.Unix()),
		})
	}
	return c.JSON(http.StatusOK, userInfos)
}
