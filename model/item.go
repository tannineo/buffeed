package model

import "time"

// Item 每一篇文章的简要实体
type FeedItem struct {
	ID      int64     `xorm:"pk autoincr 'id'"`
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`

	FeedID   int64  `xorm:"'feed_id'"`
	FeedLink string `xorm:"varchar(256) 'feed_link'"`

	Title     string    `xorm:"varchar(64) 'title'"`     // Title 标题
	Published time.Time `xorm:"'published'"`             // Published 发布时间
	Author    string    `xorm:"var(64) 'author'"`        // Author 作者
	Link      string    `xorm:"varchar(256) 'link'"`     // Link 跳转链接
	Content   string    `xorm:"varchar(4096) 'content'"` // Content 内容
}

// InsertIn 插入FeedItem
func (f *FeedItem) InsertIn() (int, error) {
	affected, err := engine.Insert(f)
	return int(affected), err
}

// GetItem 根据条件获取单挑item
func (i *FeedItem) GetItem() (has bool, err error) {
	has, err = engine.Get(i)
	return
}

// FindAllItemsLimit 分页获取所有feed的items 以published作为排序
func FindAllItemsLimit(startRow, limitRow int) (items []FeedItem, err error) {
	err = engine.Desc("published").Limit(limitRow, startRow).Find(&items)
	return
}

// FindAllItemsByFeedIDLimit 根据feedID分页获取items 以published作为排序
func (i *FeedItem) FindAllItemsByFeedIDLimit(startRow, limitRow int) (items []FeedItem, err error) {
	err = engine.Where("feed_item.feed_id = ?", i.FeedID).Desc("published").Limit(limitRow, startRow).Find(&items)
	return
}
