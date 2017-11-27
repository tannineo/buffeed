package model

// User 用户
type User struct {
	ID int64 `xorm:"pk autoincr"`

	Salt string `xorm:"varchar(64) notnull"` // Salt 密码加密用
	Pwd  string `xorm:"varchar(64) notnull"` // Pwd 密码

	Name  string `xorm:"varchar(64) notnull unique"`  // Name nickname 昵称
	Email string `xorm:"varchar(128) notnull unique"` // Email 邮件地址

	Access string `xorm:"varchar(64) notnull"` // Access 访问等级 TODO:字符串常量定义

	BasicMeta
}
