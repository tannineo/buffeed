package model

import (
	"errors"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // go-sqlite3 注册sqlite3到sql
	"github.com/tannineo/buffeed/setting"
)

// Engine 数据库连接engine
var engine *xorm.Engine

func init() {
	NewDB()
}

// TearDownDB 直接销毁db 测试用
// 危险 DANGER 危险 DANGER
func TearDownDB() {
	if _, err := os.Stat(setting.Config.DataPath); err == nil {
		os.Remove(setting.Config.DataPath)
	}
}

// RenewDBIfNotExist 如果db文件不存在则直接新建db 测试用
// 危险 DANGER 危险 DANGER
func RenewDBIfNotExist() {
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
	engine.Sync2(new(Feed))
	engine.Sync2(new(Item))
	engine.Sync2(new(Tag))
}
