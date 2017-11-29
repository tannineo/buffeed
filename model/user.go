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
	count, err := engine.Count(new(User))
	return int(count), err
}

// GetUser 根据user中的条件获取用户
func (u *User) GetUser() (has bool, err error) {
	has, err = engine.Get(u)
	return
}

// DeleteUserByID 根据id软删除用户
// 因为删除是个较危险的操作(哪怕是软删除) 会置空其他条件使用id
func (u *User) DeleteUserByID() (int, error) {
	affected, err := engine.Id(u.ID).Delete(u)
	// TODO: affected row 强转 int 在很多很多很多很多...行时会出bug...?
	// 其他也是同理
	return int(affected), err
}

// ModifyUserByID 根据id修改用户
func (u *User) ModifyUserByID() (int, error) {
	affected, err := engine.Id(u.ID).Update(u)
	return int(affected), err
}

// FindUserByNameOrEmail 根据name或Email查找用户
func (u *User) FindUserByNameOrEmail() (*[]User, error) {
	users := &[]User{}
	err := engine.Where("user.name = ?", u.Name).Or("user.email = ?", u.Email).Find(users)
	return users, err
}

// AllUsers 获取所有用户的信息
func AllUsers() (*[]User, error) {
	allUsers := &[]User{}
	err := engine.Find(allUsers)
	return allUsers, err
}

// ModifyUserByName 根据用户昵称(唯一)修改用户信息
// TODO: 存在根据id修改的方法 这么做合理不合理?
func (u *User) ModifyUserByName() (int, error) {
	affected, err := engine.Where("user.name = ?", u.Name).Update(u)
	return int(affected), err
}
