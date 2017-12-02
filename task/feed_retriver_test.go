package task_test

import (
	"testing"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/model"
	"github.com/tannineo/buffeed/task"
)

const (
	feedURL1 = `https://github.com/tannineo.private.atom?token=AfzTKcax1NyoT0AdoI71TRuMGVusfYGhks64F_8EwA==`
)

var testModel1 = model.Feed{
	Hash:     "test",
	UserID:   233,
	UserName: "test_user",
	Alias:    "test_feed",
	Link:     feedURL1,
	JumpLink: "https://github.com/tannineo",
}

func Test_GetFeedItems(t *testing.T) {
	// setup
	model.NewDB()
	testModel1.InsertIn()
	e := echo.New()

	Convey("Test get feed items for "+feedURL1, t, func() {
		task.GetFeedItems(testModel1.ID, testModel1.Link, e.Logger)

		// db验证
		itemCond := &model.Item{
			FeedID: 1,
		}
		items, err := itemCond.FindAllItemsByFeedID()

		So(err, ShouldBeNil)
		So(items, ShouldNotBeEmpty)
	})

	model.TearDownDB()
}
