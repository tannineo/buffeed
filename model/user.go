package model

import (
	"errors"

	"github.com/tannineo/buffeed/setting"
	"github.com/tannineo/buffeed/util"
)

// User 用户
type User struct {
	ID int64 `xorm:"pk autoincr 'id'"`

	Salt string `xorm:"varchar(64) notnull 'salt'"` // Salt 密码加密用
	Pwd  string `xorm:"varchar(64) notnull 'pwd'"`  // Pwd 密码

	Name  string `xorm:"varchar(64) notnull unique 'name'"`   // Name nickname 昵称
	Email string `xorm:"varchar(128) notnull unique 'email'"` // Email 邮件地址

	Access AccessGrade `xorm:"varchar(64) notnull 'access'"` // Access 访问等级 TODO:字符串常量定义

	BasicMeta `xorm:"extends"`
}

// AccessGrade 访问级别
type AccessGrade string

// 访问级别常量
const (
	Admin  AccessGrade = "admin"  // Admin 管理员
	Senior             = "senior" // Senior 高级用户
	Novice             = ""       // Novice 普通用户
	Banned             = "banned" // banned 被ban用户
)

// InsertIn 插入user
// pwd此时未摘要 在当中摘要
func (u *User) InsertIn() (err error) {
	// validate
	switch {
	case !util.IsNickname(u.Name):
		return errors.New("invalid nickname")
	case !util.IsPwd(u.Pwd):
		return errors.New("invalid password")
	case !util.IsEmail(u.Email):
		return errors.New("invalid email")
	}

	u.Salt = setting.Config.Salt
	u.Pwd = util.GetMd5String(u.Pwd)
	count := 0
	u.Access = Novice
	if count, err = CountUsers(); err != nil {
		return
	} else if count == 0 {
		u.Access = Admin
	}

	// insert
	engine.Insert(u)
	return
}

// CountUsers 未删除的用户数
func CountUsers() (int, error) {
	user := new(User)
	count, err := engine.Count(user)
	return int(count), err
}

// GetUserByID 根据id获取用户
func (u *User) GetUserByID() (has bool, err error) {
	has, err = engine.Get(u)
	return
}

// GetUserByName 根据名称获取用户

// GetUserByEmail 根据email获取用户

// DeleteUserByID 根据id软删除用户

// ModifyUserByID 根据id修改用户
