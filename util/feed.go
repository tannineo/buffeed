package util

import (
	"github.com/mmcdole/gofeed"
)

var feedParser *gofeed.Parser

func init() {
	feedParser = gofeed.NewParser()
}

// MarshalFeed feed解析
func MarshalFeed(url string) (*gofeed.Feed, error) {
	var feed *gofeed.Feed
	var err error
	if feed, err = feedParser.ParseURL(url); err != nil {
		return nil, err
	}
	return feed, err
}
