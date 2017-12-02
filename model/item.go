package model

import "time"

// Item 每一篇文章的简要实体
type Item struct {
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

// InsertIn 插入Feed
func (i *Item) InsertIn() (int, error) {
	affected, err := engine.Insert(i)
	return int(affected), err
}

func (i *Item) GetItem() (has bool, err error) {
	has, err = engine.Get(i)
	return
}

func (i *Item) FindAllItemsByFeedID() (items []Item, err error) {
	err = engine.Where("item.feed_id = ?", i.FeedID).Find(&items)
	return
}
