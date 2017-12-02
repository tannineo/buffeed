package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/tannineo/buffeed/model"
	"github.com/tannineo/buffeed/setting"
	"github.com/tannineo/buffeed/util"
)

// RetriveFeedsWithTicker 根据设置定时更新所有的feed
func RetriveFeedsWithTicker(stopSignal chan struct{}, log echo.Logger) {
	// ticker
	ticker := time.NewTicker(time.Duration(setting.Config.Interval) * time.Minute)
	for {
		<-ticker.C // 时间到啦
		RetriveFeedsNow(log)
	}
	// goroutine拉feed
}

// RetriveFeedsNow 更新所有的feed
func RetriveFeedsNow(log echo.Logger) {
	// 获取所有feed
	feeds, err := model.AllFeeds()
	if err != nil {
		log.Error(err)
		return
	}
	for _, v := range *feeds {
		log.Info(v.Alias + " : " + v.Link + " start to fetch & write items")
		go GetFeedItems(v.ID, v.Link, log)
	}
}

// GetFeedItems 根据link获取
func GetFeedItems(feedID int64, link string, log echo.Logger) {
	// MarshalFeed feed解析
	feed, err := util.MarshalFeed(link)
	if err != nil {
		log.Error(err)
		return
	}
	lenI := len(feed.Items)
	log.Info(fmt.Sprintf("Get %d items...", lenI))
	for i, v := range feed.Items {
		log.Info(strconv.Itoa(i) + " : " + v.Title + " |:| " + v.Content)
		// 写入db TODO: 没有有效的唯一标识item方法 用发布时间 published 替代
		// 插入前的重复验证
		newItem := &model.Item{
			FeedID:    feedID,
			FeedLink:  link,
			Title:     v.Title,
			Published: *v.PublishedParsed,
		}
		has, err := newItem.GetItem()
		if err != nil {
			log.Error(err)
			continue
		} else if has {
			log.Errorf("feed item insert failed : already exists :  feedId=%v item_title=%v published_at:%d", newItem.FeedID, newItem.Title, newItem.Published.Unix())
			continue
		}
		// 插入
		newItem = &model.Item{
			FeedID:    feedID,
			FeedLink:  link,
			Title:     v.Title,
			Published: *v.PublishedParsed,
			Author:    v.Author.Name,
			Link:      v.Link,
			Content:   v.Content,
		}
		affected, err := newItem.InsertIn()
		if err != nil {
			log.Error(err)
			continue
		} else if affected != 1 {
			log.Error("feed item insert failed : affected rows =" + strconv.Itoa(affected))
			continue
		}
	}
}
