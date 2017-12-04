package control_test

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/control"
	"github.com/tannineo/buffeed/model"
	"github.com/tannineo/buffeed/task"
)

const (
	pageJSON = `{"start":0,"size":10}`
)

func Test_ItemGetAllLimit(t *testing.T) {
	// Setup
	model.NewDB()
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/items?start=0&size=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	Convey("Test get all items by page", t, func() {
		So(func() {
			task.GetFeedItems(testModel1.ID, testModel1.Link, e.Logger)
		}, ShouldNotPanic)

		So(func() {
			control.ItemGetAllLimit(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldNotBeEmpty)
		So(rec.Body.String(), ShouldNotEqual, "[]")
	})

	model.TearDownDB()
}

func Test_ItemGetAllLimitByFeed(t *testing.T) {
	// Setup
	model.NewDB()
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/feed/"+testModel1.Hash+"/items?start=0&size=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testModel1.InsertIn()

	Convey("Test get all items by feed hash by page", t, func() {
		So(func() {
			task.GetFeedItems(testModel1.ID, testModel1.Link, e.Logger)
		}, ShouldNotPanic)
		So(func() {
			control.ItemGetAllLimitByFeed(c)
		}, ShouldNotPanic)

		So(rec.Body.String(), ShouldNotBeEmpty)
		So(rec.Body.String(), ShouldNotEqual, "[]")
	})

	model.TearDownDB()
}
