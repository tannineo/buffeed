package control_test

import (
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

func TestMain(m *testing.M) {
	// before

	// run test
	m.Run()

	// after
	model.TearDownDB()
}

// 测试创建用户
func Test_CreateUser_CountUser(t *testing.T) {
	// Setup
	e := echo.New()

	// Convey
	Convey("Test create user1"+userJSON1, t, func() {
		req := httptest.NewRequest(echo.POST, "/user/new", strings.NewReader(userJSON1))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserCreateUser(c)
		}, ShouldNotPanic)

		// 验证用户插入成功
		userCond := &model.User{
			ID: 1,
		}
		has, err := userCond.GetUser()
		So(err, ShouldBeNil)
		So(has, ShouldBeTrue)
	})

	Convey("Count user 1", t, func() {
		req := httptest.NewRequest(echo.GET, "/user/count", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserGetUserCount(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldEqual, "1")

	})

	Convey("Test create user2"+userJSON1, t, func() {
		req := httptest.NewRequest(echo.POST, "/user/new", strings.NewReader(userJSON2))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserCreateUser(c)
		}, ShouldNotPanic)

		// 验证用户插入成功
		userCond := &model.User{
			ID: 2,
		}
		has, err := userCond.GetUser()
		So(err, ShouldBeNil)
		So(has, ShouldBeTrue)
	})

	Convey("Count user 2", t, func() {
		req := httptest.NewRequest(echo.GET, "/user/count", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.UserGetUserCount(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldEqual, "2")

	})

	Convey("Test create user3f"+userJSON1, t, func() {
		req := httptest.NewRequest(echo.POST, "/user/new", strings.NewReader(userJSON3f))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			So(control.UserCreateUser(c), ShouldBeError)
		}, ShouldNotPanic)

		// 验证用户插入失败
		userCond := &model.User{
			ID: 3,
		}
		has, err := userCond.GetUser()
		So(err, ShouldBeNil)
		So(has, ShouldBeFalse)
	})

}

// TestGetUser 测试获取用户信息
func TestGetUser(t *testing.T) {

}
