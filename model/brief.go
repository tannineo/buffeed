package model

// Brief 每一篇文章的简要实体
type Brief struct {
	ID int64 `xorm:"pk autoincr"`

	SubID int64

	Title   string `xorm:"varchar(64)"`  // Name nickname 昵称
	Article string `xorm:"varchar(128)"` // Email 邮件地址

	BasicMeta
}
