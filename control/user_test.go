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
	userJSON1 = `{"name":"cary","email":"Didyouhavemeal@gmail.com","pwd":"23333333"}`
	userJSON2 = `{"name":"ken","email":"bbindream@qq.com","pwd":"23333333"}`
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/user/new", strings.NewReader(userJSON1))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Convey
	Convey("Test create user", t, func() {
		So(func() {
			control.UserCreateUser(c)
		}, ShouldNotPanic)

		// 验证用户插入成功
		userCond := &model.User{
			ID: 1,
		}
		has, err := userCond.GetUserByID()
		So(err, ShouldBeNil)
		So(has, ShouldBeTrue)

		// After
		Reset(func() {
			model.TearDownDB()
		})
	})

}
