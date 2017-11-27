package model

import "time"

// Sub 订阅关系
type Sub struct {
	ID int64 `xorm:"pk autoincr"`

	UserID int64

	Alias   string `xorm:"varchar(64)"`  // Alias feed别称
	FeedURL string `xorm:"varchar(128)"` // FeedURL feed地址
	URL     string `xorm:"varchar(128)"` // URL 跳转地址

	LastFetch time.Time `xorm:"created"` // LastFetch 最后一次更新

	BasicMeta
}
