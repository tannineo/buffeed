package model

// Tag 收藏关系
type Tag struct {
	ID int64 `xorm:"pk autoincr 'id'"`

	UserID  int64 `xorm:"'user_id'"`
	SubID   int64 `xorm:"'sub_id'"`
	BriefID int64 `xorm:"'brief_id'"`

	TagName string `xorm:"varchar(64) 'tag_name'"` // TagName tag名

	BasicMeta
}
