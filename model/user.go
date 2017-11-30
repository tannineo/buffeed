package model

import (
	"github.com/tannineo/buffeed/setting"
	"github.com/tannineo/buffeed/util"
	"time"
)

// User 用户
type User struct {
	ID      int64     `xorm:"pk autoincr 'id'"`
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`

	Salt string `xorm:"varchar(64) notnull 'salt'"` // Salt 密码加密用
	Pwd  string `xorm:"varchar(64) notnull 'pwd'"`  // Pwd 密码

	Name  string `xorm:"varchar(64) notnull unique 'name'"`   // Name nickname 昵称
	Email string `xorm:"varchar(128) notnull unique 'email'"` // Email 邮件地址

	Access AccessGrade `xorm:"varchar(64) notnull 'access'"` // Access 访问等级
}

// AccessGrade 访问级别
type AccessGrade string

// 访问级别常量
const (
	Admin  AccessGrade = "admin"  // Admin 管理员
	Senior             = "senior" // Senior 高级用户
	Novice             = "novice" // Novice 普通用户
	Banned             = "banned" // banned 被ban用户
)

// InsertIn 插入user
// pwd此时未摘要 在当中摘要
func (u *User) InsertIn() (int, error) {
	// TODO: 有必要在model再次做一次验证?
	u.Salt = setting.Config.Salt
	u.Pwd = util.GetMd5String(u.Pwd)
	u.Access = Novice
	if count, err := CountUsers(); err != nil {
		return 0, err
	} else if count == 0 {
		u.Access = Admin
	}

	// insert
	affected, err := engine.Insert(u)
	return int(affected), err
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
	// 还有许多地方存在这种代码 review 和 重构要注意
	// 其他也是同理
	return int(affected), err
}

// ModifyUserByID 根据id修改用户
func (u *User) ModifyUserByID() (int, error) {
	affected, err := engine.Id(u.ID).Update(u)
	return int(affected), err
}

// ModifyUserAccessByID 根据id修改access
func (u *User) ModifyUserAccessByID() (int, error) {
	affected, err := engine.Id(u.ID).Cols("access").Update(u)
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
