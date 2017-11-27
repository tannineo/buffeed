package model

// User 用户
type User struct {
	ID int64 `xorm:"pk autoincr 'id'"`

	Salt string `xorm:"varchar(64) notnull 'salt'"` // Salt 密码加密用
	Pwd  string `xorm:"varchar(64) notnull 'pwd'"`  // Pwd 密码

	Name  string `xorm:"varchar(64) notnull unique 'name'"`   // Name nickname 昵称
	Email string `xorm:"varchar(128) notnull unique 'email'"` // Email 邮件地址

	Access string `xorm:"varchar(64) notnull 'access'"` // Access 访问等级 TODO:字符串常量定义

	BasicMeta
}
