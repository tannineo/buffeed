package control

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/mmcdole/gofeed"
	"github.com/tannineo/buffeed/model"
	"github.com/tannineo/buffeed/util"
)

// FeedGetAll 获取所有feed信息
func FeedGetAll(c echo.Context) (err error) {
	var feeds *[]model.Sub
	if feeds, err = model.AllSubs(); err != nil {
		return
	}
	// 包装
	var feedInfos = make([]map[string]string, 0)
	for _, v := range *feeds {
		c.Logger().Info("feed url: " + v.FeedURL)
		feedInfos = append(feedInfos, map[string]string{
			"hash":            v.Hash,
			"user_name":       v.UserName,
			"alias":           v.Alias,
			"feed_url":        v.FeedURL,
			"url":             v.URL,
			"last_fetch_time": fmt.Sprintf("%d", v.LastFetch.Unix()),
			"created":         fmt.Sprintf("%d", v.Created.Unix()),
		})
	}
	return c.JSON(http.StatusOK, feedInfos)
}

// FeedCreate 用户创建(上传)feed
func FeedCreate(c echo.Context) (err error) {
	// dto
	newFeed := struct { // 验证密码是否输入两边是客户端的动作
		Alias   string
		FeedURL string
	}{}
	// TODO: 获取登录用户 session机制
	userID := 233
	userName := "foobar"

	// bind
	if err = c.Bind(&newFeed); err != nil {
		return
	}
	// alias 不应该像name和pwd一样限制严格
	// TODO: 但是SQL注入咋办?
	c.Logger().Infoj(log.JSON{
		"alias":    newFeed.Alias,
		"feed_url": newFeed.FeedURL,
	})

	// 查找重复 feed url
	subCond := &model.Sub{
		FeedURL: newFeed.FeedURL,
	}
	has := false
	has, err = subCond.GetSub()
	if err != nil {
		return
	}
	if has {
		return c.String(http.StatusOK, "feed url duplicated")
	}

	// 获取并验证feed url
	var feed *gofeed.Feed
	if feed, err = util.MarshalFeed(newFeed.FeedURL); err != nil {
		return
	}

	// write
	sub := &model.Sub{
		Hash:     util.GetMd5String(newFeed.FeedURL + fmt.Sprintf("%d", time.Now().Unix())),
		UserID:   int64(userID),
		UserName: userName,
		Alias:    newFeed.Alias,
		FeedURL:  newFeed.FeedURL,
		URL:      feed.Link,
	}

	affected := 0
	if affected, err = sub.InsertIn(); err != nil {
		return
	} else if affected != 1 {
		return errors.New("sub insert affected row = " + strconv.Itoa(affected))
	}

	return c.String(http.StatusOK, "OK")
}

// FeedDelete 用户(软)删除feed
// 想不通这里多次调用时的错误处理
// 保持幂等? 还是让用户有权知道delete数次?
// TODO: 说到底幂等到底是对于外部来说还是内部来说?
func FeedDelete(c echo.Context) (err error) {
	hash := c.Param("hash")
	feedInfo := &model.Sub{
		Hash: hash,
	}
	has := false
	if has, err = feedInfo.GetSub(); err == nil && has {
		affected := 0
		if affected, err = feedInfo.DeleteByID(); err != nil {
			return err
		} else if affected != 1 {
			return errors.New("sub delete soft affected row = " + strconv.Itoa(affected))
		}
		return c.String(http.StatusOK, "OK")
	}
	return c.String(http.StatusNotFound, "")
}

// FeedGetOne 用户获取一个feed的信息
// TODO: 删除连带的feed item(没人tag的)
// TODO: 引入事务
func FeedGetOne(c echo.Context) (err error) {
	hash := c.Param("hash")
	feedInfo := &model.Sub{
		Hash: hash,
	}
	has := false
	if has, err = feedInfo.GetSub(); err == nil && has {
		// 存在feed
		// LastFetch 不可能为nil 零值是某个很早的时间
		// see https://godoc.org/time#Time
		return c.JSON(http.StatusOK, &map[string]string{
			"hash":       feedInfo.Hash,
			"alias":      feedInfo.Alias,
			"feed_url":   feedInfo.FeedURL,
			"url":        feedInfo.URL,
			"last_fetch": fmt.Sprintf("%d", feedInfo.LastFetch.Unix()),
			"created":    fmt.Sprintf("%d", feedInfo.Created.Unix()),
			"user_name":  feedInfo.UserName,
		})
	}
	// 不存在feed
	return c.String(http.StatusNotFound, "")
}

// FeedModify 用户修改feed信息
func FeedModify(c echo.Context) (err error) {
	hash := c.Param("hash")
	// dto
	feedMod := struct {
		Alias   string
		FeedURL string
	}{}

	// bind
	if err = c.Bind(&feedMod); err != nil {
		return
	}
	c.Logger().Infoj(log.JSON{
		"alias":    feedMod.Alias,
		"feed_url": feedMod.FeedURL,
	})
	// TODO: 对于一个feed hash取出动作有好几个重复 可以重构成middleware?
	// TODO: 话说Sub可以改名成Feed吧???
	feedInfo := &model.Sub{
		Hash: hash,
	}
	has := false
	if has, err = feedInfo.GetSub(); err == nil && has {
		// 存在feed
		feedInfo.Alias = feedMod.Alias
		feedInfo.FeedURL = feedMod.FeedURL
		affected := 0
		if affected, err = feedInfo.ModifyAliasAndFeedURLByID(); err != nil {
			return err
		} else if affected != 1 {
			return errors.New("sub delete soft affected row = " + strconv.Itoa(affected))
		}
		return c.String(http.StatusOK, "OK")
	}
	// 不存在feed
	return c.String(http.StatusNotFound, "")
}
