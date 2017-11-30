package control_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/tannineo/buffeed/control"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/model"
)

const (
	feedURL1 = `https://github.com/tannineo.private.atom?token=AfzTKcax1NyoT0AdoI71TRuMGVusfYGhks64F_8EwA==`
	feedURL2 = `https://github.com/carynova.private.atom?token=AG7IVadoQ-6k2wapcU1mPA5cRH1aSP8Rks64J9cpwA==`
	feedURL3 = `https://github.com/tannineo.atom`
)

var testModel1 = model.Feed{
	Hash:     "test",
	UserID:   233,
	UserName: "test_user",
	Alias:    "test_feed",
	Link:     feedURL1,
	JumpLink: "https://github.com/tannineo",
}

var testModel2 = model.Feed{
	Hash:     "test2",
	UserID:   233,
	UserName: "test_user",
	Alias:    "test_feed2",
	Link:     feedURL2,
	JumpLink: "https://github.com/carynova",
}

var testModel3 = model.Feed{
	Hash:     "test3",
	UserID:   2333,
	UserName: "test_user",
	Alias:    "test_feed3",
	Link:     feedURL3,
	JumpLink: "https://github.com/tannineo",
}

var feedJSON2 = `{"alias":"test_feed2","link":"` + feedURL2 + `"}`

func Test_FeedGetOne(t *testing.T) {
	// Setup
	model.NewDB()
	testModel1.InsertIn()

	e := echo.New()

	Convey("Get test feed", t, func() {
		req := httptest.NewRequest(echo.GET, "/feed/"+testModel1.Hash, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.FeedGetOne(c)
		}, ShouldNotPanic)

		// TODO: 测试代码中需要一个解析json然后获取限定条目的方法
		// 虽然也可像这样搜子串
		So(rec.Body.String(), ShouldContainSubstring, testModel1.Alias)
	})
	model.TearDownDB()
}

func Test_FeedCreate(t *testing.T) {
	// Setup
	model.NewDB()
	e := echo.New()

	Convey("Create test feed2", t, func() {
		req := httptest.NewRequest(echo.POST, "/feed", strings.NewReader(feedJSON2))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.FeedCreate(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldEqual, "OK")
	})
	model.TearDownDB()
}

func Test_FeedModify(t *testing.T) {
	// Setup
	model.NewDB()
	testModel1.InsertIn()
	e := echo.New()

	Convey("Create modify feed1 user feed2's info", t, func() {
		req := httptest.NewRequest(echo.POST, "/feed/"+testModel1.Hash, strings.NewReader(feedJSON2))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.FeedModify(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldEqual, "OK")

		// db 验证
		feedCond := &model.Feed{
			Hash: testModel1.Hash,
		}
		has, err := feedCond.GetFeed()
		So(err, ShouldBeNil)
		So(has, ShouldBeTrue)
		So(feedCond.Alias, ShouldEqual, "test_feed2")
		So(feedCond.Link, ShouldEqual, feedURL2)
	})
	model.TearDownDB()
}

func Test_FeedGetAll(t *testing.T) {
	model.NewDB()
	e := echo.New()

	Convey("Get all feed info", t, func() {
		req := httptest.NewRequest(echo.GET, "/feeds", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		So(func() {
			control.FeedGetAll(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldNotBeEmpty)
	})
	model.TearDownDB()
}

func Test_FeedDelete(t *testing.T) {
	// Setup
	model.NewDB()
	testModel3.InsertIn()

	e := echo.New()

	Convey("Delete testfeed3", t, func() {
		req := httptest.NewRequest(echo.DELETE, "/feed/"+testModel3.Hash, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		control.FeedDelete(c)

		So(rec.Body.String(), ShouldEqual, "OK")

		// test in db
		feedCond := &model.Feed{
			Hash: testModel3.Hash,
		}
		has, err := feedCond.GetFeed()
		So(err, ShouldBeNil)
		So(feedCond.ID, ShouldEqual, 0)
		So(has, ShouldBeFalse)
	})
	model.TearDownDB()
}
