package model

import "time"

// Tag 收藏关系
type Tag struct {
	ID      int64     `xorm:"pk autoincr 'id'"`
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`

	UserID  int64 `xorm:"'user_id'"`
	SubID   int64 `xorm:"'sub_id'"`
	BriefID int64 `xorm:"'brief_id'"`

	TagName string `xorm:"varchar(64) 'tag_name'"` // TagName tag名
}
