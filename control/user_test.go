package control_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/control"
	"github.com/tannineo/buffeed/model"
)

var (
	userJSON1 = `
{"name":"cary","email":"Didyouhavemeal@gmail.com","pwd":"23333333"}`
	userJSON2 = `
{"name":"ken","email":"bbindream@qq.com","pwd":"23333333"}`
	userJSON3f = `
{"name":"whatthepwd","email":"bbindream@qq.com","pwd":"大家308hd08h"}`
)

// TODO: 偷懒把几个接口测试揉在了一起 有空拆开吧
// 测试创建用户
// 测试计算用户总数
// 测试获取单个用户信息
func Test_CreateUser_CountUser_GetUser(t *testing.T) {
	// Setup
	model.NewDB()
	e := echo.New()

	// Convey
	Convey("Test create user1"+userJSON1, t, func() {
		req := httptest.NewRequest(echo.POST, "/sginup", strings.NewReader(userJSON1))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserCreateUser(c)
		}, ShouldNotPanic)

		// 验证用户插入成功
		userCond := &model.User{}
		userCond.ID = 1
		has, err := userCond.GetUser()
		So(err, ShouldBeNil)
		So(has, ShouldBeTrue)
	})

	Convey("Count user 1", t, func() {
		req := httptest.NewRequest(echo.GET, "/users/count", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserGetUserCount(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldEqual, "1")
	})

	Convey("Test get user1: cary", t, func() {
		req := httptest.NewRequest(echo.GET, "/user/cary", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserCreateUser(c)
		}, ShouldNotPanic)

		// 验证用户获取成功
		So(rec.Result().StatusCode, ShouldEqual, http.StatusOK)
		So(rec.Result().Body, ShouldNotBeEmpty)
	})

	Convey("Test create user2"+userJSON1, t, func() {
		req := httptest.NewRequest(echo.POST, "/sginup", strings.NewReader(userJSON2))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserCreateUser(c)
		}, ShouldNotPanic)

		// 验证用户插入成功
		userCond := &model.User{}
		userCond.ID = 2
		has, err := userCond.GetUser()
		So(err, ShouldBeNil)
		So(has, ShouldBeTrue)
	})

	Convey("Count user 2", t, func() {
		req := httptest.NewRequest(echo.GET, "/users/count", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserGetUserCount(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldEqual, "2")

	})

	Convey("Test create user3f"+userJSON1, t, func() {
		req := httptest.NewRequest(echo.POST, "/sginup", strings.NewReader(userJSON3f))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			So(control.UserCreateUser(c), ShouldBeError)
		}, ShouldNotPanic)

		// 验证用户插入失败
		userCond := &model.User{}
		userCond.ID = 3
		has, err := userCond.GetUser()
		So(err, ShouldBeNil)
		So(has, ShouldBeFalse)
	})

	Convey("Test get all users' info", t, func() {
		req := httptest.NewRequest(echo.GET, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserGetAll(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldNotBeEmpty)
	})

	model.TearDownDB()
}
