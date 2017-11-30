package model

// Item 每一篇文章的简要实体
type Item struct {
	SubID int64 `xorm:"'sub_id'"`

	Title   string `xorm:"varchar(64) 'title'"`    // Name nickname 昵称
	Article string `xorm:"varchar(128) 'article'"` // Email 邮件地址

	BasicMeta `xorm:"extends"`
}
