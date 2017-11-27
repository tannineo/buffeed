package model

import "time"

// Sub 订阅关系
type Sub struct {
	ID int64 `xorm:"pk autoincr 'id'"`

	UserID int64 `xorm:"'user_id'"`

	Alias   string `xorm:"varchar(64) 'alias'"`     // Alias feed别称
	FeedURL string `xorm:"varchar(128) 'feed_url'"` // FeedURL feed地址
	URL     string `xorm:"varchar(128) 'url'"`      // URL 跳转地址

	LastFetch time.Time `xorm:"'last_fetch'"` // LastFetch 最后一次更新

	BasicMeta `xorm:"extends"`
}
