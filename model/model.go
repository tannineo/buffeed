package model

import (
	"errors"
	"os"
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // go-sqlite3 注册sqlite3到sql
	"github.com/tannineo/buffeed/setting"
)

// Engine 数据库连接engine
var engine *xorm.Engine

func init() {
	NewDB()
}

// BasicMeta 基础元信息
// 创建 更新 删除 乐观锁
type BasicMeta struct {
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// TearDownDB 直接销毁db 测试用
// 危险 DANGER 危险 DANGER
func TearDownDB() {
	if _, err := os.Stat(setting.Config.DataPath); err == nil {
		os.Remove(setting.Config.DataPath)
	}
}

// RewDBIfNotExist 如果db文件不存在则直接新建db 测试用
// 危险 DANGER 危险 DANGER
func RewDBIfNotExist() {
	if _, err := os.Stat(setting.Config.DataPath); err != nil {
		NewDB()
	}
}

// NewDB 直接新建db 测试用
// 危险 DANGER 危险 DANGER
func NewDB() {
	var err error
	if setting.Config.DataPath == "" {
		panic(errors.New("empty sqlite3 data path"))
	} else {
		engine, err = xorm.NewEngine("sqlite3", setting.Config.DataPath)
	}
	if err != nil {
		// TODO: logging
		panic(err)
	}
	engine.Sync2(new(User))
	engine.Sync2(new(Sub))
	engine.Sync2(new(Brief))
	engine.Sync2(new(Tag))
}
