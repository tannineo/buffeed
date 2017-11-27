package control

import (
	"errors"
	"net/http"

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
