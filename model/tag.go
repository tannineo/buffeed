package model

// Tag 收藏关系
type Tag struct {
	ID int64 `xorm:"pk autoincr"`

	UserID  int64
	SubID   int64
	BriefID int64

	TagName string `xorm:"varchar(64)"` // TagName tag名

	BasicMeta
}
