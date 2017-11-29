package control

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tannineo/buffeed/model"
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
	return nil
}

// FeedDelete 用户(软)删除feed
func FeedDelete(c echo.Context) (err error) {
	return nil
}

// FeedGetOne 用户获取一个feed的信息
func FeedGetOne(c echo.Context) (err error) {
	return nil
}

// FeedModify 用户修改feed信息
func FeedModify(c echo.Context) (err error) {
	return nil
}
