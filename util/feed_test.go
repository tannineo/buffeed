package util_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tannineo/buffeed/util"
)

const (
	// atom
	feedURL1 = `https://github.com/tannineo.private.atom?token=AfzTKcax1NyoT0AdoI71TRuMGVusfYGhks64F_8EwA==`
	// rss
	feedURL2 = `http://coolshell.cn/feed`
)

func Test_MarshalFeed(t *testing.T) {
	Convey("Marshal Feed 1 - my github feed", t, func() {
		feed, err := util.MarshalFeed(feedURL1)
		So(err, ShouldBeNil)
		So(feed, ShouldNotBeNil)
		So(feed.FeedType, ShouldEqual, "atom")
		So(feed.Link, ShouldEqual, `https://github.com/tannineo`)
		So(feed.Title, ShouldEqual, "Private Feed for tannineo")
	})

	Convey("Marshal Feed 2 - codeshell", t, func() {
		feed, err := util.MarshalFeed(feedURL2)
		So(err, ShouldBeNil)
		So(feed, ShouldNotBeNil)
		So(feed.FeedType, ShouldEqual, "rss")
		So(feed.Link, ShouldEqual, `https://coolshell.cn`)
		So(feed.Title, ShouldEqual, "酷 壳 &#8211; CoolShell") // so entities in html won't be transformed...
	})
}
