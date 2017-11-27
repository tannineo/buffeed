package model

import (
	"time"

	"github.com/go-xorm/xorm"
)

// Engine 数据库连接engine
var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "../data.db")
	if err != nil {
		// TODO: logging
		panic(err)
	}
}

// BasicMeta 基础元信息
// 创建 更新 删除 乐观锁
type BasicMeta struct {
	Version int       `xorm:"version"`
	Deleted time.Time `xorm:"deleted"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
