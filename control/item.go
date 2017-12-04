package control

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/tannineo/buffeed/dto"
	"github.com/tannineo/buffeed/model"
)

// ItemGetAllLimit 根据分页获取items
func ItemGetAllLimit(c echo.Context) (err error) {
	pageDTO := &dto.Page{}
	if err = c.Bind(pageDTO); err != nil {
		return
	}
	c.Logger().Infoj(log.JSON{
		"start": pageDTO.Start,
		"size":  pageDTO.Size,
	})

	items := []model.FeedItem{}
	items, err = model.FindAllItemsLimit(pageDTO.Start, pageDTO.Size)
	if err != nil {
		return
	}
	// write
	views := []map[string]string{}
	for _, v := range items {
		view := map[string]string{
			"title":     v.Title,
			"author":    v.Author,
			"link":      v.Link,
			"content":   v.Content,
			"published": fmt.Sprintf("%d", v.Published.Unix()),
		}
		views = append(views, view)
	}
	return c.JSON(http.StatusOK, views)
}

// ItemGetAllLimitByFeed 根据feed hash获取items
func ItemGetAllLimitByFeed(c echo.Context) (err error) {
	hash := c.Param("hash")
	pageDTO := &dto.Page{}
	if err = c.Bind(pageDTO); err != nil {
		return
	}
	c.Logger().Infoj(log.JSON{
		"hash":  hash,
		"start": pageDTO.Start,
		"size":  pageDTO.Size,
	})

	feedCond := &model.Feed{
		Hash: hash,
	}
	has := false
	has, err = feedCond.GetFeed()
	if err != nil {
		return
	} else if !has {
		return errors.New("cannot find feed, hash:" + hash)
	}

	itemCond := model.FeedItem{
		FeedID: feedCond.ID,
	}
	items := []model.FeedItem{}
	items, err = itemCond.FindAllItemsByFeedIDLimit(pageDTO.Start, pageDTO.Size)
	if err != nil {
		return
	}
	// write
	views := []map[string]string{}
	for _, v := range items {
		view := map[string]string{
			"title":     v.Title,
			"author":    v.Author,
			"link":      v.Link,
			"content":   v.Content,
			"published": fmt.Sprintf("%d", v.Published.Unix()),
		}
		views = append(views, view)
	}
	return c.JSON(http.StatusOK, views)
}
