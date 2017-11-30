package model

import "time"

// Item 每一篇文章的简要实体
type Item struct {
	ID      int64     `xorm:"pk autoincr 'id'"`
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`

	SubID int64 `xorm:"'sub_id'"`

	Title   string `xorm:"varchar(64) 'title'"`    // Name nickname 昵称
	Article string `xorm:"varchar(128) 'article'"` // Email 邮件地址
}
